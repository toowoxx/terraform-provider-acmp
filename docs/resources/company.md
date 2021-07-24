# acmp_company

Manages an ACMP Company

## Example Usage

```hcl
resource "acmp_company" "example" {
  name = "mycompany"
  street = "mystreet"
  city = "mycity"
  country = "mygermany"
  zip = "myzip"
  email = "myemail@mydomain.com"
  currency = "mycurrency"
  customer_id = ""
  domain = ["mydomain1.com","mydomain2.com"]
  marketplaces = ["myMarketplace1","MyMarketplace2"]
}
```

## Argument Reference

* `name` - (Required) Name of the company.
* `street` - (Required) Street address of the company.
* `city` - (Required) City of the company.
* `country` - (Required) Countryof the company.
* `zip` - (Required) Zip code of the company.
* `email` - (Required) Contact e-mail address of the company.
* `currency` - (Required) Primary currency of the company.
* `customer_id` - (Required) Customer_id of the company.
* `domain` - (Required) Domain-Names of the company. Multiple values can be submitted.
* `marketplaces` - (Required) Marketplaces available for this company. Multiple values can be submitted.

## Attribute Reference

* `attribute_name` - List attributes that this resource exports.

## Timeouts

## Import

