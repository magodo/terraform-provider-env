package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &EnvProvider{}
var _ provider.ProviderWithFunctions = &EnvProvider{}
var _ provider.ProviderWithEphemeralResources = &EnvProvider{}
var _ provider.ProviderWithActions = &EnvProvider{}

type EnvProvider struct {
	version string
}

func (p *EnvProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "env"
	resp.Version = p.version
}

func (p *EnvProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Terraform provider to access environment variable.",
	}
}

func (p *EnvProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *EnvProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *EnvProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		NewEnvEphemeralResource,
	}
}

func (p *EnvProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *EnvProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewEnvFunction,
		NewEnvsFunction,
	}
}

func (p *EnvProvider) Actions(ctx context.Context) []func() action.Action {
	return []func() action.Action{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &EnvProvider{
			version: version,
		}
	}
}
