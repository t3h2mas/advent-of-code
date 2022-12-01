package ship

import (
	"reflect"
	"testing"
)

func TestPosition_RotateClockwise(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		degrees int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Position
	}{
		{
			name: "can rotate clockwise",
			fields: fields{
				x: 10,
				y: 4,
			},
			args: args{degrees: 90},
			want: Position{x: 4, y: -10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p.RotateClockwise(tt.args.degrees); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateClockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_RotateCounterClockwise(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		degrees int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Position
	}{
		{
			name: "can rotate counter clockwise",
			fields: fields{
				x: 4,
				y: -10,
			},
			args: args{degrees: 90},
			want: Position{10, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p.RotateCounterClockwise(tt.args.degrees); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateCounterClockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}
