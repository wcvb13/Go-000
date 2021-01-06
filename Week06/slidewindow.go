package slidewindow

import (
	"sync"
	"time"
)

type Bucket struct {
	Success float64
	Failure float64
}

type Window struct {
	Buckets map[int64]*Bucket
	internal int64
	Mux *sync.RWMutex
}

func NewWindow(internal int64) *Window {
	return &Window{
		Buckets: make(map[int64]*Bucket),
		internal: internal,
		Mux: &sync.RWMutex{},
	}
}

func (w *Window) getCurrentBucket() *Bucket {
	now := time.Now().Unix()
	var bucket *Bucket
	var ok bool
	if bucket,ok = w.Buckets[now]; !ok {
		bucket = &Bucket{}
		w.Buckets[now] = bucket
	}
	return bucket
}

func (w *Window) removeOldBucket()  {
	now := time.Now().Unix() - w.internal
	for timestamp := range w.Buckets {
		if timestamp < now {
			delete(w.Buckets,timestamp)
		}
	}
}

func (w *Window)IncrementSuccess(i float64)  {
	if i == 0 {
		return
	}
	w.Mux.Lock()
	defer w.Mux.Unlock()
	b := w.getCurrentBucket()
	b.Success += i
	w.removeOldBucket()
}

func (w *Window)IncrementFailure(i float64)  {
	if i == 0 {
		return
	}
	w.Mux.Lock()
	defer w.Mux.Unlock()
	b := w.getCurrentBucket()
	b.Failure += i
	w.removeOldBucket()
}

func (w *Window) SumSuccess() float64 {
	t := time.Now().Unix() - w.internal
	sum := float64(0)
	
	w.Mux.RLock()
	defer w.Mux.RUnlock()

	for timestamp, bucket := range w.Buckets {
		if timestamp > t {
			 sum += bucket.Success
		}
	}
	return sum
}

func (w *Window) SumFailure() float64 {
	t := time.Now().Unix() - w.internal
	sum := float64(0)

	w.Mux.RLock()
	defer w.Mux.RUnlock()

	for timestamp, bucket := range w.Buckets {
		if timestamp > t {
			sum += bucket.Failure
		}
	}
	return sum
}

func (w *Window)MaxSuccess(now time.Time) float64 {
	var max float64

	w.Mux.RLock()
	defer w.Mux.RUnlock()

	for timestamp, bucket := range w.Buckets {
		if timestamp >= now.Unix() - w.internal {
			if bucket.Success > max {
				max = bucket.Success
			}
		}
	}
	return max
}

func (w *Window)MaxFailure(now time.Time) float64 {
	var max float64

	w.Mux.RLock()
	defer w.Mux.RUnlock()

	for timestamp, bucket := range w.Buckets {
		if timestamp >= now.Unix() - w.internal {
			if bucket.Failure > max {
				max = bucket.Failure
			}
		}
	}
	return max
}