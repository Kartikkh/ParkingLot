package commands

import (
	"fmt"
	"github.com/ParkingLot/err"
	"github.com/ParkingLot/slots"
	"strconv"
	"strings"
)

type Leave struct {
	slotNumber uint64
}

func InitLeaveCommand() *Leave {
	return &Leave{}
}

func (l *Leave) Run() (error, string) {
	for _, slot := range parkingLot.GetFilledSlot() {
		if slot.GetNumber() == l.slotNumber {
			if slot.IsFree() {
				return err.ErrSlotAlreadyEmpty, ""
			}
			slot.FreeSlot()
			return nil, fmt.Sprintf("Slot number %d is free", l.slotNumber)
		}
	}
	return err.ErrNoFilledSlots, ""
}

func (l *Leave) Parse(args string) error {
	if args != "" {
		result := strings.Split(args, "")
		status, er := strconv.ParseUint(result[0], 10, 64)
		if er != nil {
			return err.ErrCommandParsing
		}
		l.slotNumber = status
		if !l.Verify() {
			return err.ErrInvalidParams
		}
		return nil
	}
	return err.ErrVehicleData
}

func (l *Leave) Verify() bool {
	if slot.IsSlotValid(l.slotNumber) {
		return true
	}
	return false
}
