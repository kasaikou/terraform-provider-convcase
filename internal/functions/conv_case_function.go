package functions

import (
	"context"
	"fmt"
	"terraform-provider-convcase/internal/convcase"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/samber/lo"
)

var ()

type convCaseFunction struct {
	funcName string
	caseName string
	convCase convcase.ConvCase
}

func ConvCaseFunctionFactory(
	funcName string,
	caseName string,
	convCase convcase.ConvCase,
) func() function.Function {
	return func() function.Function {
		return &convCaseFunction{
			funcName: funcName,
			caseName: caseName,
			convCase: convCase,
		}
	}
}

func (p *convCaseFunction) Metadata(_ context.Context, _ function.MetadataRequest, res *function.MetadataResponse) {
	res.Name = p.funcName
}

func (p *convCaseFunction) Definition(_ context.Context, _ function.DefinitionRequest, res *function.DefinitionResponse) {
	caseName := p.convCase.Convert(lo.Must(convcase.SplitWords(p.caseName)))

	res.Definition = function.Definition{
		Summary:     fmt.Sprintf("Convert to %s", caseName),
		Description: fmt.Sprintf("Given a string value, returns the same value converted to %s.", caseName),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name: "input",
			},
		},
		Return: function.StringReturn{},
	}
}

func (p *convCaseFunction) Run(ctx context.Context, req function.RunRequest, res *function.RunResponse) {
	input := ""
	res.Error = function.ConcatFuncErrors(res.Error, req.Arguments.Get(ctx, &input))

	words, err := convcase.SplitWords(input)
	if err != nil {
		res.Error = function.ConcatFuncErrors(res.Error, function.NewFuncError(fmt.Sprintf("Failed to split words: %s", err)))
	}
	res.Error = function.ConcatFuncErrors(res.Error, res.Result.Set(ctx, p.convCase.Convert(words)))
}
