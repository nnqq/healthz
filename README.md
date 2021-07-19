# healthz

HTTP healthcheck

```
go get github.com/nnqq/healthz
```

```go
package main

import "github.com/nnqq/healthz"

func main() {
    h := healthz.NewHealthz()
    err := h.Serve()
	if err != nil {
        panic(err)
	}
}
```
