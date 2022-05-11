package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Let's start the fun
	size := len(os.Args)
	for i := 1; i < size; i++ {
		data := strings.Split(os.Args[i], "=")
		if len(data) != 2 {
			println("Revisa los argumentos.")
		}
		conn, err := net.Dial("tcp", data[1])
		if err != nil {
			log.Fatal(err)
		}
		hour := make([]byte, 11)
		conn.Read(hour)
		if err != nil {
			conn.Close()
			log.Print(err)
		} else {
			fmt.Printf("%s:%s", data[0], hour)
		}

	}

}
