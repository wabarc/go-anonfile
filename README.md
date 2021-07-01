# go-anonfile

`go-anonfile` is a toolkit to help upload files to [anonfiles](https://anonfiles.com).

## Installation

The simplest, cross-platform way is to download from [GitHub Releases](https://github.com/wabarc/go-anonfile/releases) and place the executable file in your PATH.

Via Golang package get command

```sh
go get -u github.com/wabarc/go-anonfile/cmd/anonfile
```

From [gobinaries.com](https://gobinaries.com):

```sh
$ curl -sf https://gobinaries.com/wabarc/go-anonfile | sh
```

## Usage

Command-line:

```sh
$ anonfile
A CLI tool help upload files to anonfiles.

Usage:

  anonfile [options] [file1] ... [fileN]
```

Go package:
```go
import (
        "fmt"

        "github.com/wabarc/go-anonfile"
)

func main() {
        if url, err := anonfile.NewAnonfile(nil).Upload(path); err != nil {
            fmt.Fprintf(os.Stderr, "anonfiles: %v\n", err)
        } else {
            fmt.Fprintf(os.Stdout, "%s  %s\n", url, path)
        }
}
```

## License

This software is released under the terms of the MIT. See the [LICENSE](https://github.com/wabarc/go-anonfile/blob/main/LICENSE) file for details.
