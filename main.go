package main

import (
	"github.com/ParkingLot/input"
	"os"
)

func main() {
	// Process the input file and give an interactive shell
	if len(os.Args) > 1 && "" != os.Args[1] {
		input.InteractFile("file_input.txt").ProcessFile()
	} else {
		// Interactive shell will provide to user to operate
		// with file data and console process
		input.InteractShell().Process()
	}

}
