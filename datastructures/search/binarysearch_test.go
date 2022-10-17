package main

import (
	"testing"
)

func BenchmarkBinarySearchRecursive(b *testing.B) {
	var testCase []int
	for i := 0; i < 1000; i++ {
		testCase = append(testCase, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = binarySearchRecursive(testCase, i, 0, len(testCase)-1)
	}
}

func Test_binarySearchRecursive(t *testing.T) {
	type args struct {
		arr    []int
		target int
		low    int
		high   int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := binarySearchRecursive(tt.args.arr, tt.args.target, tt.args.low, tt.args.high)
			print(got, " - ", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("binarySearchRecursive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			print(got, " - ", tt.want)
			if got != tt.want {
				t.Errorf("binarySearchRecursive() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_binarySearchIterative(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := binarySearchIterative(tt.args.arr, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("binarySearchIterative() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("binarySearchIterative() got = %v, want %v", got, tt.want)
			}
		})
	}
}
