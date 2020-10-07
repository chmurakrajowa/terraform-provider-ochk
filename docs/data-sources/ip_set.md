---
page_title: "IP Set Data Source"
---

# IP Set Data Source

IP Set is a part of security group to group the IP addresses and after that can be used as sources and destinations in firewall rules. An IP Set can contain a combination of individual IP addresses, IP ranges, and subnets. 
This Data Source is used for reading IP Sets by display name. 

## Example Usage

```hcl
data "ochk_ip_set" "ips" {
  display_name = "ip-set-display-name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of IP Set.

## Attribute Reference

The following attributes are exported:
 * `display_name` - IP Set name. 
 * `addresses` - List of IP addresses that IP Set contains. Each entry has following values:
   * **address**: address, e.g. `10.10.1.1/24`, `10.10.1.1`, `10.10.1.1-10.10.1.10  `
   * **id**: address identifier
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred. 
     
 
