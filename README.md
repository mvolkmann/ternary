# ternary

This is a set of Go functions that provide
a substitute for the missing ternary operator.

## Example

```go
package main

import (
  "fmt"

  t "github.com/mvolkmann/ternary"
)

func main() {
  temperature := 70
  color := t.String(temperature > 100, "red", "blue")
  fmt.Printf("color = %s\n", color)

  color = t.Any(temperature > 100, "red", "blue").(string)
  fmt.Printf("color = %s\n", color)
  fmt.Printf("color type = %T\n", color)
}
```
