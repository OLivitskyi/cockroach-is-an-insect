package types

type Combination struct {
	paths [][]*Vertex
}

func (c *Combination) String() string {
	str := "{"
	for _, path := range c.paths {
		for _, v := range path {
			str += v.Name + " "
		}
		str += "}\n"
	}
	return str
}
