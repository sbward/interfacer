# interfacer [![GoDoc](https://godoc.org/github.com/sbward/interfacer?status.png)](https://godoc.org/github.com/sbward/interfacer)

Code generation tools for Go's interfaces.

Tools available in this repository:

- [cmd/interfacer](#cmdinterfacer-)

### cmd/interfacer [![GoDoc](https://godoc.org/github.com/sbward/interfacer/cmd/interfacer?status.png)](https://godoc.org/github.com/rjeczalik/interfaces/cmd/interfacer)

Generates an interface for a named type.

*Installation*
```bash
~ $ go get github.com/rjeczalik/interfaces/cmd/interfacer
```

*Usage*

```bash
~ $ interfacer -help
```
```
Usage of interfacer:
  -all
        Include also unexported methods.
  -as string
        Generated interface name. (default "main.Interface")
  -for string
        Type to generate an interface for.
  -o string
        Output file. (default "-")
```

*Example*
- generate by manually
```bash
~ $ interfacer -for os.File -as mock.File
```
- generate by go generate
```go
//go:generate interfacer -for os.File -as mock.File -o file_iface.go
```
```bash
~ $ go generate  ./...
```
- output
```go
// Created by interfacer; DO NOT EDIT

package mock

import (
        "os"
)

// File is an interface generated for "os".File.
type File interface {
        Chdir() error
        Chmod(os.FileMode) error
        Chown(int, int) error
        Close() error
        Fd() uintptr
        Name() string
        Read([]byte) (int, error)
        ReadAt([]byte, int64) (int, error)
        Readdir(int) ([]os.FileInfo, error)
        Readdirnames(int) ([]string, error)
        Seek(int64, int) (int64, error)
        Stat() (os.FileInfo, error)
        Sync() error
        Truncate(int64) error
        Write([]byte) (int, error)
        WriteAt([]byte, int64) (int, error)
        WriteString(string) (int, error)
}
```
