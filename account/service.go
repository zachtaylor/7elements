package account

// Service provides Accounts
type Service interface {
	// Test returns an account from cache only
	Test(string) *T
	// Cache stores an account
	Cache(*T)
	// Forget uncaches an account
	Forget(string)
	// Find uses Test/Get&Cache best effort to provide account
	Find(string) (*T, error)
	// Get loads an account from back end
	Get(string) (*T, error)
	// GetCount returns a number of registered accounts from back end
	GetCount() (int, error)
	// Insert creates an account on back end
	Insert(*T) error
	// UpdateCoins updates an accounts coin count on back end
	UpdateCoins(*T) error
	// UpdateEmail updates an accounts email
	UpdateEmail(*T) error
	// UpdateLogin updates an accounts login time on back end
	UpdateLogin(*T) error
	// UpdatePassword updates an accounts password on back end
	UpdatePassword(*T) error
	// Delete removes an account
	Delete(string) error
}
