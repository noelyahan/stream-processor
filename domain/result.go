package domain

// ResultMat represents the Materialized state domain for output
type ResultMat struct {
	Id    string
	Name  string
	Count int
}

// ByCount is the basic sorter interface implementation
type ByCount []ResultMat

func (p ByCount) Len() int      { return len(p) }
func (p ByCount) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p ByCount) Less(i, j int) bool {
	return p[i].Count > p[j].Count
}