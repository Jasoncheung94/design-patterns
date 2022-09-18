package main

import "fmt"

type Request struct {
	method         string
	controllerDone bool
	validationDone bool
}

type Service interface {
	execute(r *Request)
	setNext(Service)
}

type RequestHandler struct {
	next Service
}

func (r *RequestHandler) execute(req *Request) {
	fmt.Println("Handling the request")
	r.next.execute(req)
}

func (r *RequestHandler) setNext(s Service) {
	r.next = s
}

type Controller struct {
	next Service
}

func (c *Controller) execute(r *Request) {
	fmt.Println("Controller executed")
	r.controllerDone = true
	c.next.execute(r)
}

func (c *Controller) setNext(s Service) {
	c.next = s
}

type Validation struct {
	next Service
}

func (v *Validation) execute(r *Request) {
	fmt.Println("Validation executed")
	r.validationDone = true
	if v.next != nil {
		v.next.execute(r)
	} else {
		fmt.Println("Request completed")
	}
}

func (v *Validation) setNext(s Service) {
	v.next = s
}

// GO111MODULE="off" go run main.go
// Request -> handler -> controller -> validation -> request completed
func main() {
	r := &Request{
		method: "GET",
	}
	handler := &RequestHandler{}
	controller := &Controller{}
	handler.setNext(controller)
	validation := &Validation{}
	controller.setNext(validation)

	handler.execute(r)
}
