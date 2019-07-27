package commands

import (
	"github.com/ParkingLot/err"
	"strings"
)

type commandGetRegNumWithColour struct {
	colour string
}

func InitGetCarRegNumberByColour() *commandGetRegNumWithColour {
	return &commandGetRegNumWithColour{}
}

func (rc *commandGetRegNumWithColour) Run() (error, string) {
	var regNums []string
	for _, slot := range parkingLot.GetFilledSlot() {
		if slot.GetVehicle().GetColour() == rc.colour {
			regNums = append(regNums, slot.GetVehicle().GetRegistrationNumber())
		}
	}
	if len(regNums) == 0 {
		return err.ErrNoFilledSlots, ""
	}

	return nil, strings.Join(regNums, ", ")
}

func (rc *commandGetRegNumWithColour) Parse(args string) error {
	if args != "" {
		rc.colour = args
		if rc.Verify() {
			return nil
		}
	}
	return err.ErrInvalidParams
}

func (rc *commandGetRegNumWithColour) Verify() bool {
	result := strings.Split(rc.colour, " ")
	if len(result) > 1 {
		return false
	}
	return true
}
