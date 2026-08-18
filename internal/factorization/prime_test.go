package factorization

import (
	"slices"
	"testing"
)

// returing true means the number is odd, false mean the number is even
func TestIsPrime(t *testing.T) {
	resultOdd := IsPrime(120)
	if resultOdd == true {
		t.Fatalf("got true, want false")
	}
	resultEven := IsPrime(53)
	if resultEven == false {
		t.Fatalf("got false, want true")
	}

}

func TestSieve(t *testing.T) {
	got := Sieve(10)
	want := []int{2, 3, 5, 7}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
