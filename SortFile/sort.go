package sortFile

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func SortFromFile(filepath string) {
	// Read input file
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Sort lines line by line
	fileContent := string(b)

	list := strings.Split(fileContent,"\n")
	
	uniq := make(map[string]bool)
	for _, line := range list {
        uniq[line] = true
    }
	var SortLines []string
	for line := range uniq {
        SortLines = append(SortLines, line)
    }
	sort.Strings(SortLines) 
	SortLinesStr := strings.Join(SortLines, "\n")

	fmt.Println(SortLinesStr)
	fmt.Println("Sorting and removing duplicate lines is complete!")

}
