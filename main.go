package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func find_ip() {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			fmt.Println(err.Error())
		}
	}()

	const URL = "https://api.ipify.org?format=json"
	const KEY = "ip"

	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
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
	if ip, ok := ipInfo[KEY]; ok {
		fmt.Println("Public IP Address:", ip)
		return
	}
	panic(errors.New("Unable to determine the public IP address."))
}

func main() {
	find_ip()
}
