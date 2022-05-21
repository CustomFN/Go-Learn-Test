package cron

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"testing"
	"time"
)

func TestGoCron(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(3).Second().Do(func() {
		fmt.Println(time.Now())
	})
	s.RunAll()
}
