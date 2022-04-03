# orca-webhook-go

Example of how to build an [Orca Scan WebHook](https://orcascan.com/docs/api/webhooks) endpoint and [Orca Scan WebHook In](https://orcascan.com/guides/how-to-update-orca-scan-from-your-system-4b249706) in using [Go](https://go.dev/).

## Install & Run

First ensure you have [Go](https://go.dev/) installed. If not, follow [this guide](https://go.dev/doc/install).

```bash
# start the project
go run server.go
```

Your WebHook receiver will now be running on port 3000.

You can emulate an Orca Scan WebHook using [cURL](https://dev.to/ibmdeveloper/what-is-curl-and-why-is-it-all-over-api-docs-9mh) by running the following:

```bash
curl --location --request POST 'http://127.0.0.1:3000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "___orca_action": "add",
    "___orca_sheet_name": "Vehicle Checks",
    "___orca_user_email": "hidden@requires.https",
    "___orca_row_id": "5cf5c1efc66a9681047a0f3d",
    "Barcode": "4S3BMHB68B3286050",
    "Make": "SUBARU",
    "Model": "Legacy",
    "Model Year": "2011",
    "Vehicle Type": "PASSENGER CAR",
    "Plant City": "Lafayette",
    "Trim": "Premium",
    "Location": "52.2034823, 0.1235817",
    "Notes": "Needs new tires"
}'
```

### Important things to note

1. Only Orca Scan system fields start with `___`
2. Properties in the JSON payload are an exact match to the  field names in your sheet _(case and space)_
3. WebHooks are never retried, regardless of the HTTP response

## How this example works

This [example](server.go) work as follows:

### WebHook Out 

[Orca Scan WebHook](https://orcascan.com/docs/api/webhooks)

```go
func webHookOutHandler(w http.ResponseWriter, r *http.Request) {
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
```
### WebHook In 

and [Orca Scan WebHook In](https://orcascan.com/guides/how-to-update-orca-scan-from-your-system-4b249706)

```go
func webhookIn() {
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
}
```
## Troubleshooting

If you run into any issues not listed here, please [open a ticket](https://github.com/orca-scan/orca-webhook-go/issues).

## Examples in other langauges
* [orca-webhook-dotnet](https://github.com/orca-scan/orca-webhook-dotnet)
* [orca-webhook-python](https://github.com/orca-scan/orca-webhook-python)
* [orca-webhook-go](https://github.com/orca-scan/orca-webhook-go)
* [orca-webhook-java](https://github.com/orca-scan/orca-webhook-java)
* [orca-webhook-php](https://github.com/orca-scan/orca-webhook-php)

## History

For change-log, check [releases](https://github.com/orca-scan/orca-webhook-node/releases).

## License

&copy; Orca Scan, the [Barcode Scanner app for iOS and Android](https://orcascan.com).