# terraform-provider-ipam

_Toying with the idea of IP Address Management (IPAM) backed by AWS DynamoDB or similar_

## Roadmap

- [ ] Define `network` resource
  - [ ] Description
  - [ ] Designate: [supernet, subnet (usable)]
  - [ ] Optional parent network
  - [ ] Tags
- [ ] Define `iprange` resource
  - [ ] Description
  - [ ] Required parent `network`
  - [ ] Tags
- [ ] Define `association` resource
  - [ ] Description
  - [ ] Required parent `network` or `iprange`
  - [ ] Tags

## Build and Develop

_Eventually this will be auto-built using Github Actions, but for now you can build locally._

### Pre-requisites

- [Go](https://www.terraform.io/downloads.html) (or install using your usual package manager)
- [task](https://taskfile.dev/#/installation)
- [terraform](https://www.terraform.io/downloads.html)

### Develop

Continuously run tests as you develop

```sh
task test --watch
```

### Install

Install the plugin locally `~/.terraform.d/plugins/`

```sh
task install
```

### External Test

Runs a `terraform plan` against the module in `example/`

```sh
task validate
```