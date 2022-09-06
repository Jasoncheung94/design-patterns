package main

import "fmt"

type Iphone struct {
}

func (i *Iphone) Charge() {
	fmt.Println("Charging Iphone")
}

type Android struct {
}

func (a *Android) Charge() {
	fmt.Println("Charging Android")
}

type Charger interface {
	Charge()
}

type Client struct {
}

func (c *Client) ChargingWith(charger Charger) {
	charger.Charge()
}

type IphoneAdapter struct {
	iphone *Iphone
}

func (i *IphoneAdapter) Charge() {
	fmt.Println("Using adapter to charge Iphone:")
	i.iphone.Charge()
}

func main() {
	iphone := &Iphone{}
	android := &Android{}
	client := &Client{}
	adapter := &IphoneAdapter{iphone: iphone}

	client.ChargingWith(iphone)
	client.ChargingWith(android)
	client.ChargingWith(adapter)
}
