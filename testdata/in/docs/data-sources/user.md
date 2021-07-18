# foo\_user

Provides blah.

Also blah.

## Example Usage

```hcl
data "foo_user" "example" {
  username = "myuser"
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Optional) The username of the user.

* `user_id` - (Optional) The ID of the user.

**Note**: only one of user_id or username must be provided.

## Attributes Reference

* `email` - The e-mail address of the user.

* `name` - The name of the user.

**Note**: some attributes might not be returned.

Some more text.
