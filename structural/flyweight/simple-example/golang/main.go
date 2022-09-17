package main

import "fmt"

type CheeseBrand struct {
	brand string
	cost  float64
}

var (
	// cheeseMenu is a map of cheese brands that are available in the shop. aka shared memory
	cheeseMenu = map[string]*CheeseBrand{}
)

type CheeseShop struct {
	orders map[string]int
}

func NewCheeseShop() *CheeseShop {
	return &CheeseShop{
		orders: map[string]int{},
	}
}

func (c *CheeseShop) StockCheese(brand string, cost float64) {
	cheese := &CheeseBrand{brand: brand, cost: cost}
	cheeseMenu[brand] = cheese
}

func (c *CheeseShop) SellCheese(brand string, units int) {
	c.orders[brand] += units
}

func (c *CheeseShop) TotalUnitsSold() int {
	total := 0
	for _, units := range c.orders {
		total += units
	}
	return total
}

func (c *CheeseShop) TotalIncome() float64 {
	income := 0.0
	for brand, units := range c.orders {
		income += cheeseMenu[brand].cost * float64(units)
	}
	return income
}

func main() {
	// One dedicated cheese menu
	// Each shop has its own orders based on the dedicated menu.
	shop := NewCheeseShop()
	shop2 := NewCheeseShop()
	shop.StockCheese("Gouda", 10.0)
	shop.StockCheese("Edam", 5.0)
	shop.StockCheese("Brie", 20.0)

	shop.SellCheese("Gouda", 2)
	shop.SellCheese("Edam", 3)
	shop.SellCheese("Brie", 1)

	// Shop 2 didn't stock any cheese but it's from the shared menu and allowed to sell now.
	shop2.SellCheese("Gouda", 2)
	shop2.SellCheese("Edam", 3)

	fmt.Printf("Total units sold: %d\n", shop.TotalUnitsSold())
	fmt.Printf("Total income: %.2f\n", shop.TotalIncome())
	fmt.Printf("Total units sold: %d\n", shop2.TotalUnitsSold())
	fmt.Printf("Total income: %.2f\n", shop2.TotalIncome())

}
