package main

import (
	"github.com/goinaction/code/chapter7/patterns/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 1
)

type dbConnection struct {
	ID int32
}

func (dc *dbConnection) Close() error {
	log.Println("Close Connection", dc.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create:New Connection", id)
	return &dbConnection{id}, nil
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Microsecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown Program")
	p.Close()
}
