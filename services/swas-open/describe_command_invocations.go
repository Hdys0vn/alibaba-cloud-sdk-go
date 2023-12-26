package swas_open

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeCommandInvocations invokes the swas_open.DescribeCommandInvocations API synchronously
func (client *Client) DescribeCommandInvocations(request *DescribeCommandInvocationsRequest) (response *DescribeCommandInvocationsResponse, err error) {
	response = CreateDescribeCommandInvocationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCommandInvocationsWithChan invokes the swas_open.DescribeCommandInvocations API asynchronously
func (client *Client) DescribeCommandInvocationsWithChan(request *DescribeCommandInvocationsRequest) (<-chan *DescribeCommandInvocationsResponse, <-chan error) {
	responseChan := make(chan *DescribeCommandInvocationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCommandInvocations(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeCommandInvocationsWithCallback invokes the swas_open.DescribeCommandInvocations API asynchronously
func (client *Client) DescribeCommandInvocationsWithCallback(request *DescribeCommandInvocationsRequest, callback func(response *DescribeCommandInvocationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCommandInvocationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCommandInvocations(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeCommandInvocationsRequest is the request struct for api DescribeCommandInvocations
type DescribeCommandInvocationsRequest struct {
	*requests.RpcRequest
	CommandId        string `position:"Query" name:"CommandId"`
	PageNumber       string `position:"Query" name:"PageNumber"`
	PageSize         string `position:"Query" name:"PageSize"`
	InvokeId         string `position:"Query" name:"InvokeId"`
	CommandName      string `position:"Query" name:"CommandName"`
	CommandType      string `position:"Query" name:"CommandType"`
	InstanceId       string `position:"Query" name:"InstanceId"`
	InvocationStatus string `position:"Query" name:"InvocationStatus"`
}

// DescribeCommandInvocationsResponse is the response struct for api DescribeCommandInvocations
type DescribeCommandInvocationsResponse struct {
	*responses.BaseResponse
	RequestId          string    `json:"RequestId" xml:"RequestId"`
	TotalCount         int       `json:"TotalCount" xml:"TotalCount"`
	PageNumber         int       `json:"PageNumber" xml:"PageNumber"`
	PageSize           int       `json:"PageSize" xml:"PageSize"`
	CommandInvocations []Command `json:"CommandInvocations" xml:"CommandInvocations"`
}

// CreateDescribeCommandInvocationsRequest creates a request to invoke DescribeCommandInvocations API
func CreateDescribeCommandInvocationsRequest() (request *DescribeCommandInvocationsRequest) {
	request = &DescribeCommandInvocationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DescribeCommandInvocations", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCommandInvocationsResponse creates a response to parse from DescribeCommandInvocations response
func CreateDescribeCommandInvocationsResponse() (response *DescribeCommandInvocationsResponse) {
	response = &DescribeCommandInvocationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
