// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviLicenseStatus() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviLicenseStatusRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"essentials_enforced_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"saas_status": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSaasLicensingStatusSchema(),
			},
			"service_update": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceLicenseServiceUpdateSchema(),
			},
			"tenant_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
