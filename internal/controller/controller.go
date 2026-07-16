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
	BenchmarkOn      bool
	SelectedMethod   string
	CurrentInputText string
	LatestFactors    []int
	mu               sync.Mutex
	Running          bool
	Done             chan []int
	ctx              context.Context
	cancel           context.CancelFunc
	Duration         time.Duration
	FactoredNumber   int
	SieveOn          bool
	Totient          int
	E                int
	D                int
	RsaDuration      time.Duration
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
		appController.Totient = rsa.Totient(appController.LatestFactors[0], appController.LatestFactors[1])
		appController.E = rsa.E(appController.Totient)
		appController.D = rsa.D(appController.E, appController.Totient)
		appController.RsaDuration = time.Since(start)
	}()
	return appController.Totient, appController.E, appController.D
}
