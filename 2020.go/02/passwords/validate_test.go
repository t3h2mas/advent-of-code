package passwords

import (
	"reflect"
	"testing"
)

func Test_corporatePasswordFrom(t *testing.T) {
	passwordEntry := "1-3 b: cdefg"

	expected := &CorporatePassword{
		password: "cdefg",
		policy: &PasswordPolicy{
			min:  1,
			max:  3,
			char: 'b',
		},
	}

	got := corporatePasswordFrom(passwordEntry)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("corporatePolicyFrom got() = %+v, want %+v", got, expected)
	}
}
