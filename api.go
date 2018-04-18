package main

import (
	"net/http"
	"fmt"
	"strings"
	"os"
	"encoding/json"
	"api.hcabosantos.cat/types"
)

type Response struct {
	Response interface{} `json:"response"`
	Errors   []string    `json:"errors"`
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if CheckAuthenticity(r) {
		// Create return string
		var request []string

		for name, headers := range r.Header {
			name = strings.ToLower(name)
			for _, h := range headers {
				request = append(request, fmt.Sprintf("%v: %v", name, h))
			}
		}

		w.Write([]byte(strings.Join(request, "\n")))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden"))
	}

	//db, err := GetDb()
	//if err != nil {
	//	log.Error(err)
	//} else {
	//	defer db.Close()
	//	var patients []types.Patient
	//	err := db.Select(&patients, "SELECT * FROM patients ORDER BY surname ASC")
	//	if err != nil {
	//		log.Error(err)
	//	} else {
	//		for _, patient := range patients {
	//			validDNI, err := patient.ValidateDni()
	//			log.Info(fmt.Sprintf("%d -> %s with valid DNI %s -> %t", patient.Id, patient.Name, patient.Nif, validDNI))
	//			if err != nil {
	//				log.Error(err)
	//			}
	//		}
	//	}
	//}
}

func PatientHandler(w http.ResponseWriter, r *http.Request) {
	if CheckAuthenticity(r) {
		var response Response
		db, err := GetDb()
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			jsonStruct, _ := StructToJson(response)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(jsonStruct))
		}
		defer db.Close()
		var patients []types.Patient
		err = db.Select(&patients, "SELECT * FROM patients ORDER BY surname ASC")
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			jsonStruct, _ := StructToJson(response)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(jsonStruct))
		}
		response.Response = patients
		jsonStruct, _ := StructToJson(response)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(jsonStruct))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden"))
	}
}

func CheckAuthenticity(r *http.Request) bool {
	allowedXForwardedFor := strings.Join(strings.Split(os.Getenv("ALLOWEDXFORWARDEDFOR"), " "), ", ")
	return r.Header.Get("X-Forwarded-For") == allowedXForwardedFor
}

func StructToJson(item interface{}) (string, error) {
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
