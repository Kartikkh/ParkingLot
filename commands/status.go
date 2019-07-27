package commands

import (
	"fmt"
	"github.com/ParkingLot/err"
	"strings"
)

type Status struct{}

func InitStatusCommand() *Status {
	return &Status{}
}

func (cp *Status) Run() (error, string) {
	var outPutList = []string{
		fmt.Sprintf("%-12s%-20s%-10s",
			"Slot No.",
			"Registration No",
			"Colour",
		),
		fmt.Sprintf("%-12v%-20v%-10v",
			"------------",
			"--------------------",
			"----------",
		),
	}

	for _, slot := range parkingLot.GetFilledSlot() {
		v := slot.GetVehicle()
		outPutList = append(outPutList,
			fmt.Sprintf(
				"%-12v%-20v%-10v",
				slot.GetNumber(),
				v.GetRegistrationNumber(),
				v.GetColour(),
			))
	}
	return nil, strings.Join(outPutList, "\n")
}

func (cp *Status) Parse(args string) error {
	if args == "" {
		return nil
	}
	return err.ErrInvalidCommand
}

func (cp *Status) Verify() bool {
	return true
}
