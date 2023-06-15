// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the URL of an assessment report in Audit Manager.
func (c *Client) GetAssessmentReportUrl(ctx context.Context, params *GetAssessmentReportUrlInput, optFns ...func(*Options)) (*GetAssessmentReportUrlOutput, error) {
	if params == nil {
		params = &GetAssessmentReportUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssessmentReportUrl", params, optFns, c.addOperationGetAssessmentReportUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssessmentReportUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssessmentReportUrlInput struct {

	// The unique identifier for the assessment.
	//
	// This member is required.
	AssessmentId *string

	// The unique identifier for the assessment report.
	//
	// This member is required.
	AssessmentReportId *string

	noSmithyDocumentSerde
}

type GetAssessmentReportUrlOutput struct {

	// Short for uniform resource locator. A URL is used as a unique identifier to
	// locate a resource on the internet.
	PreSignedUrl *types.URL

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAssessmentReportUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAssessmentReportUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAssessmentReportUrl{}, middleware.After)
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
	if err = addOpGetAssessmentReportUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssessmentReportUrl(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetAssessmentReportUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "GetAssessmentReportUrl",
	}
}
