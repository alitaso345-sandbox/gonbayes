package gonbayes_test

import (
	"github.com/alitaso345-sandbox/gonbayes"
	"reflect"
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
			got := gonbayes.IsStopWords(tt.word)
			if got != tt.want {
				t.Errorf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}

func TestStem(t *testing.T) {
	tests := []struct {
		name, word, want string
	}{
		{
			name: "stemming swin",
			word: "swims",
			want: "swim",
		},
		{
			name: "stemminh ceiling",
			word: "ceiling",
			want: "ceil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gonbayes.Stem(tt.word)
			if got != tt.want {
				t.Errorf("want = %v, got = %v\n", tt.want, got)
			}
		})
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		name, word, want string
	}{
		{
			name: "clean",
			word: "You are the best friend!!!!!!(^^)",
			want: "you are the best friend",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gonbayes.Clean(tt.word)
			if got != tt.want {
				t.Errorf("want = %v, got = %v\n", tt.want, got)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name string
		word string
		want map[string]int
	}{
		{
			name: "count words",
			word: "I say hello. You say hi.",
			want: map[string]int{
				"say":   2,
				"hello": 1,
				"hi":    1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gonbayes.CountWords(tt.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want = %v, got = %v\n", tt.want, got)
			}
		})
	}
}
