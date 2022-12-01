package expenses

import (
	"reflect"
	"testing"
)

func Test_sumsTo(t *testing.T) {
	type args struct {
		target int
		search []int
	}
	tests := []struct {
		name    string
		args    args
		want    [2]int
		wantErr bool
	}{
		{
			name: "pair found for target",
			args: args{
				target: 2020,
				search: []int{
					1721,
					979,
					366,
					299,
					675,
					1456,
				},
			},
			want: [2]int{
				299, 1721,
			},
			wantErr: false,
		},
		{
			name: "pair not found for target",
			args: args{
				target: 2020,
				search: []int{
					979,
					366,
					299,
					675,
					1456,
				},
			},
			want:    [2]int{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sumsTo(tt.args.target, tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("sumsTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumsTo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tripletSumsTo(t *testing.T) {
	type args struct {
		target int
		search []int
	}
	tests := []struct {
		name    string
		args    args
		want    [3]int
		wantErr bool
	}{
		{
			name: "triplet found for target",
			args: args{
				target: 2020,
				search: []int{
					1721,
					979,
					366,
					299,
					675,
					1456,
				},
			},
			want:    [3]int{979, 675, 366},
			wantErr: false,
		},
		{
			name: "triplet not found for target",
			args: args{
				target: 2020,
				search: []int{
					1721,
					366,
					299,
					675,
					1456,
				},
			},
			want:    [3]int{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tripletSumsTo(tt.args.target, tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("tripletSumsTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tripletSumsTo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
