package input

import (
	"bufio"
	"github.com/ParkingLot/commands"
	"log"
	"os"
)

type FileProcessingInfo struct {
	fileName string
}

func InteractFile(fileName string) *FileProcessingInfo {
	return &FileProcessingInfo{
		fileName: fileName,
	}
}

func (f *FileProcessingInfo) ProcessFile() {

	commandsFile, er := os.Open("./" + f.fileName)
	if er != nil {
		log.Println("error while reading file", er)
		return
	}
	defer commandsFile.Close()
	commandScanner := bufio.NewScanner(commandsFile)
	commands.InitCommands()
	var commandString string
	for commandScanner.Scan() {
		commandString = commandScanner.Text()
		output, er := commands.ProcessCommand(commandString)
		if er != nil {
			log.Println(er)
		} else {
			log.Println(output)
		}
	}

	if err := commandScanner.Err(); err != nil {
		log.Println("error while scanning")
	}

}
