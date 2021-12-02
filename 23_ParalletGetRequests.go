package main

import (
	
	"net/http"
	"time"
	"log"

)

type Result struct {
	url string
	err error
    latency time.Duration
}

func getData(url string , ch chan<- Result) {
	startTime := time.Now()

	if resp, err := http.Get(url) ; err != nil {
		ch <- Result{ url, err, 0}
	} else {
		duration := time.Since(startTime).Round(time.Millisecond)
		ch<-Result{url, nil, duration}
		resp.Body.Close()
	}

}


func main(){

	ch := make(chan Result)
	urls := [] string {
		"https://yahoo.com",
		"https://amarujala.com",
		
		"https://jagran.com",

		"https://timesofindia.com",

	}

	for _,url := range urls {
		go getData(url, ch)

	}

	for range urls {

		rs := <-ch

		if rs.err != nil {
			log.Printf("URL => %30s Time %s   Err!!!", rs.url, rs.latency)	

		} else {
			log.Printf("URL => %-30s Time %s ", rs.url, rs.latency)	
		}


	}



}