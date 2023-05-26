package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var filepath string
	flag.StringVar(&filepath, "c", "goatman.conf", "config file to specify hostnames and ip addresses")
	flag.Parse()
	getTargetList(filepath)
}

func getTargetList(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := trimString(scanner.Text())
		if line == "" {
			continue
		}

		fmt.Println(line)
	}
}
