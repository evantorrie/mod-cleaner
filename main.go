package main

import (
	"fmt"

	"go.opentelemetry.io/otel/api/kv"
)

func main() {
	fmt.Println("hello world");

	x := kv.Bool("just-a-test", true);
	fmt.Printf("output a key: %v\n", x);
}
