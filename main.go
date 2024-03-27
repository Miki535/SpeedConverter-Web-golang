package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", osn)
	http.ListenAndServe(":8080", nil)
}

func osn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		kmh1 := r.FormValue("kmh")
		mph1 := r.FormValue("mph")

		kmh, err1 := strconv.Atoi(kmh1)
		mph, err2 := strconv.Atoi(mph1)
		//convert mph and kmh to float64
		kmh2 := float64(kmh)
		mph2 := float64(mph)

		if err1 != nil || err2 != nil {

		}
		//if user search kmh:
		if kmh1 == "" {
			mphh := 1.6

			result := fmt.Sprint("You result: ", mph2*mphh, "\nKmh")

			data := struct {
				Result string
			}{
				Result: result,
			}
			tpl.Execute(w, data)
			return

			//if user search mph:
		} else if mph1 == "" {
			mphh := 0.6

			result := fmt.Sprint("You result: ", kmh2*mphh, "\nMph")

			data := struct {
				Result string
			}{
				Result: result,
			}
			tpl.Execute(w, data)
			return

		}
	}
	tpl.Execute(w, nil)
}
