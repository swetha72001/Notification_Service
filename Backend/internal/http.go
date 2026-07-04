package internal

import (
	"Backend/dto"
	"encoding/json"
	"log"
	"net/http"
)

type HttpSerivice struct {
	EventChn chan *dto.Events
}

func (h *HttpSerivice) NotifyHandler(w http.ResponseWriter, r *http.Request) {
	// need to set header
	w.Header().Set("content-Type", "application/json")
	// decode req to struct
	var event dto.Events
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error in decode", err)
		return
	}

	// use the value according to logic
	h.EventChn <- &event

	// form resp (encode)
	resp := []byte("response sent succesfully")
	w.WriteHeader(http.StatusAccepted)
	w.Write(resp)
}
