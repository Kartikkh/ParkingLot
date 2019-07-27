package vehicle

type Vehicle struct {
	registrationNumber string
	colour             string
}

func NewVehicle(number, colour string) *Vehicle {
	return &Vehicle{
		registrationNumber: number,
		colour:             colour,
	}
}

func (v *Vehicle) GetRegistrationNumber() string {
	return v.registrationNumber

}

func (v *Vehicle) GetColour() string {
	return v.colour
}
