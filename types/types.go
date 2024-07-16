package types

import "sync"

type User struct {
	Name   string
	Gender string
	Email  string
}

type PubSub struct {
	Subscribers []chan User
	Mu          sync.Mutex
}
