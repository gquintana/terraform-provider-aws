// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package bcmdataexports

import (
	"context"
	"fmt"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bcmdataexports"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
)

var _ bcmdataexports.EndpointResolverV2 = resolverSDKv2{}

type resolverSDKv2 struct {
	defaultResolver bcmdataexports.EndpointResolverV2
}

func newEndpointResolverSDKv2() resolverSDKv2 {
	return resolverSDKv2{
		defaultResolver: bcmdataexports.NewDefaultEndpointResolverV2(),
	}
}

func (r resolverSDKv2) ResolveEndpoint(ctx context.Context, params bcmdataexports.EndpointParameters) (endpoint smithyendpoints.Endpoint, err error) {
	params = params.WithDefaults()
	useFIPS := aws.ToBool(params.UseFIPS)

	if eps := params.Endpoint; aws.ToString(eps) != "" {
		tflog.Debug(ctx, "setting endpoint", map[string]any{
			"tf_aws.endpoint": endpoint,
		})

		if useFIPS {
			tflog.Debug(ctx, "endpoint set, ignoring UseFIPSEndpoint setting")
			params.UseFIPS = aws.Bool(false)
		}

		return r.defaultResolver.ResolveEndpoint(ctx, params)
	} else if useFIPS {
		ctx = tflog.SetField(ctx, "tf_aws.use_fips", useFIPS)

		endpoint, err = r.defaultResolver.ResolveEndpoint(ctx, params)
		if err != nil {
			return endpoint, err
		}

		tflog.Debug(ctx, "endpoint resolved", map[string]any{
			"tf_aws.endpoint": endpoint.URI.String(),
		})

		hostname := endpoint.URI.Hostname()
		_, err = net.LookupHost(hostname)
		if err != nil {
			if dnsErr, ok := errs.As[*net.DNSError](err); ok && dnsErr.IsNotFound {
				tflog.Debug(ctx, "default endpoint host not found, disabling FIPS", map[string]any{
					"tf_aws.hostname": hostname,
				})
				params.UseFIPS = aws.Bool(false)
			} else {
				err = fmt.Errorf("looking up bcmdataexports endpoint %q: %s", hostname, err)
				return
			}
		} else {
			return endpoint, err
		}
	}

	return r.defaultResolver.ResolveEndpoint(ctx, params)
}

func withBaseEndpoint(endpoint string) func(*bcmdataexports.Options) {
	return func(o *bcmdataexports.Options) {
		if endpoint != "" {
			o.BaseEndpoint = aws.String(endpoint)
		}
	}
}
