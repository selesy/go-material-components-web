package generator

type Block struct {
	Class string
	Tag   string
}

type Element struct {
	Name      string
	Component Component
}

type Modifier struct {
	Name    string
	Variant bool
}

type Component struct {
	Name  string
	Block Block
	// Elements  []Element
	Modifiers []Modifier
}
