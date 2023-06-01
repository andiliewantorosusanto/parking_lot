package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

//

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
