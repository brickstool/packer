// Code generated by "hcl2-schema -type Config"; DO NOT EDIT.\n

package saltmasterless

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"SkipBootstrap":     &hcldec.AttrSpec{Name: "skip_bootstrap", Type: cty.Bool, Required: false},
		"BootstrapArgs":     &hcldec.AttrSpec{Name: "bootstrap_args", Type: cty.String, Required: false},
		"DisableSudo":       &hcldec.AttrSpec{Name: "disable_sudo", Type: cty.Bool, Required: false},
		"CustomState":       &hcldec.AttrSpec{Name: "custom_state", Type: cty.String, Required: false},
		"MinionConfig":      &hcldec.AttrSpec{Name: "minion_config", Type: cty.String, Required: false},
		"GrainsFile":        &hcldec.AttrSpec{Name: "grains_file", Type: cty.String, Required: false},
		"LocalStateTree":    &hcldec.AttrSpec{Name: "local_state_tree", Type: cty.String, Required: false},
		"LocalPillarRoots":  &hcldec.AttrSpec{Name: "local_pillar_roots", Type: cty.String, Required: false},
		"RemoteStateTree":   &hcldec.AttrSpec{Name: "remote_state_tree", Type: cty.String, Required: false},
		"RemotePillarRoots": &hcldec.AttrSpec{Name: "remote_pillar_roots", Type: cty.String, Required: false},
		"TempConfigDir":     &hcldec.AttrSpec{Name: "temp_config_dir", Type: cty.String, Required: false},
		"NoExitOnFailure":   &hcldec.AttrSpec{Name: "no_exit_on_failure", Type: cty.Bool, Required: false},
		"LogLevel":          &hcldec.AttrSpec{Name: "log_level", Type: cty.String, Required: false},
		"SaltCallArgs":      &hcldec.AttrSpec{Name: "salt_call_args", Type: cty.String, Required: false},
		"SaltBinDir":        &hcldec.AttrSpec{Name: "salt_bin_dir", Type: cty.String, Required: false},
		"CmdArgs":           &hcldec.AttrSpec{Name: "cmd_args", Type: cty.String, Required: false},
		"GuestOSType":       &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
	}
	return s
}
