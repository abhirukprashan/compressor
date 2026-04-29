package main

import (
    "fmt"
    "os"
)

func compress(){
   fmt.Println("Compressing")
}
func decompress(){
   fmt.Println("Decompressing")
}


func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: compressor <file>")
        os.Exit(1)
    }
    command := os.Args[1]
/*
    file, err := os.ReadFile(os.Args[2])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

*/
    switch command {
    case "compress":
        compress()
    case "decompress":
        decompress()
    default:
        fmt.Println("Unknown command")
        os.Exit(1)
    }
}
