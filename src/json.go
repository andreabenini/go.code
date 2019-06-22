// readJSONFile - Unmarshalling json information from file
func readJSONFile(filename string) map[string]interface{} {
    fmt.Printf("Reading information from [%s]", filepath.Base(filename))
    var jsdata map[string]interface{}
    jsdata = make(map[string]interface{})
    fileContents, err := ioutil.ReadFile(filename)
    if err == nil {
        json.Unmarshal(fileContents, &jsdata)
    }
    return jsdata
} /**/

// writeJSONFile - Marshalling json information to file
func writeJSONFile(filename string, contents map[string]interface{}) {
    fmt.Printf("Writing information to [%s]", filepath.Base(filename))
    jsonData, err := json.Marshal(contents)
    if err == nil {
        ioutil.WriteFile(filename, jsonData, 0640)
    }
} /**/
