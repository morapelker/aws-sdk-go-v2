// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a profiling group.
func (c *Client) CreateProfilingGroup(ctx context.Context, params *CreateProfilingGroupInput, optFns ...func(*Options)) (*CreateProfilingGroupOutput, error) {
	if params == nil {
		params = &CreateProfilingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProfilingGroup", params, optFns, c.addOperationCreateProfilingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProfilingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the createProfiliingGroupRequest.
type CreateProfilingGroupInput struct {

	// Amazon CodeGuru Profiler uses this universally unique identifier (UUID) to
	// prevent the accidental creation of duplicate profiling groups if there are
	// failures and retries.
	//
	// This member is required.
	ClientToken *string

	// The name of the profiling group to create.
	//
	// This member is required.
	ProfilingGroupName *string

	// Specifies whether profiling is enabled or disabled for the created profiling
	// group.
	AgentOrchestrationConfig *types.AgentOrchestrationConfig

	// The compute platform of the profiling group. Use AWSLambda if your application
	// runs on AWS Lambda. Use Default if your application runs on a compute platform
	// that is not AWS Lambda, such an Amazon EC2 instance, an on-premises server, or a
	// different platform. If not specified, Default is used.
	ComputePlatform types.ComputePlatform

	// A list of tags to add to the created profiling group.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The structure representing the createProfilingGroupResponse.
type CreateProfilingGroupOutput struct {

	// The returned ProfilingGroupDescription
	// (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ProfilingGroupDescription.html)
	// object that contains information about the created profiling group.
	//
	// This member is required.
	ProfilingGroup *types.ProfilingGroupDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProfilingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProfilingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProfilingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateProfilingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProfilingGroup(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateProfilingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-profiler",
		OperationName: "CreateProfilingGroup",
	}
}
