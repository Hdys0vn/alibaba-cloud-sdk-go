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

// DescribeInvocationResult invokes the swas_open.DescribeInvocationResult API synchronously
func (client *Client) DescribeInvocationResult(request *DescribeInvocationResultRequest) (response *DescribeInvocationResultResponse, err error) {
	response = CreateDescribeInvocationResultResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInvocationResultWithChan invokes the swas_open.DescribeInvocationResult API asynchronously
func (client *Client) DescribeInvocationResultWithChan(request *DescribeInvocationResultRequest) (<-chan *DescribeInvocationResultResponse, <-chan error) {
	responseChan := make(chan *DescribeInvocationResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInvocationResult(request)
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

// DescribeInvocationResultWithCallback invokes the swas_open.DescribeInvocationResult API asynchronously
func (client *Client) DescribeInvocationResultWithCallback(request *DescribeInvocationResultRequest, callback func(response *DescribeInvocationResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInvocationResultResponse
		var err error
		defer close(result)
		response, err = client.DescribeInvocationResult(request)
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

// DescribeInvocationResultRequest is the request struct for api DescribeInvocationResult
type DescribeInvocationResultRequest struct {
	*requests.RpcRequest
	InvokeId   string `position:"Query" name:"InvokeId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeInvocationResultResponse is the response struct for api DescribeInvocationResult
type DescribeInvocationResultResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	InvocationResult InvocationResult `json:"InvocationResult" xml:"InvocationResult"`
}

// CreateDescribeInvocationResultRequest creates a request to invoke DescribeInvocationResult API
func CreateDescribeInvocationResultRequest() (request *DescribeInvocationResultRequest) {
	request = &DescribeInvocationResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DescribeInvocationResult", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInvocationResultResponse creates a response to parse from DescribeInvocationResult response
func CreateDescribeInvocationResultResponse() (response *DescribeInvocationResultResponse) {
	response = &DescribeInvocationResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}