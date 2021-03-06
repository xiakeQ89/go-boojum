package aggregator

import "time"

// MockAggregator ..
type MockAggregator struct {}

// MakeExample ..
func (mock *MockAggregator) MakeExample() []byte {
	return []byte{0, 1, 2, 3, 4, 5, 6}
}

// AggregateTrees ..
func (mock *MockAggregator) AggregateTrees(left, right []byte) []byte {
	time.Sleep(time.Duration(200) * time.Millisecond)
	return []byte{0, 1, 2, 3, 4, 5, 6}
}

// Verify ..
func (mock *MockAggregator) Verify(buff []byte) bool {
	return true
}