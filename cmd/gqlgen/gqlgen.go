//go:generate go run ./gqlgen.go
package main

import (
	"fmt"
	"github.com/99designs/gqlgen/cmd"
)

func main() {
	fmt.Println("Building Graphql schema")
	cmd.Execute()
}
