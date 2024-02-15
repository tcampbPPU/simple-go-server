package helpers

import (
	"reflect"
	"testing"
)

type T int

func TestReverse(t *testing.T) {
	type args struct {
		s []T
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		{
			name: "Test Reverse",
			args: args{
				s: []T{1, 2, 3, 4, 5},
			},
			want: []T{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
