package smallint


func Convert(val int) []interface{} {
    var slice = make([]interface{}, 100)
    for i := 0; i < 100; i++ {
        slice[i] = val
    }

    return slice
}
