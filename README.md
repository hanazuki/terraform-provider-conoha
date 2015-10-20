# terraform-provider-conoha
Terraform plugin for ConoHa -- this project is still a work in progress.

## Install
Just `go get && go build` and copy `terraform-provider-conoha` into the same directory as your `terraform` executable.

## Example

```hcl
provider "conoha" {
  region = "tyo1"
}

resource "conoha_compute_keypair" "anzu" {
  name = "anzu"
  public_key = "ssh-rsa XXXXXXXXXXXXXXX anzu@localhost"
}

resource "conoha_dns_domain" "example_com" {
  name = "exmaple.com."
  email = "webmaster@example.com"
}

```
