package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format("Mon Jan _2"))
	return []byte(stamp), nil
}

type Flight struct {
	ID          string   `json:"id"`
	Airline     string   `json:"airline"`
	Departure   string   `json:"departure"`
	Destination string   `json:"destination"`
	Flytime     JSONTime `json:"flytime"`
	Class       string   `json:"class"`
	Off         float32  `json:"off"`
	Original    int      `json:"original"`
}

var flights []Flight

func GetFlightEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range flights {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Flight{})
}
func GetFlightsEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(flights)
}
func CreateFlightEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var flight Flight
	_ = json.NewDecoder(r.Body).Decode(&flight)
	flight.ID = params["id"]
	flights = append(flights, flight)
	json.NewEncoder(w).Encode(flights)
}
func DeleteFlightEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range flights {
		if item.ID == params["id"] {
			flights = append(flights[:index], flights[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(flights)
	}
}

func main() {
	router := mux.NewRouter()
	flights = append(flights, Flight{ID: "1", Airline: "SZ", Departure: "ShenZhen", Destination: "Beijing", Flytime: JSONTime{time.Now()}, Class: "First", Off: 0.1, Original: 2202})
	flights = append(flights, Flight{ID: "3", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "4", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "5", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "6", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "7", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "8", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "9", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "10", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "11", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "12", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "13", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "14", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "15", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "16", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "17", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "18", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "19", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "20", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "21", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "22", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "23", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})
	flights = append(flights, Flight{ID: "24", Airline: "CA", Departure: "Bangkok", Destination: "Beijing", Class: "Bussiness", Off: 0.3, Original: 1290})

	router.HandleFunc("/flights", GetFlightsEndpoint).Methods("GET")
	router.HandleFunc("/flights/{id}", GetFlightEndpoint).Methods("GET")
	router.HandleFunc("/flights/{id}", CreateFlightEndpoint).Methods("POST")
	router.HandleFunc("/flights/{id}", DeleteFlightEndpoint).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	log.Fatal(http.ListenAndServe(":8001", c.Handler(router)))
}
