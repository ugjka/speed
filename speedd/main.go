// Speed test server
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	log.SetPrefix("speedd ")
	p := flag.Uint("p", 44444, "listener port")
	flag.Parse()

	if *p == 0 || *p > 65535 {
		fmt.Fprintf(os.Stderr, "Error: invalid port number: %d\n", *p)
		return
	}
	l, err := net.Listen("tcp4", fmt.Sprintf(":%d", *p))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on :%d", *p)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go client(conn)
	}
}

func client(conn net.Conn) {
	defer conn.Close()
	var upcommand int64 = 10
	var downcommand int64 = 20
	log.Printf("client connected: %s", conn.RemoteAddr())
	defer log.Printf("client disconnected: %s", conn.RemoteAddr())
	var cmd int64
	err := binary.Read(conn, binary.LittleEndian, &cmd)
	if err != nil {
		return
	}
	if cmd == upcommand {
		log.Printf("client sent the up command: %s", conn.RemoteAddr())
		read(conn)
		return
	}
	if cmd == downcommand {
		log.Printf("client sent the down command: %s", conn.RemoteAddr())
		write(conn)
		return
	}
	log.Printf("client sent an invalid command: %s", conn.RemoteAddr())
}

func read(conn net.Conn) {
	null, err := os.OpenFile("/dev/null", os.O_WRONLY, 0666)
	if err != nil {
		log.Print(err)
		return
	}
	defer null.Close()
	_, err = io.Copy(null, conn)
	if err != nil {
		return
	}
}

func write(conn net.Conn) {
	rand, err := os.Open("/dev/urandom")
	if err != nil {
		log.Print(err)
		return
	}
	const upsize = 1 << 20
	b := make([]byte, upsize)
	n, err := rand.Read(b)
	rand.Close()
	if err != nil {
		log.Print(err)
		return
	}
	if n != upsize {
		log.Print("error: n != upsize")
		return
	}
	for {
		_, err := conn.Write(b)
		if err != nil {
			return
		}
	}
}
