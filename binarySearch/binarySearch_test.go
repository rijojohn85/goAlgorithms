package binarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3}, 2},
		{"case2", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
