## Below Prompt is used to generate the README file



type Config struct {
    Encrypt           bool
    Decrypt           bool
    Zip               bool
    Verbose           bool
    Rename            bool
    TestMode          bool
    IgnoreHiddenFiles bool
    KeepOriginal      bool
    MaxDepth          int
    Password          string
    IgnorePattern     string
    IncludePattern    string
}


func printUsage() {
    fmt.Println("Usage: ./gohide [commands] [options]")
    fmt.Println("\nCommands (can be combined):")
    fmt.Println("  e : Encrypt")
    fmt.Println("  d : Decrypt")
    fmt.Println("  z : Zip result")
    fmt.Println("  v : Verbose")
    fmt.Println("  n : No rename")
    fmt.Println("  t : Run In test mode")
    fmt.Println("  h : Include hidden files & directories")
    fmt.Println("  k : Keep original files after encryption or descryption")

    fmt.Println("\nOptions:")
    fmt.Println("  -i [pattern] : Include pattern")
    fmt.Println("  -g [pattern] : Ignore pattern")
    fmt.Println("  -p [password]: Provide password (insecure: shows in history)")
    fmt.Println("  -m [depth]   : Provide max-recursive depth, default is 3")
    fmt.Println()
}

am releasing v2, write me a README.md for the project

zip functionality still in the works

my app is a tool to help encrypt any kind file in a drectory in a recursive manner, it currently works on the CWD (specifiy src and dst are also in the works)

conf := Config{Rename: true, IgnoreHiddenFiles: true, MaxDepth: 3}
default config sets Rename,ignoreHiddenFiles and maxdepths above	

Make sure to add install instractions from sources and from prebuilt binary provided by CI/CD pipeline)
Installation from binary at the top 
  GOOS=linux   GOARCH=amd64 go build -o dist/gohide-linux-amd64   .
          GOOS=windows GOARCH=amd64 go build -o dist/gohide-windows-amd64.exe .
          GOOS=darwin  GOARCH=amd64 go build -o dist/gohide-darwin-amd64  .
          GOOS=darwin  GOARCH=arm64 go build -o dist/gohide-darwin-arm64  .

Examples can use -p [password] format but NOTE: pattern is not recommended, the app will prompt form password securely if -p is not used

- Use emojis, eg on the titles for markdown