package decorator

import "testing"

func TestExampleDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)
	res := c.Calc()

	t.Logf("res %d\n", res)
	// Output:
	// res 80

	var c2 Component = &ConcreteComponent{}
	c2 = WarpMulDecorator(c2, 8)
	c2 = WarpAddDecorator(c2, 10)
	res2 := c2.Calc()
	t.Logf("res2 %d\n", res2)

	// Output:
	// res2 10
}
