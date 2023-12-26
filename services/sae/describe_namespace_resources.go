package sae

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

// DescribeNamespaceResources invokes the sae.DescribeNamespaceResources API synchronously
func (client *Client) DescribeNamespaceResources(request *DescribeNamespaceResourcesRequest) (response *DescribeNamespaceResourcesResponse, err error) {
	response = CreateDescribeNamespaceResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNamespaceResourcesWithChan invokes the sae.DescribeNamespaceResources API asynchronously
func (client *Client) DescribeNamespaceResourcesWithChan(request *DescribeNamespaceResourcesRequest) (<-chan *DescribeNamespaceResourcesResponse, <-chan error) {
	responseChan := make(chan *DescribeNamespaceResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNamespaceResources(request)
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

// DescribeNamespaceResourcesWithCallback invokes the sae.DescribeNamespaceResources API asynchronously
func (client *Client) DescribeNamespaceResourcesWithCallback(request *DescribeNamespaceResourcesRequest, callback func(response *DescribeNamespaceResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNamespaceResourcesResponse
		var err error
		defer close(result)
		response, err = client.DescribeNamespaceResources(request)
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

// DescribeNamespaceResourcesRequest is the request struct for api DescribeNamespaceResources
type DescribeNamespaceResourcesRequest struct {
	*requests.RoaRequest
	NamespaceId      string `position:"Query" name:"NamespaceId"`
	NameSpaceShortId string `position:"Query" name:"NameSpaceShortId"`
}

// DescribeNamespaceResourcesResponse is the response struct for api DescribeNamespaceResources
type DescribeNamespaceResourcesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeNamespaceResourcesRequest creates a request to invoke DescribeNamespaceResources API
func CreateDescribeNamespaceResourcesRequest() (request *DescribeNamespaceResourcesRequest) {
	request = &DescribeNamespaceResourcesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeNamespaceResources", "/pop/v1/sam/namespace/describeNamespaceResources", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeNamespaceResourcesResponse creates a response to parse from DescribeNamespaceResources response
func CreateDescribeNamespaceResourcesResponse() (response *DescribeNamespaceResourcesResponse) {
	response = &DescribeNamespaceResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
