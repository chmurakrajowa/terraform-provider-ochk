---
page_title: "IP Collection Data Source"
---

# IP Collections Data Source

IP Collection is a part of security group to group the ip addresses and after that can be used as sources and destinations in firewall rules. An ip collection can contain a combination of individual ip addresses, ip ranges, and subnets.
This Data Source is used for reading list of all ip collections.

## Example Usage

```hcl
data "ochk_ip_collections" "ipcs" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `ip_collections` - ip collections list. Each entry has the following values:
    * **display_name** - IP collection name.
    * **ip_collection_id** - IP collection id.
    * **project_id** - Project id to which ip collection is assigned.
    * **ip_addresses** - List of ip addresses that ip collection contains. Each element can be an individual ip address (e.g: `10.10.1.1`), CIDR (e.g. `10.10.1.1/24`) or ip range (e.g. `10.10.1.1-10.10.1.10`).

