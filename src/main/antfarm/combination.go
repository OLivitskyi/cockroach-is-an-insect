package antfarm

import "strings"

type Combination struct {
	paths [][]*GraphVertex
}

func (c *Combination) String() string {
	if len(c.paths) == 0 {
		return ""
	}

	str := "{"
	for _, path := range c.paths {
		for _, v := range path {
			str += v.Name + " "
		}
		str = strings.TrimSpace(str)
		str += "}"
		str += "\n"
	}
	return str
}
