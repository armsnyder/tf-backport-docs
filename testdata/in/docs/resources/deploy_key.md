# foo\_deploy\_key

This resource allows you to blah.

## Example Usage

**By doot**

```hcl
resource "foo_deploy_key" "example_1" {
  project = "example/deploying"
  title   = "Example deploy key"
}
```

### By blah

```hcl
resource "foo_deploy_key" "example_2" {
  project = "example/deploying2"
  title   = "Example deploy key 2"
}
```

## Argument Reference

The following arguments are supported:

* `project` - (Required, string) The name or id of the project to add the deploy key to.

* `title` - (Required, string) A title to describe the deploy key with.

## Import

You can import using `terraform import <resource> <id>`. The
`id` can be whatever the [get single project api][get_single_project] takes for
its `:id` value, so for example:

```
$ terraform import foo_deploy_key.test 1:3
```
