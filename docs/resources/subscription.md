# acmp_subscription

Manages an ACMP Subscription or service

## Example Usage

```hcl
resource "acmp_subscription" "example" {
    name = "mymarketplace1"
    account_id = ""
}
```

## Argument Reference

* `name` - (Required) Name of the subscription/service.
* `account_id` - (Required) Parent account ID of the subscription.

## Attribute Reference

* `attribute_name` - List attributes that this resource exports.

## Timeouts

## Import

