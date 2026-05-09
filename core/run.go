package core

import (
	"fmt"
	"sync"

	"com.gosafe/utils"
)

func Run() error {
	c := ParseArgs()

	if c.Verbose {
		fmt.Printf("Mode: Encrypt=%v, Decrypt=%v, Zip=%v, Rename=%v\n",
			c.Encrypt, c.Decrypt, c.Zip, c.Rename)
	}

	if c.Decrypt || c.Encrypt {
		password, err := utils.GetPassword()
		if err != nil {
			fmt.Printf("%v\n", err)
			return err
		}
		c.Password = password
	}

	actionName := utils.Ternary(c.Encrypt, "Encrypt", "Decrypt")
	errs_count := 0

	var wg sync.WaitGroup
	// Limit concurrency to 10 files at a time to avoid crashing the OS
	semaphore := make(chan struct{}, 10)

	// Mutex to safely increment error count from multiple goroutines
	var mu sync.Mutex

	for file := range GetFiles(c) {
		if c.TestMode {
			fmt.Printf("[TEST] Would %s: %s\n", actionName, file)
			continue
		}

		wg.Add(1)
		semaphore <- struct{}{} // Block if 10 goroutines are already running

		go func(f string) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release slot when finished

			action := utils.Ternary(c.Encrypt, EncryptFile, DecryptFile)
			err := action(f, c)

			if err != nil {
				mu.Lock()
				errs_count++
				mu.Unlock()
				if c.Verbose {
					fmt.Printf("Error processing %s: %v\n", f, err)
				}
			}
		}(file)
	}

	wg.Wait()

	if errs_count > 0 {
		fmt.Printf("%s completed with %d errors\n", actionName, errs_count)
	}

	return nil
}
