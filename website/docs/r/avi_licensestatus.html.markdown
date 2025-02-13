<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_licensestatus"
sidebar_current: "docs-avi-resource-licensestatus"
description: |-
  Creates and manages Avi LicenseStatus.
---

# avi_licensestatus

The LicenseStatus resource allows the creation and management of Avi LicenseStatus

## Example Usage

```hcl
resource "avi_licensestatus" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.3. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `essentials_enforced_at` - (Optional) License enforcement date when we upgrade controller to 31.1.1 and license tier is essential before upgrade. Field introduced in 31.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `saas_status` - (Optional) Saas licensing status. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `service_update` - (Optional) Pulse license service update. Field introduced in 21.1.4. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_uuid` - (Optional) Tenant uuid. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.

