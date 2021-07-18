# tf-backport-docs

Tool for backporting Terraform provider documentation into code. Useful during transitional period from ad-hoc managed
documentation to terraform-plugin-docs generated documentation.

**NOTE:** Backported code is not guaranteed to be perfect. False negatives may occur as certain attributes are not
supported.

### Usage

Run the command from the root directory of a Terraform provider repository.

```shell
go run github.com/armsnyder/tf-backport-docs
```
