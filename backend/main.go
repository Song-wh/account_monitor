package main

import "net/http"

func main() {
	InitDB()

	http.HandleFunc("/webhook/deposit", WebhookDepositHandler)
	http.HandleFunc("/api/deposits", GetDepositsHandler)
	http.HandleFunc("/api/auth", AuthHandler)

	// Handle CORS preflight requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			enableCORS(w)
			w.WriteHeader(http.StatusOK)
			return
		}
		http.NotFound(w, r)
	})
	http.ListenAndServe(":7001", nil)
}
