package ordenacao

type Person struct {
	Name string
	Age  int
}

type ByName []Person
type ByAge []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps ByAge) Len() int {
	return len(ps)
}

func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
