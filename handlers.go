package main

import (
	"encoding/json"
	"net/http"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetClientProfile(w, r)
	case http.MethodPatch:
		UpdateClientProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetClientProfile(w http.ResponseWriter, r *http.Request) {

	clientProfile := r.Context().Value("clientProfile").(ClientProfile)

	response := ClientProfile{
		Id:    clientProfile.Id,
		Name:  clientProfile.Name,
		Email: clientProfile.Email,
		Token: clientProfile.Token,
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateClientProfile(w http.ResponseWriter, r *http.Request) {

	clientProfile := r.Context().Value("clientProfile").(ClientProfile)

	var payloadData ClientProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	clientProfile.Name = payloadData.Name
	clientProfile.Email = payloadData.Email
	database[clientProfile.Id] = clientProfile

	w.WriteHeader(http.StatusOK)

}
