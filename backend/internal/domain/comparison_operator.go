package domain

type ComparisonOperator struct {
	ID      int
	CompOpe string
}

func NewComparisonOperatorWithID(id int, compOpe string) *ComparisonOperator {
	return &ComparisonOperator{
		ID:      id,
		CompOpe: compOpe,
	}
}

func NewComparisonOperator(id int, compOpe string) *ComparisonOperator {
	return &ComparisonOperator{
		CompOpe: compOpe,
	}
}
