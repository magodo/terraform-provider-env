package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ephemeral.EphemeralResource = &EnvEphemeralResource{}

func NewEnvEphemeralResource() ephemeral.EphemeralResource {
	return &EnvEphemeralResource{}
}

type EnvEphemeralResource struct{}

type EnvEphemeralResourceModel struct {
	Variables types.List `tfsdk:"variables"`
	Values    types.Map  `tfsdk:"values"`
}

func (r *EnvEphemeralResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName
}

func (r *EnvEphemeralResource) Schema(ctx context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Ephemeral resource to access environment variables.",
		Attributes: map[string]schema.Attribute{
			"variables": schema.ListAttribute{
				MarkdownDescription: "The list of environment variable names to retrieve.",
				Required:            true,
				ElementType:         types.StringType,
			},
			"values": schema.MapAttribute{
				MarkdownDescription: "The map of the environment variable values retrieved.",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func (r *EnvEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data EnvEphemeralResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var names []string
	resp.Diagnostics.Append(data.Variables.ElementsAs(ctx, &names, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	m := map[string]attr.Value{}
	for _, name := range names {
		v, ok := os.LookupEnv(name)
		vv := types.StringValue(v)
		if !ok {
			vv = types.StringNull()
		}
		m[name] = vv
	}

	values, diags := types.MapValue(types.StringType, m)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Values = values

	resp.Diagnostics.Append(resp.Result.Set(ctx, &data)...)
}
