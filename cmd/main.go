package main

import (
	"fmt"

	"github.com/AxiomSamarth/go-release-consume/pkg/operations"
)

func main() {
	fmt.Printf("Sum: %d", operations.FindSum(1, 2))
}
