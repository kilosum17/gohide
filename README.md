# 🔐 gohide v2

Recursive file encryption for entire directories — fast, simple, and written in Go.

`gohide` helps encrypt files inside the current working directory recursively, making it useful for backups, archival protection, portable encrypted storage, and bulk file security workflows.

> ⚠️ Zip functionality is currently under development in v2.

---

# ✨ Features

* 🔒 Encrypt files recursively
* 🔓 Decrypt encrypted files
* 📂 Recursive directory traversal
* 🙈 Ignore hidden files/directories by default
* 🏷️ Optional filename randomization/renaming
* 🧪 Test mode support
* 📢 Verbose logging
* 🎯 Include and ignore patterns
* 📏 Configurable recursion depth
* 🖥️ Cross-platform builds (Linux, Windows, macOS)

---

# ⚙️ Default Configuration

```go
conf := Config{
    Rename: true,
    IgnoreHiddenFiles: true,
    MaxDepth: 3,
}
```

Default behavior:

| Setting                | Default   |
| ---------------------- | --------- |
| Rename encrypted files | ✅ Enabled |
| Ignore hidden files    | ✅ Enabled |
| Max recursion depth    | `3`       |

---

# 📦 Installation

# 🚀 Install Prebuilt Binary (Recommended)

Prebuilt binaries are provided through the CI/CD pipeline releases.

## Available Builds

| Platform                    | Binary                     |
| --------------------------- | -------------------------- |
| Linux AMD64                 | `gohide-linux-amd64`       |
| Windows AMD64               | `gohide-windows-amd64.exe` |
| macOS AMD64                 | `gohide-darwin-amd64`      |
| macOS ARM64 (Apple Silicon) | `gohide-darwin-arm64`      |

## Linux / macOS

```bash
chmod +x gohide-linux-amd64
./gohide-linux-amd64
```

## Windows

```powershell
gohide-windows-amd64.exe
```

---

# 🛠️ Build From Source

## Requirements

* Go 1.22+ recommended

## Clone Repository

```bash
git clone https://github.com/yourusername/gohide.git
cd gohide
```

## Build

### Linux AMD64

```bash
GOOS=linux GOARCH=amd64 go build -o dist/gohide-linux-amd64 .
```

### Windows AMD64

```bash
GOOS=windows GOARCH=amd64 go build -o dist/gohide-windows-amd64.exe .
```

### macOS AMD64

```bash
GOOS=darwin GOARCH=amd64 go build -o dist/gohide-darwin-amd64 .
```

### macOS ARM64 (Apple Silicon)

```bash
GOOS=darwin GOARCH=arm64 go build -o dist/gohide-darwin-arm64 .
```

---

# 🚀 Usage

```bash
./gohide [commands] [options]
```

---

# 🎮 Commands

Commands can be combined together.

| Command | Description                                     |
| ------- | ----------------------------------------------- |
| `e`     | Encrypt                                         |
| `d`     | Decrypt                                         |
| `z`     | Zip result *(WIP)*                              |
| `v`     | Verbose mode                                    |
| `n`     | Disable renaming                                |
| `t`     | Test mode                                       |
| `h`     | Include hidden files & directories              |
| `k`     | Keep original files after encryption/decryption |

---

# 🧩 Options

| Option          | Description             |
| --------------- | ----------------------- |
| `-i [pattern]`  | Include pattern         |
| `-g [pattern]`  | Ignore pattern          |
| `-p [password]` | Provide password        |
| `-m [depth]`    | Maximum recursive depth |

---

# 🔐 Password Handling

You can provide a password directly:

```bash
./gohide ev -p mypassword
```

However:

> ⚠️ Using `-p` is NOT recommended because the password may appear in:
>
> * shell history
> * process lists
> * terminal logs

If `-p` is omitted, `gohide` securely prompts for the password interactively.

Recommended usage:

```bash
./gohide ev
```

---

# 📚 Examples

## 🔒 Encrypt Current Directory

```bash
./gohide e
```

---

## 🔒 Encrypt Verbosely

```bash
./gohide ev
```

---

## 🔓 Decrypt Files

```bash
./gohide d
```

---

## 🔒 Encrypt While Keeping Originals

```bash
./gohide ek
```

---

## 🔒 Encrypt Including Hidden Files

```bash
./gohide eh
```

---

## 🔒 Encrypt Only `.txt` Files

```bash
./gohide e -i "*.txt"
```

---

## 🔒 Ignore `node_modules`

```bash
./gohide e -g "node_modules"
```

---

## 🔒 Set Max Recursive Depth

```bash
./gohide e -m 5
```

---

## 🧪 Test Mode

```bash
./gohide etv
```

Test mode simulates actions without modifying files.

---

# 📂 Current Limitations

* Currently operates on the **Current Working Directory (CWD)** only
* Source and destination directory selection is still in development
* Zip/archive support is still being implemented
* Recursive traversal depth is limited by `MaxDepth`

---

# 🗺️ Planned Features

* 📁 Custom source directory
* 📤 Custom destination directory
* 📦 Stable zip/archive support
* ⚡ Parallel encryption
* 🧾 Better logging and reporting
* 🔑 Keyfile support
* 🧠 Smarter pattern matching

---

# ⚠️ Important Notes

* Always test on backup/sample files first
* Keep your password safe — encrypted data cannot be recovered without it
* Recursive encryption can affect many files quickly
* Be careful when using broad include patterns

---

# ❤️ Contributing

Contributions, bug reports, feature requests, and pull requests are welcome.

---

# 📜 License

MIT License — see the LICENSE file for details.

---

# 🔥 Example Help Output

```text
Usage: ./gohide [commands] [options]

Commands (can be combined):
  e : Encrypt
  d : Decrypt
  z : Zip result
  v : Verbose
  n : No rename
  t : Run In test mode
  h : Include hidden files & directories
  k : Keep original files after encryption or descryption

Options:
  -i [pattern] : Include pattern
  -g [pattern] : Ignore pattern
  -p [password]: Provide password (insecure: shows in history)
  -m [depth]   : Provide max-recursive depth, default is 3
```
