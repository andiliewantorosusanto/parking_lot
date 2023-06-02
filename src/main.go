package main

import (
	"github.com/gorilla/mux"
	"gitlab.mapan.io/playground/parking-lot-golang/src/handler"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/create_parking_lot/{numberOfSlot}", handler.CreateParkingSlot).Methods("POST")
	r.HandleFunc("/park/{regNumber}/{colour}", handler.Reserve).Methods("POST")
	r.HandleFunc("/leave/{slotNumber}", handler.Leave).Methods("POST")
	r.HandleFunc("/status", handler.Status).Methods("GET")
	r.HandleFunc("/cars_registration_numbers/colour/{colour}", handler.GetRegNumbersByColour).Methods("GET")
	r.HandleFunc("/cars_slot/colour/{colour}", handler.GetReservedSlotsByColour).Methods("GET")
	r.HandleFunc("/slot_number/car_registration_number/{regNumber}", handler.GetReservedSlotByRegNumber).Methods("GET")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		return
	}
}
