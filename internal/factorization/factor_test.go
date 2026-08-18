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
func TestFactorizationTrialDivisionSieve(t *testing.T) {
	ch := make(chan []int, 1)
	ctx := context.Background()
	FactorTrialDivisionSieve(12014123, ch, ctx)
	gotResult := <-ch
	wantResult := []int{11, 71, 15383}
	if !slices.Equal(gotResult, wantResult) {
		t.Fatalf("got %v, want %v", gotResult, wantResult)
	}
}

func TestFactorTrialDivisionSIMD(t *testing.T) {
	ch := make(chan []int, 1)
	ctx := context.Background()
	FactorTrialDivisionSIMD(12014123, ch, ctx)
	gotResult := <-ch
	wantResult := []int{11, 71, 15383}
	if !slices.Equal(gotResult, wantResult) {
		t.Fatalf("got %v, want %v", gotResult, wantResult)
	}
}

func TestFactorTrialDivisionSieveSIMD(t *testing.T) {
	ch := make(chan []int, 1)
	ctx := context.Background()
	FactorTrialDivisionSieveSIMD(12014123, ch, ctx)
	gotResult := <-ch
	wantResult := []int{11, 71, 15383}
	if !slices.Equal(gotResult, wantResult) {
		t.Fatalf("got %v, want %v", gotResult, wantResult)
	}
}

func TestFactorSqrt(t *testing.T) {
	ch := make(chan []int, 1)
	ctx := context.Background()
	FactorSqrt(12014123, ch, ctx)
	gotResult := <-ch
	wantResult := []int{11, 71, 15383}
	if !slices.Equal(gotResult, wantResult) {
		t.Fatalf("got %v, want %v", gotResult, wantResult)
	}
}

func TestFactorFermat(t *testing.T) {
	ctx := context.Background()
	gotResult := FactorFermat(12014123, ctx)
	wantResult := []int{11, 71, 15383}
	if !slices.Equal(gotResult, wantResult) {
		t.Fatalf("got %v, want %v", gotResult, wantResult)
	}
}
