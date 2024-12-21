package main

import "errors"

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

func splitEmail(email string) (string, string) {
	username, domain := "", ""

	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}

	return username, domain
}

// closure
func adder() func(int) int {
	num_sum := 0

	return func(num int) int {
		num_sum += num
		return num_sum
	}
}

// structs
type t133 struct {
	phoneNumber int
	message     string
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	return mToSend.sender.number != 0 && mToSend.sender.name != "" && mToSend.recipient.number != 0 && mToSend.recipient.name != ""
}

type car struct {
	make  string
	model string
}

type truck struct {
	car
	towable_weight int
}

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	weight, height float64
}

func (r rect) area() float64 {
	return r.weight * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.weight + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14159 * c.radius
}

type sendable interface {
	cost() float64
}

type email struct {
	body          string
	is_subscribed bool
}

func (e email) cost() float64 {
	body_len := float64(len(e.body))
	if !e.is_subscribed {
		return 0.05 * body_len
	}
	return 0.01 * body_len
}
