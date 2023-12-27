this library use data [APIHariLibur_V2](https://github.com/guangrei/APIHariLibur_V2).

installation

```
go get github.com/cirebon-dev/go_holiday
```

example usage:

```go
package main

import (
    "fmt"
    "github.com/cirebon-dev/go_holiday"
)

func main() {
	result, err := go_holiday.Holiday("2023-06-01")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if result == "" {
		fmt.Println("Tidak libur!")
	} else {
		fmt.Println(result)
	}
}
``` 