// Speed test client
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	u := flag.Bool("u", false, "upload test")
	d := flag.Bool("d", true, "download test")
	s := flag.String("s", ":44444", "server adress")
	p := flag.Uint64("p", 0, "piece size in KB")
	flag.Parse()

	if *u == true {
		*d = false
	}

	if *p<<10 > 100<<20 {
		fmt.Fprintln(os.Stderr, "Error: piece size cannot be bigger than 100MB")
		return
	}

	if srv := os.Getenv("SPEEDCSRV"); srv != "" {
		*s = srv
	}

	var upcommand int64 = 10
	var downcommand int64 = 20

	var upsize = 512 << 10
	var downsize = 1 << 20
	if *p != 0 {
		upsize = int(*p) << 10
		downsize = int(*p) << 10
	}
	conn, err := net.Dial("tcp4", *s)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	wg := &sync.WaitGroup{}

	if *u {
		err := binary.Write(conn, binary.LittleEndian, upcommand)
		if err != nil {
			log.Print(err)
			return
		}
		wg.Add(1)
		go write(wg, conn, upsize)
	}
	if *d {
		err := binary.Write(conn, binary.LittleEndian, downcommand)
		if err != nil {
			log.Print(err)
			return
		}
		wg.Add(1)
		go read(wg, conn, downsize)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	go func() {
		<-stop
		conn.Close()
	}()
	wg.Wait()
	fmt.Fprintln(os.Stderr, "\nspeedc terminated ")
}

func tput(v ...string) error {
	cmd := exec.Command("tput", v...)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func write(wg *sync.WaitGroup, conn net.Conn, size int) {
	defer wg.Done()
	rand, err := os.Open("/dev/urandom")
	if err != nil {
		log.Print(err)
		return
	}
	b := make([]byte, size)
	n, err := rand.Read(b)
	rand.Close()
	if err != nil {
		log.Print(err)
		return
	}
	if n != size {
		log.Print("error: buff size do not match")
		return
	}
	now := time.Now()
	t := 0
	i := 0
	var total float64
	for {
		n, err := conn.Write(b)
		if err != nil {
			return
		}
		if time.Since(now) > time.Second {
			i++
			curr := float64(t) * 8 / 1024 / 1024 / time.Since(now).Seconds()
			total += curr
			fmt.Print("\r")
			err := tput("el")
			if err != nil {
				fmt.Print("\033[K")
			}
			fmt.Printf("Upload: %f Mbit/s | Avg: %f Mbit/s | %dKB piece ",
				curr, total/float64(i), size/1024)
			t = 0
			now = time.Now()
		}
		t += n
	}
}

func read(wg *sync.WaitGroup, conn net.Conn, size int) {
	defer wg.Done()
	b := make([]byte, size)
	now := time.Now()
	t := 0
	i := 0
	var total float64
	for {
		n, err := conn.Read(b)
		if err != nil {
			return
		}
		if time.Since(now) > time.Second {
			i++
			curr := float64(t) * 8 / 1024 / 1024 / time.Since(now).Seconds()
			total += curr
			fmt.Print("\r")
			err := tput("el")
			if err != nil {
				fmt.Print("\033[K")
			}
			fmt.Printf("Download: %f Mbit/s | Avg: %f Mbit/s | %dKB piece ",
				curr, total/float64(i), size/1024)
			t = 0
			now = time.Now()
		}
		t += n
	}
}
