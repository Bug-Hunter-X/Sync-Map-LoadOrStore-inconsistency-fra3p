func main() {
    var m sync.Map
    m.Store("key", "value")
    fmt.Println(m.Load("key")) // Output: value
    m.Delete("key")
    fmt.Println(m.Load("key")) // Output: <nil>
    fmt.Println(m.LoadOrStore("key", "new value")) // Output: new value, but the map is not updated
}