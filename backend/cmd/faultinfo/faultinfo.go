package main

import (
	"log"
	"os"

	"github.com/charakoba-com/fault_info/backend"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	if err := faultinfo.Run(`:8080`); err != nil {
		log.Printf("%s", err)
		return 1
	}
	return 0
}
