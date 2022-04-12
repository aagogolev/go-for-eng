package main

import (
	"fmt"
)

type Avatar struct {
	URL  string
	Size int64
}

type Client struct {
	ID   int64
	Name string
	Age  int
	IMG  *Avatar
}

func (c Client) HesAvatar() bool {
	if c.IMG != nil && c.IMG.URL != "" {
		return true
	}

	return false
}

func main() {
	client := &Client{}
	fmt.Printf("%#v\n", client)

	// updateAvater(client)
	// fmt.Printf("%#v\n", client)

	client2 := Client{
		ID:   7,
		Name: "admin",
		Age:  20,
	}

	fmt.Println(client2.HesAvatar())

}

func updateAvater(client *Client) {
	client.IMG.URL = "update_url"
}
