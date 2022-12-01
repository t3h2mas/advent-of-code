package docking

import (
	"reflect"
	"testing"
)

func Test_DecodeAddress(t *testing.T) {
	type args struct {
		address int
		mask    string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should generate floating bit addresses",
			args: args{
				address: 42,
				mask:    "000000000000000000000000000000X1001X",
			},
			want: []int{
				26,
				27,
				58,
				59,
			},
		},
		{
			name: "should generate floating bit addresses 2",
			args: args{
				address: 26,
				mask:    "00000000000000000000000000000000X0XX",
			},
			want: []int{
				16, 17, 18, 19, 24, 25, 26, 27,
			},
		},
		{
			name: "should generate floating bit addresses 3",
			args: args{
				address: 7,
				mask:    "00000000000000000000000000000010X10X",
			},
			want: []int{
				38, 39, 46, 47,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeAddress(tt.args.address, tt.args.mask); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
