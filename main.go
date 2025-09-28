package main

import (
	"flag"
	"fmt"
	"github.com/tlopo-go/flat/lib/flat"
	"io"
	"os"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func main() {
	flag.Usage = func() {
		fmt.Println("flat [OPTS]")
		flag.PrintDefaults()
	}

	separator := flag.String("s", " | ", "separator")
	flag.Parse()

	content := string(Must(io.ReadAll(os.Stdin)))

	r := flat.New().Content(content).Separator(*separator).Run()
	for _, k := range r.Keys() {
		v, _ := r.Get(k)
		fmt.Printf("%s = %s\n", k, v)
	}
}
