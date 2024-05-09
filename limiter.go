/**
 * Package main
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/9
 */

package main

import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate   Rate
	mu     sync.Mutex
	last   time.Time
	counts int
}

func NewLimiter(rate Rate) RateLimiter {
	return RateLimiter{
		rate: rate,
	}
}

func (rl *RateLimiter) Allow() bool {
	if rl.rate.Number == 0 {
		return true
	}

	now := time.Now()
	switch rl.rate.Unit {
	case RateUnitSecond:
		return rl.allow(now, time.Second)
	case RateUnitMinute:
		return rl.allow(now, time.Minute)
	case RateUnitHour:
		return rl.allow(now, time.Hour)
	}
	return false
}

func (rl *RateLimiter) allow(now time.Time, period time.Duration) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if now.Sub(rl.last) >= period {
		rl.resetCounts(now)
	}

	rl.counts++

	return rl.counts <= rl.rate.Number
}

func (rl *RateLimiter) resetCounts(t time.Time) {
	rl.counts = 0
	rl.last = t
}
