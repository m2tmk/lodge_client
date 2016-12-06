package main

import (
  "log"
  "net"
  "time"
  "strconv"
)

func main() {
  serverAddr, err := net.ResolveUDPAddr("udp", "localhost:54321")
  fatalError(err)

  clientAddr, err := net.ResolveUDPAddr("udp", "localhost:54300")
  fatalError(err)

  conn, err := net.DialUDP("udp", clientAddr, serverAddr)
  fatalError(err)

  defer conn.Close()

  i := 0
  for {
    msg := strconv.Itoa(i)

    i++
    buffer := []byte(msg)
    _, err := conn.Write(buffer)
    fatalError(err)

    log.Printf("Send: [%v]: %v\n", serverAddr, msg)

    time.Sleep(time.Duration(300) * time.Millisecond)
  }
}

func fatalError(err error){
  if err != nil {
    log.Fatal("error: ", err.Error())
  }
}
