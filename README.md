# irsem

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/llir/irsem)

### Usage

```go
import (
    "github.com/llir/llvm/ir"
    "github.com/llir/irsem"
)

mod := ir.NewModule()
// mod ...
irsem.Verify(mod)
```
