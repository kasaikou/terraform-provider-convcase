package convcase

import "testing"

func TestConvCase_Convert(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name     string
		convcase ConvCase
		args     args
		want     string
	}{
		{
			name:     "to camel case",
			convcase: CamelCase,
			args:     args{words: []string{"abc", "def", "xyz"}},
			want:     "abcDefXyz",
		},
		{
			name:     "to pascal case",
			convcase: PascalCase,
			args:     args{words: []string{"abc", "def", "xyz"}},
			want:     "AbcDefXyz",
		},
		{
			name:     "to snake case",
			convcase: SnakeCase,
			args:     args{words: []string{"abc", "def", "xyz"}},
			want:     "abc_def_xyz",
		},
		{
			name:     "to constant case",
			convcase: ConstantCase,
			args:     args{words: []string{"abc", "def", "xyz"}},
			want:     "ABC_DEF_XYZ",
		},
		{
			name:     "to kebab case",
			convcase: KebabCase,
			args:     args{words: []string{"abc", "def", "xyz"}},
			want:     "abc-def-xyz",
		},
		{
			name:     "to train case",
			convcase: TrainCase,
			args:     args{words: []string{"abc", "def", "xyz"}},
			want:     "Abc-Def-Xyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.convcase.Convert(tt.args.words); got != tt.want {
				t.Errorf("ConvCase.Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
