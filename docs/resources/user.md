# acmp_user

Manages an ACMP User Account

## Example Usage

```hcl
resource "acmp_user" "example" {
    name = "mymarketplace1"
    first_name = "John"
    last_name = "Doe"
    email = "john.doe@example.com"
    parent_account_id = "100001"
}
```

## Argument Reference

* `name` - (Required) Username of the user.
* `first_name` - (Required) First name of the user.
* `last_name` - (Required) Last name of the user.
* `email` - (Required) E-mail address of the user.
* `parent_account_id` - (Required) Parent Account ID of the User

## Attribute Reference

* `attribute_name` - List attributes that this resource exports.

## Timeouts

## Import

