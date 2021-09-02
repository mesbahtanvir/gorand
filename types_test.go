package gorand

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func TestRune(t *testing.T) {
	for i := 0; i < 100000; i++ {
		assert.True(t, utf8.ValidRune(Rune()))
	}
}

func BenchmarkRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rune()
	}
}

func TestString_Empty(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "WHEN len negative THEN return empty string",
			args: args{len: -1},
			want: "",
		}, {
			name: "WHEN len is equal 0 THEN return empty string",
			args: args{len: 0},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.len); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_NonEmpty(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "WHEN len 10 THEN return string sized 10",
			args: args{len: 10},
		}, {
			name: "WHEN len is equal 51234 THEN return string sized 51234",
			args: args{len: 51234},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := String(tt.args.len)
			if utf8.RuneCountInString(got) != tt.args.len {
				t.Errorf("String() = %v, want %v", len(got), tt.args.len)
			}
			assert.True(t, utf8.Valid([]byte(got)))
		})
	}
}
