package task

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"testing"
	"time"
)

func GetNumber() (int, error) {
	time.Sleep(time.Millisecond * 1000)
	return rand.Int(), nil
}

func TestBuilder_ForkJoin(t *testing.T) {
	var future1, future2 *Task[int]

	tb := &Builder{
		MaxWorkers: 2,
	}

	start := time.Now()
	tb.ForkJoin(func(c *Awaitable) {
		future1 = Await[int](c, GetNumber)
		future2 = Await[int](c, GetNumber)
	})

	assert.NotNil(t, future1.Result)
	assert.NoError(t, future1.Err)
	log.Println(future1.Result)

	assert.NotNil(t, future2.Result)
	assert.NoError(t, future2.Err)
	log.Println(future2.Result)

	end := time.Since(start)
	log.Println(end)

	assert.Greater(t, time.Millisecond*(1000*1.01), end)
}
