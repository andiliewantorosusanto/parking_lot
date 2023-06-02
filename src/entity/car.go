package entity

type Car struct {
	RegNumber string
	Colour    string
}

func GetNearestAvailableNumber(numberOfSlot int, cars map[int]Car) int {
	for i := 1; i <= numberOfSlot; i++ {
		if _, ok := cars[i]; !ok {
			return i
		}
	}

	return len(cars) + 1
}

func FindCarByRegNumber(regNumber string, cars map[int]Car) (*int, *Car) {
	for index, car := range cars {
		if car.RegNumber == regNumber {
			return &index, &car
		}
	}

	return nil, nil
}
