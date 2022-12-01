package passwords

type ContainsValidator struct{}

func (v *ContainsValidator) Validate(corporatePassword *CorporatePassword) bool {
	charCount := 0

	for _, ch := range corporatePassword.password {
		if ch == corporatePassword.policy.char {
			charCount++
		}

		if charCount > corporatePassword.policy.max {
			return false
		}
	}

	return charCount >= corporatePassword.policy.min
}
