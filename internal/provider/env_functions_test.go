package provider_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/magodo/terraform-provider-env/internal/provider"
)

func TestEnvsFunction(t *testing.T) {
	os.Setenv("foo", "bar")
	defer os.Unsetenv("foo")

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "envs" {
					value = provider::env::envs()
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"envs",
						knownvalue.MapPartial(map[string]knownvalue.Check{
							"foo": knownvalue.StringExact("bar"),
						}),
					),
				},
			},
		},
	})
}
