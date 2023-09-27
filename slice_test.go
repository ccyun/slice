package slice

import "testing"

func TestInSlice(t *testing.T) {
	type args[T comparable] struct {
		value     T
		sliceData []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}

	v := int64(0)
	vv := int64(0)
	tests := []testCase[*int64]{
		{
			name: "uint64---0",
			args: args[*int64]{
				value:     &v,
				sliceData: []*int64{&vv},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSlice(tt.args.value, tt.args.sliceData); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
	tests2 := []testCase[string]{
		{
			name: "string---0",
			args: args[string]{
				value:     "s",
				sliceData: []string{"aa", "bb", "cc"},
			},
			want: false,
		},
		{
			name: "string---1",
			args: args[string]{
				value:     "s",
				sliceData: []string{"aa", "bb", "s", "cc"},
			},
			want: true,
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSlice(tt.args.value, tt.args.sliceData); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
