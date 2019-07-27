package commands

import (
	"fmt"
	"github.com/ParkingLot/err"
	"strings"
)

type commandGetSlotFromRegNumber struct {
	colour string
}

func InitGetSlotFromRegNumber() *commandGetSlotFromRegNumber {
	return &commandGetSlotFromRegNumber{}
}

func (sr *commandGetSlotFromRegNumber) Run() (error, string) {
	var slotNums []string
	for _, slot := range parkingLot.GetFilledSlot() {
		if slot.GetVehicle().GetColour() == sr.colour {
			slotNums = append(slotNums, fmt.Sprint(slot.GetNumber()))
		}
	}
	if len(slotNums) == 0 {
		return err.ErrNoFilledSlots, ""
	}

	return nil, strings.Join(slotNums, ", ")
}

func (sr *commandGetSlotFromRegNumber) Parse(args string) error {
	if args != "" {
		sr.colour = args
		if sr.Verify() {
			return nil
		}
	}
	return err.ErrInvalidParams
}

func (sr *commandGetSlotFromRegNumber) Verify() bool {
	result := strings.Split(sr.colour, " ")
	if len(result) > 1 {
		return false
	}
	return true
}
