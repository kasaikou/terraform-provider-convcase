package convcase

import (
	"reflect"
	"testing"
)

func TestSplitWords(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name      string
		args      args
		wantWords []string
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name:      "snake case",
			args:      args{text: "abc_def_xyz"},
			wantWords: []string{"abc", "def", "xyz"},
		},
		{
			name:      "constant case",
			args:      args{text: "ABC_DEF_XYZ"},
			wantWords: []string{"abc", "def", "xyz"},
		},
		{
			name:      "kebab case",
			args:      args{text: "abc-def-xyz"},
			wantWords: []string{"abc", "def", "xyz"},
		},
		{
			name:      "train case",
			args:      args{text: "Abc-Def-Xyz"},
			wantWords: []string{"abc", "def", "xyz"},
		},
		{
			name:      "path style",
			args:      args{text: "abc/def/xyz"},
			wantWords: []string{"abc", "def", "xyz"},
		},
		{
			name:      "dot style",
			args:      args{text: "abc.def.xyz"},
			wantWords: []string{"abc", "def", "xyz"},
		},
		{
			name:      "text style",
			args:      args{text: "abc def xyz"},
			wantWords: []string{"abc", "def", "xyz"},
		},
		{
			name:      "camel case",
			args:      args{"abcDEFXyz"},
			wantWords: []string{"abc", "DEF", "Xyz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWords, err := SplitWords(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("SplitWords() = %v, want %v", gotWords, tt.wantWords)
			}
		})
	}
}
