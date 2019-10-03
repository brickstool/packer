// Code generated by "hcl2-schema -type Config,nicConfig,diskConfig"; DO NOT EDIT.\n

package proxmox

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"RawBootKeyInterval":  &hcldec.AttrSpec{Name: "boot_key_interval", Type: cty.String, Required: false},
		"BootKeyInterval":     &hcldec.AttrSpec{Name: "boot_key_interval", Type: cty.String, Required: false},
		"ProxmoxURLRaw":       &hcldec.AttrSpec{Name: "proxmox_url", Type: cty.String, Required: false},
		"ProxmoxURL":          &hcldec.AttrSpec{Name: "proxmox_url", Type: cty.String, Required: false},
		"SkipCertValidation":  &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"Username":            &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"Password":            &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"Node":                &hcldec.AttrSpec{Name: "node", Type: cty.String, Required: false},
		"Pool":                &hcldec.AttrSpec{Name: "pool", Type: cty.String, Required: false},
		"VMName":              &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"VMID":                &hcldec.AttrSpec{Name: "vm_id", Type: cty.Number, Required: false},
		"Memory":              &hcldec.AttrSpec{Name: "memory", Type: cty.Number, Required: false},
		"Cores":               &hcldec.AttrSpec{Name: "cores", Type: cty.Number, Required: false},
		"Sockets":             &hcldec.AttrSpec{Name: "sockets", Type: cty.Number, Required: false},
		"OS":                  &hcldec.AttrSpec{Name: "os", Type: cty.String, Required: false},
		"ISOFile":             &hcldec.AttrSpec{Name: "iso_file", Type: cty.String, Required: false},
		"Agent":               &hcldec.AttrSpec{Name: "qemu_agent", Type: cty.Bool, Required: false},
		"TemplateName":        &hcldec.AttrSpec{Name: "template_name", Type: cty.String, Required: false},
		"TemplateDescription": &hcldec.AttrSpec{Name: "template_description", Type: cty.String, Required: false},
		"UnmountISO":          &hcldec.AttrSpec{Name: "unmount_iso", Type: cty.Bool, Required: false},
		"network_adapters":    &hcldec.BlockListSpec{TypeName: "[]nicConfig", Nested: hcldec.ObjectSpec((&Config{}).NICs[0].HCL2Spec())},
		"disks":               &hcldec.BlockListSpec{TypeName: "[]diskConfig", Nested: hcldec.ObjectSpec((&Config{}).Disks[0].HCL2Spec())},
	}
	for k, v := range (&Config{}).HTTPConfig.HCL2Spec() {
		s[k] = v
	}
	for k, v := range (&Config{}).BootConfig.HCL2Spec() {
		s[k] = v
	}
	return s
}

func (*nicConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Model":      &hcldec.AttrSpec{Name: "model", Type: cty.String, Required: false},
		"MACAddress": &hcldec.AttrSpec{Name: "mac_address", Type: cty.String, Required: false},
		"Bridge":     &hcldec.AttrSpec{Name: "bridge", Type: cty.String, Required: false},
		"VLANTag":    &hcldec.AttrSpec{Name: "vlan_tag", Type: cty.String, Required: false},
	}
	return s
}

func (*diskConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Type":            &hcldec.AttrSpec{Name: "type", Type: cty.String, Required: false},
		"StoragePool":     &hcldec.AttrSpec{Name: "storage_pool", Type: cty.String, Required: false},
		"StoragePoolType": &hcldec.AttrSpec{Name: "storage_pool_type", Type: cty.String, Required: false},
		"Size":            &hcldec.AttrSpec{Name: "disk_size", Type: cty.String, Required: false},
		"CacheMode":       &hcldec.AttrSpec{Name: "cache_mode", Type: cty.String, Required: false},
		"DiskFormat":      &hcldec.AttrSpec{Name: "format", Type: cty.String, Required: false},
	}
	return s
}
