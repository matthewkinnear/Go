package pubsub

import (
	"my-pubsub-app/types"
)

type PubSub struct {
	types.PubSub
}

func (ps *PubSub) Publish(user types.User) {
	ps.Mu.Lock()
	defer ps.Mu.Unlock()

	for _, subscriber := range ps.Subscribers {
		subscriber <- user
	}
}

func (ps *PubSub) Subscribe() <-chan types.User {
	ps.Mu.Lock()
	defer ps.Mu.Unlock()

	ch := make(chan types.User, 1)
	ps.Subscribers = append(ps.Subscribers, ch)
	return ch
}

var pubSub = &PubSub{}

func Publish(user types.User) {
	pubSub.Publish(user)
}

func Subscribe() <-chan types.User {
	return pubSub.Subscribe()
}
