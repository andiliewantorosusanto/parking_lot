package main

import (
	"github.com/gorilla/mux"
	"gitlab.mapan.io/playground/parking-lot-golang/src/entity"
	"net/http"
	"strconv"
	"strings"
)

var cars = map[int]entity.Car{}
var numberOfSlot = 0

func assignNumberOfSlot(numberOfSlotRequestStr string) []byte {
	numberOfSlotRequest, err := strconv.Atoi(numberOfSlotRequestStr)
	if err != nil {
		return []byte("Error converting string to int. please check your param")
	}

	if numberOfSlotRequest < numberOfSlot {
		return []byte("Cannot decrease number of slot. Loss of data may occur")
	}

	numberOfSlot = numberOfSlotRequest
	return []byte("Created a parking lot with " + strconv.Itoa(numberOfSlot) + " slots")
}

func reserveParkingLo(regNumber string, colour string) []byte {
	if numberOfSlot <= len(cars) {
		return []byte("Sorry, parking lot is full")
	}

	lotNumber := entity.GetNearestAvailableNumber(numberOfSlot, cars)
	cars[lotNumber] = entity.Car{Colour: colour, RegNumber: regNumber}

	return []byte("Allocated slot number: " + strconv.Itoa(lotNumber))
}

func createParkingLot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Write(assignNumberOfSlot(params["numberOfSlot"]))
}

func deleteCar(slotNumberStr string) []byte {
	slotNumber, err := strconv.Atoi(slotNumberStr)
	if err != nil {
		return []byte("Error converting string to int. please check your param")
	}

	delete(cars, slotNumber)

	return []byte("Slot number " + slotNumberStr + " is free")
}

func reserveParkingLot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	regNumber := params["regNumber"]
	colour := params["colour"]

	w.Write(reserveParkingLo(regNumber, colour))
}

func leaveParkingLot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Write(deleteCar(params["slotNumber"]))
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
