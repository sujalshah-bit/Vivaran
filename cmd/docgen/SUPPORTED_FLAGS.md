## **Supported Flags**

| Flag | Name | Description | Example |
| ---- | ---- | ----------- | ------- |
| `-c` | Size | Outputs the file size in bytes | `go run ./cmd/main.go -c test.txt` |
| `-l` | Lines | Counts the number of lines | `go run ./cmd/main.go -l test.txt` |
| `-w` | Words | Counts words separated by whitespace | `go run ./cmd/main.go -w test.txt` |
| `-m` | Characters | Counts Unicode characters (multibyte-safe) | `go run ./cmd/main.go -m test.txt` |
| `-bs` | Buffer size | Set buffer size | `go run ./cmd/main.go -bs test.txt` |
