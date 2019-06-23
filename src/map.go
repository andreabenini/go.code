// Detect if an item exists in a map
if element, exist := myMap[key]; exist {
    fmt.Println("Element "+element)
} else {
    fmt.Println("Element does not exists")
}


// interface{} -> map
// content  : interface{}
// mapValue : map[string]interface{}
mapValue := content.(map[string]interface{})
