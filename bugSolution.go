func main() {
    var m sync.Map
    m.Store("key", "value")
    fmt.Println(m.Load("key")) // Output: value
    m.Delete("key")
    fmt.Println(m.Load("key")) // Output: <nil>
    val, ok := m.Load("key")
    if !ok {
        m.Store("key", "new value")
    }
    fmt.Println(m.Load("key")) // Output: new value
}