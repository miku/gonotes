package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// os.Stdout.Close()
	if _, err := fmt.Fprintln(os.Stdout, "hello"); err != nil {
		log.Printf("%#v %T\n", err, err) // &fs.PathError{Op:"write", Path:"/dev/stdout", Err:(*errors.errorString)(0xc000010120)} *fs.PathError
		log.Fatal(err)
	}
}
