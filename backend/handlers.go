package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CORS middleware
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// Handle OPTIONS requests for CORS
func handleCORS(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	w.WriteHeader(http.StatusOK)
}

// Webhook 수신
func WebhookDepositHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var payload map[string]interface{}
	json.NewDecoder(r.Body).Decode(&payload)

	va := payload["virtual_account"].(string)
	remitter := payload["remitter_name"].(string)
	remitterAccount := payload["remitter_account"].(string)
	amount := payload["amount"].(float64)

	// Create new deposit record
	deposit := DepositRecord{
		ID:               nextID,
		VirtualAccountNo: va,
		RemitterName:     remitter,
		RemitterAccount:  remitterAccount,
		Amount:           amount,
		PgSource:         "TOSPAYMENTS",
		Payload:          string(mustMarshalJSON(payload)),
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
	}

	// Save to INI file
	err := saveDeposit(deposit)
	if err != nil {
		http.Error(w, "Failed to save deposit", 500)
		return
	}

	// Add to memory
	deposits = append(deposits, deposit)
	nextID++

	w.Write([]byte("ok"))
}

// 특정 가상계좌 조회
func GetDepositsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	va := r.URL.Query().Get("virtual_account")
	if va == "" {
		http.Error(w, "missing virtual_account", 400)
		return
	}

	// Get deposits for this virtual account
	depositRecords := getDepositsByVirtualAccount(va)

	type Deposit struct {
		RemitterName    string `json:"remitter_name"`
		RemitterAccount string `json:"remitter_account"`
		Amount          string `json:"amount"`
		CreatedAt       string `json:"created_at"`
	}

	var res []Deposit
	for _, record := range depositRecords {
		d := Deposit{
			RemitterName:    record.RemitterName,
			RemitterAccount: record.RemitterAccount,
			Amount:          fmt.Sprintf("%.2f", record.Amount),
			CreatedAt:       record.CreatedAt,
		}
		res = append(res, d)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"data": res})
}

// Helper function to marshal JSON without error handling
func mustMarshalJSON(v interface{}) []byte {
	data, _ := json.Marshal(v)
	return data
}

// AuthHandler handles account authentication
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var authRequest struct {
		VirtualAccount string `json:"virtual_account"`
		AuthKey        string `json:"auth_key"`
	}

	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Simple authentication logic - you can modify this
	// For demo purposes, we'll use a simple key validation
	validKeys := map[string]string{
		"123-456-7890": "key123",
		"987-654-3210": "key456",
		"555-123-4567": "key789",
	}

	expectedKey, exists := validKeys[authRequest.VirtualAccount]
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "존재하지 않는 가상계좌번호입니다.",
		})
		return
	}

	if authRequest.AuthKey != expectedKey {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "인증키가 올바르지 않습니다.",
		})
		return
	}

	// Authentication successful
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":         "인증 성공",
		"virtual_account": authRequest.VirtualAccount,
	})
}
