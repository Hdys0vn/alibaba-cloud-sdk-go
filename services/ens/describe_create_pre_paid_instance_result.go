package ens

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

// DescribeCreatePrePaidInstanceResult invokes the ens.DescribeCreatePrePaidInstanceResult API synchronously
func (client *Client) DescribeCreatePrePaidInstanceResult(request *DescribeCreatePrePaidInstanceResultRequest) (response *DescribeCreatePrePaidInstanceResultResponse, err error) {
	response = CreateDescribeCreatePrePaidInstanceResultResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCreatePrePaidInstanceResultWithChan invokes the ens.DescribeCreatePrePaidInstanceResult API asynchronously
func (client *Client) DescribeCreatePrePaidInstanceResultWithChan(request *DescribeCreatePrePaidInstanceResultRequest) (<-chan *DescribeCreatePrePaidInstanceResultResponse, <-chan error) {
	responseChan := make(chan *DescribeCreatePrePaidInstanceResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCreatePrePaidInstanceResult(request)
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

// DescribeCreatePrePaidInstanceResultWithCallback invokes the ens.DescribeCreatePrePaidInstanceResult API asynchronously
func (client *Client) DescribeCreatePrePaidInstanceResultWithCallback(request *DescribeCreatePrePaidInstanceResultRequest, callback func(response *DescribeCreatePrePaidInstanceResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCreatePrePaidInstanceResultResponse
		var err error
		defer close(result)
		response, err = client.DescribeCreatePrePaidInstanceResult(request)
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

// DescribeCreatePrePaidInstanceResultRequest is the request struct for api DescribeCreatePrePaidInstanceResult
type DescribeCreatePrePaidInstanceResultRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeCreatePrePaidInstanceResultResponse is the response struct for api DescribeCreatePrePaidInstanceResult
type DescribeCreatePrePaidInstanceResultResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	InstanceCreateResult InstanceCreateResult `json:"InstanceCreateResult" xml:"InstanceCreateResult"`
}

// CreateDescribeCreatePrePaidInstanceResultRequest creates a request to invoke DescribeCreatePrePaidInstanceResult API
func CreateDescribeCreatePrePaidInstanceResultRequest() (request *DescribeCreatePrePaidInstanceResultRequest) {
	request = &DescribeCreatePrePaidInstanceResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeCreatePrePaidInstanceResult", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCreatePrePaidInstanceResultResponse creates a response to parse from DescribeCreatePrePaidInstanceResult response
func CreateDescribeCreatePrePaidInstanceResultResponse() (response *DescribeCreatePrePaidInstanceResultResponse) {
	response = &DescribeCreatePrePaidInstanceResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
