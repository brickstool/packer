// Code generated by "hcl2-schema -type ShutdownConfig"; DO NOT EDIT.\n

package shutdowncommand

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*ShutdownConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"ShutdownCommand":    &hcldec.AttrSpec{Name: "shutdown_command", Type: cty.String, Required: false},
		"RawShutdownTimeout": &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"ShutdownTimeout":    &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
	}
	return s
}
