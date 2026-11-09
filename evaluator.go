// evaluator.go
package main

import "fmt"

// Expr is a tiny AST node for the Peace expression language.
type Expr interface{}

type NumberExpr struct {
	Value int
}

type VariableExpr struct {
	Name string
}

type AddExpr struct {
	Left  Expr
	Right Expr
}

type SubExpr struct {
	Left  Expr
	Right Expr
}

type MulExpr struct {
	Left  Expr
	Right Expr
}

type DivExpr struct {
	Left  Expr
	Right Expr
}

type SetExpr struct {
	Name  string
	Value Expr
}

type Env map[string]int

func Eval(env Env, expr Expr) (int, error) {
	switch e := expr.(type) {
	case NumberExpr:
		return e.Value, nil
	case *NumberExpr:
		if e == nil {
			return 0, fmt.Errorf("nil number expression")
		}
		return e.Value, nil
	case VariableExpr:
		value, ok := env[e.Name]
		if !ok {
			return 0, fmt.Errorf("undefined variable: %s", e.Name)
		}
		return value, nil
	case *VariableExpr:
		if e == nil {
			return 0, fmt.Errorf("nil variable expression")
		}
		value, ok := env[e.Name]
		if !ok {
			return 0, fmt.Errorf("undefined variable: %s", e.Name)
		}
		return value, nil
	case AddExpr:
		left, err := Eval(env, e.Left)
		if err != nil {
			return 0, err
		}
		right, err := Eval(env, e.Right)
		if err != nil {
			return 0, err
		}
		return left + right, nil
	case *AddExpr:
		if e == nil {
			return 0, fmt.Errorf("nil add expression")
		}
		return Eval(env, AddExpr{Left: e.Left, Right: e.Right})
	case SubExpr:
		left, err := Eval(env, e.Left)
		if err != nil {
			return 0, err
		}
		right, err := Eval(env, e.Right)
		if err != nil {
			return 0, err
		}
		return left - right, nil
	case *SubExpr:
		if e == nil {
			return 0, fmt.Errorf("nil sub expression")
		}
		return Eval(env, SubExpr{Left: e.Left, Right: e.Right})
	case MulExpr:
		left, err := Eval(env, e.Left)
		if err != nil {
			return 0, err
		}
		right, err := Eval(env, e.Right)
		if err != nil {
			return 0, err
		}
		return left * right, nil
	case *MulExpr:
		if e == nil {
			return 0, fmt.Errorf("nil mul expression")
		}
		return Eval(env, MulExpr{Left: e.Left, Right: e.Right})
	case DivExpr:
		left, err := Eval(env, e.Left)
		if err != nil {
			return 0, err
		}
		right, err := Eval(env, e.Right)
		if err != nil {
			return 0, err
		}
		if right == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return left / right, nil
	case *DivExpr:
		if e == nil {
			return 0, fmt.Errorf("nil div expression")
		}
		return Eval(env, DivExpr{Left: e.Left, Right: e.Right})
	case SetExpr:
		value, err := Eval(env, e.Value)
		if err != nil {
			return 0, err
		}
		env[e.Name] = value
		return value, nil
	case *SetExpr:
		if e == nil {
			return 0, fmt.Errorf("nil set expression")
		}
		return Eval(env, SetExpr{Name: e.Name, Value: e.Value})
	default:
		return 0, fmt.Errorf("unknown expression type: %T", expr)
	}
}
