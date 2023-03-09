package main

type State interface {
	setState(id int)
	setName(name string)
}
