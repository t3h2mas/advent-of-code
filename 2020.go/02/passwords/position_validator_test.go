package passwords

import "testing"

func TestPositionValidator_Validate(t *testing.T) {
	type args struct {
		corporatePassword *CorporatePassword
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "password with no valid positions",
			args: args{
				&CorporatePassword{
					policy: &PasswordPolicy{
						min:  1,
						max:  2,
						char: 'b',
					},
					password: "asdbb",
				},
			},
			want: false,
		},
		{
			name: "password with two valid positions",
			args: args{
				&CorporatePassword{
					policy: &PasswordPolicy{
						min:  1,
						max:  2,
						char: 'b',
					},
					password: "bbdcc",
				},
			},
			want: false,
		},
		{
			name: "password with one valid position",
			args: args{
				&CorporatePassword{
					policy: &PasswordPolicy{
						min:  1,
						max:  2,
						char: 'b',
					},
					password: "bsdbb",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &PositionValidator{}
			if got := v.Validate(tt.args.corporatePassword); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
