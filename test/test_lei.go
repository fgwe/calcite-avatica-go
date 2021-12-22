package main

import (
	"fmt"
	avatica "github.com/fgwe/calcite-avatica-go/v5"
)

func main() {
	c, config, _ := avatica.ParseDSN("http://localhost:30060/test?user=root&password=root")
	fmt.Println(config)
	fmt.Println(c)
}
