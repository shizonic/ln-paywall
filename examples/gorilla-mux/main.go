package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philippgille/ln-paywall/pay"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

func main() {
	r := mux.NewRouter()

	// Configure and use middleware
	invoiceOptions := pay.DefaultInvoiceOptions // Price: 1 Satoshi; Memo: "API call"
	lndOptions := pay.DefaultLNDoptions         // Address: "localhost:10009", CertFile: "tls.cert", MacaroonFile: "invoice.macaroon"
	storageClient := pay.NewGoMap()
	r.Use(pay.NewHandlerMiddleware(invoiceOptions, lndOptions, storageClient))

	r.HandleFunc("/ping", PingHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}