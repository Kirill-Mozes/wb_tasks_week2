package main

// element
type shape interface {
	getType() string
	accept(visitor)
}
