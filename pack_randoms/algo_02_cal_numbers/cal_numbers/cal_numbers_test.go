package calnumbers

import "testing"

func TestCalInStandardArrZeroToN(t *testing.T) {
	type args struct {
		n   int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test-case-1",
			args: args{
				n:   2,
				key: 2,
			},
			want: 1,
		},
		{
			name: "test-case-2",
			args: args{
				n:   6,
				key: 4,
			},
			want: 1,
		},
		{
			name: "test-case-3",
			args: args{
				n:   12,
				key: 1,
			},
			want: 5,
		},
		{
			name: "test-case-4",
			args: args{
				n:   17,
				key: 1,
			},
			want: 10,
		},
		{
			name: "test-case-5",
			args: args{
				n:   23,
				key: 3,
			},
			want: 3,
		},
		{
			name: "test-case-6",
			args: args{
				n:   29,
				key: 2,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalInStandardArrZeroToN(tt.args.n, tt.args.key); got != tt.want {
				t.Errorf("CalInStandardArrZeroToN() = %v, want %v", got, tt.want)
			}
		})
	}
}
