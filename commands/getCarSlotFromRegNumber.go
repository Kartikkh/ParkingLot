package commands

import (
	"fmt"
	"github.com/ParkingLot/err"
	"strings"
)

type commandGetCarSlotFromRegNumber struct {
	regNumber string
}

func InitGetCarSlotFromRegNumber() *commandGetCarSlotFromRegNumber {
	return &commandGetCarSlotFromRegNumber{}
}

func (csr *commandGetCarSlotFromRegNumber) Run() (error, string) {
	for _, slot := range parkingLot.GetFilledSlot() {
		if slot.GetVehicle().GetRegistrationNumber() == csr.regNumber {
			return nil, fmt.Sprint(slot.GetNumber())
		}
	}
	return nil, fmt.Sprint("Not found")
}

func (csr *commandGetCarSlotFromRegNumber) Parse(args string) error {
	if args != "" {
		csr.regNumber = args
		if csr.Verify() {
			return nil
		}
	}
	return err.ErrInvalidParams
}

func (csr *commandGetCarSlotFromRegNumber) Verify() bool {
	result := strings.Split(csr.regNumber, " ")
	if len(result) > 1 {
		return false
	}
	return true
}
