package controller

import (
	"context"
	"fmt"
	"goRSAlite/internal/factorization"
	"goRSAlite/internal/rsa"
	"sync"
	"time"
)

type RSAResult struct {
	Totient         int
	PublicExponent  int
	PrivateExponent int
}

type LogEntry struct {
	Time     time.Time
	Kind     string
	Input    string
	Result   string
	Method   string
	Duration time.Duration
}

type Controller struct {
	BenchmarkOn        bool
	SelectedMethod     string
	CurrentInputText   string
	CurrentMessageText string
	LatestFactors      []int
	mu                 sync.Mutex
	Running            bool
	Done               chan []int
	EncryptionDone     chan []int
	DecryptionDone     chan []int
	RsaDone            chan RSAResult
	ctx                context.Context
	cancel             context.CancelFunc
	Duration           time.Duration
	FactoredNumber     int
	SieveOn            bool
	Totient            int
	PublicExponent     int
	PrivateExponent    int
	RsaDuration        time.Duration
	RsaRunning         bool
	Encryption         int
	cipherText         []int
	input              string
	EncryptionDuration time.Duration
	DecryptionDuration time.Duration
	DecryptedMessage   []int
	SimdOn             bool
	ExtendedGcdOn      bool
	History            []LogEntry
}

func (CancelFunction *Controller) CancelRunningFunction() {
	if CancelFunction.cancel == nil {
		return
	}
	CancelFunction.cancel()
}

func (appController *Controller) DoFactorization(numberToFactor int) {
	appController.ctx, appController.cancel = context.WithCancel(context.Background())
	factorResult := make(chan []int, 1)
	appController.FactoredNumber = numberToFactor
	switch appController.SelectedMethod {
	case "Trial Division":
		appController.Running = true
		start := time.Now()
		go func() {
			if appController.SimdOn && appController.SieveOn == true {
				factorization.FactorTrialDivisionSieveSIMD(numberToFactor, factorResult, appController.ctx)
				factors := <-factorResult
				appController.mu.Lock()
				appController.LatestFactors = factors
				appController.mu.Unlock()
				appController.Duration = time.Since(start)
				appController.Running = false
			} else if appController.SieveOn == true {
				factorization.FactorTrialDivisionSieve(numberToFactor, factorResult, appController.ctx)
				factors := <-factorResult
				appController.mu.Lock()
				appController.LatestFactors = factors
				appController.mu.Unlock()
				appController.Duration = time.Since(start)
				appController.Running = false
			} else if appController.SimdOn == true {
				factorization.FactorTrialDivisionSieveSIMD(numberToFactor, factorResult, appController.ctx)
				factors := <-factorResult
				appController.mu.Lock()
				appController.LatestFactors = factors
				appController.mu.Unlock()
				appController.Duration = time.Since(start)
				appController.Running = false
			} else {
				factorization.FactorTrialDivision(numberToFactor, factorResult, appController.ctx)
				factors := <-factorResult
				appController.mu.Lock()
				appController.LatestFactors = factors
				appController.mu.Unlock()
				appController.Duration = time.Since(start)
				appController.Running = false
			}
			appController.Done <- appController.LatestFactors

		}()
	case "Square Root":
		appController.Running = true
		start := time.Now()
		go func() {
			factorization.FactorSqrt(numberToFactor, factorResult, appController.ctx)
			factors := <-factorResult
			appController.mu.Lock()
			appController.LatestFactors = factors
			appController.mu.Unlock()
			appController.Duration = time.Since(start)
			appController.Running = false
			appController.Done <- appController.LatestFactors
		}()
	case "Fermat Factorization":
		appController.ctx, appController.cancel = context.WithCancel(context.Background())
		appController.Running = true
		start := time.Now()
		go func() {
			factors := factorization.FactorFermat(numberToFactor, appController.ctx)
			factorResult <- factors
			appController.mu.Lock()
			appController.LatestFactors = factors
			appController.mu.Unlock()
			appController.Duration = time.Since(start)
			appController.Running = false
			appController.Done <- appController.LatestFactors
		}()
	default:
		fmt.Println("Invalid method selected")
	}

}

func (appController *Controller) DoRsa() {
	start := time.Now()
	go func() {
		appController.RsaRunning = true
		t1 := time.Now()
		totient := rsa.Totient(appController.LatestFactors[0], appController.LatestFactors[1])
		fmt.Println("Totient took:", time.Since(t1), "value:", totient)
		t2 := time.Now()
		publicExponent := rsa.E(totient, appController.ExtendedGcdOn)
		fmt.Println("E took:", time.Since(t2), "value:", publicExponent, "extended:", appController.ExtendedGcdOn)
		t3 := time.Now()
		privateExponent := rsa.D(publicExponent, totient, appController.ExtendedGcdOn)
		fmt.Println("D took:", time.Since(t3), "value:", privateExponent)
		appController.mu.Lock()
		appController.Totient = totient
		appController.PublicExponent = publicExponent
		appController.PrivateExponent = privateExponent
		appController.RsaDuration = time.Since(start)
		fmt.Println("Total RSA time:", time.Since(start))
		appController.RsaRunning = false
		appController.mu.Unlock()
		appController.RsaDone <- RSAResult{
			Totient:         totient,
			PublicExponent:  publicExponent,
			PrivateExponent: privateExponent,
		}
	}()
}

func (appController *Controller) DoEncrypt(input string) {
	start := time.Now()
	go func() {
		appController.RsaRunning = true
		n := appController.LatestFactors[0] * appController.LatestFactors[1]
		appController.cipherText = make([]int, 0, len(input))
		for i := 0; i < len(input); i++ {
			appController.cipherText = append(
				appController.cipherText,
				rsa.Encrypt(int(input[i]),
					appController.PublicExponent,
					n))
		}
		appController.EncryptionDone <- appController.cipherText
		appController.RsaRunning = false
		appController.EncryptionDuration = time.Since(start)
	}()
}

func (appController *Controller) DoDecrypt(input string) {
	start := time.Now()
	go func() {
		appController.RsaRunning = true
		n := appController.LatestFactors[0] * appController.LatestFactors[1]
		appController.DecryptedMessage = make([]int, 0, len(input))
		for i := 0; i < len(appController.cipherText); i++ {
			appController.DecryptedMessage = append(appController.DecryptedMessage,
				rsa.Decrypt(appController.cipherText[i], appController.PrivateExponent, n))
		}
		for i := 0; i < len(appController.DecryptedMessage); i++ {
			fmt.Printf("%c", appController.DecryptedMessage[i])
		}
		appController.DecryptionDone <- appController.DecryptedMessage
		appController.DecryptionDuration = time.Since(start)
		appController.RsaRunning = false

	}()
}
