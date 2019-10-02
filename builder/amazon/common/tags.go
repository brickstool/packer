package common

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
	"github.com/zclconf/go-cty/cty"
)

type TagMap map[string]string
type EC2Tags []*ec2.Tag

func (t EC2Tags) Report(ui packer.Ui) {
	for _, tag := range t {
		ui.Message(fmt.Sprintf("Adding tag: \"%s\": \"%s\"",
			aws.StringValue(tag.Key), aws.StringValue(tag.Value)))
	}
}

func (t TagMap) IsSet() bool {
	return len(t) > 0
}

func (t TagMap) EC2Tags(ictx interpolate.Context, region string, state multistep.StateBag) (EC2Tags, error) {
	var ec2Tags []*ec2.Tag
	ictx.Data = extractBuildInfo(region, state)

	for key, value := range t {
		interpolatedKey, err := interpolate.Render(key, &ictx)
		if err != nil {
			return nil, fmt.Errorf("Error processing tag: %s:%s - %s", key, value, err)
		}
		interpolatedValue, err := interpolate.Render(value, &ictx)
		if err != nil {
			return nil, fmt.Errorf("Error processing tag: %s:%s - %s", key, value, err)
		}
		ec2Tags = append(ec2Tags, &ec2.Tag{
			Key:   aws.String(interpolatedKey),
			Value: aws.String(interpolatedValue),
		})
	}
	return ec2Tags, nil
}

func (*TagMap) HCL2Spec() *hcldec.BlockAttrsSpec {
	return &hcldec.BlockAttrsSpec{
		TypeName:    "TagMap",
		ElementType: cty.String,
		Required:    false,
	}
}
