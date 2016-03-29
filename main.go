package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Index function handles the / and /json URLs and returns either a string or
// JSON containing the ipaddress
func Index(w http.ResponseWriter, r *http.Request) {
	output := mux.Vars(r)["output"]
	ip := r.Header.Get("X-Forwarded-For")

	if ip == "" {
		if ip = r.Header.Get("REMOTE_ADDR"); ip == "" {
			ip = "127.0.01"
		}
	}

	switch output {
	case "json":
		json_structure := map[string]string{"ipaddress": ip}
		fout, _ := json.Marshal(json_structure)
		fmt.Fprintf(w, string(fout))
	default:
		fmt.Fprintf(w, ip)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/{output}", Index).Methods("GET")
	http.Handle("/", r)

	if err := http.ListenAndServe("127.0.0.1:3000", nil); err != nil {
		fmt.Println("Encountered error:")
		fmt.Println(err)
	}
}
