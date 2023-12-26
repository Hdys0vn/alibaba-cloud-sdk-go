package oceanbasepro

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

// DescribeProjectComponents invokes the oceanbasepro.DescribeProjectComponents API synchronously
func (client *Client) DescribeProjectComponents(request *DescribeProjectComponentsRequest) (response *DescribeProjectComponentsResponse, err error) {
	response = CreateDescribeProjectComponentsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProjectComponentsWithChan invokes the oceanbasepro.DescribeProjectComponents API asynchronously
func (client *Client) DescribeProjectComponentsWithChan(request *DescribeProjectComponentsRequest) (<-chan *DescribeProjectComponentsResponse, <-chan error) {
	responseChan := make(chan *DescribeProjectComponentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProjectComponents(request)
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

// DescribeProjectComponentsWithCallback invokes the oceanbasepro.DescribeProjectComponents API asynchronously
func (client *Client) DescribeProjectComponentsWithCallback(request *DescribeProjectComponentsRequest, callback func(response *DescribeProjectComponentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProjectComponentsResponse
		var err error
		defer close(result)
		response, err = client.DescribeProjectComponents(request)
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

// DescribeProjectComponentsRequest is the request struct for api DescribeProjectComponents
type DescribeProjectComponentsRequest struct {
	*requests.RpcRequest
	Id string `position:"Body" name:"Id"`
}

// DescribeProjectComponentsResponse is the response struct for api DescribeProjectComponents
type DescribeProjectComponentsResponse struct {
	*responses.BaseResponse
}

// CreateDescribeProjectComponentsRequest creates a request to invoke DescribeProjectComponents API
func CreateDescribeProjectComponentsRequest() (request *DescribeProjectComponentsRequest) {
	request = &DescribeProjectComponentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeProjectComponents", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeProjectComponentsResponse creates a response to parse from DescribeProjectComponents response
func CreateDescribeProjectComponentsResponse() (response *DescribeProjectComponentsResponse) {
	response = &DescribeProjectComponentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
