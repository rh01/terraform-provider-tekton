# Terraform Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">


## Requirements


-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.14 (to build the provider plugin)

## Building The Provider

1. Clone repository

2. Change directory to repository root folder

3. Build the project by running `make build`

```sh
$ make build
```

## Using the provider

In order to use it in your computer, run install command inside the repository:

```sh
$ make install
```

You can find examples how to use the provider in [_example directory](https://github.com/rh01/terraform-provider-tekton/tree/master/_examples)

## Contributing to the Provider

### Code structure

* All code located under [tekton directory](https://github.com/nirarg/terraform-provider-tekton/tree/master/tekton)
* Backend tekton client wrapper located in [client](https://github.com/nirarg/terraform-provider-tekton/tree/master/tekton/client/client.go)
* All terraform schema definitions located under [schema directory](https://github.com/nirarg/terraform-provider-tekton/tree/master/tekton/schema)
* Terraform resource is defined (operations and schema) in `tekton/resource_*.go` for example: virtualmachine resource defined [here](https://github.com/rh01/terraform-provider-tekton/tree/master/tekton/resource_task.go)
* The main file, which define the provider's flags and structures is [provider](https://github.com/nirarg/terraform-provider-tekton/tree/master/tekton/provider.go)

### Development Environment

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.14+ is *required*).

To compile the provider, run `make build`. This will build the provider and put the provider binary in the repository root directory.

```sh
$ make build
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run terraform on your computer with `terraform-provider-tekton`, run `make install`

```sh
$ make install
```
That would build the binary and copy it to `$(HOME)/.terraform.d/plugins/$(GOOS)_$(GOARCH)`
