// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	DisableHTTPS bool
}

// Resolver EFS endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &aws.MissingRegionError{}
	}

	opt := endpoints.Options{
		DisableHTTPS: options.DisableHTTPS,
	}
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var partitionRegexp = struct {
	Aws      *regexp.Regexp
	AwsCn    *regexp.Regexp
	AwsIso   *regexp.Regexp
	AwsIsoB  *regexp.Regexp
	AwsUsGov *regexp.Regexp
}{

	Aws:      regexp.MustCompile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$"),
	AwsCn:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
	AwsIso:   regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
	AwsIsoB:  regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
	AwsUsGov: regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: endpoints.Endpoint{
			Hostname:          "elasticfilesystem.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.Aws,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"af-south-1":     endpoints.Endpoint{},
			"ap-east-1":      endpoints.Endpoint{},
			"ap-northeast-1": endpoints.Endpoint{},
			"ap-northeast-2": endpoints.Endpoint{},
			"ap-northeast-3": endpoints.Endpoint{},
			"ap-south-1":     endpoints.Endpoint{},
			"ap-southeast-1": endpoints.Endpoint{},
			"ap-southeast-2": endpoints.Endpoint{},
			"ca-central-1":   endpoints.Endpoint{},
			"eu-central-1":   endpoints.Endpoint{},
			"eu-north-1":     endpoints.Endpoint{},
			"eu-south-1":     endpoints.Endpoint{},
			"eu-west-1":      endpoints.Endpoint{},
			"eu-west-2":      endpoints.Endpoint{},
			"eu-west-3":      endpoints.Endpoint{},
			"fips-af-south-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.af-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "af-south-1",
				},
			},
			"fips-ap-east-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.ap-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-east-1",
				},
			},
			"fips-ap-northeast-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.ap-northeast-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-northeast-1",
				},
			},
			"fips-ap-northeast-2": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.ap-northeast-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-northeast-2",
				},
			},
			"fips-ap-northeast-3": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.ap-northeast-3.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-northeast-3",
				},
			},
			"fips-ap-south-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.ap-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-south-1",
				},
			},
			"fips-ap-southeast-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.ap-southeast-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-southeast-1",
				},
			},
			"fips-ap-southeast-2": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.ap-southeast-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-southeast-2",
				},
			},
			"fips-ca-central-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.ca-central-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ca-central-1",
				},
			},
			"fips-eu-central-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.eu-central-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-central-1",
				},
			},
			"fips-eu-north-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.eu-north-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-north-1",
				},
			},
			"fips-eu-south-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.eu-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-south-1",
				},
			},
			"fips-eu-west-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.eu-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-1",
				},
			},
			"fips-eu-west-2": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.eu-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-2",
				},
			},
			"fips-eu-west-3": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.eu-west-3.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-3",
				},
			},
			"fips-me-south-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.me-south-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "me-south-1",
				},
			},
			"fips-sa-east-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.sa-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "sa-east-1",
				},
			},
			"fips-us-east-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			"fips-us-east-2": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.us-east-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
			},
			"fips-us-west-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.us-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
			},
			"fips-us-west-2": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.us-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
			},
			"me-south-1": endpoints.Endpoint{},
			"sa-east-1":  endpoints.Endpoint{},
			"us-east-1":  endpoints.Endpoint{},
			"us-east-2":  endpoints.Endpoint{},
			"us-west-1":  endpoints.Endpoint{},
			"us-west-2":  endpoints.Endpoint{},
		},
	},
	{
		ID: "aws-cn",
		Defaults: endpoints.Endpoint{
			Hostname:          "elasticfilesystem.{region}.amazonaws.com.cn",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsCn,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"cn-north-1":     endpoints.Endpoint{},
			"cn-northwest-1": endpoints.Endpoint{},
			"fips-cn-north-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.cn-north-1.amazonaws.com.cn",
				CredentialScope: endpoints.CredentialScope{
					Region: "cn-north-1",
				},
			},
			"fips-cn-northwest-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.cn-northwest-1.amazonaws.com.cn",
				CredentialScope: endpoints.CredentialScope{
					Region: "cn-northwest-1",
				},
			},
		},
	},
	{
		ID: "aws-iso",
		Defaults: endpoints.Endpoint{
			Hostname:          "elasticfilesystem.{region}.c2s.ic.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsIso,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"fips-us-iso-east-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.us-iso-east-1.c2s.ic.gov",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-iso-east-1",
				},
			},
			"us-iso-east-1": endpoints.Endpoint{},
		},
	},
	{
		ID: "aws-iso-b",
		Defaults: endpoints.Endpoint{
			Hostname:          "elasticfilesystem.{region}.sc2s.sgov.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsIsoB,
		IsRegionalized: true,
	},
	{
		ID: "aws-us-gov",
		Defaults: endpoints.Endpoint{
			Hostname:          "elasticfilesystem.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsUsGov,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"fips-us-gov-east-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.us-gov-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-east-1",
				},
			},
			"fips-us-gov-west-1": endpoints.Endpoint{
				Hostname: "elasticfilesystem-fips.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
			"us-gov-east-1": endpoints.Endpoint{},
			"us-gov-west-1": endpoints.Endpoint{},
		},
	},
}
