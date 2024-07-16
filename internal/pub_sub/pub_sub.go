package pubsub

import (
	types "my-pubsub-app/utils"
	"sync"
)

type PubSub struct {
	subscribers []chan types.User
	mu          sync.Mutex
}

func (ps *PubSub) Publish(user types.User) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, subscriber := range ps.subscribers {
		subscriber <- user
	}
}

func (ps *PubSub) Subscribe() <-chan types.User {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan types.User, 1)
	ps.subscribers = append(ps.subscribers, ch)
	return ch
}

var pubSub = &PubSub{}

func Publish(user types.User) {
	pubSub.Publish(user)
}

func Subscribe() <-chan types.User {
	return pubSub.Subscribe()
}
