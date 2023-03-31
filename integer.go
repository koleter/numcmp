package numcmp

type integer struct {
	i string
}

func (n1 *integer) compare(n2 value) int {

	switch n := n2.(type) {
	case *float:
		return n1.compareFloat(n)
	case *integer:
		return n1.compareInteger(n)
	default:
		panic("unhandle value type")
	}
}

func (n1 *integer) compareFloat(n2 *float) int {
	return -n2.compareInteger(n1)
}

func (n1 *integer) compareInteger(n2 *integer) int {
	return cmpStringOfInteger(n1.i, n2.i)
}
