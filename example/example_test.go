package example

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	argss := []args{
		{1, 2},
		{2, 2},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2=3",
			args: argss[0],
			want: 3,
		},
		{
			name: "2+2=4",
			args: argss[1],
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
