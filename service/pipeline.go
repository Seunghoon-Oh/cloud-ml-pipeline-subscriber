package service

import (
	"encoding/json"
	"fmt"

	"github.com/Seunghoon-Oh/cloud-ml-pipeline-subscriber/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var cb *circuit.Breaker
var httpClient *circuit.HTTPClient

func SetupPipelineCircuitBreaker() {
	httpClient, cb = network.GetHttpClient()
}

func CreatePipeline() {
	if cb.Ready() {
		resp, err := httpClient.Post("http://cloud-ml-pipeline-manager.cloud-ml-pipeline:8082/pipeline", "", nil)
		if err != nil {
			fmt.Println(err)
			cb.Fail()
			return
		}
		cb.Success()
		defer resp.Body.Close()
		rsData := network.ResponseData{}
		json.NewDecoder(resp.Body).Decode(&rsData)
		fmt.Println(rsData.Data)
		return
	}
}
