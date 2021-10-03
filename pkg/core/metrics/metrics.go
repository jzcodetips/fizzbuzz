package metrics

import (
	"fizzbuzz/pkg/entity"
	"sync"

	"github.com/pkg/errors"
)

var (
	mtx  sync.Mutex
	data map[string]Collected = make(map[string]Collected)

	// errors
	ErrHasNotBeenCollected = errors.New("has not been collected")
)

// Collected holds collected data.
type Collected struct {
	Endpoint string                        `json:"endpoint"`
	Count    int                           `json:"count"`
	Params   map[entity.FizzBuzzParams]int `json:"params"`
}

// ResponseCollected struct.
type ReponseCollected struct {
	Endpoint string                `json:"endpoint"`
	Count    int                   `json:"count"`
	Params   entity.FizzBuzzParams `json:"params"`
}

// Collect for the given name.
func CollectFizzBuzz(endpoint string, params entity.FizzBuzzParams) {
	mtx.Lock()
	defer mtx.Unlock()

	collected, ok := data[endpoint]
	if !ok {
		collected = Collected{Endpoint: endpoint, Params: make(map[entity.FizzBuzzParams]int)}
	}

	collected.Count++

	if _, ok = collected.Params[params]; !ok {
		collected.Params[params] = 1
	} else {
		collected.Params[params]++
	}

	data[endpoint] = collected
}

// GetCollectedByName returns ReponseCollected for the given name.
func GetCollectedByName(endpoint string) (ReponseCollected, error) {
	mtx.Lock()
	defer mtx.Unlock()

	resp := ReponseCollected{}

	collected, ok := data[endpoint]
	if !ok {
		return resp, errors.Wrapf(ErrHasNotBeenCollected, "endpoint %s", endpoint)
	}

	idx := 0
	mostHits := entity.FizzBuzzParams{}

	for param, count := range collected.Params {
		if count > idx {
			idx = count
			mostHits = param
		}
	}

	resp = ReponseCollected{
		Endpoint: collected.Endpoint,
		Count:    idx,
		Params:   mostHits,
	}

	return resp, nil
}
