package main

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

// name the 
type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

// type assertions
// c is a new circle cast from s
// ok is a bool and is true if s was a circle or not
c, ok := s.(circle)

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}
	s, ok := e.(sms)
	if ok {
		return s.toPhoneNumber, sms.cost()
	}
	return "", 0.0
}

// go type switch
// use if you're checking for multiple possible types 
func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}