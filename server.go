package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func webHookOutHandler(w http.ResponseWriter, r *http.Request) {
	// read request body
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// parse JSON into a map to handle any sheet fields dynamically
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// debug purpose: show in console raw data received
	formatted, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("Request received:\n%s\n", formatted)

	// get the name of the event that triggered this request (rows:add, rows:update, rows:delete, test)
	eventName, _ := data["___orca_event"].(string)

	// get the name of the sheet this action impacts
	sheetName, _ := data["___orca_sheet_name"].(string)

	// get the email of the user who performed the action (empty if not HTTPS)
	userEmail, _ := data["___orca_user_email"].(string)

	// NOTE:
	// orca system fields start with ___
	// you can access the value of each field using the field name (data["Name"], data["Barcode"], data["Location"])
	_, _ = sheetName, userEmail

	switch eventName {
	case "rows:add":
		// TODO: do something when a row has been added
	case "rows:update":
		// TODO: do something when a row has been updated
	case "rows:delete":
		// TODO: do something when a row has been deleted
	case "test":
		// TODO: do something when the user in the web app hits the test button
	}

	w.WriteHeader(http.StatusOK)
}

func webHookInHandler(w http.ResponseWriter, r *http.Request) {
	// The following example adds a new row to a sheet, setting the value of Barcode, Name, Quantity and Description
	payload := map[string]interface{}{
		"___orca_action": "add",
		"Barcode":        "0123456789",
		"Name":           "New 1",
		"Quantity":       12,
		"Description":    "Add new row example",
	}
	jsonValue, _ := json.Marshal(payload)

	// TODO: change url to https://api.orcascan.com/sheets/{id}
	response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/orca-webhook-out", webHookOutHandler)
	http.HandleFunc("/trigger-webhook-in", webHookInHandler)

	fmt.Println("Server started at port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

