package partyHandler

import (
	"encoding/json"
	"net/http"
)

func HandleGetParties(w http.ResponseWriter, r *http.Request) {
	var searchedParty partyRequest

	err := json.NewDecoder(r.Body).Decode(&searchedParty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parties, _ := GetParties(searchedParty)

	resp := partyResponse{
		Parties: parties,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
