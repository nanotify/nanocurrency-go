package nanocurrency

// IsValidAccount returns a bool describing whether or not the account
// string input is a valid Nano account.
func IsValidAccount(a string) bool {
	_, err := NewAccount(a)
	return err == nil
}
