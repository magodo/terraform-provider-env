package provider

import (
	"context"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = EnvsFunction{}
)

func NewEnvsFunction() function.Function {
	return EnvsFunction{}
}

type EnvsFunction struct{}

func (r EnvsFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "envs"
}

func (r EnvsFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Read all the environment variables",
		MarkdownDescription: "Read all the environment variables.",
		Parameters:          []function.Parameter{},
		Return: function.MapReturn{
			ElementType: types.StringType,
		},
	}
}

func (r EnvsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	envs := map[string]string{}
	for _, expr := range os.Environ() {
		k, v, ok := strings.Cut(expr, "=")
		if !ok {
			continue
		}
		envs[k] = v
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, envs))
}
