package main

import (
	"flag"
	setBanner "zagrosboy/Banner"
	readFile "zagrosboy/ReadFile"
	saveFile "zagrosboy/SaveFile"
	sortFile "zagrosboy/SortFile"
)

func main() {
	setBanner.CreateBanner()
	Filename := flag.String("f", "", "filename for input")
	OutPut := flag.String("o", "", "filename for output")
	Sorting := flag.String("s", "", "sort order by filename")
	flag.Parse()
	if *Filename != "" {
		readFile.ReaderFromFile(*Filename)
	}
	if *OutPut != "" {
		saveFile.Saver(*OutPut)

	}
	if *Sorting!= "" {
   	    sortFile.SortFromFile(*Sorting)
    }

}
