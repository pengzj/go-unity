package main

import (
	"fmt"
	"net"
	"log"
	"container/list"
)

func main()  {
	createServer()
}

func createServer()  {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":2000")
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	connects := list.New()
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go func(c *net.TCPConn) {
			//broadcast
			go func(c *net.TCPConn, connects *list.List) {
				for e := connects.Front(); e != nil; e = e.Next() {
					adr := c.RemoteAddr()
					oneConn, err := e.Value.(*net.TCPConn)
					if err  != true {
						c.Close()
						log.Fatal(err)
					}
					oneConn.Write([]byte("welcome " + adr.String() + " login into server\n"))
				}
			}(c, connects)

			connects.PushBack(c)


			for {
				data := make([]byte, 1024)

				_, err = c.Read(data)
				if err != nil {
					c.Close()
					log.Fatal(err)
				}

				go func(c *net.TCPConn,connects *list.List, data []byte) {
					for e := connects.Front(); e != nil; e = e.Next() {
						oneConn, err := e.Value.(*net.TCPConn)
						if oneConn == c {
							continue
						}
						if err  != true {
							c.Close()
							log.Fatal(err)
						}
						adr := c.RemoteAddr()
						oneConn.Write([]byte([]byte(adr.String() + " say: " + string(data) + "\n")))
					}
				}(c, connects, data)
			}

		}(conn)
	}
}
