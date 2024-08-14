package main

import "time"

// TimeProvider

//To test an object or function that relies on the current time in Go,
//you can use dependency injection to provide a mockable time source.
//Instead of calling time.Now() directly in your code, you can use an
//interface that abstracts the time retrieval. This way, you can inject
//a mock implementation during testing.

// TimeProvider
// Create an interface that represents the time retrieval
type TimeProvider interface {
	Now() time.Time
}

// RealTimeProvider
// Create a concrete implementation of this interface that uses the actual time.Now()
type RealTimeProvider struct{}

func (rtp *RealTimeProvider) Now() time.Time {
	return time.Now()
}
