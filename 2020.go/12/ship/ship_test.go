package ship

import (
	"reflect"
	"testing"
)

func TestNewShip_starts_facing_East(t *testing.T) {
	got := NewShip().facing
	expected := East

	if got != expected {
		t.Errorf("NewShip().facing = %s, want %s", got.String(), expected.String())
	}
}

func TestShip_turnRight(t *testing.T) {
	type fields struct {
		facing CardinalDirection
	}
	tests := []struct {
		name   string
		fields fields
		want   CardinalDirection
	}{
		{
			name:   "turning right from North",
			fields: fields{facing: North},
			want:   East,
		},
		{
			name:   "turning right from East",
			fields: fields{facing: East},
			want:   South,
		},
		{
			name:   "turning right from South",
			fields: fields{facing: South},
			want:   West,
		},
		{
			name:   "turning right from West",
			fields: fields{facing: West},
			want:   North,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Ship{
				facing: tt.fields.facing,
			}

			s.turnRight()

			got := s.facing

			if got != tt.want {
				t.Errorf("Ship.facing = %s, want %s", got.String(), tt.want.String())
			}
		})
	}
}

func TestShip_turnLeft(t *testing.T) {
	type fields struct {
		facing CardinalDirection
	}
	tests := []struct {
		name   string
		fields fields
		want   CardinalDirection
	}{
		{
			name:   "turning left from North",
			fields: fields{North},
			want:   West,
		},
		{
			name:   "turning left from West",
			fields: fields{West},
			want:   South,
		},
		{
			name:   "turning left from South",
			fields: fields{South},
			want:   East,
		},
		{
			name:   "turning left from East",
			fields: fields{East},
			want:   North,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Ship{
				facing: tt.fields.facing,
			}

			s.turnLeft()

			got := s.facing

			if got != tt.want {
				t.Errorf("Ship.facing = %s, want %s", got.String(), tt.want.String())
			}
		})
	}
}

func TestShip_turn(t *testing.T) {
	type fields struct {
		facing   CardinalDirection
		position Position
	}
	type args struct {
		direction Direction
		degrees   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   CardinalDirection
	}{
		{
			name: "should turn left 90 degrees",
			fields: fields{
				facing: East,
			},
			args: args{
				direction: Left,
				degrees:   90,
			},
			want: North,
		},
		{
			name: "should turn left 270 degrees",
			fields: fields{
				facing: East,
			},
			args: args{
				direction: Left,
				degrees:   270,
			},
			want: South,
		},
		{
			name: "should turn right 90 degrees",
			fields: fields{
				facing: East,
			},
			args: args{
				direction: Right,
				degrees:   90,
			},
			want: South,
		},
		{
			name: "should turn right 270 degrees",
			fields: fields{
				facing: East,
			},
			args: args{
				direction: Right,
				degrees:   270,
			},
			want: North,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Ship{
				facing:   tt.fields.facing,
				position: tt.fields.position,
			}
			s.turn(tt.args.direction, tt.args.degrees)
			got := s.facing

			if got != tt.want {
				t.Errorf("Ship.facing = %s, want %s", got.String(), tt.want.String())
			}
		})
	}
}

func TestShip_move(t *testing.T) {
	type args struct {
		direction CardinalDirection
		units     int
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{
			name: "should move north",
			args: args{
				direction: North,
				units:     10,
			},
			want: Position{
				x: 0,
				y: 10,
			},
		},
		{
			name: "should move East",
			args: args{
				direction: East,
				units:     10,
			},
			want: Position{
				x: 10,
				y: 0,
			},
		},
		{
			name: "should move West",
			args: args{
				direction: West,
				units:     10,
			},
			want: Position{
				x: -10,
				y: 0,
			},
		},
		{
			name: "should move South",
			args: args{
				direction: South,
				units:     10,
			},
			want: Position{
				x: 0,
				y: -10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewShip()

			s.move(tt.args.direction, tt.args.units)

			got := s.position

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %#v, want %#v", got, tt.want)
			}
		})
	}
}
