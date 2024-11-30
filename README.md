[![Go Report Card](https://goreportcard.com/badge/github.com/mskrha/gototp)](https://goreportcard.com/report/github.com/mskrha/gototp)

## gototp

### Description
Very simple Go library for generating TOTP codes. Simplified version of [https://github.com/xlzd/gotp](https://github.com/xlzd/gotp) supporting just TOTP.

### Installation
`go get github.com/mskrha/gototp`

### Example usage
```go
package main

import (
	"fmt"

	"github.com/mskrha/gototp"
)

func main() {
	t, err := gototp.NewDefault("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Generate())
}
```
