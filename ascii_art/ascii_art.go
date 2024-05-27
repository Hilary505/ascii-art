package ascii_art 

import (
	"os"
	"fmt"
	"strings"
)
// functiom that reads from the file 
func GetFile(filename  string) ([]string,error) {
	file , err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("An error", err)
		os.Exit(0)
	}
	myfile := string(file)
	contents := strings.Split(myfile, "\n")
	return contents, nil
}
// functions that accepts our commandline arguments
func ProcessInput(contents  []string, input string) {
	newInput := strings.Split(input, "\\n")
	for _, arg := range newInput {
		for i := 0; i <= 8; i++ {
			for _, ch := range arg {
				index := int(ch-32) *9 + i
				if index >= 0 && index < len(contents) {
					fmt.Print(contents[index])
				}
			}
			fmt.Println()
		}
	}
}