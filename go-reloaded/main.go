package main

import(
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(`Usage: go run . sample.txt result.txt`)
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error: unable to readFile", err)
		os.Exit(1)
	}
	result := Processor(string(data))
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	fmt.Println(result)
}
