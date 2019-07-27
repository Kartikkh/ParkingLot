package input

import (
	"bufio"
	"github.com/ParkingLot/commands"
	"log"
	"os"
)

type ShellInfo struct {
}

func InteractShell() *ShellInfo {
	return &ShellInfo{}
}

func (s *ShellInfo) Process() {
	//reader := bufio.NewReader(os.Stdin)
	log.Println("Simple Shell")
	log.Println("---------------------")
	commands.InitCommands()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		log.Println("text", text)
		log.Println([]byte(text))
		log.Println([]byte("exit"))
		if text == `exit` {
			os.Exit(0)
		}
		output, er := commands.ProcessCommand(text)
		if er != nil {
			log.Println(er)
		} else {
			log.Println(output)
		}

	}

	//for {
	//	log.Println("-> ")
	//	text, _ := reader.ReadString('\n')
	//	log.Println("text", text)
	//	log.Println([]byte(text))
	//	strings.Trim(text, "")
	//	//log.Println([]byte(text))
	//	//log.Println([]byte("exit"))
	//	//log.Println("sadjksadk", strings.Compare("exitA", text) == 0)
	//	//log.Println(text == `exit`)
	//	if text == `exit` {
	//		os.Exit(0)
	//	}
	//	output, er := commands.ProcessCommand(text)
	//	if er != nil {
	//		log.Println(er)
	//	} else {
	//		log.Println(output)
	//	}
	//
	//}
}
