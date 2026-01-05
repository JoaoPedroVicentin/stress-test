package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	numRequests := 10000
	url := "http://google.com"
	timeout := time.Duration(5 * time.Second)
	var wg sync.WaitGroup
	wg.Add(numRequests)

	client := http.Client{Timeout: timeout}
	startTime := time.Now()

	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			resp, err := client.Get(url)
			if err != nil {
				// Tratar erros (ex: timeouts, conexões perdidas)
				return
			}
			defer resp.Body.Close()
			// Processar a resposta se necessário
		}()
	}

	wg.Wait()
	duration := time.Since(startTime)
	fmt.Printf("Total de %d requisições concluídas em %s\n", numRequests, duration)
}
