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
