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

// DescribeEnsResourceUsage invokes the ens.DescribeEnsResourceUsage API synchronously
func (client *Client) DescribeEnsResourceUsage(request *DescribeEnsResourceUsageRequest) (response *DescribeEnsResourceUsageResponse, err error) {
	response = CreateDescribeEnsResourceUsageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEnsResourceUsageWithChan invokes the ens.DescribeEnsResourceUsage API asynchronously
func (client *Client) DescribeEnsResourceUsageWithChan(request *DescribeEnsResourceUsageRequest) (<-chan *DescribeEnsResourceUsageResponse, <-chan error) {
	responseChan := make(chan *DescribeEnsResourceUsageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEnsResourceUsage(request)
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

// DescribeEnsResourceUsageWithCallback invokes the ens.DescribeEnsResourceUsage API asynchronously
func (client *Client) DescribeEnsResourceUsageWithCallback(request *DescribeEnsResourceUsageRequest, callback func(response *DescribeEnsResourceUsageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEnsResourceUsageResponse
		var err error
		defer close(result)
		response, err = client.DescribeEnsResourceUsage(request)
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

// DescribeEnsResourceUsageRequest is the request struct for api DescribeEnsResourceUsage
type DescribeEnsResourceUsageRequest struct {
	*requests.RpcRequest
	ExpiredStartTime string `position:"Query" name:"ExpiredStartTime"`
	ExpiredEndTime   string `position:"Query" name:"ExpiredEndTime"`
}

// DescribeEnsResourceUsageResponse is the response struct for api DescribeEnsResourceUsage
type DescribeEnsResourceUsageResponse struct {
	*responses.BaseResponse
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	EnsResourceUsage []EnsResourceUsageItem `json:"EnsResourceUsage" xml:"EnsResourceUsage"`
}

// CreateDescribeEnsResourceUsageRequest creates a request to invoke DescribeEnsResourceUsage API
func CreateDescribeEnsResourceUsageRequest() (request *DescribeEnsResourceUsageRequest) {
	request = &DescribeEnsResourceUsageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeEnsResourceUsage", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeEnsResourceUsageResponse creates a response to parse from DescribeEnsResourceUsage response
func CreateDescribeEnsResourceUsageResponse() (response *DescribeEnsResourceUsageResponse) {
	response = &DescribeEnsResourceUsageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
