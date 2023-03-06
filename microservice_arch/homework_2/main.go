package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	router := mux.NewRouter()
	router.HandleFunc("/health", healthCheck)
	router.HandleFunc("/", defaultRoute)
	router.HandleFunc("/k8s_env", getEnv)

	log.Println("app started on port ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal(err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	type resp struct {
		Status string `json:"status"`
	}
	bytes, err := json.Marshal(resp{Status: "OK"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	return
}

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello from %s", os.Getenv("HOSTNAME"))))
	return
}

func getEnv(w http.ResponseWriter, r *http.Request) {
	envValues := os.Environ()
	outPut := ""
	for _, val := range envValues {
		if outPut != "" {
			outPut = fmt.Sprintf("%s\n%s", outPut, val)
			continue
		}
		outPut = fmt.Sprintf("%s", val)
	}

	w.Write([]byte(outPut))
	return
}
