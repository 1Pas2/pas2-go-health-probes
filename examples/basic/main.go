package main

import (
	"context"
	"fmt"
	lib "github.com/1Pas2/pas2-go-health-probes"
)

func main() {
	results, err := lib.Run(context.Background(), map[string]lib.Probe{"database": func(context.Context) error { return nil }}, 2)
	fmt.Println(results, err)
}
