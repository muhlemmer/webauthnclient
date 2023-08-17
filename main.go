/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/muhlemmer/webauthnclient/client"

	"github.com/eiannone/keyboard"
)

func main() {
	c := client.NewClient("zitadel", "localhost", "http://localhost:8080")
	fmt.Println("Prepare attestation.json, press key to continue")
	keyboard.GetSingleKey()
	attestation, err := os.ReadFile("attestation.json")
	if err != nil {
		panic(err)
	}
	resp, err := c.CreateAttestationResponse(string(attestation))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	fmt.Println("Prepare assertion.json, press key to continue")
	keyboard.GetSingleKey()
	assertion, err := os.ReadFile("assertion.json")
	if err != nil {
		panic(err)
	}
	resp, err = c.CreateAssertionResponse(string(assertion))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
