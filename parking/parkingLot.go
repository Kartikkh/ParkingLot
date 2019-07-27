package parking

import (
	"github.com/ParkingLot/slots"
)

type Parking struct {
	size  uint64
	slots []*slot.Slot
}

func InitParking(size uint64) *Parking {
	slots := slot.InitSlots(size)
	return &Parking{
		size:  size,
		slots: slots,
	}
}

func (p *Parking) GetFreeSlot() *slot.Slot {
	for _, sl := range p.slots {
		if sl.IsFree() {
			return sl
		}
	}
	return nil
}

func (p *Parking) GetFilledSlot() []*slot.Slot {
	slots := make([]*slot.Slot, 0)
	for _, sl := range p.slots {
		if !sl.IsFree() {
			slots = append(slots, sl)
		}
	}
	return slots
}
