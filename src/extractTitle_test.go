package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractTitle(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantTitle   string
		wantContent string
		expectErr   bool
	}{
		{
			name:        "Simple H1 at top",
			input:       "# My Title\n\nThis is a paragraph.",
			wantTitle:   "My Title",
			wantContent: "This is a paragraph.",
			expectErr:   false,
		},
		{
			name:        "H1 with leading/trailing whitespace",
			input:       "   #   Trim me   \n\nAnother paragraph.",
			wantTitle:   "Trim me",
			wantContent: "Another paragraph.",
			expectErr:   false,
		},
		{
			name:        "H1 not at top",
			input:       "Intro text\n\n# The Real Title\n\nMore text.",
			wantTitle:   "The Real Title",
			wantContent: "Intro text\n\nMore text.",
			expectErr:   false,
		},
		{
			name:        "No H1 present",
			input:       "## Not a title\n\nRegular paragraph.",
			wantTitle:   "",
			wantContent: "",
			expectErr:   true,
		},
		{
			name:        "Only H1",
			input:       "# Standalone Title",
			wantTitle:   "Standalone Title",
			wantContent: "",
			expectErr:   false,
		},
		{
			name:        "H1 with markdown kept intact",
			input:       "# **Bold** _Title_",
			wantTitle:   "**Bold** _Title_",
			wantContent: "",
			expectErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			title, content, err := extractTitle(tt.input)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantTitle, title)
				assert.Equal(t, tt.wantContent, content)
			}
		})
	}
}
