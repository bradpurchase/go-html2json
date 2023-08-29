# go-html2json

Go package to convert an HTML string to JSON.

## Usage

```go
package main

import (
  "fmt"
  "github.com/bradpurchase/go-html2json"
)

func main() {
  html := "<div><p>Hello World</p></div>"
  json, err := html2json.Convert(html)
  if err != nil {
    panic(err)
  }
  fmt.Println(json) // {"tag":"html","attrs":{},"children":[{"tag":"head","attrs":{},"children":null},{"tag":"body","attrs":{},"children":[{"tag":"div","attrs":{},"children":[{"tag":"p","attrs":{},"children":[{"tag":"text","attrs":{"text":"Hello World"},"children":[{"tag":"","attrs":null,"children":null}]}]}]}]}]}
}
```
