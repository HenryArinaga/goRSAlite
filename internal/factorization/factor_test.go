package factorization

import (
	"context"
	"slices"
	"testing"
)

func TestFactorizationTrialDivision(t *testing.T) {
	ch := make(chan []int, 1)
	ctx := context.Background()
	FactorTrialDivision(77, ch, ctx)
	gotResult := <-ch
	wantResult := []int{7, 11}
	if !slices.Equal(gotResult, wantResult) {
		t.Fatalf("got %v, want %v", gotResult, wantResult)
	}
}
