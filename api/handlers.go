package api

import (
	"encoding/json"
	"log"
	"net/http"

	"log/slog"

	"booking_api/internal/repository"
)

func ComputerSpecs(w http.ResponseWriter, r *http.Request) {
	computerSpecs, err := repository.GetAllComputerSpecs(db)
	if err != nil {
		log.Println("Error fetching computer specs:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(computerSpecs); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ComputerStatuses(w http.ResponseWriter, r *http.Request) {
	computerStatuses, err := repository.GetAllComputerStatuses(db)
	if err != nil {
		log.Println("Error fetching computer statuses:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(computerStatuses); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetAllUsers(db)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error, Please message administrators", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func Packages(w http.ResponseWriter, r *http.Request) {
	packages, err := repository.GetPackages(db)
	if err != nil {
		log.Println("Error fetching packages:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(packages); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func PendingBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := repository.GetAllPendingBookings(db)
	if err != nil {
		log.Println("Error fetching pending bookings:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bookings); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func FinishedBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := repository.GetAllFinishedBookings(db)
	if err != nil {
		log.Println("Error fetching finished bookings:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bookings); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
