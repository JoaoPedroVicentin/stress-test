package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Index          int
	HttpStatusCode int
}

func makeRequest(url string, index int, resultsThreads chan<- Result, waitGroup *sync.WaitGroup, threads int) {
	defer waitGroup.Done()
	resp, err := http.Get(url)
	if err != nil {
		if resp != nil {
			resultsThreads <- Result{Index: index, HttpStatusCode: resp.StatusCode}
			return
		}
		resultsThreads <- Result{Index: index, HttpStatusCode: 0}
		return
	}
	defer resp.Body.Close()
	resultsThreads <- Result{Index: index, HttpStatusCode: resp.StatusCode}
}

func makeReport(results []Result, duration time.Duration, url string, numRequests int, concurrency int) {
	total := len(results)
	statusDist := make(map[int]int)
	for _, r := range results {

		statusDist[r.HttpStatusCode]++
	}

	fmt.Println("\n=============== RELATORIO ===============")
	fmt.Println("Parâmetros utilizados:")
	fmt.Printf("URL do serviço testado: %s\n", url)
	fmt.Printf("Número total de requests: %d\n", numRequests)
	fmt.Printf("Número de chamadas simultâneas: %d\n", concurrency)
	fmt.Println("\nResultados:")
	fmt.Printf("Tempo total de execução: %s\n", duration)
	fmt.Printf("Quantidade total de requests: %d\n", total)
	for code, count := range statusDist {
		if code == 0 {
			fmt.Printf("- Erros de requisição (sem status): %d\n", count)
		} else {
			fmt.Printf("- Quantidade de requests com status HTTP %d: %d\n", code, count)
		}
	}
	fmt.Print("=========================================\n")
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	numRequests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 0, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" || *numRequests == 0 || *concurrency == 0 {
		fmt.Println("Error: Os parâmetros --url, --requests, e --concurrency são obrigatórios")
		flag.PrintDefaults()
		return
	}

	resultsThreads := make(chan Result, *numRequests)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(*numRequests)

	threads := make(chan struct{}, *concurrency)

	startTime := time.Now()

	for i := 0; i < *numRequests; i++ {
		threads <- struct{}{}
		threadNum := i % *concurrency
		go func(idx int, tnum int) {
			defer func() { <-threads }()
			makeRequest(*url, idx, resultsThreads, &waitGroup, tnum)
		}(i, threadNum)
	}

	waitGroup.Wait()
	close(resultsThreads)
	duration := time.Since(startTime)

	results := make([]Result, 0, *numRequests)
	for r := range resultsThreads {
		results = append(results, r)
	}

	makeReport(results, duration, *url, *numRequests, *concurrency)
}
