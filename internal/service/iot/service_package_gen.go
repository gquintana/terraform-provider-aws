// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceEndpoint,
			TypeName: "aws_iot_endpoint",
			Name:     "Endpoint",
		},
		{
			Factory:  dataSourceRegistrationCode,
			TypeName: "aws_iot_registration_code",
			Name:     "Registration Code",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAuthorizer,
			TypeName: "aws_iot_authorizer",
			Name:     "Authorizer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceBillingGroup,
			TypeName: "aws_iot_billing_group",
			Name:     "Billing Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCACertificate,
			TypeName: "aws_iot_ca_certificate",
			Name:     "CA Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCertificate,
			TypeName: "aws_iot_certificate",
			Name:     "Certificate",
		},
		{
			Factory:  resourceDomainConfiguration,
			TypeName: "aws_iot_domain_configuration",
			Name:     "Domain Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceEventConfigurations,
			TypeName: "aws_iot_event_configurations",
			Name:     "Event Configurations",
		},
		{
			Factory:  resourceIndexingConfiguration,
			TypeName: "aws_iot_indexing_configuration",
			Name:     "Indexing Configuration",
		},
		{
			Factory:  resourceLoggingOptions,
			TypeName: "aws_iot_logging_options",
			Name:     "Logging Options",
		},
		{
			Factory:  resourcePolicy,
			TypeName: "aws_iot_policy",
			Name:     "Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourcePolicyAttachment,
			TypeName: "aws_iot_policy_attachment",
		},
		{
			Factory:  resourceProvisioningTemplate,
			TypeName: "aws_iot_provisioning_template",
			Name:     "Provisioning Template",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceRoleAlias,
			TypeName: "aws_iot_role_alias",
			Name:     "Role Alias",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceThing,
			TypeName: "aws_iot_thing",
			Name:     "Thing",
		},
		{
			Factory:  resourceThingGroup,
			TypeName: "aws_iot_thing_group",
			Name:     "Thing Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceThingGroupMembership,
			TypeName: "aws_iot_thing_group_membership",
			Name:     "Thing Group Membership",
		},
		{
			Factory:  resourceThingPrincipalAttachment,
			TypeName: "aws_iot_thing_principal_attachment",
			Name:     "Thing Principal Attachment",
		},
		{
			Factory:  resourceThingType,
			TypeName: "aws_iot_thing_type",
			Name:     "Thing Type",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTopicRule,
			TypeName: "aws_iot_topic_rule",
			Name:     "Topic Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTopicRuleDestination,
			TypeName: "aws_iot_topic_rule_destination",
			Name:     "Topic Rule Destination",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.IoT
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*iot.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return iot.NewFromConfig(cfg,
		iot.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
