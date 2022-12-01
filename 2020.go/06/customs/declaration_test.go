package customs

import (
	"reflect"
	"testing"
)

func TestParseGroupAnswers(t *testing.T) {
	type args struct {
		answers []string
	}
	tests := []struct {
		name string
		args args
		want FormAnswers
	}{
		{
			name: "group with 6 questions",
			args: args{
				answers: []string{
					"abcx",
					"abcy",
					"abcz",
				},
			},
			want: map[rune]bool{
				'a': true,
				'b': true,
				'c': true,
				'x': true,
				'y': true,
				'z': true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseGroupAnswers(tt.args.answers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseGroupAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}
