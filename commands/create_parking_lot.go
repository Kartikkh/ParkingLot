package commands

import (
	"fmt"
	"github.com/ParkingLot/err"
	"github.com/ParkingLot/parking"
	"strconv"
)

type CreateParkingLot struct {
	size uint64
}

func InitParkingLot() *CreateParkingLot {
	return &CreateParkingLot{}
}

func (cp *CreateParkingLot) Run() (error, string) {
	parkingLot = parking.InitParking(uint64(cp.size))
	return nil, fmt.Sprintf("Created a parking lot with %d slots", cp.size)
}

func (cp *CreateParkingLot) Verify() bool {
	if cp.size < 1 {
		return false
	}
	return true
}

func (cp *CreateParkingLot) Parse(args string) error {
	if args != "" {
		val, er := strconv.ParseUint(args, 0, 64)
		if nil != er {
			return err.ErrInvalidParams
		}
		cp.size = val
		if !cp.Verify() {
			return err.ErrInvalidParkingSize
		}
		return nil

	}
	return err.ErrCreateParkingCenter
}
