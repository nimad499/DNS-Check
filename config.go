package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadConfig(filepath string) ([]DNS, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return scanConfig(file)
}

func scanConfig(config_file *os.File) ([]DNS, error) {
	var config []DNS

	scanner := bufio.NewScanner(config_file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid config: %s", line)
		}

		name := strings.TrimSpace(parts[0])
		dnsServer := strings.TrimSpace(parts[1])

		valid := ValidateDNS(dnsServer)
		if !valid {
			return nil, fmt.Errorf("invalid dns server: %s", line)
		}

		config = append(config, DNS{name, dnsServer})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}
