// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified account from the organization. The removed account
// becomes a standalone account that isn't a member of any organization. It's no
// longer subject to any policies and is responsible for its own bill payments. The
// organization's management account is no longer charged for any expenses accrued
// by the member account after it's removed from the organization. This operation
// can be called only from the organization's management account. Member accounts
// can remove themselves with LeaveOrganization instead.
//   - You can remove an account from your organization only if the account is
//     configured with the information required to operate as a standalone account.
//     When you create an account in an organization using the Organizations console,
//     API, or CLI commands, the information required of standalone accounts is not
//     automatically collected. For more information, see Considerations before
//     removing an account from an organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_account-before-remove.html)
//     in the Organizations User Guide.
//   - The account that you want to leave must not be a delegated administrator
//     account for any Amazon Web Services service enabled for your organization. If
//     the account is a delegated administrator, you must first change the delegated
//     administrator account to another account that is remaining in the organization.
//   - After the account leaves the organization, all tags that were attached to
//     the account object in the organization are deleted. Amazon Web Services accounts
//     outside of an organization do not support tags.
func (c *Client) RemoveAccountFromOrganization(ctx context.Context, params *RemoveAccountFromOrganizationInput, optFns ...func(*Options)) (*RemoveAccountFromOrganizationOutput, error) {
	if params == nil {
		params = &RemoveAccountFromOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveAccountFromOrganization", params, optFns, c.addOperationRemoveAccountFromOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveAccountFromOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveAccountFromOrganizationInput struct {

	// The unique identifier (ID) of the member account that you want to remove from
	// the organization. The regex pattern (http://wikipedia.org/wiki/regex) for an
	// account ID string requires exactly 12 digits.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

type RemoveAccountFromOrganizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveAccountFromOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveAccountFromOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveAccountFromOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RemoveAccountFromOrganization"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRemoveAccountFromOrganizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveAccountFromOrganization(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opRemoveAccountFromOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RemoveAccountFromOrganization",
	}
}
