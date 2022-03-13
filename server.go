package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
)

func WebHookHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Parse JSON data
	data := map[string]string{}
	jsonErr := json.Unmarshal([]byte(body), &data)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), 500)
		return
	}

    // dubug purpose: show in console raw data received
	fmt.Println(data)

	// get the name of the action that triggered this request (add, update, delete, test)
	action := data["___orca_action"]

	// get the name of the sheet this action impacts
	sheetName := data["___orca_sheet_name"]
	fmt.Println(sheetName)

	// get the email of the user who preformed the action (empty if not HTTPS)
	userEmail := data["___orca_user_email"]
	fmt.Println(userEmail)

	// NOTE:
    // orca system fields start with ___
    // you can access the value of each field using the field name (data["Name"], data["Barcode"], data["Location"])

	switch action {
    case "add":
        // TODO: do something when a row has been added
    case "update":
        // TODO: do something when a row has been updated
    case "delete":
        // TODO: do something when a row has been deleted
	case "test":
		// TODO: do something when the user in the web app hits the test button
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	return
}

func main() {

    http.HandleFunc("/", WebHookHandler)

    fmt.Println("Server started at port 3000")
    log.Fatal(http.ListenAndServe(":3000", nil))
}

