// Code generated by "hcl2-schema -type Config,SharedImageGallery,SharedImageGalleryDestination,PlanInformation"; DO NOT EDIT.\n

package arm

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"CaptureNamePrefix":                   &hcldec.AttrSpec{Name: "capture_name_prefix", Type: cty.String, Required: false},
		"CaptureContainerName":                &hcldec.AttrSpec{Name: "capture_container_name", Type: cty.String, Required: false},
		"SharedGalleryTimeout":                &hcldec.AttrSpec{Name: "shared_image_gallery_timeout", Type: cty.String, Required: false},
		"ImagePublisher":                      &hcldec.AttrSpec{Name: "image_publisher", Type: cty.String, Required: false},
		"ImageOffer":                          &hcldec.AttrSpec{Name: "image_offer", Type: cty.String, Required: false},
		"ImageSku":                            &hcldec.AttrSpec{Name: "image_sku", Type: cty.String, Required: false},
		"ImageVersion":                        &hcldec.AttrSpec{Name: "image_version", Type: cty.String, Required: false},
		"ImageUrl":                            &hcldec.AttrSpec{Name: "image_url", Type: cty.String, Required: false},
		"CustomManagedImageResourceGroupName": &hcldec.AttrSpec{Name: "custom_managed_image_resource_group_name", Type: cty.String, Required: false},
		"CustomManagedImageName":              &hcldec.AttrSpec{Name: "custom_managed_image_name", Type: cty.String, Required: false},
		"Location":                            &hcldec.AttrSpec{Name: "location", Type: cty.String, Required: false},
		"VMSize":                              &hcldec.AttrSpec{Name: "vm_size", Type: cty.String, Required: false},
		"ManagedImageResourceGroupName":       &hcldec.AttrSpec{Name: "managed_image_resource_group_name", Type: cty.String, Required: false},
		"ManagedImageName":                    &hcldec.AttrSpec{Name: "managed_image_name", Type: cty.String, Required: false},
		"ManagedImageStorageAccountType":      &hcldec.AttrSpec{Name: "managed_image_storage_account_type", Type: cty.String, Required: false},
		"ManagedImageOSDiskSnapshotName":      &hcldec.AttrSpec{Name: "managed_image_os_disk_snapshot_name", Type: cty.String, Required: false},
		"ManagedImageDataDiskSnapshotPrefix":  &hcldec.AttrSpec{Name: "managed_image_data_disk_snapshot_prefix", Type: cty.String, Required: false},
		"ManagedImageZoneResilient":           &hcldec.AttrSpec{Name: "managed_image_zone_resilient", Type: cty.Bool, Required: false},
		"AzureTags":                           &hcldec.BlockAttrsSpec{TypeName: "azure_tags", ElementType: cty.String, Required: false},
		"ResourceGroupName":                   &hcldec.AttrSpec{Name: "resource_group_name", Type: cty.String, Required: false},
		"StorageAccount":                      &hcldec.AttrSpec{Name: "storage_account", Type: cty.String, Required: false},
		"TempComputeName":                     &hcldec.AttrSpec{Name: "temp_compute_name", Type: cty.String, Required: false},
		"TempResourceGroupName":               &hcldec.AttrSpec{Name: "temp_resource_group_name", Type: cty.String, Required: false},
		"BuildResourceGroupName":              &hcldec.AttrSpec{Name: "build_resource_group_name", Type: cty.String, Required: false},
		"PrivateVirtualNetworkWithPublicIp":   &hcldec.AttrSpec{Name: "private_virtual_network_with_public_ip", Type: cty.Bool, Required: false},
		"VirtualNetworkName":                  &hcldec.AttrSpec{Name: "virtual_network_name", Type: cty.String, Required: false},
		"VirtualNetworkSubnetName":            &hcldec.AttrSpec{Name: "virtual_network_subnet_name", Type: cty.String, Required: false},
		"VirtualNetworkResourceGroupName":     &hcldec.AttrSpec{Name: "virtual_network_resource_group_name", Type: cty.String, Required: false},
		"CustomDataFile":                      &hcldec.AttrSpec{Name: "custom_data_file", Type: cty.String, Required: false},
		"OSType":                              &hcldec.AttrSpec{Name: "os_type", Type: cty.String, Required: false},
		"OSDiskSizeGB":                        &hcldec.AttrSpec{Name: "os_disk_size_gb", Type: cty.Number, Required: false},
		"AdditionalDiskSize":                  &hcldec.AttrSpec{Name: "disk_additional_size", Type: cty.List(cty.Number), Required: false},
		"DiskCachingType":                     &hcldec.AttrSpec{Name: "disk_caching_type", Type: cty.String, Required: false},
		"UserName":                            &hcldec.AttrSpec{Name: "user_name", Type: cty.String, Required: false},
		"Password":                            &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"AsyncResourceGroupDelete":            &hcldec.AttrSpec{Name: "async_resourcegroup_delete", Type: cty.Bool, Required: false},
		"shared_image_gallery":                &hcldec.BlockObjectSpec{TypeName: "SharedImageGallery", LabelNames: []string(nil), Nested: hcldec.ObjectSpec((&Config{}).SharedGallery.HCL2Spec())},
		"shared_image_gallery_destination":    &hcldec.BlockObjectSpec{TypeName: "SharedImageGalleryDestination", LabelNames: []string(nil), Nested: hcldec.ObjectSpec((&Config{}).SharedGalleryDestination.HCL2Spec())},
		"plan_info":                           &hcldec.BlockObjectSpec{TypeName: "PlanInformation", LabelNames: []string(nil), Nested: hcldec.ObjectSpec((&Config{}).PlanInfo.HCL2Spec())},
	}
	for k, v := range (&Config{}).ClientConfig.HCL2Spec() {
		s[k] = v
	}
	return s
}

func (*SharedImageGallery) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Subscription":  &hcldec.AttrSpec{Name: "subscription", Type: cty.String, Required: false},
		"ResourceGroup": &hcldec.AttrSpec{Name: "resource_group", Type: cty.String, Required: false},
		"GalleryName":   &hcldec.AttrSpec{Name: "gallery_name", Type: cty.String, Required: false},
		"ImageName":     &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"ImageVersion":  &hcldec.AttrSpec{Name: "image_version", Type: cty.String, Required: false},
	}
	return s
}

func (*SharedImageGalleryDestination) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"SigDestinationResourceGroup":      &hcldec.AttrSpec{Name: "resource_group", Type: cty.String, Required: false},
		"SigDestinationGalleryName":        &hcldec.AttrSpec{Name: "gallery_name", Type: cty.String, Required: false},
		"SigDestinationImageName":          &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"SigDestinationImageVersion":       &hcldec.AttrSpec{Name: "image_version", Type: cty.String, Required: false},
		"SigDestinationReplicationRegions": &hcldec.AttrSpec{Name: "replication_regions", Type: cty.List(cty.String), Required: false},
	}
	return s
}

func (*PlanInformation) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"PlanName":          &hcldec.AttrSpec{Name: "plan_name", Type: cty.String, Required: false},
		"PlanProduct":       &hcldec.AttrSpec{Name: "plan_product", Type: cty.String, Required: false},
		"PlanPublisher":     &hcldec.AttrSpec{Name: "plan_publisher", Type: cty.String, Required: false},
		"PlanPromotionCode": &hcldec.AttrSpec{Name: "plan_promotion_code", Type: cty.String, Required: false},
	}
	return s
}
