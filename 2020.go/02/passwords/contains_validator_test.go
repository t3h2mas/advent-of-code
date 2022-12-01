package passwords

import "testing"

func TestContainsValidator_Validate(t *testing.T) {
	type args struct {
		corporatePassword *CorporatePassword
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid corporate password",
			args: args{
				&CorporatePassword{
					password: "abcde",
					policy: &PasswordPolicy{
						min:  1,
						max:  3,
						char: 'a',
					},
				},
			},
			want: true,
		},
		{
			name: "invalid corporate password",
			args: args{
				&CorporatePassword{
					password: "cdefg",
					policy: &PasswordPolicy{
						min:  1,
						max:  3,
						char: 'b',
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &ContainsValidator{}
			if got := v.Validate(tt.args.corporatePassword); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
