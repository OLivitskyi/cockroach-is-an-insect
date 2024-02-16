package antfarm

type Combination struct {
	paths [][]*GraphVertex
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
