---
page_title: "IP Collection Data Source"
---

# IP Collection Data Source

IP Collection is a part of security group to group the IP addresses and after that can be used as sources and destinations in firewall rules. An IP Collection can contain a combination of individual IP addresses, IP ranges, and subnets. 
This Data Source is used for reading IP Collections by display name. 

## Example Usage

```hcl
data "ochk_ip_collection" "ips" {
  display_name = "ip-set-display-name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of IP Collection.

## Attribute Reference

The following attributes are exported:
 * `display_name` - IP Collection name. 
 * `addresses` - List of IP addresses that IP Collection contains. Each element can be an individual ip address (e.g: `10.10.1.1`), CIDR (e.g. `10.10.1.1/24`) or IP range (e.g. `10.10.1.1-10.10.1.10`)
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred. 
     
 
