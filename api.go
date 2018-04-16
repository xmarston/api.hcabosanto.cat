package main

import (
	"net/http"
	"fmt"
	"strings"
	"os"
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

	//db, err := GetDb()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	defer db.Close()
	//	rows, err := db.Query("SELECT id, name FROM patients")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	for rows.Next() {
	//		var id int
	//		var name string
	//		err = rows.Scan(&id, &name)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		fmt.Println(fmt.Sprintf("%d -> %s", id, name))
	//	}
	//}
}
