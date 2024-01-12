package saveFile

import (
	"fmt"
	"log"
	"os"
)

func Saver(filepath string) {
	file , err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "Hello, World!")
}