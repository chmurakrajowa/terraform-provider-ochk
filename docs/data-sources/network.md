---
page_title: "Network Data Source"
---

# Network Data Source

Data Source for getting network by name.

## Example Usage

```hcl
data "ochk_network" "sg" {
  name = "network-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Exact name of Network.

## Attribute Reference

The following attributes are exported:
The following attributes are exported:
 * `name` - Name of the network.
 * `network` - Identifier of the network
 * `network_type` - The type (STANDARD_PORTGROUP, DISTRIBUTED_PORTGROUP, OPAQUE_NETWORK) of a server network
 
    
 
