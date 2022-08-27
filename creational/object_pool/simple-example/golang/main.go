package main

import (
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"
)

const getResMaxTime = 3 * time.Second

var (
	ErrPoolNotExist  = errors.New("pool not exist")
	ErrGetResTimeout = errors.New("get resource time out")
)

// Resource
type Resource struct {
	resId int
}

// NewResource Simulate slow resource initialization creation
// (e.g., TCP connection, SSL symmetric key acquisition, auth authentication are time-consuming)
func NewResource(id int) *Resource {
	time.Sleep(500 * time.Millisecond)
	return &Resource{resId: id}
}

// Do Simulation resources are time consuming and random consumption is 0~400ms
func (r *Resource) Do(workId int) {
	time.Sleep(time.Duration(rand.Intn(5)) * 100 * time.Millisecond)
	log.Printf("using resource #%d finished work %d finish\n", r.resId, workId)
}

// Pool based on Go channel implementation, to avoid resource race state problem
type Pool chan *Resource

// New a resource pool of the specified size
// Resources are created concurrently to save resource initialization time
func New(size int) Pool {
	p := make(Pool, size)
	wg := new(sync.WaitGroup)
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(resId int) {
			p <- NewResource(resId)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return p
}

// GetResource based on channel, resource race state is avoided and resource acquisition timeout is set for empty pool
func (p Pool) GetResource() (r *Resource, err error) {
	select {
	case r := <-p:
		return r, nil
	case <-time.After(getResMaxTime):
		return nil, ErrGetResTimeout
	}
}

// GiveBackResource returns resources to the resource pool
func (p Pool) GiveBackResource(r *Resource) error {
	if p == nil {
		return ErrPoolNotExist
	}
	p <- r
	return nil
}

func main() {
	// Initialize a pool of five resources,
	// which can be adjusted to 1 or 10 to see the difference
	size := 5
	p := New(size)

	// Invokes a resource to do the id job
	doWork := func(workId int, wg *sync.WaitGroup) {
		defer wg.Done()
		// Get the resource from the resource pool
		res, err := p.GetResource()
		if err != nil {
			log.Println(err)
			return
		}
		// Resources to return
		defer p.GiveBackResource(res)
		// Use resources to handle work
		res.Do(workId)
	}

	// Simulate 100 concurrent processes to get resources from the asset pool
	num := 100
	wg := new(sync.WaitGroup)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go doWork(i, wg)
	}
	wg.Wait()
}
