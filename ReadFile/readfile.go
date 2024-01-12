package readFile

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReaderFromFile(file string){
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		fmt.Printf("%s", buf[:n])
	}
}
