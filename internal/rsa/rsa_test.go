package rsa

import (
	"slices"
	"testing"
)

func TestTotient(t *testing.T) {
	got := Totient(8, 32)
	want := 217
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestGCD(t *testing.T) {
	got := GCD(2432, 1442)
	want := 2
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestExtendedGCD(t *testing.T) {
	got, got2, got3 := ExtendedGCD(5, 217)
	want := 1
	want2 := 87
	want3 := -2
	if got != want {
		t.Fatalf("got %v,  want %v", got, want)
	}
	if got2 != want2 {
		t.Fatalf("got %v, want %v", got2, want)
	}
	if got3 != want3 {
		t.Fatalf("got %v, want %v", got3, want3)
	}
}

func TestE(t *testing.T) {
	got := E(35360412, false)
	want := 5
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	gotExtended := E(35360412, true)
	wantExtended := 5
	if gotExtended != wantExtended {
		t.Fatalf("got %v, want %v", gotExtended, wantExtended)
	}
}

func TestD(t *testing.T) {
	got := D(5, 35360412, false)
	want := 14144165
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	gotExtended := D(5, 35360412, true)
	wantExtended := 14144165
	if gotExtended != wantExtended {
		t.Fatalf("got %v, want %v", gotExtended, wantExtended)
	}
}

func TestEncryptSingleByte(t *testing.T) {
	got := Encrypt(72, 5, 35421341)
	want := 22165218
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestEncryptMessage(t *testing.T) {
	inputString := "Hello"
	cipherText := make([]int, 0, len(inputString))
	for i := 0; i < len(inputString); i++ {
		cipherText = append(
			cipherText,
			Encrypt(int(inputString[i]), 5, 35421341))
	}
	want := []int{22165218, 25383565, 28845594, 28845594, 25444576}
	if !slices.Equal(cipherText, want) {
		t.Fatalf("got %v, want %v", cipherText, want)
	}
}

func TestDecryptSingleByte(t *testing.T) {
	got := Decrypt(22165218, 14144165, 35421341)
	want := 72
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestDecryptMessage(t *testing.T) {
	cipherText := []int{22165218, 25383565, 28845594, 28845594, 25444576}
	decryptedMessage := make([]int, 0, len(cipherText))
	for i := 0; i < len(cipherText); i++ {
		decryptedMessage = append(decryptedMessage,
			Decrypt(cipherText[i], 14144165, 35421341))
	}
	want := []int{72, 101, 108, 108, 111}
	lenBytes := len(decryptedMessage)
	wantLenBytes := 5
	if !slices.Equal(decryptedMessage, want) {
		t.Fatalf("got %v, want %v", decryptedMessage, want)
	}
	if lenBytes != wantLenBytes {
		t.Fatalf("got %v, want %v", lenBytes, wantLenBytes)
	}

	decryptedText := make([]byte, 0, len(cipherText))
	for i := 0; i < len(cipherText); i++ {
		decryptedText = append(decryptedText,
			byte(decryptedMessage[i]))
	}
	gotText := string(decryptedText)
	wantText := "Hello"
	if gotText != wantText {
		t.Fatalf("got %v, want %v", gotText, wantText)
	}
}
