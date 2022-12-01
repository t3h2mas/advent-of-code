package boarding

import (
	"reflect"
	"testing"
)

func TestNewPass(t *testing.T) {
	type args struct {
		seat string
	}
	tests := []struct {
		name      string
		args      args
		want      *Pass
		wantError bool
	}{
		{
			name: "example one",
			args: args{
				seat: "BFFFBBFRRR",
			},
			want: &Pass{
				row: 70, col: 7,
			},
			wantError: false,
		},
		{
			name: "example two",
			args: args{
				seat: "FFFBBBFRRR",
			},
			want: &Pass{
				row: 14, col: 7,
			},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PassFrom(tt.args.seat)

			haveError := err != nil

			if haveError != tt.wantError {
				t.Errorf("PassFrom() unexpected error: %s", err.Error())
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PassFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPass_SeatID(t *testing.T) {
	type fields struct {
		row int
		col int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "it should calculate the seat id",
			fields: fields{
				row: 44,
				col: 5,
			},
			want: 357,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pass{
				row: tt.fields.row,
				col: tt.fields.col,
			}
			if got := p.SeatID(); got != tt.want {
				t.Errorf("SeatID() = %v, want %v", got, tt.want)
			}
		})
	}
}
