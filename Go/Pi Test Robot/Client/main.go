package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func FindHost() string {
	url := ""
	max := 130
	timeout := 1 * time.Second
	for x := 100; x < max; x++ {
		fmt.Printf("Checking %s\n", url)
		url = "192.168.1." + strconv.Itoa(x) + ":8080"
		_, err := net.DialTimeout("tcp", url, timeout)
		if err != nil {
			fmt.Printf(" %s - Host not Found\n ", url)
		} else {
			x = max
			fmt.Printf("Host Found at %s\n", url)
			url = "http://" + url
			err = os.WriteFile("ip.cfg", []byte(url), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return url
}

//------------------------------------------------- Main

func main() {

	xip := FindHost()
	port := "8080"

	fmt.Printf("Outbound IP  : %s Port : %s\n", xip, port)
	resp, err := http.Get(xip)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
