package controller

import (
	"context"
	"fmt"
	"goRSAlite/internal/factorization"
	"goRSAlite/internal/rsa"
	"sync"
	"time"
)

type Controller struct {
	BenchmarkOn        bool
	SelectedMethod     string
	CurrentInputText   string
	LatestFactors      []int
	mu                 sync.Mutex
	Running            bool
	Done               chan []int
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
	ciphertext         []int
	input              string
	EncryptionDuration time.Duration
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
			if appController.SieveOn == true {
				factorization.FactorTrialDivisionSieve(numberToFactor, factorResult, appController.ctx)
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

func (appController *Controller) DoRsa(numberToFactor int) (int, int, int) {
	start := time.Now()
	go func() {
		appController.RsaRunning = true
		appController.Totient = rsa.Totient(appController.LatestFactors[0], appController.LatestFactors[1])
		appController.PublicExponent = rsa.E(appController.Totient)
		appController.PrivateExponent = rsa.D(appController.PublicExponent, appController.Totient)
		appController.RsaRunning = false
		appController.RsaDuration = time.Since(start)
	}()
	return appController.Totient, appController.PublicExponent, appController.PrivateExponent
}

func (appController *Controller) DoEncrypt() []int {
	start := time.Now()
	go func() {
		appController.RsaRunning = true
		n := appController.LatestFactors[0] * appController.LatestFactors[1]
		appController.ciphertext = make([]int, 0, len(appController.input))
		for i := 0; i < len(appController.input); i++ {
			appController.ciphertext = append(
				appController.ciphertext,
				rsa.Encrypt(int(appController.input[i]),
					appController.PublicExponent,
					n))
		}
		appController.RsaRunning = false
		appController.EncryptionDuration = time.Since(start)
	}()
	return appController.ciphertext
}
