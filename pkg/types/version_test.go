package types

import (
	"testing"
)

func TestVersion_IntVersion(t *testing.T) {
	tests := []struct {
		name string
		v    Version
		want int
	}{
		{
			name: "v0",
			v:    "v0",
			want: 0,
		},
		{
			name: "v0.0",
			v:    "v0.0",
			want: 0,
		},
		{
			name: "v0.0.0",
			v:    "v0.0.0",
			want: 0,
		},
		{
			name: "v0.1",
			v:    "v0.1",
			want: 1000,
		},
		{
			name: "v0.1.0",
			v:    "v0.1.0",
			want: 1000,
		},
		{
			name: "v0.1.1",
			v:    "v0.1.1",
			want: 1001,
		},
		{
			name: "v1",
			v:    "v1",
			want: 1000000,
		},
		{
			name: "v1.0",
			v:    "v1.0",
			want: 1000000,
		},
		{
			name: "v1.0.0",
			v:    "v1.0.0",
			want: 1000000,
		},
		{
			name: "v1.0.1",
			v:    "v1.0.1",
			want: 1000001,
		},
		{
			name: "v1.1",
			v:    "v1.1",
			want: 1001000,
		},
		{
			name: "v1.1.1",
			v:    "v1.1.1",
			want: 1001001,
		},
		{
			name: "v1.2.3",
			v:    "v1.2.3",
			want: 1002003,
		},
		{
			name: "v999.999.999",
			v:    "v999.999.999",
			want: 999999999,
		},
		{
			name: "v0.0.1000",
			v:    "v0.0.1000",
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.IntVersion(); got != tt.want {
				t.Errorf("IntVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
