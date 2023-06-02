package handler

import (
	"github.com/gorilla/mux"
	"gitlab.mapan.io/playground/parking-lot-golang/src/entity"
	"net/http"
)

var parkingLot = entity.NewParkingLot()

func CreateParkingSlot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Write(parkingLot.SetNumberOfSlot(params["numberOfSlot"]))
}

func Reserve(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	regNumber := params["regNumber"]
	colour := params["colour"]

	w.Write(parkingLot.Reserve(regNumber, colour))
}

func Leave(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Write(parkingLot.Leave(params["slotNumber"]))
}

func Status(w http.ResponseWriter, _ *http.Request) {
	w.Write(parkingLot.Status())
}

func GetRegNumbersByColour(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	colour := params["colour"]

	w.Write(parkingLot.GetRegNumbersByColour(colour))
}

func GetReservedSlotsByColour(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	colour := params["colour"]

	w.Write(parkingLot.GetReservedSlotsByColour(colour))
}

func GetReservedSlotByRegNumber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	regNumber := params["regNumber"]
	w.Write(parkingLot.GetReservedSlotsByColour(regNumber))
}
