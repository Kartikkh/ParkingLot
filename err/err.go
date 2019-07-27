package err

import "errors"

var (
	ErrParkingUnExpected       = errors.New("parking: Some unexpected error occured")
	ErrCreateParkingCenter     = errors.New("parking: Some unexpected error while creating parking lot")
	ErrInvalidParkingSize      = errors.New("parking: You cannot create a very parking slot with this size")
	ErrParkingFull             = errors.New("Sorry, parking lot is full")
	ErrSlotNumberInvalid       = errors.New("slot: Please provide valid slot number 1 or greater")
	ErrSlotAlreadyEmpty        = errors.New("slot: Slot number is already empty")
	ErrSlotAlreadyAllocated    = errors.New("slot: Slots already allocated")
	ErrSlotDuplicateNumber     = errors.New("slot: Please use unique number for slots")
	ErrInvalidGetSlotNumber    = errors.New("slot: Trying to fetch invalid slot")
	ErrNoFilledSlots           = errors.New("slot: All slots are empty")
	ErrVehicleData             = errors.New("vehicle:Please enter a valid vehicle Number & Colour")
	ErrVehicleDataAlreadyAdded = errors.New("vehicle: Data already added")
	ErrInvalidSlotAllocate     = errors.New("vehicle: Cannot allocate vehicle in invalid slot")
	ErrCommandParsing          = errors.New("command: Error while parsing command and arguments")
	ErrInvalidParams           = errors.New("command: Invalid parameters please provide valid input")
	ErrInvalidCommand          = errors.New("command: Please input valid command")
	ErrNotFound                = errors.New("Not found")
	ErrCannotPark              = errors.New("parking:Vehicle Number you are trying to park is already parked")
)
