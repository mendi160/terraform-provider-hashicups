---
page_title: "athenz_sub_domain Resource - terraform-provider-hashicups"
subcategory: ""
description: |-
The athenz_sub_domain resource allows you to create athenz sub-domain.
---

##Resource: athenz_sub_domain

`athenz_sub_domain` provides an Athenz sub-domain resource.

###Important Note: Use this resource only for create new sub-domain, update not supported. For import existing one, pls use terraform import.

### Example Usage

```terraform
resource "athenz_sub_domain" "sub_domain-test" {
parent_name="home.some_user"
name = "test"
admin_users = ["user.someone"]
audit_ref = "create domain"
}
```

### Argument Reference

The following arguments are supported:

- `parnet_name` - (Required) name of the parent domain.


- `name` - (Required) name of the domain.


- `admin_users` - (Required) list of domain administrators. must be in this format: `user.<userid/> or <domain/>.<service/>`.


            - `audit_ref` - (Optional Default = "done by terraform provider")  string containing audit specification or ticket number.

