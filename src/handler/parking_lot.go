package handler

import (
	"github.com/gorilla/mux"
	"gitlab.mapan.io/playground/parking-lot-golang/src/entity"
	"io"
	"log"
	"net/http"
	"strings"
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
	w.Write(parkingLot.GetReservedSlotByRegNumber(regNumber))
}

func Bulk(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error reading commands. Please check commands inputted."))
		return
	}

	body := string(b)
	commands := strings.Split(body, "\n")

	var resp []byte
	for index, command := range commands {
		params := strings.Split(command, " ")
		if params[0] == "create_parking_lot" {
			resp = append(resp, parkingLot.SetNumberOfSlot(params[1])...)
		} else if params[0] == "park" {
			resp = append(resp, parkingLot.Reserve(params[1], params[2])...)
		} else if params[0] == "leave" {
			resp = append(resp, parkingLot.Leave(params[1])...)
		} else if params[0] == "status" {
			resp = append(resp, parkingLot.Status()...)
		} else if params[0] == "registration_numbers_for_cars_with_colour" {
			resp = append(resp, parkingLot.GetRegNumbersByColour(params[1])...)
		} else if params[0] == "slot_numbers_for_cars_with_colour" {
			resp = append(resp, parkingLot.GetReservedSlotsByColour(params[1])...)
		} else if params[0] == "slot_number_for_registration_number" {
			resp = append(resp, parkingLot.GetReservedSlotByRegNumber(params[1])...)
		} else {
			log.Println("Command not found! Ignore!")
		}

		if index != len(commands)-1 {
			resp = append(resp, []byte("\n")...)
		}
	}

	w.Write(resp)
}
