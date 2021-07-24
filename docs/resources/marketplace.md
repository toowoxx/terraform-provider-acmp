# acmp_marketplace

Manages an ACMP Marketplace

## Example Usage

```hcl
resource "acmp_marketplace" "example" {
    name = "mymarketplace1"
    services = ["myservice1","myservice2"]
}
```

## Argument Reference

* `name` - (Required) Name of the marketplace.
* `services` - (Optional) List of services that should be available in this marketplace.

## Attribute Reference

* `attribute_name` - List attributes that this resource exports.

## Timeouts

## Import

