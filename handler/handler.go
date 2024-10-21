package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"GO-TRAINING/average"
	"GO-TRAINING/utils"
)

func AverageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	content, _ := io.ReadAll(r.Body)
	fmt.Println(string(content))
	// if err != nil {
	// 	http.Error(w, "Failed to read request", http.StatusInternalServerError)
	// 	fmt.Println("fhjgiuutt", err)
	// 	return
	// }

	defer r.Body.Close()

	// Check if the body is empty after reading it
	if len(content) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "no numbers provided", http.StatusBadRequest)
		return
	}

	calci := &average.Entities{}
	numstring := string(content)

	if !utils.IsValidInput(numstring) {
		http.Error(w, "Invalid input: please provide space-separated digits only", http.StatusBadRequest)
		return
	}

	for _, numStr := range strings.Fields(numstring) {
		num, err := strconv.ParseFloat(numStr, 64)
		if err == nil {
			calci.Input(num)
		}
	}

	numStr := r.URL.Query().Get("n")
	var avg float64
	var err error
	switch numStr {
	case "1":
		avg, err = calci.CalculateEvenAvg()
	case "2":
		avg, err = calci.CalculateOddAvg()
	case "3":
		avg, err = calci.CalculateAll()
	default:
		http.Error(w, "Invalid choice", http.StatusBadRequest)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		fmt.Fprintf(w, "The average is %.2f\n", avg)
	}
}
