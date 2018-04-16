package main

import (
	"net/http"
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var (
	PORT = "8443"
	log  = logrus.New()
)

func main() {
	log.Out = os.Stdout
	gotenv.Load()
	parseFlags()

	r := mux.NewRouter()
	r.HandleFunc("/", WelcomeHandler)

	http.Handle("/", r)
	log.Infof("Listen in :%s", PORT)
	err := http.ListenAndServeTLS(":"+PORT, "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func parseFlags() {
	addrPort := flag.String("port", PORT, "a string")
	flag.Parse()

	if *addrPort != PORT {
		PORT = *addrPort
	}
}
