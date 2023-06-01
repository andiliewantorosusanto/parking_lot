package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Car struct {
	regNumber string
	colour    string
}

var cars map[int]Car
var numberOfSlot = 0

func getNearestAvailableNumber() int {
	for i := 1; i <= numberOfSlot; i++ {
		if _, ok := cars[i]; !ok {
			return i
		}
	}

	return len(cars) + 1
}

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

	if numberOfSlot >= len(cars) {
		w.Write([]byte("Sorry, parking lot is full"))
		return
	}

	lotNumber := getNearestAvailableNumber()
	cars[lotNumber] = Car{colour: colour, regNumber: regNumber}

	w.Write([]byte("Allocated slot number: " + strconv.Itoa(lotNumber)))
}

func leaveParkingLot(w http.ResponseWriter, r *http.Request) {

}

func getParkingLotStatus(w http.ResponseWriter, r *http.Request) {

}

func getCarsRegNumberByColour(w http.ResponseWriter, r *http.Request) {

}

func getReservedLotNumberByColour(w http.ResponseWriter, r *http.Request) {

}

func getLotNumberByRegNumber(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/create_parking_lot/{numberOfSlot}", createParkingLot).Methods("POST")
	r.HandleFunc("/park/{regNumber}/{colour}", reserveParkingLot).Methods("POST")
	r.HandleFunc("/leave/{slotNumber}", leaveParkingLot).Methods("POST")
	r.HandleFunc("/status", getParkingLotStatus).Methods("GET")
	r.HandleFunc("/cars_registration_numbers/colour/{colour}", getCarsRegNumberByColour).Methods("GET")
	r.HandleFunc("/cars_slot/colour/{colour}", getReservedLotNumberByColour).Methods("GET")
	r.HandleFunc("/slot_number/car_Registration_number/{regNumber}", getLotNumberByRegNumber).Methods("GET")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		return
	}
}
