// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deprovisions your Autonomous System Number (ASN) from your Amazon Web Services
// account. This action can only be called after any BYOIP CIDR associations are
// removed from your Amazon Web Services account with DisassociateIpamByoasn (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DisassociateIpamByoasn.html)
// . For more information, see Tutorial: Bring your ASN to IPAM (https://docs.aws.amazon.com/vpc/latest/ipam/tutorials-byoasn.html)
// in the Amazon VPC IPAM guide.
func (c *Client) DeprovisionIpamByoasn(ctx context.Context, params *DeprovisionIpamByoasnInput, optFns ...func(*Options)) (*DeprovisionIpamByoasnOutput, error) {
	if params == nil {
		params = &DeprovisionIpamByoasnInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeprovisionIpamByoasn", params, optFns, c.addOperationDeprovisionIpamByoasnMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeprovisionIpamByoasnOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeprovisionIpamByoasnInput struct {

	// An ASN.
	//
	// This member is required.
	Asn *string

	// The IPAM ID.
	//
	// This member is required.
	IpamId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type DeprovisionIpamByoasnOutput struct {

	// An ASN and BYOIP CIDR association.
	Byoasn *types.Byoasn

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeprovisionIpamByoasnMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeprovisionIpamByoasn{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeprovisionIpamByoasn{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeprovisionIpamByoasn"); err != nil {
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
	if err = addOpDeprovisionIpamByoasnValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeprovisionIpamByoasn(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeprovisionIpamByoasn(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeprovisionIpamByoasn",
	}
}
