---
page_title: "IP Collection Data Source"
---

# IP Collection Data Source

IP Collection is a part of security group to group the ip addresses and after that can be used as sources and destinations in firewall rules. An ip collection can contain a combination of individual ip addresses, ip ranges, and subnets. 
This Data Source is used for reading ip collections by display name. 

## Example Usage

```hcl
data "ochk_ip_collection" "{{ .DataSourceName}}" {
  display_name = "ip_collection_name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of ip collection.

## Attribute Reference

The following attributes are exported:
 * `display_name` - IP collection name. 
 * `project_id` - Project id to which ip collection is assigned.
 * `ip_addresses` - List of ip addresses that ip collection contains. Each element can be an individual ip address (e.g: `10.10.1.1`), CIDR (e.g. `10.10.1.1/24`) or ip range (e.g. `10.10.1.1-10.10.1.10`).
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred. 
     
 
