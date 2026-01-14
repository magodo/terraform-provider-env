package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/magodo/terraform-provider-env/internal/provider"
)

func TestEnvFunction(t *testing.T) {
	t.Setenv("foo", "bar")
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "terraform_data" "test" {
					input = provider::env::env("non-exist")
				}
				output "foo" {
					value = provider::env::env("foo")
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"foo",
						knownvalue.StringExact("bar"),
					),
					statecheck.ExpectKnownValue(
						"terraform_data.test",
						tfjsonpath.New("output"),
						knownvalue.Null(),
					),
				},
			},
		},
	})
}
