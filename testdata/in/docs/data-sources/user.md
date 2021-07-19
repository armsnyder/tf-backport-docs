# foo\_user

Provides blah for `thing`.

Also blah for "stuff".

## Example Usage - Foo Bar

```hcl
data "foo_user" "example_1" {
  username = "myuser"
}
```

## Example Usage - Baz

```hcl
data "foo_user" "example_2" {
  username = "otheruser"
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Optional) The username of the user.

* `user_id` - (Optional) The ID of the user, for example `foo\:\/\/`.

**Note**: only one of user_id or username must be provided.

## Attributes Reference

* `email` - The e-mail address of the user.

* `name` - The name of the user.

* `dupe_same_desc` - A description.

* `dupe_same_desc` - A description.

* `dupe_diff_desc` - One description.

* `dupe_diff_desc` - Another description.

**Note**: some attributes might not be returned.

Some more text.
