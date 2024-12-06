package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", TempeHandler)
	http.ListenAndServe(":3000", nil)
}

func convertTempe(value float64, fromUnit, toUnit string) (float64, error) {

	switch fromUnit + " to " + toUnit {
	case "celsius to kelvin":
		return (value + 273), nil
	case "celsius to fahrenheit":
		return (value*(9/5) + 32), nil
	case "celsius to celsius":
		return value, nil
	case "kelvin to celsius":
		return (value - 273), nil
	case "kelvin to fahrenheit":
		return (9*(value-273.15)/5 + 32), nil
	case "kelvin to kelvin":
		return value, nil
	case "fahrenheit to celsius":
		return 5 * (value - 32) / 9, nil
	case "fahrenheit to kelvin":
		return (5*(value-32)/9 + 273.15), nil
	default:
		return value, fmt.Errorf("not supporte from %s to %s", fromUnit, toUnit)
	}
}

func TempeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		value, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		formUnit := r.FormValue("fromUnit")
		toUnit := r.FormValue("toUnit")
		result, err := convertTempe(value, formUnit, toUnit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl := template.Must(template.ParseFiles("templates/tempe.html"))
		tmpl.Execute(w, map[string]interface{}{
			"Result": fmt.Sprintf("%.2f", result),
		})
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/tempe.html"))
	tmpl.Execute(w, nil)
}
