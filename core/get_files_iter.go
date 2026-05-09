package core

import (
	"fmt"
	"iter"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetFiles(conf *Config) iter.Seq[string] {
	return func(yield func(string) bool) {
		root := "."

		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if conf.IgnoreHiddenFiles && strings.HasPrefix(d.Name(), ".") && d.Name() != "." {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}

			rel, _ := filepath.Rel(root, path)
			depth := 0
			if rel != "." {
				depth = len(strings.Split(rel, string(os.PathSeparator)))
			}

			if d.IsDir() {
				if conf.MaxDepth > 0 && depth >= conf.MaxDepth {
					return filepath.SkipDir
				}
				return nil
			}

			if !conf.matchesRules(path) {
				return nil
			}

			if !yield(path) {
				return filepath.SkipAll
			}

			return nil
		})

		if err != nil && conf.Verbose {
			fmt.Printf("Error walking directory: %v\n", err)
		}
	}
}

func (c *Config) matchesRules(path string) bool {
	if c.IncludePattern != "" {
		matched, err := regexp.MatchString(c.IncludePattern, path)
		if err != nil {
			if c.Verbose {
				fmt.Printf("Regex error (include): %v\n", err)
			}
			return false
		}
		if !matched {
			return false
		}
	}

	if c.IgnorePattern != "" {
		matched, err := regexp.MatchString(c.IgnorePattern, path)
		if err != nil {
			if c.Verbose {
				fmt.Printf("Regex error (ignore): %v\n", err)
			}
			return true
		}
		if matched {
			return false
		}
	}

	return true
}
