package commands

import (
	"errors"
	"github.com/ParkingLot/err"
	"github.com/ParkingLot/parking"
	"strings"
)

var commandMap map[string]commandInterface

var parkingLot *parking.Parking

type commandString struct {
	command, args, userInputString string
}

type commandInterface interface {
	Run() (error, string)
	Parse(string) error
	Verify() bool
}

func InitCommands() {
	commandMap = make(map[string]commandInterface, 0)
	commandMap["create_parking_lot"] = InitParkingLot()
	commandMap["leave"] = InitLeaveCommand()
	commandMap["status"] = InitStatusCommand()
	commandMap["park"] = InitParkCommand()
	commandMap["registration_numbers_for_cars_with_colour"] = InitGetCarRegNumberByColour()
	commandMap["slot_number_for_registration_number"] = InitGetCarSlotFromRegNumber()
	commandMap["slot_numbers_for_cars_with_colour"] = InitGetSlotFromRegNumber()
}

func (cs *commandString) IsCommandValid() bool {
	if _, ok := commandMap[cs.command]; ok {
		return true
	}
	return false
}

// parse return command and arguments
func (cs *commandString) Parse() error {
	command := strings.Trim(cs.userInputString, " \n\t")
	results := strings.SplitN(command, " ", 2)
	cs.command = strings.ToLower(results[0])
	if len(results) > 1 {
		cs.args = results[1]
	}
	if cs.command == "" {
		return errors.New("command not found ")
	}
	return nil
}

func ProcessCommand(text string) (string, error) {
	var commandString commandString
	var output string
	commandString.userInputString = text
	er := commandString.Parse()
	if er != nil {
		return "", err.ErrCommandParsing
	}
	if !commandString.IsCommandValid() {
		return "", err.ErrInvalidCommand
	}

	if runner, ok := commandMap[commandString.command]; ok {
		er = runner.Parse(commandString.args)
		if er != nil {
			return "", er
		}
		er, output = runner.Run()
		if er != nil {
			return "", er
		}
	}
	return output, nil
}
