package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

const UUIDS_LEN = 1000
const SCHEMA_ID = "a77b283d-1823-497c-b34c-01b0f07e03cf"
const URL = "http://localhost:8428/write"

func main() {

	var base_bodies [UUIDS_LEN]string
	for i := 0; i < UUIDS_LEN; i++ {
		base_bodies[i] = SCHEMA_ID + ",secondUUID=" + uuid.New().String()
	}
	r := rand.New(rand.NewSource(99))
	var client fasthttp.Client = fasthttp.Client{MaxConnsPerHost: 1024, DisablePathNormalizing: true, DisableHeaderNamesNormalizing: true}
	var wg sync.WaitGroup
	for {
		//start := time.Now()
		wg.Add(UUIDS_LEN)
		for _, base_body := range base_bodies {
			var body = base_body + fmt.Sprintf(" value=%f", r.Float64()*300) + ",value2=2,value3=0,value4=0,value5=220,value6=220,value7=true,value8=14951,value9=20008387"
			go SendPostAsync(URL, &body, &wg, &client)
		}
		wg.Wait()
		//log.Printf("%f req/s", float64(UUIDS_LEN)/time.Since(start).Seconds())
	}
}

func SendPostAsync(url string, body *string, wg *sync.WaitGroup, client *fasthttp.Client) {
	defer wg.Done()
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURI(url)
	req.SetBodyString(*body)
	client.Do(req, resp)
}
