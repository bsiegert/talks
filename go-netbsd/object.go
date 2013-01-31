package foo

type Number int

func (n Number) Square() Number {
	return n * n
}
