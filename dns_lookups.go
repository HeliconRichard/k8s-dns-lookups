package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide ip list file")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Printf("Domain names for svc:\n")
	for scanner.Scan() {
		ip := scanner.Text()

		hostnames, err := net.LookupAddr(ip)
		if err != nil {
			fmt.Println("Error looking up hostname for %s: %s\n", ip, err)
			continue
		}

		for _, hostname := range hostnames {
			// paddedHostname := fmt.Sprintf("%-30s", hostname)
			fmt.Printf("%-20s | %s\n", ip, hostname)
			// fmt.Printf("%-70s | %s\n", paddedHostname, ip)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
	fmt.Println()
}
