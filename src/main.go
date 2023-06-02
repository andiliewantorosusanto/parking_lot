package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.mapan.io/playground/parking-lot-golang/src/entity"
	"net/http"
	"strconv"
	"strings"
)

var cars = map[int]entity.Car{}
var numberOfSlot = 0

func createParkingLot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	numberOfSlotRequest, err := strconv.Atoi(params["numberOfSlot"])
	if err != nil {
		fmt.Println("Error converting string to int")
		return
	}

	if numberOfSlotRequest < numberOfSlot {
		fmt.Println("Cannot decrease number of slot. Loss of data may occur")
		return
	}

	numberOfSlot = numberOfSlotRequest

	w.Write([]byte("Created a parking lot with " + params["numberOfSlot"] + " slots"))
}

func reserveParkingLot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	regNumber := params["regNumber"]
	colour := params["colour"]

	if numberOfSlot <= len(cars) {
		w.Write([]byte("Sorry, parking lot is full"))
		return
	}

	lotNumber := entity.GetNearestAvailableNumber(numberOfSlot, cars)
	cars[lotNumber] = entity.Car{Colour: colour, RegNumber: regNumber}

	w.Write([]byte("Allocated slot number: " + strconv.Itoa(lotNumber)))
}

func leaveParkingLot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	slotNumber, err := strconv.Atoi(params["slotNumber"])
	if err != nil {
		fmt.Println("Error converting string to int")
		return
	}

	delete(cars, slotNumber)
	w.Write([]byte("Slot number " + params["slotNumber"] + " is free"))
}

func getParkingLotStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Slot No. Registration No Colour\n"))

	var parkingLotDetail []string
	for i := 1; i <= numberOfSlot; i++ {
		if car, ok := cars[i]; ok {
			slotNumber := strconv.Itoa(i)
			body := slotNumber + " " + car.RegNumber + " " + car.Colour
			parkingLotDetail = append(parkingLotDetail, body)
		}
	}

	w.Write([]byte(strings.Join(parkingLotDetail, "\n")))
}

func getCarsRegNumberByColour(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	colour := params["colour"]
	var regNumbers []string

	for i := 1; i <= numberOfSlot; i++ {
		if car, ok := cars[i]; ok {
			if car.Colour == colour {
				regNumbers = append(regNumbers, car.RegNumber)
			}
		}
	}

	w.Write([]byte(strings.Join(regNumbers, ", ")))
}

func getReservedLotNumberByColour(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	colour := params["colour"]
	var reservedSlots []string

	for i := 1; i <= numberOfSlot; i++ {
		if car, ok := cars[i]; ok {
			if car.Colour == colour {
				reservedSlots = append(reservedSlots, strconv.Itoa(i))
			}
		}
	}

	w.Write([]byte(strings.Join(reservedSlots, ", ")))
}

func getLotNumberByRegNumber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	regNumber := params["regNumber"]

	slotNumber, _ := entity.FindCarByRegNumber(regNumber, cars)
	if slotNumber == nil {
		w.Write([]byte("Not found"))
		return
	}
	body := *slotNumber

	w.Write([]byte(strconv.Itoa(body)))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/create_parking_lot/{numberOfSlot}", createParkingLot).Methods("POST")
	r.HandleFunc("/park/{regNumber}/{colour}", reserveParkingLot).Methods("POST")
	r.HandleFunc("/leave/{slotNumber}", leaveParkingLot).Methods("POST")
	r.HandleFunc("/status", getParkingLotStatus).Methods("GET")
	r.HandleFunc("/cars_registration_numbers/colour/{colour}", getCarsRegNumberByColour).Methods("GET")
	r.HandleFunc("/cars_slot/colour/{colour}", getReservedLotNumberByColour).Methods("GET")
	r.HandleFunc("/slot_number/car_registration_number/{regNumber}", getLotNumberByRegNumber).Methods("GET")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		return
	}
}
