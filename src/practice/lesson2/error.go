package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./slice.go")
	if err != nil {
		log.Fatalln("an error in os.Open")
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("an error in file.Read")
	}
	fmt.Println(count, string(data))

	if os.Chdir("test"); err != nil {
		log.Fatalln("an error in os.Chdir")
	}
}
