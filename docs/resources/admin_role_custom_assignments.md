---
page_title: "Resource: okta_admin_role_custom_assignments"
description: |-
  Resource to manage the assignment and unassignment of Custom Roles
---

# Resource: okta_admin_role_custom_assignments

Resource to manage the assignment and unassignment of Custom Roles



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `custom_role_id` (String) ID of the Custom Role
- `resource_set_id` (String) ID of the target Resource Set

### Optional

- `members` (Set of String) The hrefs that point to User(s) and/or Group(s) that receive the Role

### Read-Only

- `id` (String) The ID of this resource.

