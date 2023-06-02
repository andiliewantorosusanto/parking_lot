package main

import (
	"github.com/gorilla/mux"
	"gitlab.mapan.io/playground/parking-lot-golang/src/entity"
	"net/http"
)

var parkingLot = entity.NewParkingLot()

func createParkingSlot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Write(parkingLot.SetNumberOfSlot(params["numberOfSlot"]))
}

func reserve(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	regNumber := params["regNumber"]
	colour := params["colour"]

	w.Write(parkingLot.Reserve(regNumber, colour))
}

func leave(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Write(parkingLot.Leave(params["slotNumber"]))
}

func status(w http.ResponseWriter, _ *http.Request) {
	w.Write(parkingLot.Status())
}

func getRegNumbersByColour(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	colour := params["colour"]

	w.Write(parkingLot.GetRegNumbersByColour(colour))
}

func getReservedSlotsByColour(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	colour := params["colour"]

	w.Write(parkingLot.GetReservedSlotsByColour(colour))
}

func getReservedSlotByRegNumber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	regNumber := params["regNumber"]
	w.Write(parkingLot.GetReservedSlotsByColour(regNumber))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/create_parking_lot/{numberOfSlot}", createParkingSlot).Methods("POST")
	r.HandleFunc("/park/{regNumber}/{colour}", reserve).Methods("POST")
	r.HandleFunc("/leave/{slotNumber}", leave).Methods("POST")
	r.HandleFunc("/status", status).Methods("GET")
	r.HandleFunc("/cars_registration_numbers/colour/{colour}", getRegNumbersByColour).Methods("GET")
	r.HandleFunc("/cars_slot/colour/{colour}", getReservedSlotsByColour).Methods("GET")
	r.HandleFunc("/slot_number/car_registration_number/{regNumber}", getReservedSlotByRegNumber).Methods("GET")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		return
	}
}
