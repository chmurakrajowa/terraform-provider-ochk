---
page_title: "IP Collection Resource"
---

# IP Collection Resource

IP Collection is a part of security group to group the IP addresses and after that can be used as sources and destinations in firewall rules. An IP Collection can contain a combination of individual IP addresses, IP ranges, and subnets.

## Example Usage

```hcl
data "ochk_projects" "project" {
  display_name = "project-example"
}

resource "ochk_ip_collection" "dns-servers" {
  display_name = "dns-servers"
  project_id = data.ochk.project.project.id
  ip_addresses = ["1.1.1.1", "1.0.0.1", "8.8.8.8"]
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Display name of IP Collection.
* `project_id` - (Required) Project id to which ip collection is assigned.
* `ip_addresses` - List of IP addresses that IP Collection contains. Each element can be an individual ip address (e.g: `10.10.1.1`), CIDR (e.g. `10.10.1.1/24`) or IP range (e.g. `10.10.1.1-10.10.1.10`)
  
## Attribute Reference

The following attributes are exported in addition to above arguments:
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred. 
