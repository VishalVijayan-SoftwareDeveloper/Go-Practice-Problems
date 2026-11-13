package problems

type RateLimiter struct {
	requests map[string][]int64
	limit    int
	window   int64
}

func NewRateLimiter(limit int, window int64) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]int64),
		limit:    limit,
		window:   window,
	}
}

func (r *RateLimiter) Allow(
	deviceID string,
	timestamp int64,
) bool {
	requests := r.requests[deviceID]

	// Remove expired timestamps.
	start := 0

	for start < len(requests) &&
		requests[start] <= timestamp-r.window {
		start++
	}

	requests = requests[start:]

	if len(requests) >= r.limit {
		r.requests[deviceID] = requests
		return false
	}

	requests = append(requests, timestamp)
	r.requests[deviceID] = requests

	return true
}
