package task_test

import (
	task "github.com/arielsrv/taskpool"
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
	var task1, task2 *task.Task[int]

	tb := &task.Builder{
		MaxWorkers: 2,
	}

	start := time.Now()
	tb.ForkJoin(func(c *task.Awaitable) {
		task1 = task.Await[int](c, GetNumber)
		task2 = task.Await[int](c, GetNumber)
	})

	assert.NotNil(t, task1.Result)
	assert.NoError(t, task1.Err)
	log.Println(task1.Result)

	assert.NotNil(t, task2.Result)
	assert.NoError(t, task2.Err)
	log.Println(task2.Result)

	end := time.Since(start)
	log.Println(end)

	assert.Greater(t, time.Millisecond*(1000*1.01), end)
}
