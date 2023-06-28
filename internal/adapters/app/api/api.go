package api

import (
	"github.com/mansoorceksport/hexarch/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DbPort
}

func NewAdapter(arith ports.ArithmeticPort, db ports.DbPort) *Adapter {
	return &Adapter{
		arith: arith,
		db:    db,
	}
}

func (adapter Adapter) GetAdd(a, b int32) (int32, error) {
	answer, err := adapter.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (adapter Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := adapter.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (adapter Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := adapter.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (adapter Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := adapter.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
