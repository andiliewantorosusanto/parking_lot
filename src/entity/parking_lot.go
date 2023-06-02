package entity

import (
	"strconv"
	"strings"
)

type ParkingLot struct {
	cars         map[int]Car
	numberOfSpot int
}

func (p *ParkingLot) getNearestAvailableNumber() int {
	for i := 1; i <= p.numberOfSpot; i++ {
		if _, ok := p.cars[i]; !ok {
			return i
		}
	}

	return len(p.cars) + 1
}

func (p *ParkingLot) findCarByRegNumber(regNumber string) (*int, *Car) {
	for index, car := range p.cars {
		if car.RegNumber == regNumber {
			return &index, &car
		}
	}

	return nil, nil
}

func (p *ParkingLot) SetNumberOfLot(numberOfSlotRequestStr string) []byte {
	numberOfSlotRequest, err := strconv.Atoi(numberOfSlotRequestStr)
	if err != nil {
		return []byte("Error converting string to int. please check your param")
	}

	if numberOfSlotRequest < p.numberOfSpot {
		return []byte("Cannot decrease number of slot. Loss of data may occur")
	}

	p.numberOfSlot = numberOfSlotRequest
	return []byte("Created a parking lot with " + numberOfSlotRequestStr + " slots")
}

func (p *ParkingLot) Reserve(regNumber string, colour string) []byte {
	if p.numberOfSlot <= len(p.cars) {
		return []byte("Sorry, parking lot is full")
	}

	lotNumber := p.getNearestAvailableNumber()
	p.cars[lotNumber] = Car{Colour: colour, RegNumber: regNumber}

	return []byte("Allocated slot number: " + strconv.Itoa(lotNumber))
}

func (p *ParkingLot) Departure(slotNumberStr string) []byte {
	slotNumber, err := strconv.Atoi(slotNumberStr)
	if err != nil {
		return []byte("Error converting string to int. please check your param")
	}

	delete(p.cars, slotNumber)

	return []byte("Slot number " + slotNumberStr + " is free")
}

func (p *ParkingLot) Status() []byte {
	msg := "Slot No. Registration No Colour\n"

	var parkingLotDetail []string
	for i := 1; i <= numberOfSlot; i++ {
		if car, ok := cars[i]; ok {
			slotNumber := strconv.Itoa(i)
			body := slotNumber + " " + car.RegNumber + " " + car.Colour
			parkingLotDetail = append(parkingLotDetail, body)
		}
	}

	msg = msg + strings.Join(parkingLotDetail, "\n")

	return []byte(msg)
}
