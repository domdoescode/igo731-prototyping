package inventory_test

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"

	"github.com/domdoescode/IGO731/inventory"
)

func TestAddItem(t *testing.T) {
	var tests = []struct {
		name string
		inventory [][]bool
		in inventory.Item
		err error
	}{
		{
			"Add 1 height 1 width item to 1 height 1 width inventory at position 0,0",
			[][]bool{
				{true},
			},
			inventory.Item{
				ID: "1",
				Name: "Item 1",
				Shape: [][]bool{
					{true},
				},
			},
			nil,
		},
		{
			"Add 2 height 1 width item to 1 height 1 width inventory at position 0,0",
			[][]bool{
				{true},
				{true},
			},
			inventory.Item{
				ID: "1",
				Name: "Item 1",
				Shape: [][]bool{
					{true},
				},
			},
			errors.New("item does not fit in inventory"),
		},
	}
	
    for _, tt := range tests {
		invt := inventory.NewInventory(tt.inventory)
        
		t.Run(tt.name, func(t *testing.T) {
			err := invt.AddItem(tt.in, [2]int{0, 0})
			assert.Equal(t, tt.err, err)
        })
    }
}
