package gonbayes_test

import (
	"github.com/alitaso345-sandbox/gonbayes"
	"testing"
)

func TestIsStopWord(t *testing.T) {
	tests := []struct {
		name, word string
		want       bool
	}{
		{
			name: "stop word",
			word: "you",
			want: true,
		},
		{
			name: "not stop word",
			word: "great",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gonbayes.IsStopWord(tt.word)
			if got != tt.want {
				t.Errorf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}
