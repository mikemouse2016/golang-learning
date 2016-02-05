package main

import (
    "fmt"
    "github.com/gansidui/geohash"
)

func main() {
    latitude := 31.187084
    longitude := 121.413873
    precision := 5

    hash, box := geohash.Encode(latitude, longitude, precision)

    fmt.Println(hash)
    fmt.Println(box.MinLat, box.MaxLat, box.MinLng, box.MaxLng)

    neighbors := geohash.GetNeighbors(latitude, longitude, precision)
    for _, hash = range neighbors {
        fmt.Print(hash, " ")
    }
}
