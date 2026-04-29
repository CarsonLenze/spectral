package congestion

import (
	"testing"
	"time"
)

func TestRTTInitialSampleInitialisesEstimator(t *testing.T) {
	rtt := NewRTT()

	rtt.Add(50*time.Millisecond, 0)

	if !rtt.measured {
		t.Fatal("expected estimator to be marked as measured after first sample")
	}
	if got := rtt.LatestRTT(); got != 50*time.Millisecond {
		t.Fatalf("expected latest RTT to be 50ms, got %v", got)
	}
	if got := rtt.SRTT(); got != 50*time.Millisecond {
		t.Fatalf("expected smoothed RTT to match the first sample, got %v", got)
	}
	if got := rtt.RTTVAR(); got != 25*time.Millisecond {
		t.Fatalf("expected RTT variance to be initialised from the first sample, got %v", got)
	}
}
