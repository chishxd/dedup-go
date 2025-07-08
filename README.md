# dedup-go

**dedup-go** is a simple, fast, idiomatic Go tool to detect and remove duplicate files in a directory by computing their content hashes.  
It’s designed to be a lightweight, real-world example of writing robust Go CLIs and learning file I/O, hashing, and concurrency.

---

## ✨ Features

- Scan any target directory recursively
- Compute SHA-256 hashes for file comparison
- Report duplicates and optionally delete them
- Clean, modular Go project structure
- Ready for use as a CLI or as a library

---

## 📦 Installation

```bash
# Clone the repo
git clone https://github.com/chishxd/dedup-go.git

# Initialize dependencies
cd dedup-go
go mod tidy

# Run the CLI
go run ./cmd/dedup.go ./path/to/your/target_dir
````

---

## ⚙️ Usage

Basic usage:

```bash
dedup-go --path ./target_dir --delete
```

| Flag       | Description                     |
| ---------- | ------------------------------- |
| `--path`   | Path to the directory to scan   |
| `--delete` | Remove duplicates automatically |

---

## 🗂️ Project Structure

```
dedup-go/
├── cmd/            # CLI entry point
│   └── dedup.go
├── internal/       # Core logic
│   ├── hash.go
│   ├── dedup.go
│   └── utils.go
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

---

## 🚀 Roadmap

* [ ] Add unit tests
* [ ] Add concurrency for faster hashing
* [ ] Add logging and verbose output
* [ ] Build cross-platform binaries

---

## 🤝 Contributing

Pull requests are welcome! Please open an issue first to discuss major changes.
Make sure to follow Go best practices and run `go fmt` before submitting.

---

## 📄 License

This project is licensed under the MIT License — see [LICENSE](./LICENSE) for details.

---

**Built with 💚 by [Chinmay Shet](https://github.com/chishxd).**
