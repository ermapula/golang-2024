package cars

import "errors"

type Car struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Number     string `json:"number"`
	Speed      int    `json:"speed"`
	Country    string `json:"country"`
	RacingType string `json:"racing_type"`
}

var cars = []Car{
	{
		Id:         "1",
		Name:       "Lightning McQueen",
		Number:     "95",
		Speed:      200,
		Country:    "USA",
		RacingType: "Piston Cup",
	},
	{
		Id:         "2",
		Name:       "Francesco Bernoulli",
		Number:     "1",
		Speed:      220,
		Country:    "Italy",
		RacingType: "Formula Racer",
	},
	{
		Id:         "3",
		Name:       "Jeff Gorvette",
		Number:     "24",
		Speed:      200,
		Country:    "USA",
		RacingType: "GTS2",
	},
	{
		Id:         "4",
		Name:       "Lewis Hamilton",
		Number:     "2",
		Speed:      190,
		Country:    "United Kingdom",
		RacingType: "GTS1",
	},
	{
		Id:         "5",
		Name:       "Max Schnell",
		Number:     "4",
		Speed:      160,
		Country:    "Germany",
		RacingType: "World Torque Champion League",
	},
	{
		Id:         "6",
		Name:       "Raoul ÇaRoule",
		Number:     "06",
		Speed:      125,
		Country:    "France",
		RacingType: "Global Rally Championship",
	},
	{
		Id:         "7",
		Name:       "Carla Veloso",
		Number:     "8",
		Speed:      206,
		Country:    "Brazil",
		RacingType: "Le Motor Prototype",
	},
	{
		Id:         "8",
		Name:       "Shu Todoroki",
		Number:     "7",
		Speed:      203,
		Country:    "Japan",
		RacingType: "Le Motor Prototype",
	},
	{
		Id:         "9",
		Name:       "Nigel Gearsley",
		Number:     "9",
		Speed:      180,
		Country:    "United Kingdom",
		RacingType: "GTS1",
	},
	{
		Id:         "10",
		Name:       "Miguel Camino",
		Number:     "5",
		Speed:      200,
		Country:    "Spain",
		RacingType: "GTS2",
	},
	{
		Id:         "11",
		Name:       "Rip Clutchgoneski",
		Number:     "10",
		Speed:      199,
		Country:    "Spain",
		RacingType: "Formula 6000",
	},
}

func GetCars() []Car {
	return cars
}

func GetCar(id string) (*Car, error) {
	for _, c := range cars {
		if c.Id == id {
			return &c, nil
		}
	}
	return nil, errors.New("Car not found")
}
