package main

import (
  "log"
  "net"
  "time"
  "strconv"
)

func main() {
  serverAddr, err := net.ResolveUDPAddr("udp","localhost:54300")
  fatalError(err)

  conn, err := net.DialUDP("udp", nil, serverAddr)
  fatalError(err)

  defer conn.Close()
}

func fatalError(err error){
  if err != nil {
    log.Fatal("error: ", err.Error())
  }
}
