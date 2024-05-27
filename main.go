package main 

import (
 "fmt" 
 "os"
 a "ascii/ascii_art"
)

func main() {
	// check if the length of arguments are required
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Enter 2 or 3 arguments but not more:")
		return
	}
	contents, err := a.GetFile("standard.txt")
	if err != nil {
		fmt.Println("An error:", err) 
		return
	}
	input := os.Args[1]
	a.ProcessInput(contents, input)
}