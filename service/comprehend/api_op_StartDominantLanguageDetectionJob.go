// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous dominant language detection job for a collection of
// documents. Use the operation to track the status of a job.
func (c *Client) StartDominantLanguageDetectionJob(ctx context.Context, params *StartDominantLanguageDetectionJobInput, optFns ...func(*Options)) (*StartDominantLanguageDetectionJobOutput, error) {
	if params == nil {
		params = &StartDominantLanguageDetectionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDominantLanguageDetectionJob", params, optFns, c.addOperationStartDominantLanguageDetectionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDominantLanguageDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDominantLanguageDetectionJobInput struct {

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that grants Amazon Comprehend read access to your input data. For more
	// information, see
	// https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions.html#auth-role-permissions
	// (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions.html#auth-role-permissions).
	//
	// This member is required.
	DataAccessRoleArn *string

	// Specifies the format and location of the input data for the job.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// Specifies where to send the output files.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// A unique identifier for the request. If you do not set the client request token,
	// Amazon Comprehend generates one.
	ClientRequestToken *string

	// An identifier for the job.
	JobName *string

	// Tags to be associated with the dominant language detection job. A tag is a
	// key-value pair that adds metadata to a resource used by Amazon Comprehend. For
	// example, a tag with "Sales" as the key might be added to a resource to indicate
	// its use by the sales department.
	Tags []types.Tag

	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to
	// encrypt data on the storage volume attached to the ML compute instance(s) that
	// process the analysis job. The VolumeKmsKeyId can be either of the following
	// formats:
	//
	// * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	// * Amazon
	// Resource Name (ARN) of a KMS Key:
	// "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	VolumeKmsKeyId *string

	// Configuration parameters for an optional private Virtual Private Cloud (VPC)
	// containing the resources you are using for your dominant language detection job.
	// For more information, see Amazon VPC
	// (https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type StartDominantLanguageDetectionJobOutput struct {

	// The Amazon Resource Name (ARN) of the dominant language detection job. It is a
	// unique, fully qualified identifier for the job. It includes the AWS account,
	// Region, and the job ID. The format of the ARN is as follows:
	// arn::comprehend:::dominant-language-detection-job/ The following is an example
	// job ARN:
	// arn:aws:comprehend:us-west-2:111122223333:dominant-language-detection-job/1234abcd12ab34cd56ef1234567890ab
	JobArn *string

	// The identifier generated for the job. To get the status of a job, use this
	// identifier with the operation.
	JobId *string

	// The status of the job.
	//
	// * SUBMITTED - The job has been received and is queued
	// for processing.
	//
	// * IN_PROGRESS - Amazon Comprehend is processing the job.
	//
	// *
	// COMPLETED - The job was successfully completed and the output is available.
	//
	// *
	// FAILED - The job did not complete. To get details, use the operation.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDominantLanguageDetectionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartDominantLanguageDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDominantLanguageDetectionJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartDominantLanguageDetectionJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartDominantLanguageDetectionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDominantLanguageDetectionJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartDominantLanguageDetectionJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartDominantLanguageDetectionJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartDominantLanguageDetectionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartDominantLanguageDetectionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartDominantLanguageDetectionJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartDominantLanguageDetectionJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartDominantLanguageDetectionJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartDominantLanguageDetectionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "StartDominantLanguageDetectionJob",
	}
}
