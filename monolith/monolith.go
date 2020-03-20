package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type myDate struct {
	Date string
	Id int
}

type myTime struct {
	Time string
	Id int
}

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("i am big fat server, will call micro services")
	client := &http.Client{}
	go callDateMicroService(client, &wg)
	go callTimeMicroService(client, &wg)
	wg.Wait()
	fmt.Print("done\n")

}

func callMicroService(client *http.Client, wg *sync.WaitGroup, serviceUrl, serviceName string, target interface{}) {
	defer func () {
		fmt.Printf("call to %s Micro Service is done\n", serviceName)
		wg.Done()
	}()
	if microServiceResponse , err := infrastructure(client, serviceUrl); err != nil {
		fmt.Printf("%s infrastructure error: %v\n",serviceName, err)
	} else {
		fmt.Printf("%s raw response: %s\n",serviceName, string(microServiceResponse))
		if err := json.Unmarshal(microServiceResponse, target); err != nil {
			fmt.Printf("%s unmarshal error: %v\n",serviceName, err)
		} else {
			fmt.Printf("%s response: %+v\n",serviceName, target)
		}

	}
}

func callDateMicroService(client *http.Client, wg *sync.WaitGroup) {
	target := myDate{}
	callMicroService(client, wg, "http://127.0.0.1:8014/getDate", "Date", &target)
}


func callTimeMicroService(client *http.Client, wg *sync.WaitGroup) {
	target := myTime{}
	callMicroService(client, wg, "http://127.0.0.1:8015/getTime", "Time", &target)
}

// create http request
// send http request
// convert byte[] response to string
func infrastructure(client *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client req error: " + err.Error())
	} else {
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("body read error: " + err.Error())
		} else {
			return responseBody, err
		}
	}
}

