package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 65535; i++ {
		wg.Add(1)
		fmt.Println(i)
		go func(j int) {
			defer wg.Done()

			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}

			conn.Close()

			fmt.Printf("%d open\n", j)

		}(i)
	}

	wg.Wait()
}
