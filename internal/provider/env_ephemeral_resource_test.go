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

func TestAccExampleEphemeralResource(t *testing.T) {
	t.Setenv("foo", "bar")
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_10_0),
		},
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactoriesWithEcho,
		Steps: []resource.TestStep{
			{
				Config: testAccExampleEphemeralResourceConfig(),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"echo.test",
						tfjsonpath.New("data").AtMapKey("foo"),
						knownvalue.StringExact("bar"),
					),
					statecheck.ExpectKnownValue(
						"echo.test",
						tfjsonpath.New("data").AtMapKey("non-exist"),
						knownvalue.Null(),
					),
				},
			},
		},
	})
}

func testAccExampleEphemeralResourceConfig() string {
	return `
ephemeral "env" "test" {
	variables = ["foo", "non-exist"]
}

provider "echo" {
	data = ephemeral.env.test.values
}

resource "echo" "test" {}
`
}
