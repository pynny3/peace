package main

import "testing"

func TestEvalNumberExpr(t *testing.T) {
	got, err := Eval(Env{}, NumberExpr{Value: 7})
	if err != nil {
		t.Fatalf("Eval(NumberExpr) returned error: %v", err)
	}
	if got != 7 {
		t.Fatalf("Eval(NumberExpr) = %d, want 7", got)
	}
}

func TestEvalVariableExpr(t *testing.T) {
	env := Env{"x": 10}
	got, err := Eval(env, VariableExpr{Name: "x"})
	if err != nil {
		t.Fatalf("Eval(VariableExpr) returned error: %v", err)
	}
	if got != 10 {
		t.Fatalf("Eval(VariableExpr) = %d, want 10", got)
	}
}

func TestEvalAddExpr(t *testing.T) {
	env := Env{"x": 2}
	got, err := Eval(env, AddExpr{
		Left:  VariableExpr{Name: "x"},
		Right: NumberExpr{Value: 3},
	})
	if err != nil {
		t.Fatalf("Eval(AddExpr) returned error: %v", err)
	}
	if got != 5 {
		t.Fatalf("Eval(AddExpr) = %d, want 5", got)
	}
}

func TestEvalNestedAddExpr(t *testing.T) {
	got, err := Eval(Env{}, AddExpr{
		Left: NumberExpr{Value: 2},
		Right: AddExpr{
			Left:  NumberExpr{Value: 3},
			Right: NumberExpr{Value: 4},
		},
	})
	if err != nil {
		t.Fatalf("Eval(nested AddExpr) returned error: %v", err)
	}
	if got != 9 {
		t.Fatalf("Eval(nested AddExpr) = %d, want 9", got)
	}
}

func TestEvalSetExpr(t *testing.T) {
	env := Env{}
	got, err := Eval(env, SetExpr{Name: "x", Value: NumberExpr{Value: 4}})
	if err != nil {
		t.Fatalf("Eval(SetExpr) returned error: %v", err)
	}
	if got != 4 {
		t.Fatalf("Eval(SetExpr) = %d, want 4", got)
	}
	if env["x"] != 4 {
		t.Fatalf("env[x] = %d, want 4", env["x"])
	}
}

func TestEvalArithmeticExpressions(t *testing.T) {
	env := Env{"x": 20}
	tests := []struct {
		name string
		expr Expr
		want int
	}{
		{name: "sub", expr: SubExpr{Left: VariableExpr{Name: "x"}, Right: NumberExpr{Value: 3}}, want: 17},
		{name: "mul", expr: MulExpr{Left: VariableExpr{Name: "x"}, Right: NumberExpr{Value: 2}}, want: 40},
		{name: "div", expr: DivExpr{Left: VariableExpr{Name: "x"}, Right: NumberExpr{Value: 4}}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Eval(env, tt.expr)
			if err != nil {
				t.Fatalf("Eval(%s) returned error: %v", tt.name, err)
			}
			if got != tt.want {
				t.Fatalf("Eval(%s) = %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}

func TestEvalUndefinedVariable(t *testing.T) {
	_, err := Eval(Env{}, VariableExpr{Name: "missing"})
	if err == nil {
		t.Fatal("Eval(VariableExpr) error = nil, want undefined-variable error")
	}
}
