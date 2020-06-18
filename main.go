package main

import (
	"fmt"

	"go.opentelemetry.io/otel/api/key"
)

func main() {
	fmt.Println("hello world");

	x := key.Bool("just-a-test", true);
	fmt.Printf("output a key: %v\n", x);
}
