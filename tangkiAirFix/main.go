package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"tank-water-simulation/algorithms"
)

type Result struct {
	N             int     `json:"n"`
	IterativeTime float64 `json:"iterative"`
	RecursiveTime float64 `json:"recursive"`
}

var dataset map[int]bool // ini

func loadDataset() error {
	dataset = make(map[int]bool)

	file, err := os.Open("dataset.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// skip header
	for i := 1; i < len(rows); i++ {
		n, err := strconv.Atoi(rows[i][0])
		if err == nil {
			dataset[n] = true
		}
	}
	return nil
} // sampe ini

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)

	if err != nil || n <= 0 {
		http.Error(w, "Input tidak valid", http.StatusBadRequest)
		return
	}

	if !dataset[n] {
		http.Error(w, "Nilai n tidak terdapat pada dataset", http.StatusBadRequest)
		return
	} // ini oi

	repeat := 1000

	start := time.Now()
	for i := 0; i < repeat; i++ {
		algorithms.TotalWaterIterative(n)
	}
	iterTime := float64(time.Since(start).Nanoseconds()) / 1e6 / float64(repeat)

	start = time.Now()
	for i := 0; i < repeat; i++ {
		algorithms.TotalWaterRecursive(n)
	}
	recTime := float64(time.Since(start).Nanoseconds()) / 1e6 / float64(repeat)

	// res := Result{n, iterTime, recTime} // tadinya ini

	res := Result{
		N:             n,
		IterativeTime: iterTime,
		RecursiveTime: recTime,
	} // ini jg

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	err := loadDataset() /// ni
	if err != nil {
		log.Fatal("Gagal load dataset:", err)
	} // ni

	http.HandleFunc("/api/analyze", analyzeHandler)
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	fmt.Println("Server running at http://localhost:2121")
	log.Fatal(http.ListenAndServe(":2121", nil))
}
