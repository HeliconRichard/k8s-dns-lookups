package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide ip list file")
		return
	}

	var results []string
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ip := scanner.Text()

		hostnames, err := net.LookupAddr(ip)
		if err != nil {
			fmt.Printf("Error looking up hostname for %s: %s\n", ip, err)
			continue
		}

		for _, hostname := range hostnames {
			result := fmt.Sprintf("%-15s | %s\n", ip, hostname)
			results = append(results, result)
		}
	}

	sort.Strings(results)

	for _, result := range results {
		fmt.Printf("%s", result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
}
