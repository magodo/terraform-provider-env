# terraform-provider-env

`terraform-provider-env` is a logical provider to access the environment variable(s) in the running environment through the following ways:

- (**Recommended**) [Ephemeral Resource](https://developer.hashicorp.com/terraform/language/block/ephemeral) `env`: Access a list of specified environment variables.
- [Function](https://developer.hashicorp.com/terraform/plugin/framework/functions) `provider::env::env()`: Access a single environment variable by name.
- [Function](https://developer.hashicorp.com/terraform/plugin/framework/functions) `provider::env::envs()`: Access all environment variables.
