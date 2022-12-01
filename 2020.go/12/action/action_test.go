package action

import (
	"reflect"
	"testing"
)

func TestNewActionFrom(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *Action
	}{
		{
			name: "should build an action from a string",
			args: args{input: "F10"},
			want: &Action{
				actionType: 'F',
				units:      10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewActionFrom(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewActionFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
