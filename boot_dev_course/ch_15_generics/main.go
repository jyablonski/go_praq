package main

type store[P product] interface {
	Sell(P)
}

type product interface {
	Price() float64
	Name() string
}
