// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package profitbricks

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"PBUsername":   &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"PBPassword":   &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"PBUrl":        &hcldec.AttrSpec{Name: "url", Type: cty.String, Required: false},
		"Region":       &hcldec.AttrSpec{Name: "location", Type: cty.String, Required: false},
		"Image":        &hcldec.AttrSpec{Name: "image", Type: cty.String, Required: false},
		"SSHKey":       &hcldec.AttrSpec{Name: "ssh_key", Type: cty.String, Required: false},
		"SnapshotName": &hcldec.AttrSpec{Name: "snapshot_name", Type: cty.String, Required: false},
		"DiskSize":     &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"DiskType":     &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"Cores":        &hcldec.AttrSpec{Name: "cores", Type: cty.Number, Required: false},
		"Ram":          &hcldec.AttrSpec{Name: "ram", Type: cty.Number, Required: false},
		"Retries":      &hcldec.AttrSpec{Name: "retries", Type: cty.Number, Required: false},
	}
	return s
}
