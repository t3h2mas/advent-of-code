package passwords

type PositionValidator struct{}

func (v *PositionValidator) Validate(corporatePassword *CorporatePassword) bool {
	correctPositionCount := 0

	if rune(corporatePassword.password[corporatePassword.policy.min-1]) == corporatePassword.policy.char {
		correctPositionCount++
	}

	if rune(corporatePassword.password[corporatePassword.policy.max-1]) == corporatePassword.policy.char {
		correctPositionCount++
	}

	// valid if only 1 position is correct
	return correctPositionCount == 1
}
