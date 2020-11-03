package shared

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var Instance string
var Mu sync.Mutex

func GetInstance() string {
	Mu.Lock()
	if Instance == "" {
		fmt.Println("Set instance")
		Instance = uuid.New().String()
	} else {
		fmt.Println("Instance is already set")
	}
	Mu.Unlock()

	return Instance
}
