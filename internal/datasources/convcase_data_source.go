package datasources

import (
	"context"
	"terraform-provider-convcase/internal/convcase"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type convcaseDataSource struct{}

func NewConvcaseDataSource() datasource.DataSource {
	return &convcaseDataSource{}
}

func (c *convcaseDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName
}

func (c *convcaseDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"input":  schema.StringAttribute{Required: true},
			"camel":  schema.StringAttribute{Computed: true},
			"pascal": schema.StringAttribute{Computed: true},
			"snake":  schema.StringAttribute{Computed: true},
			"const":  schema.StringAttribute{Computed: true},
			"kebab":  schema.StringAttribute{Computed: true},
			"train":  schema.StringAttribute{Computed: true},
			"path":   schema.StringAttribute{Computed: true},
			"dot":    schema.StringAttribute{Computed: true},
		},
	}
}

type convcaseModel struct {
	Input  types.String `tfsdk:"input"`
	Camel  types.String `tfsdk:"camel"`
	Pascal types.String `tfsdk:"pascal"`
	Snake  types.String `tfsdk:"snake"`
	Const  types.String `tfsdk:"const"`
	Kebab  types.String `tfsdk:"kebab"`
	Train  types.String `tfsdk:"train"`
	Path   types.String `tfsdk:"path"`
	Dot    types.String `tfsdk:"dot"`
}

func (c *convcaseModel) WithResult(
	camel types.String,
	pascal types.String,
	snake types.String,
	constant types.String,
	kebab types.String,
	train types.String,
	path types.String,
	dot types.String,
) *convcaseModel {
	return &convcaseModel{
		Input:  c.Input,
		Camel:  camel,
		Pascal: pascal,
		Snake:  snake,
		Const:  constant,
		Kebab:  kebab,
		Train:  train,
		Path:   path,
		Dot:    dot,
	}
}

func (c *convcaseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	model := &convcaseModel{}
	diags := req.Config.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	words, err := convcase.SplitWords(model.Input.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to split words", err.Error())
		return
	}

	model = model.WithResult(
		types.StringValue(convcase.CamelCase.Convert(words)),
		types.StringValue(convcase.PascalCase.Convert(words)),
		types.StringValue(convcase.SnakeCase.Convert(words)),
		types.StringValue(convcase.ConstantCase.Convert(words)),
		types.StringValue(convcase.KebabCase.Convert(words)),
		types.StringValue(convcase.TrainCase.Convert(words)),
		types.StringValue(convcase.PathStyle.Convert(words)),
		types.StringValue(convcase.DotStyle.Convert(words)),
	)

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
