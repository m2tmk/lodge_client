package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
  zmq "github.com/zeromq/goczmq"
)

type LatLng struct {
	Lat float64
	Lng float64
}

func main() {
	rand.Seed(time.Now().UnixNano())

  tcpClient()
  //udpClient()

}

func tcpClient() {
  req, _ := zmq.NewReq("tcp://localhost:54321")
  defer req.Destroy()

  log.Println("Req created and bound.")

  i := 0
  for {
    data := fmt.Sprintf("%4d:%v", i, createData())

    req.SendFrame([]byte(data), 0)

    reply, _ := req.RecvMessage()
	  log.Printf("Reply: %v\n", string(reply[0]))

    i++
		time.Sleep(time.Duration(1) * time.Microsecond)
  }
}

func udpClient() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:54321")
	fatalError(err)

	clientAddr, err := net.ResolveUDPAddr("udp", "localhost:54300")
	fatalError(err)

	conn, err := net.DialUDP("udp", clientAddr, serverAddr)
	fatalError(err)

	defer conn.Close()

  i := 0
	for {

    data := fmt.Sprintf("%4d:%v", i, createData())
		send(conn, serverAddr, data)

    i++

		time.Sleep(time.Duration(300) * time.Microsecond)
	}
}

func createData() string {
	latLng := latlng()
  carId := fmt.Sprintf("%04d", rand.Intn(1000))
  status := "empty"

	return fmt.Sprintf("0001:%v:%v:%v:%v", carId, status, latLng.Lat, latLng.Lng)
}

func send(conn net.Conn, serverAddr *net.UDPAddr, data string) {
	buffer := []byte(data)
	_, err := conn.Write(buffer)
	fatalError(err)

	log.Printf("Send: [%v]: %v\n", serverAddr, data)
}

func fatalError(err error) {
	if err != nil {
		log.Fatal("error: ", err.Error())
	}
}

func latlng() LatLng {
	// 東端：緯度 35.46.50  経度 139.53.41
	//        35,763888888888886,139.8947222222222
	// 西端：緯度 35.45.43  経度 139.33.46
	//        35.761944444444445,139.5627777777778
	// 北端：緯度 35.49.04  経度 139.46.03
	//        35.817777777777778,139.7675
	// 南端：緯度 35.31.16  経度 139.48.04
	//        35.52111111111111,139.8011111111111
	minLat := 35.52111111111111
	maxLat := 35.817777777777778
	minLng := 139.5627777777778
	maxLng := 139.8947222222222

	latLng := LatLng{}
	latLng.Lat = minLat + rand.Float64()*(maxLat-minLat)
	latLng.Lng = minLng + rand.Float64()*(maxLng-minLng)

	return latLng
}
