// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListNodegroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon EKS cluster that you would like to list node groups
	// in.
	//
	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// The maximum number of node group results returned by ListNodegroups in paginated
	// output. When you use this parameter, ListNodegroups returns only maxResults
	// results in a single page along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another ListNodegroups
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If you don't use this parameter, ListNodegroups returns up to 100 results
	// and a nextToken value if applicable.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated ListNodegroups request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListNodegroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNodegroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNodegroupsInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNodegroupsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListNodegroupsOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListNodegroups request. When the
	// results of a ListNodegroups request exceed maxResults, you can use this value
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of all of the node groups associated with the specified cluster.
	Nodegroups []string `locationName:"nodegroups" type:"list"`
}

// String returns the string representation
func (s ListNodegroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNodegroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Nodegroups != nil {
		v := s.Nodegroups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "nodegroups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opListNodegroups = "ListNodegroups"

// ListNodegroupsRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Lists the Amazon EKS node groups associated with the specified cluster in
// your AWS account in the specified Region.
//
//    // Example sending a request using ListNodegroupsRequest.
//    req := client.ListNodegroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/ListNodegroups
func (c *Client) ListNodegroupsRequest(input *ListNodegroupsInput) ListNodegroupsRequest {
	op := &aws.Operation{
		Name:       opListNodegroups,
		HTTPMethod: "GET",
		HTTPPath:   "/clusters/{name}/node-groups",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListNodegroupsInput{}
	}

	req := c.newRequest(op, input, &ListNodegroupsOutput{})
	return ListNodegroupsRequest{Request: req, Input: input, Copy: c.ListNodegroupsRequest}
}

// ListNodegroupsRequest is the request type for the
// ListNodegroups API operation.
type ListNodegroupsRequest struct {
	*aws.Request
	Input *ListNodegroupsInput
	Copy  func(*ListNodegroupsInput) ListNodegroupsRequest
}

// Send marshals and sends the ListNodegroups API request.
func (r ListNodegroupsRequest) Send(ctx context.Context) (*ListNodegroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListNodegroupsResponse{
		ListNodegroupsOutput: r.Request.Data.(*ListNodegroupsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListNodegroupsRequestPaginator returns a paginator for ListNodegroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListNodegroupsRequest(input)
//   p := eks.NewListNodegroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListNodegroupsPaginator(req ListNodegroupsRequest) ListNodegroupsPaginator {
	return ListNodegroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListNodegroupsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListNodegroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListNodegroupsPaginator struct {
	aws.Pager
}

func (p *ListNodegroupsPaginator) CurrentPage() *ListNodegroupsOutput {
	return p.Pager.CurrentPage().(*ListNodegroupsOutput)
}

// ListNodegroupsResponse is the response type for the
// ListNodegroups API operation.
type ListNodegroupsResponse struct {
	*ListNodegroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListNodegroups request.
func (r *ListNodegroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}