# **Vivaran CLI — Flag & Command Guide**

Vivaran is a `wc`-like command-line utility written in Go that reads either **a file path** or **standard input (stdin)** and prints text statistics.

## **Supported Flags**

| Flag | Name       | Description                                                                                  | Example               |
| ---- | ---------- | -------------------------------------------------------------------------------------------- | --------------------- |
| `-c` | Size       | Outputs the file size in bytes                                                               | `go run ./cmd/main.go -c test.txt` |
| `-l` | Lines      | Counts the number of lines                                                                   | `go run ./cmd/main.go -l test.txt` |
| `-w` | Words      | Counts words separated by any whitespace                                                     | `go run ./cmd/main.go -w test.txt` |
| `-m` | Characters | Counts Unicode characters (multibyte-safe). Same as `-c` when input has no multibyte support | `go run ./cmd/main.go -m test.txt` |

You can combine multiple flags.

## **Usage Patterns**

### **1. Read from a file**

Provide a file path as a positional argument:

```sh
go run ./cmd/main.go test.txt
```

Output (default all stats):

```
342190 7145 58164 339292
```

### **2. Read from stdin**

Pipe input from another command:

```sh
echo "hello world" | go run ./cmd/main.go
```

Output:

```
12 1 2 12
```

### **3. Use specific flags with file**

```sh
go run ./cmd/main.go -l -w test.txt
```

Output:

```
7145 58164
```

### **4. Use flags with stdin**

```sh
cat test.txt | go run ./cmd/main.go -c -l -w -m
```

Output:

```
342190 7145 58164 339292
```

### **5. Character counting (`-m`)**

Unicode-aware mode counts runes, not raw bytes:

```sh
wc -m test.txt
go run ./cmd/main.go -m test.txt
```

Both should match exactly ✅


## **Important Rules Followed by Vivaran**

* Accepts **only one input source** (file **or** stdin).
* Panics if both are missing.
* Panics if more than one file argument is passed together with stdin.


## **Sample Commands**

```sh
# Size in bytes
go run ./cmd/main.go -c test.txt

# Line count
go run ./cmd/main.go -l test.txt

# Word count
go run ./cmd/main.go -w test.txt

# Unicode character count
go run ./cmd/main.go -m test.txt

# Combined flags
go run ./cmd/main.go -c -l -w -m test.txt

# With stdin
cat test.txt | go run ./cmd/main.go -l -m
```

## **Expected Output Formatting**

Vivaran prints numbers space-separated and adds the filename only when a file is provided:

```
<size> <lines> <words> <chars>
```

Example:

```bash
─ go run ./cmd/main.go test.txt                                                                        

size    Lines  words   char 
342190  7145   58164   339292
```
