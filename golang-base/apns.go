package main

import (
  "fmt"
  apns "github.com/Coccodrillo/apns"
)

func main() {
  payload := apns.NewPayload()
  payload.Alert = "Hello, world!"
  payload.Badge = 42
  payload.Sound = "bingbong.aiff"

  pn := apns.NewPushNotification()
  pn.AddPayload(payload)
  pn.Expiry = 3000
  pn.Set("foo", "bar")
  pn.Set("doctor", "who?")
  pn.Set("the_ultimate_answer", 42)

  alert, _ := pn.PayloadString()
  fmt.Println(alert)
}
