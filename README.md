# terraform-provider-env

**NOTE**: Accessing environment variables in a Terraform module is a violation of *Terraform's language design* where it *tries to avoid situations where a module's definition depends on anything other than its own source code and its input variables* (See [this](https://github.com/hashicorp/terraform/issues/26477#issuecomment-703845535) for details). Therefore, only use this provider as a last resort.

`terraform-provider-env` is a logical provider to access the environment variable(s) in the running environment through the following ways:

- [Ephemeral Resource](https://developer.hashicorp.com/terraform/language/block/ephemeral) `env`: Access a list of specified environment variables.
- [Function](https://developer.hashicorp.com/terraform/plugin/framework/functions) `provider::env::env()`: Access a single environment variable by name.
- [Function](https://developer.hashicorp.com/terraform/plugin/framework/functions) `provider::env::envs()`: Access all environment variables.

