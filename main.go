package main

import (
	"fmt"
	"os"
)

// ToDo : Handle prefix

func checkHostInDnsList(host string, dnsList []DNS) {
	fmt.Printf("%s:\n", host)
	for _, dns := range dnsList {
		status, err := dns.Check(host)
		fmt.Print("\t")

		if status != 0 {
			var color string
			if status == 200 {
				color = Green
			} else {
				color = Red
			}
			fmt.Printf("%s: %s%d%s", dns.name, color, status, NoColor)
		} else {
			fmt.Printf("%s: %sError%s (%s)", dns.name, Red, NoColor, err)
		}
		fmt.Println()
	}

	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("At least one address is required")
	}

	dnsList, err := loadConfig("dns.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, host := range os.Args[1:] {
		checkHostInDnsList(host, dnsList)
	}
}
