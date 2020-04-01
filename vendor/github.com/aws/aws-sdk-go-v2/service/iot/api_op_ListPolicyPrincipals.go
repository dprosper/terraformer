// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the ListPolicyPrincipals operation.
type ListPolicyPrincipalsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the order for results. If true, the results are returned in ascending
	// creation order.
	AscendingOrder *bool `location:"querystring" locationName:"isAscendingOrder" type:"boolean"`

	// The marker for the next set of results.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The result page size.
	PageSize *int64 `location:"querystring" locationName:"pageSize" min:"1" type:"integer"`

	// The policy name.
	//
	// PolicyName is a required field
	PolicyName *string `location:"header" locationName:"x-amzn-iot-policy" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListPolicyPrincipalsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPolicyPrincipalsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPolicyPrincipalsInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPolicyPrincipalsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PolicyName != nil {
		v := *s.PolicyName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amzn-iot-policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AscendingOrder != nil {
		v := *s.AscendingOrder

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "isAscendingOrder", protocol.BoolValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "pageSize", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The output from the ListPolicyPrincipals operation.
type ListPolicyPrincipalsOutput struct {
	_ struct{} `type:"structure"`

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string `locationName:"nextMarker" type:"string"`

	// The descriptions of the principals.
	Principals []string `locationName:"principals" type:"list"`
}

// String returns the string representation
func (s ListPolicyPrincipalsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPolicyPrincipalsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Principals != nil {
		v := s.Principals

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "principals", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opListPolicyPrincipals = "ListPolicyPrincipals"

// ListPolicyPrincipalsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the principals associated with the specified policy.
//
// Note: This API is deprecated. Please use ListTargetsForPolicy instead.
//
//    // Example sending a request using ListPolicyPrincipalsRequest.
//    req := client.ListPolicyPrincipalsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListPolicyPrincipalsRequest(input *ListPolicyPrincipalsInput) ListPolicyPrincipalsRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, ListPolicyPrincipals, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opListPolicyPrincipals,
		HTTPMethod: "GET",
		HTTPPath:   "/policy-principals",
	}

	if input == nil {
		input = &ListPolicyPrincipalsInput{}
	}

	req := c.newRequest(op, input, &ListPolicyPrincipalsOutput{})
	return ListPolicyPrincipalsRequest{Request: req, Input: input, Copy: c.ListPolicyPrincipalsRequest}
}

// ListPolicyPrincipalsRequest is the request type for the
// ListPolicyPrincipals API operation.
type ListPolicyPrincipalsRequest struct {
	*aws.Request
	Input *ListPolicyPrincipalsInput
	Copy  func(*ListPolicyPrincipalsInput) ListPolicyPrincipalsRequest
}

// Send marshals and sends the ListPolicyPrincipals API request.
func (r ListPolicyPrincipalsRequest) Send(ctx context.Context) (*ListPolicyPrincipalsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPolicyPrincipalsResponse{
		ListPolicyPrincipalsOutput: r.Request.Data.(*ListPolicyPrincipalsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPolicyPrincipalsResponse is the response type for the
// ListPolicyPrincipals API operation.
type ListPolicyPrincipalsResponse struct {
	*ListPolicyPrincipalsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPolicyPrincipals request.
func (r *ListPolicyPrincipalsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}