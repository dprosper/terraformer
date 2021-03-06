// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetDeploymentInstance operation.
type GetDeploymentInstanceInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a deployment.
	//
	// DeploymentId is a required field
	DeploymentId *string `locationName:"deploymentId" type:"string" required:"true"`

	// The unique ID of an instance in the deployment group.
	//
	// InstanceId is a required field
	InstanceId *string `locationName:"instanceId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeploymentInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeploymentInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeploymentInstanceInput"}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetDeploymentInstance operation.
type GetDeploymentInstanceOutput struct {
	_ struct{} `type:"structure"`

	// Information about the instance.
	InstanceSummary *InstanceSummary `locationName:"instanceSummary" deprecated:"true" type:"structure"`
}

// String returns the string representation
func (s GetDeploymentInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDeploymentInstance = "GetDeploymentInstance"

// GetDeploymentInstanceRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets information about an instance as part of a deployment.
//
//    // Example sending a request using GetDeploymentInstanceRequest.
//    req := client.GetDeploymentInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetDeploymentInstance
func (c *Client) GetDeploymentInstanceRequest(input *GetDeploymentInstanceInput) GetDeploymentInstanceRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, GetDeploymentInstance, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opGetDeploymentInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDeploymentInstanceInput{}
	}

	req := c.newRequest(op, input, &GetDeploymentInstanceOutput{})
	return GetDeploymentInstanceRequest{Request: req, Input: input, Copy: c.GetDeploymentInstanceRequest}
}

// GetDeploymentInstanceRequest is the request type for the
// GetDeploymentInstance API operation.
type GetDeploymentInstanceRequest struct {
	*aws.Request
	Input *GetDeploymentInstanceInput
	Copy  func(*GetDeploymentInstanceInput) GetDeploymentInstanceRequest
}

// Send marshals and sends the GetDeploymentInstance API request.
func (r GetDeploymentInstanceRequest) Send(ctx context.Context) (*GetDeploymentInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeploymentInstanceResponse{
		GetDeploymentInstanceOutput: r.Request.Data.(*GetDeploymentInstanceOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeploymentInstanceResponse is the response type for the
// GetDeploymentInstance API operation.
type GetDeploymentInstanceResponse struct {
	*GetDeploymentInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeploymentInstance request.
func (r *GetDeploymentInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
