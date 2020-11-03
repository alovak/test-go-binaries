package packageb

import (
	"fmt"
	"parallel/pkg/shared"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("from %s shared instance: %s\n", Package(), shared.GetInstance())
		fmt.Printf("[%s] from %s mutext: %p instance: %p\n", time.Now(), Package(), &shared.Mu, &shared.Instance)
		time.Sleep(1 * time.Second)
	}
}
