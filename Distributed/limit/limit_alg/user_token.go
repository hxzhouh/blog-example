package limit_alg

import (
	"sync"
)

// UserQuota
type UserQuota struct {
	Quanity int
	Limit   int
	mutex   sync.Mutex
}

// QuotaManager Managing user credits
type QuotaManager struct {
	users map[string]*UserQuota
	mutex sync.Mutex
}

// NewQuotaManager Create a new quota manager
func NewQuotaManager() *QuotaManager {
	return &QuotaManager{
		users: make(map[string]*UserQuota),
	}
}

// GetQuota Get the user's quota and assign an initial quota if the user has no quota
func (qm *QuotaManager) GetQuota(isVip bool, userID string, initialQuota int) *UserQuota {
	qm.mutex.Lock()
	defer qm.mutex.Unlock()
	if isVip {
		initialQuota = initialQuota * 5
	}
	if _, ok := qm.users[userID]; !ok {
		qm.users[userID] = &UserQuota{
			Quanity: initialQuota,
			Limit:   initialQuota,
		}
	}

	return qm.users[userID]
}

// ConsumeQuota Consume the user's quota. If the quota is insufficient, return false.
func (qm *QuotaManager) ConsumeQuota(isVip bool, userID string, cost int) bool {
	quota := qm.GetQuota(isVip, userID, 100)

	quota.mutex.Lock()
	defer quota.mutex.Unlock()
	if quota.Limit < cost {
		return false
	}
	quota.Limit -= cost
	return true
}

// RefillQuota Reapply for quota for users
func (qm *QuotaManager) RefillQuota(isVip bool, userID string, newQuota int) {

	quota := qm.GetQuota(isVip, userID, 100)

	quota.mutex.Lock()
	defer quota.mutex.Unlock()
	quota.Limit = newQuota
}
