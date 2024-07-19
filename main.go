package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

const (
	URL    string = "https://api.ipify.org?format=json"
	IP_KEY string = "ip"
)

func find_local_ip() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r.(error).Error())
		}
	}()

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("Local IPv4 address:", ipNet.IP.String())
			}
		}
	}
}

func find_ip() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r.(error).Error())
		}
	}()

	resp, err := http.Get(URL)
	if err != nil {
		panic(errors.New("lost network connection"))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ipInfo := make(map[string]string, 0)
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		panic(err)
	}
	if ip, ok := ipInfo[IP_KEY]; ok {
		fmt.Println("Public IP Address:", ip)
		return
	}
	panic(errors.New("Unable to determine the public IP address."))
}

func main() {
	find_local_ip()
	find_ip()
}
