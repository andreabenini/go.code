// Credits to: http://joneisen.me/development/code/2013/06/23/golang-and-default-values.html
// Some sort of default/optional parameters offered to a function with the three dots syntax

func a(log ...bool) {
    if len(log)>0 {
        fmt.Println(log[0])
    } else {
        fmt.Println("No log")
    }
}
func main() {
    a()
    a(true)
    a(false)
}
