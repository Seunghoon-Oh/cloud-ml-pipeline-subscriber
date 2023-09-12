package main

import (
	"log"
	"sync"

	"github.com/Seunghoon-Oh/cloud-ml-pipeline-subscriber/service"
	"github.com/nats-io/nats.go"
)

func main() { 
	log.SetFlags(0)
	service.SetupNotebookCircuitBreaker()
	for {
		nc, err := nats.Connect("nats://nats.cloud-ml-mgmt:4222")
		if err != nil {
			log.Fatal(err)
		}
		defer nc.Close()

		wg := sync.WaitGroup{}
		wg.Add(1)

		if _, err := nc.Subscribe("pipeline", func(m *nats.Msg) {
			log.Printf("Reply: %s", m.Data)
			msg := string(m.Data[:])
			if msg == "create" {
				service.CreateNotebook()
			} else if msg == "delete " {
				log.Println("Deleted")
			}
			// wg.Done()
		}); err != nil {
			log.Fatal(err)
		}
		wg.Wait()
	}
}
