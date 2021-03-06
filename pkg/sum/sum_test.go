package sum_test

import (
	sum "github.com/gonzalompp/go-api/pkg/sum"
	"testing"
)

func TestSum(t *testing.T) {
	total := sum.Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
