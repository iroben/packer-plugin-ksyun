//go:generate struct-markdown

package kec

import (
	"fmt"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"regexp"
)

type KingcloudDiskDevice struct {
	// The instance needs to create a snapshot ID of the image, which must contain a system disk snapshot ID
	// Can be default: Yes, this parameter cannot be default when creating image based on snapshot
	SnapshotId string `mapstructure:"snapshot_id" required:"false"`
	// The ID of the data disk that the instance needs to mirror
	// Default: Yes
	DataDiskId string `mapstructure:"data_disk_id" required:"false"`
}

type KingcloudDiskDevices struct {
	SnapshotIds []KingcloudDiskDevice `mapstructure:"snapshot_ids" required:"false"`
	DataDiskIds []KingcloudDiskDevice `mapstructure:"data_disk_ids" required:"false"`
}

type KingcloudImageConfig struct {
	// The name of the user-defined image, [2, 64] English or Chinese
	// characters. It must begin with an uppercase/lowercase letter or a
	// Chinese character, and may contain numbers, `_` or `-`. It cannot begin
	// with `http://` or `https://`.
	KingcloudImageName string `mapstructure:"image_name" required:"true"`
	// The type of image
	// LocalImage (ebs) or CommonImage (ks3)
	KingcloudImageType string `mapstructure:"image_type" required:"false"`

	KingcloudDiskDevices `mapstructure:",squash"`
}

func (c *KingcloudImageConfig) Prepare(ctx *interpolate.Context) []error {
	var errs []error
	if c.KingcloudImageName == "" {
		errs = append(errs, fmt.Errorf("image_name must be specified"))
	} else if len(c.KingcloudImageName) < 2 || len(c.KingcloudImageName) > 64 {
		errs = append(errs, fmt.Errorf("image_name must less than 64 letters and more than 1 letters"))
	}
	match, _:=regexp.MatchString("^([\\w-@#.\\p{L}]){2,64}$",c.KingcloudImageName)
	if ! match {
		errs = append(errs, fmt.Errorf("image_name can't matched"))
	}
	return errs
}