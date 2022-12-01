package passwords

type Validator interface {
	Validate(corporatePassword *CorporatePassword) bool
}

func ValidPasswordsByPolicy(entries []string, validator Validator) int {
	validCount := 0

	for _, entry := range entries {
		corporatePassword := corporatePasswordFrom(entry)

		if validator.Validate(corporatePassword) {
			validCount++
		}
	}

	return validCount
}
