package provider

import (
	"context"
	"terraform-provider-convcase/internal/convcase"
	"terraform-provider-convcase/internal/datasources"
	"terraform-provider-convcase/internal/functions"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type convcaseProvider struct {
	version string
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &convcaseProvider{
			version: version,
		}
	}
}

func (c *convcaseProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "convcase"
	resp.Version = c.version
}

func (c *convcaseProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (c *convcaseProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

func (c *convcaseProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewConvcaseDataSource,
	}
}

func (c *convcaseProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

func (c *convcaseProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		functions.ConvCaseFunctionFactory("conv_camel", "camel case", convcase.CamelCase),
		functions.ConvCaseFunctionFactory("conv_pascal", "pascal case", convcase.PascalCase),
		functions.ConvCaseFunctionFactory("conv_snake", "snake case", convcase.SnakeCase),
		functions.ConvCaseFunctionFactory("conv_const", "const case", convcase.ConstantCase),
		functions.ConvCaseFunctionFactory("conv_kebab", "kebab case", convcase.KebabCase),
		functions.ConvCaseFunctionFactory("conv_train", "train case", convcase.TrainCase),
		functions.ConvCaseFunctionFactory("conv_path", "path style", convcase.PathStyle),
		functions.ConvCaseFunctionFactory("conv_dot", "dot style", convcase.DotStyle),
	}
}
