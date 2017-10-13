package main

import "testing"

func Test_sliceElementInString(t *testing.T) {
	type args struct {
		slc []string
		inp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceElementInString(tt.args.slc, tt.args.inp); got != tt.want {
				t.Errorf("sliceElementInString() = %v, want %v", got, tt.want)
			}
		})
	}
}
