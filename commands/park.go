package commands

import (
	"fmt"
	"github.com/ParkingLot/err"
	"github.com/ParkingLot/vehicle"
	"strings"
)

type Park struct {
	parkArgs []string
	number   string
	color    string
}

func InitParkCommand() *Park {
	return &Park{}
}

func (cp *Park) Run() (error, string) {
	vehicleObj := vehicle.NewVehicle(cp.number, cp.color)
	emptySlot := parkingLot.GetFreeSlot()
	if emptySlot == nil {
		return err.ErrParkingFull, ""
	}
	allocatedSlot, er := emptySlot.AllocateSlot(vehicleObj)
	if er != nil {
		return err.ErrSlotAlreadyAllocated, ""
	}
	return nil, fmt.Sprintf("Allocated slot number: %d", allocatedSlot.GetNumber())
}

func (cp *Park) Verify() bool {
	if len(cp.parkArgs) < 2 {
		return false
	}
	return true
}

func (cp *Park) Parse(args string) error {
	if args != "" {
		result := strings.Split(args, " ")
		cp.parkArgs = result
		if !cp.Verify() {
			return err.ErrInvalidParams
		}
		cp.number = cp.parkArgs[0]
		cp.color = cp.parkArgs[1]
		return nil
	}
	return err.ErrVehicleData

}
