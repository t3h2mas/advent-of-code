package docking

import "testing"

func TestProgram_Write(t *testing.T) {
	type fields struct {
		program Program
	}
	type args struct {
		address int
		value   int
		mask    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "should write a value to memory",
			fields: fields{NewProgram()},
			args: args{
				address: 8,
				value:   11,
				mask:    "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			},
			want: 11,
		},
		{
			name:   "should write a modified value to memory",
			fields: fields{NewProgram()},
			args: args{
				address: 8,
				value:   11,
				mask:    "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			},
			want: 73,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields.program
			p.Write(tt.args.address, tt.args.value, tt.args.mask)

			got := p.memory[tt.args.address]

			if got != tt.want {
				t.Errorf("got = %v, want: %v\n", got, tt.want)
			}
		})
	}
}
