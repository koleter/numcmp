package numcmp

type float struct {
	i int
	decimal string
}

func (n1 *float) compare(n2 value) int {
	switch n := n2.(type) {
	case *float:
		return n1.compareFloat(n)
	case *integer:
		return n1.compareInteger(n)
	default:
		panic("unhandle value type")
	}
}

func (n1 *float) compareFloat(n2 *float) int {
	if n1.i > n2.i || n1.i == n2.i && n1.decimal > n2.decimal {
		return 1
	} else if n1.i == n2.i && n1.decimal == n2.decimal {
		return 0
	}
	return -1
}

func (n1 *float) compareInteger(n2 *integer) int {
	if n1.i > n2.i || n1.i == n2.i && n1.decimal != "" {
		return 1
	} else if n1.i == n2.i && n1.decimal == "" {
		return 0
	}
	return -1
}
