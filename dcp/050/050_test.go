package dcp050

import (
	"strconv"
	"testing"
)

func TestEval(t *testing.T) {
	var tests = []struct {
		expr   expr
		result int
	}{
		{constExpr(1234), 1234},
		{binaryExpr{constExpr(42), constExpr(7), opDiv}, 6},
		{
			binaryExpr{
				binaryExpr{constExpr(3), constExpr(2), opAdd},
				binaryExpr{constExpr(4), constExpr(5), opAdd},
				opMul,
			},
			45,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := tt.expr.eval()
			if result != tt.result {
				t.Fatalf("eval returned %d, want %d", result, tt.result)
			}
		})
	}
}
