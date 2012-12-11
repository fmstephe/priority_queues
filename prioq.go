package prioq

import ()

type Q interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
	Remove(interface{})
}
