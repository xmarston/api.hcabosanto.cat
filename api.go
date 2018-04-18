package main

import (
	"net/http"
	"fmt"
	"strings"
	"os"
	"api.hcabosantos.cat/types"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	allowedXForwardedFor := strings.Join(strings.Split(os.Getenv("ALLOWEDXFORWARDEDFOR"), " "), ", ")
	if r.Header.Get("X-Forwarded-For") == allowedXForwardedFor {
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

	db, err := GetDb()
	if err != nil {
		log.Error(err)
	} else {
		defer db.Close()
		var patients []types.Patient
		err := db.Select(&patients, "SELECT * FROM patients ORDER BY surname ASC")
		if err != nil {
			log.Error(err)
		} else {
			for _, patient := range patients {
				validDNI, err := patient.ValidateDni()
				log.Info(fmt.Sprintf("%d -> %s with valid DNI %s -> %t", patient.Id, patient.Name, patient.Nif, validDNI))
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
}
