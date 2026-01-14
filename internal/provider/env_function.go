package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = EnvFunction{}
)

func NewEnvFunction() function.Function {
	return EnvFunction{}
}

type EnvFunction struct{}

func (r EnvFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "env"
}

func (r EnvFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Read a single environment variable",
		MarkdownDescription: "Read a single environment variable.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "variable",
				MarkdownDescription: "The name of the environment variable.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r EnvFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &data))
	if resp.Error != nil {
		return
	}

	v, ok := os.LookupEnv(data)
	result := types.StringValue(v)
	if !ok {
		result = types.StringNull()
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
