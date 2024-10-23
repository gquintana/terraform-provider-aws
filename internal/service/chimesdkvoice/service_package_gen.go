// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package chimesdkvoice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceGlobalSettings,
			TypeName: "aws_chimesdkvoice_global_settings",
		},
		{
			Factory:  ResourceSipMediaApplication,
			TypeName: "aws_chimesdkvoice_sip_media_application",
			Name:     "Sip Media Application",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceSipRule,
			TypeName: "aws_chimesdkvoice_sip_rule",
			Name:     "Sip Rule",
		},
		{
			Factory:  ResourceVoiceProfileDomain,
			TypeName: "aws_chimesdkvoice_voice_profile_domain",
			Name:     "Voice Profile Domain",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ChimeSDKVoice
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*chimesdkvoice.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return chimesdkvoice.NewFromConfig(cfg,
		chimesdkvoice.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
