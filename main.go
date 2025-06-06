package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"log"
)

func isPrime(n int) bool{
	if n <= 1 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func heavy_calculation () int {
	count := 0
	for num := 2; num < 1000; num++ {
		if isPrime(num) {
			count++
		}
	}
	return(count)
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	if os.Getenv("HEAVY_BUILD") == "true" {
		fmt.Println("Performing heavy calculation...")
		result := heavy_calculation()
		fmt.Fprintf(w, "Hello from %s! (calculated %d prime numbers during build)", hostname, result)
	} else {
		fmt.Fprintf(w, "Hello from %s!", hostname)
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := "80"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	fmt.Printf("Server is running on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}