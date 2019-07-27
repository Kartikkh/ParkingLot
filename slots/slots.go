package slot

import (
	"github.com/ParkingLot/err"
	"github.com/ParkingLot/vehicle"
)

const MinNumberOfSlot = 1

type Slot struct {
	vehicle *vehicle.Vehicle
	number  uint64
}

func InitSlots(size uint64) []*Slot {
	slots := make([]*Slot, 0)
	for i := uint64(0); i < size; i++ {
		slots = append(slots, &Slot{
			vehicle: nil,
			number:  i + 1,
		})
	}
	return slots
}

func (s *Slot) GetNumber() uint64 {
	return s.number
}

func (s *Slot) GetVehicle() *vehicle.Vehicle {
	return s.vehicle
}

func (s *Slot) AllocateSlot(vehicle *vehicle.Vehicle) (*Slot, error) {
	if s.vehicle != nil {
		return s, err.ErrSlotAlreadyAllocated
	}
	s.vehicle = vehicle
	return s, nil
}

func IsSlotValid(slotNumber uint64) bool {
	return slotNumber >= MinNumberOfSlot
}

func (s *Slot) FreeSlot() {
	s.vehicle = nil
}

func (s *Slot) IsFree() bool {
	if s.vehicle == nil {
		return true
	}
	return false
}
