package entity

import (
	"fmt"
	"strconv"
	"testing"
)

var parkingLot = NewParkingLot()
var expectedNumberOfLot = 6

func status() string {
	return `Slot No. Registration No Colour
4 B-4-ABC Black
5 B-5-ABC Black
6 B-6-ABC Black`
}

func TestParkingLot_SetNumberOfSlot(t *testing.T) {
	response := parkingLot.SetNumberOfSlot(strconv.Itoa(expectedNumberOfLot))

	if string(response) != fmt.Sprintf("Created a parking lot with %d slots", expectedNumberOfLot) {
		t.Errorf("Set number of slot FAILED. Expected 'Created a parking lot with %d slots' got %s\n", expectedNumberOfLot, string(response))
	} else {
		t.Logf("Set number of slot PASSED.")
	}
}

func TestParkingLot_Reserve(t *testing.T) {
	for i := 1; i <= expectedNumberOfLot; i++ {
		response := parkingLot.Reserve("B-"+strconv.Itoa(i)+"-ABC", "Black")

		if string(response) != fmt.Sprintf("Allocated slot number: %d", i) {
			t.Errorf("Reserve FAILED. Expected 'Allocated slot number: %d' got %s\n", i, string(response))
		}
	}

	response := parkingLot.Reserve("B-321-ABC", "Blue")
	if string(response) != fmt.Sprintf("Sorry, parking lot is full") {
		t.Errorf("Reserve FAILED. Expected 'Sorry, parking lot is full' got %s\n", string(response))
	} else {
		t.Logf("Reserve PASSED.")
	}
}

func TestParkingLot_Leave(t *testing.T) {
	for i := 1; i <= expectedNumberOfLot/2; i++ {
		response := parkingLot.Leave(strconv.Itoa(i))

		if string(response) != fmt.Sprintf("Slot number %d is free", i) {
			t.Errorf("Leave FAILED. Expected 'Slot number %d is free' got %s\n", i, string(response))
		}
	}

	t.Logf("Reserve PASSED.")
}

func TestParkingLot_Status(t *testing.T) {
	response := parkingLot.Status()

	if string(response) != status() {
		t.Errorf("Status FAILED. Expected %s got %s\n", status(), string(response))
	} else {
		t.Logf("Status PASSED.")
	}
}

func TestParkingLot_GetRegNumbersByColour(t *testing.T) {
	response := parkingLot.GetRegNumbersByColour("Black")

	if string(response) != "B-4-ABC, B-5-ABC, B-6-ABC" {
		t.Errorf("Get reg numbers by colour FAILED. Expected 'B-4-ABC, B-5-ABC, B-6-ABC' got %s\n", string(response))
	} else {
		t.Logf("Get reg numbers by colour PASSED.")
	}
}

func TestParkingLot_GetReservedSlotsByColour(t *testing.T) {
	response := parkingLot.GetReservedSlotsByColour("Black")

	if string(response) != "4, 5, 6" {
		t.Errorf("Get reseved slots by colour FAILED. Expected '4, 5, 6' got %s\n", string(response))
	} else {
		t.Logf("Get reseved slots by colour PASSED.")
	}
}

func TestParkingLot_GetReservedSlotByRegNumber(t *testing.T) {
	response := parkingLot.GetReservedSlotByRegNumber("B-4-ABC")

	if string(response) != "4" {
		t.Errorf("Get reseved slots by reg number FAILED. Expected '4' got %s\n", string(response))
	}

	response = parkingLot.GetReservedSlotByRegNumber("B-3-ABC")

	if string(response) != "Not found" {
		t.Errorf("Get reseved slots by reg number FAILED. Expected 'Not found' got %s\n", string(response))
	} else {
		t.Logf("Get reseved slots by reg number PASSED.")
	}
}
