package main

import (
    "fmt"
    "os"
)

func compress(data []byte) []byte {
    if len(data) == 0 {
        return nil
    }
    var encoded []byte
    count := 1
    prev := data[0]

    for i := 1; i < len(data); i ++ {
        if data[i] == prev && count < 255 {
            count++
        } else {
            encoded = append(encoded, byte(count), prev)
            prev = data[i]
            count = 1
        }
    }
    encoded = append(encoded, byte(count), prev)
    return encoded
}
func decompress(data []byte) []byte {
    var decoded []byte

    for i := 0; i < len(data); i += 2 {

        if i+1 >= len(data) {
            fmt.Println("Invalid data")
            break
        }
        count := int(data[i])
        h := data[i+1]
        for j := 0; j < count; j++ {
            decoded = append(decoded, h)
        }
    }
    return decoded
}


func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: compressor <file>")
        os.Exit(1)
    }

    command := os.Args[1]
    var results []byte


    file, err := os.ReadFile(os.Args[2])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    switch command {
    case "compress":
       results = compress(file)
       fmt.Printf(string(results))
    case "decompress":
        results = decompress(file)
        fmt.Println(string(results))
    default:
        fmt.Println("Unknown command")
        os.Exit(1)
    }
}
