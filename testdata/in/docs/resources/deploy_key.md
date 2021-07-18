# foo\_deploy\_key

This resource allows you to blah.

Be careful with it.

## Example Usage

```hcl
resource "foo_deploy_key" "example" {
  project = "example/deploying"
  title   = "Example deploy key"
}
```

## Argument Reference

The following arguments are supported:

* `project` - (Required, string) The name or id of the project to add the deploy key to.

* `title` - (Required, string) A title to describe the deploy key with.

## Import

Deploy keys can be imported using an id made up of `{project_id}:{deploy_key_id}`, e.g.

```
$ terraform import foo_deploy_key.test 1:3
```
