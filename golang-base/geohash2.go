package main

import (
    "fmt"
    "github.com/sillydong/geohash"
)

func main() {
    latitude := 39.92324
    longitude := 116.3906
    precision := 6

    loopneighbors := geohash.LoopNeighbors(latitude, longitude, precision, 3)
    for loop, hashs := range loopneighbors {
        fmt.Printf("loop: %d\n", loop)
        for _, hash := range hashs {
            fmt.Println("\t"+hash)
        }
    }
}
