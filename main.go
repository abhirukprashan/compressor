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
    count := 0
    prev := data[0]

    for i := 1; i < len(data); i++ {
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
func decompress(){
   fmt.Println("Decompressing")
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
       fmt.Println(results)
    case "decompress":
        decompress()
    default:
        fmt.Println("Unknown command")
        os.Exit(1)
    }
}
