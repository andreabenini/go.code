// HTTP GET/JSON Example
func httpGet(url string) jsonResult {
	request, error := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer SoMeThInG")
	clientRequest := &http.Client{}
	response, error := clientRequest.Do(request)
	if error != nil {
		fmt.Printf("Error")
		os.Exit(1)
	}
	body, error := ioutil.ReadAll(response.Body)
	if error != nil || response.StatusCode != 200 {
		fmt.Printf("Error. HTTP %d\n%s\n", response.StatusCode, body)
		os.Exit(1)
	}
	var result jsonResult
	json.Unmarshal(body, &result)
	return result
} /**/


// HTTP POST Example
func httpPost(url string) string {
    // HTTP POST
    jsonString := []byte(`{"jsonVar1":"value1","jsonVar2":"value2"}`)
    request, error := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
    request.Header.Set("Content-Type", "application/json")
    clientRequest := &http.Client{}
    response, error := clientRequest.Do(request)
    if error != nil {
        fmt.Printf("Something wrong here...\n")
        os.Exit(1)
    }
    if response.StatusCode != 200 {
        fmt.Printf("HTTP Error %d", response.StatusCode)
        os.Exit(1)
    }
    // reading result and unmarshalling
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)
    type jsonResult struct {
        Result string `json:"resultField"`
    }
    var result jsonResult
    json.Unmarshal(body, &result)
    return result.Result
} /**/
