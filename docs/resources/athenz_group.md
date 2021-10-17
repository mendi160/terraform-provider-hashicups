---
page_title: "athenz_sub_domain Resource - terraform-provider-hashicups"
subcategory: ""
description: |-
The athenz_group resource allows you to create athenz group.
---

##Resource: athenz_group

`athenz_group` provides an Athenz group resource.

### Example Usage

```terraform
variable "group_name" {
  type = string
}

data "athenz_group" "selected" {
  name = var.group_name
  domain = "some_domain"
}
```

### Argument Reference

The arguments of this data source act as filters for querying the available groups in the current Athenz domain.
The given filters must match exactly one group whose data will be exported as attributes.

- `name` - (Required) The name of the specific Athenz group.

- `domain` - (Required) The Athenz domain name.

- `audit_ref` - (Optional Default = "done by terraform provider")  string containing audit specification or ticket number.

