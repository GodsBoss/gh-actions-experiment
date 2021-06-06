package domain_test

import (
	"github.com/GodsBoss/gh-actions-experiment/pkg/domain"

	"testing"
)

func TestConst(t *testing.T) {
	cnst := domain.Const(100)

	result, err := cnst.Calculate()

	if err != nil {
		t.Errorf("expected no error, got %+v", err)
	}

	if result != 100 {
		t.Errorf("expected result to be 100, got %d", result)
	}
}
