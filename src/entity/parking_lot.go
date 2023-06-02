package entity

import (
	"strconv"
	"strings"
)

type ParkingLot struct {
	cars         map[int]Car
	numberOfSlot int
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{
		cars:         map[int]Car{},
		numberOfSlot: 0,
	}
}

func (p *ParkingLot) getNearestAvailableNumber() int {
	for i := 1; i <= p.numberOfSlot; i++ {
		if _, ok := p.cars[i]; !ok {
			return i
		}
	}

	return len(p.cars) + 1
}

func (p *ParkingLot) findReservedSlotByRegNumber(regNumber string) int {
	for index, car := range p.cars {
		if car.RegNumber == regNumber {
			return index
		}
	}
	return -1
}

func (p *ParkingLot) SetNumberOfSlot(numberOfSlotRequestStr string) []byte {
	numberOfSlotRequest, err := strconv.Atoi(numberOfSlotRequestStr)
	if err != nil {
		return []byte("Error converting string to int. please check your param")
	}

	if numberOfSlotRequest < p.numberOfSlot {
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

func (p *ParkingLot) Leave(slotNumberStr string) []byte {
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
	for i := 1; i <= p.numberOfSlot; i++ {
		if car, ok := p.cars[i]; ok {
			slotNumber := strconv.Itoa(i)
			body := slotNumber + " " + car.RegNumber + " " + car.Colour
			parkingLotDetail = append(parkingLotDetail, body)
		}
	}

	msg = msg + strings.Join(parkingLotDetail, "\n")

	return []byte(msg)
}

func (p *ParkingLot) GetRegNumbersByColour(colour string) []byte {
	var regNumbers []string

	for i := 1; i <= p.numberOfSlot; i++ {
		if car, ok := p.cars[i]; ok {
			if car.Colour == colour {
				regNumbers = append(regNumbers, car.RegNumber)
			}
		}
	}

	return []byte(strings.Join(regNumbers, ", "))
}

func (p *ParkingLot) GetReservedSlotsByColour(colour string) []byte {
	var reservedSlots []string

	for i := 1; i <= p.numberOfSlot; i++ {
		if car, ok := p.cars[i]; ok {
			if car.Colour == colour {
				reservedSlots = append(reservedSlots, strconv.Itoa(i))
			}
		}
	}

	return []byte(strings.Join(reservedSlots, ", "))
}

func (p *ParkingLot) GetReservedSlotByRegNumber(regNumber string) []byte {
	slotNumber := p.findReservedSlotByRegNumber(regNumber)
	if slotNumber == -1 {
		return []byte("Not found")
	}

	return []byte(strconv.Itoa(slotNumber))
}
