package simi

import (
	"reflect"
	"testing"
)

func TestAggr(t *testing.T) {
	type args struct {
		data   []map[string]string
		stKeys []string
	}
	tests := []struct {
		name string
		args args
		want map[int][]map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "test-case-1",
			args: args{
				data: []map[string]string{
					{"a": "2", "b": "4"},
					{"a": "2", "b": "7"},
				},
				stKeys: []string{"a"},
			},
			want: map[int][]map[string]string{
				1: []map[string]string{
					{"a": "2", "b": "4"},
					{"a": "2", "b": "7"},
				},
			},
		},

		{
			name: "test-case-2",
			args: args{
				data: []map[string]string{
					{"a": "2", "b": "1"}, // 1
					{"a": "3", "b": "6"}, // 2
					{"a": "7", "b": "8"}, // 3 -> 2
					{"a": "4", "b": "9"}, // 4
					{"a": "3", "b": "8"}, // 5 -> 2
				},
				stKeys: []string{"a", "b"},
			},
			want: map[int][]map[string]string{
				1: []map[string]string{
					{"a": "2", "b": "1"},
				},
				2: []map[string]string{
					{"a": "3", "b": "6"},
					{"a": "7", "b": "8"},
					{"a": "3", "b": "8"},
				},
				4: []map[string]string{
					{"a": "4", "b": "9"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Aggr(tt.args.data, tt.args.stKeys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aggr() = %v, want %v", got, tt.want)
			}
		})
	}
}
