package grid

const Size = 9

type Grid [Size]int

func (g Grid) Depth() int {
	depth := -1
	for i, v := range g {
		if v > 0 {
			depth = max(depth, i)
		}
	}
	return depth + 1
}

func (g Grid) Contains(n int) bool {
	for _, v := range g {
		if v == n {
			return true
		}
	}
	return false
}

func (g Grid) Row(num int) []int {
	if num < 0 || num > 2 {
		return nil
	}

	return g[(3 * num) : (3*num)+3]
}

func (g Grid) Column(num int) []int {
	if num < 0 || num > 2 {
		return nil
	}
	return []int{g[num], g[num+3], g[num+6]}
}
