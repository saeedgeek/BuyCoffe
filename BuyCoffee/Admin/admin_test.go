package Admin

import (
	"testing"
	"time"
)

func TestReporter(t *testing.T) {
	tracker := NewTracker(true, time.Second*10)
	tracker.Users.Add(10)
	tracker.Payers.Add(10)
	tracker.Requests.Add(10)
	tracker.Views.Add(10)
}
