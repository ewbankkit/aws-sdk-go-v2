// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an analytics category. Amazon Transcribe applies the conditions
// specified by your analytics categories to your call analytics jobs. For each
// analytics category, you specify one or more rules. For example, you can specify
// a rule that the customer sentiment was neutral or negative within that category.
// If you start a call analytics job, Amazon Transcribe applies the category to the
// analytics job that you've specified.
func (c *Client) CreateCallAnalyticsCategory(ctx context.Context, params *CreateCallAnalyticsCategoryInput, optFns ...func(*Options)) (*CreateCallAnalyticsCategoryOutput, error) {
	if params == nil {
		params = &CreateCallAnalyticsCategoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCallAnalyticsCategory", params, optFns, c.addOperationCreateCallAnalyticsCategoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCallAnalyticsCategoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCallAnalyticsCategoryInput struct {

	// The name that you choose for your category when you create it.
	//
	// This member is required.
	CategoryName *string

	// To create a category, you must specify between 1 and 20 rules. For each rule,
	// you specify a filter to be applied to the attributes of the call. For example,
	// you can specify a sentiment filter to detect if the customer's sentiment was
	// negative or neutral.
	//
	// This member is required.
	Rules []types.Rule

	noSmithyDocumentSerde
}

type CreateCallAnalyticsCategoryOutput struct {

	// The rules and associated metadata used to create a category.
	CategoryProperties *types.CategoryProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCallAnalyticsCategoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCallAnalyticsCategory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCallAnalyticsCategory{}, middleware.After)
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
	if err = addOpCreateCallAnalyticsCategoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCallAnalyticsCategory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCallAnalyticsCategory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "CreateCallAnalyticsCategory",
	}
}