// internal/pubsub/simple.go
package pubsub

import (
	"sync"
	"time"
)

type Subscriber chan interface{}

type Publisher struct {
	msgs    chan interface{}
	filter  map[Subscriber]func(any) bool
	mu      sync.RWMutex
	timeout time.Duration
}

func NewPublisher(timeout time.Duration, buf int) *Publisher {
	p := &Publisher{
		msgs:    make(chan interface{}, buf),
		filter:  make(map[Subscriber]func(any) bool),
		timeout: timeout,
	}
	go p.run()
	return p
}

func (p *Publisher) run() {
	for msg := range p.msgs {
		p.mu.RLock()
		for sub, filter := range p.filter {
			if filter(msg) {
				select {
				case sub <- msg:
				case <-time.After(p.timeout):
				}
			}
		}
		p.mu.RUnlock()
	}
}

func (p *Publisher) Subscribe(filter func(interface{}) bool) Subscriber {
	ch := make(Subscriber, 10)
	p.mu.Lock()
	p.filter[ch] = filter
	p.mu.Unlock()
	return ch
}

func (p *Publisher) Publish(msg interface{}) {
	p.msgs <- msg
}

func (p *Publisher) Close() {
	close(p.msgs)
}
