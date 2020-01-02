// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package scheduler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type e struct {
	count int
	at    time.Time
}

type tickRecorder struct {
	scheduler Scheduler
	count     int
	done      chan struct{}
	recorder  chan e
}

func (m *tickRecorder) Start() {
	for {
		select {
		case t := <-m.scheduler.WaitTick():
			m.count = m.count + 1
			m.recorder <- e{count: m.count, at: t}
		case <-m.done:
			return
		}
	}
}

func (m *tickRecorder) Stop() {
	close(m.done)
}

func TestScheduler(t *testing.T) {
	t.Run("Step scheduler", testStepScheduler)
}

func newTickRecorder(scheduler Scheduler) *tickRecorder {
	return &tickRecorder{
		scheduler: scheduler,
		done:      make(chan struct{}),
		recorder:  make(chan e),
	}
}

func testStepScheduler(t *testing.T) {
	t.Run("Trigger the Tick manually", func(t *testing.T) {
		scheduler := NewStepper()
		defer scheduler.Stop()

		recorder := newTickRecorder(scheduler)
		go recorder.Start()
		defer recorder.Stop()

		scheduler.Next()
		nE := <-recorder.recorder
		require.Equal(t, 1, nE.count)
		scheduler.Next()
		nE = <-recorder.recorder
		require.Equal(t, 2, nE.count)
		scheduler.Next()
		nE = <-recorder.recorder
		require.Equal(t, 3, nE.count)
	})
}

func testPeriodic(t *testing.T) {
	t.Run("tick than wait", func(t *testing.T) {
		duration := 1 * time.Minute
		scheduler := NewPeriodic(duration)
		defer scheduler.Stop()

		startedAt := time.Now()
		recorder := newTickRecorder(scheduler)
		go recorder.Start()
		defer recorder.Stop()

		nE := <-recorder.recorder

		require.True(t, nE.at.Sub(startedAt) < duration)
	})

	t.Run("multiple tick", func(t *testing.T) {
		duration := 1 * time.Millisecond
		scheduler := NewPeriodic(duration)
		defer scheduler.Stop()

		recorder := newTickRecorder(scheduler)
		go recorder.Start()
		defer recorder.Stop()

		nE := <-recorder.recorder
		require.Equal(t, 1, nE.count)
		nE = <-recorder.recorder
		require.Equal(t, 2, nE.count)
		nE = <-recorder.recorder
		require.Equal(t, 3, nE.count)
	})
}