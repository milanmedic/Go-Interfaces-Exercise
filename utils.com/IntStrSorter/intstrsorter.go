package Intstrsorter

type Sorter interface {
	length() int
	less(i, j int) bool
	swap(i, j int)
}

type Xi []int
type Xs []string

func (p Xi) length() int            { return len(p) }
func (p Xi) less(i int, j int) bool { return p[j] < p[i] }
func (p Xi) swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func (p Xs) length() int            { return len(p) }
func (p Xs) less(i int, j int) bool { return p[j] < p[i] }
func (p Xs) swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func Sort(s Sorter) {
	for i := 0; i < s.length()-1; i++ {
		for j := i + 1; j < s.length(); j++ {
			if s.less(i, j) {
				s.swap(i, j)
			}
		}
	}
}
