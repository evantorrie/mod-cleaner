package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"go.opentelemetry.io/otel/api/kv"
)

func main() {
	fmt.Println("hello world");

	x := kv.Bool("just-a-test", true);
	w := false;
	fmt.Printf("output a key: %v\n", x);
	y := true;
	if diff := cmp.Diff(w, y); diff != "" {
		fmt.Printf("diff:\n%s", diff)
	}
}
