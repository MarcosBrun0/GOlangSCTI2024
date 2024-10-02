package main

import (
	"fmt"
	"log"
	"os"

	"go.mod/app"
)

func main() {
	fmt.Printf("Initializing APP")
	r := app.Start()
	if err := r.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}