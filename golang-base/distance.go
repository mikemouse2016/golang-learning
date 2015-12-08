package main

import (
    "fmt"
    "math"
    "time"
)

func main() {
    lat1 := 29.490295
    lng1 := 106.486654

    lat2 := 29.615467
    lng2 := 106.581515
    fmt.Println(time.Now())
    for i:=0; i< 1000000; i++{
        //fmt.Println(EarthDistance(lat1, lng1, lat2, lng2))
        EarthDistance(lat1, lng1, lat2, lng2)
    }
    fmt.Println(EarthDistance(39.941, 116.45, 39.94, 116.451))
    fmt.Println(EarthDistance(39.26, 115.25, 41.04, 117.30))
    //fmt.Println(EarthDistance(lat1, lng1, lat2, lng2))
    fmt.Println(time.Now())
}

func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
    var radius float64 = 6367000 // 6378137
    rad := math.Pi/180.0

    lat1 = lat1 * rad
    lng1 = lng1 * rad
    lat2 = lat2 * rad
    lng2 = lng2 * rad

    theta := lng2 - lng1
    dist := math.Acos(math.Sin(lat1) * math.Sin(lat2) + math.Cos(lat1) * math.Cos(lat2) * math.Cos(theta))
    return dist * radius
}
