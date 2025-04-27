package inventory


type Inventory struct {
	Shape [][]bool
	Items  [][]Item
}

func NewInventory(shape [][]bool) *Inventory {
	return &Inventory{
		Shape: shape,
	}
}

func (i *Inventory) AddItem(item Item, origin [2]int) error {
	return nil
}

type Item struct {
	ID string
	Name  string
	Shape [][]bool
}