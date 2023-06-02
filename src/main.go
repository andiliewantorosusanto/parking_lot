package main

import (
	"github.com/gorilla/mux"
	"gitlab.mapan.io/playground/parking-lot-golang/src/entity"
	"net/http"
	"strconv"
	"strings"
)

func createParkingLot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Write(assignNumberOfSlot(params["numberOfSlot"]))
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

func getParkingLotStatus(w http.ResponseWriter, _ *http.Request) {
	w.Write(getStatus())
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

	r.HandleFunc("/create_parking_lot/{numberOfSpot}", createParkingLot).Methods("POST")
	r.HandleFunc("/park/{regNumber}/{colour}", reserveParkingLot).Methods("POST")
	r.HandleFunc("/leave/{spotNumber}", leaveParkingLot).Methods("POST")
	r.HandleFunc("/status", getParkingLotStatus).Methods("GET")
	r.HandleFunc("/cars_registration_numbers/colour/{colour}", getCarsRegNumberByColour).Methods("GET")
	r.HandleFunc("/cars_slot/colour/{colour}", getReservedLotNumberByColour).Methods("GET")
	r.HandleFunc("/slot_number/car_registration_number/{regNumber}", getLotNumberByRegNumber).Methods("GET")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		return
	}
}
