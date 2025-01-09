package main

import (
	"bufio"
	"bytes"
	"fmt"
	"monitoria/class03/tcp_simple_chat/util"
	"net"
	"os"
	"sync"
)

func readWorker(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		buffer := make([]byte, util.BufferSize)
		_, err := conn.Read(buffer)
		if err != nil {
			continue
		}

		buffer = bytes.TrimSpace(buffer)

		for i := range buffer {
			fmt.Print(string(buffer[i]))
		}
	}
}

func writeWorker(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		message := make([]byte, util.BufferSize)
		_, err := bufio.NewReader(os.Stdin).Read(message)
		if err != nil {
			fmt.Print(err)
			return
		}

		message = bytes.TrimSpace(message)

		_, err = conn.Write(message)
		if err != nil {
			fmt.Print(err)
			return
		}

	}
}

func main() {

	endpoint := util.SERVER_IP + ":" + util.SERVER_PORT
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go readWorker(conn, &wg)
	go writeWorker(conn, &wg)

	wg.Wait()
}
