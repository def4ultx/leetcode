package main

// SEE: https://medium.com/@salmaan.rashid/grpc-with-curl-a65d879a18f7

import (
	"bytes"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	dat, err := ioutil.ReadFile("fname.bin")
	if err != nil {
		panic(err)
	}

	url := "https://0.0.0.0:50051/helloworld.Greeter/SayHello"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewReader(dat))
	if err != nil {
		panic(err)
	}
	req.Header.Set("TE", "trailers")
	req.Header.Set("Content-Type", "application/grpc")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("response Body:", hex.Dump(body))

}
