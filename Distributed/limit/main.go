package main

import (
	"blog-example/Distributed/limit/limit_alg"
	"fmt"
	"sync"
	"time"
)

func simulateUserRequests(quotaManager *limit_alg.QuotaManager, userID string, initialQuota int, costPerService int, isVIP bool, wg *sync.WaitGroup) {
	start := time.Now()
	defer func() {
		fmt.Printf("User %s took %v\n", userID, time.Since(start))
		wg.Done()
	}()

	for i := 1; i <= 20; i++ {
		allowed := quotaManager.ConsumeQuota(isVIP, userID, costPerService)
		if allowed {
			fmt.Printf("Request %d for user %s is allowed (Remaining Quota: %d)\n", i, userID, quotaManager.GetQuota(isVIP, userID, initialQuota).Limit)
		} else { // refill quota
			fmt.Println("Refilling quota for normal user...")
			quotaManager.RefillQuota(isVIP, userID, initialQuota)
			time.Sleep(2 * time.Second)
			fmt.Printf("Quota refilled for user %s\n", userID)
		}
	}
}

func main() {
	quotaManager := limit_alg.NewQuotaManager()

	userA := "VIP_A"
	userB := "Normal_B"
	vipInitialQuota := 500    // VIP user initial quota
	normalInitialQuota := 100 // Normal user initial quota
	costPerService := 20      // Cost per service
	var wg sync.WaitGroup
	wg.Add(2)
	// VIP
	go simulateUserRequests(quotaManager, userA, vipInitialQuota, costPerService, true, &wg)
	// Normal
	go simulateUserRequests(quotaManager, userB, normalInitialQuota, costPerService, false, &wg)
	wg.Wait()
}
