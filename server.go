package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"bytes"
)

type OrcaBarcode struct {
	Barcode					string
    Date 					string
    Description 			string
    Example					string
    Name					string
    Quantity				int
    ___autofocus			string
    ___autoselect			string
    ___lastSchemaVersion	string
    ___orca_action			string
    ___orca_row_id			string
    ___orca_sheet_name		string
    ___orca_user_email		string
    ___owner				string
    ___schemaVersion		string
}

func webHookOutHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	// Parse JSON data
	var barcode OrcaBarcode
	jsonErr := json.Unmarshal([]byte(body), &barcode)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		http.Error(w, jsonErr.Error(), 500)
		return
	}

    // dubug purpose: show in console raw data received
	fmt.Println(barcode)

	// get the name of the action that triggered this request (add, update, delete, test)
	action := barcode.___orca_action

	// get the name of the sheet this action impacts
	sheetName := barcode.___orca_sheet_name
	fmt.Println(sheetName)

	// get the email of the user who preformed the action (empty if not HTTPS)
	userEmail := barcode.___orca_user_email
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

func webHookInHandler(w http.ResponseWriter, r *http.Request) {
	values := map[string]string{
		"___orca_action": "add",
		"Barcode": "0123456789",
		"Name": "New 1",
		"Quantity": "12",
		"Description": "Add new row example",
	}
	jsonValue, _ := json.Marshal(values)
	// The following example adds a new row to a sheet, setting the value of Barcode, Name, Quantity and Description
	// TODO: change url to https://api.orcascan.com/sheets/{id}
	response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
        // read response error
		fmt.Println(err)
    } else {
		// read response body
		body, _ := ioutil.ReadAll(response.Body)
		data := map[string]string{}
		jsonErr := json.Unmarshal([]byte(body), &data)
		if jsonErr != nil {
			return
		}
		fmt.Println(data)
    }
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	return
}


func main() {

    http.HandleFunc("/orca-webhook-out", webHookOutHandler)
    http.HandleFunc("/trigger-webhook-in", webHookInHandler)

    fmt.Println("Server started at port 3000")
    log.Fatal(http.ListenAndServe(":3000", nil))
	
}

