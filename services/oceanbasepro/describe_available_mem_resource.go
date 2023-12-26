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

// DescribeAvailableMemResource invokes the oceanbasepro.DescribeAvailableMemResource API synchronously
func (client *Client) DescribeAvailableMemResource(request *DescribeAvailableMemResourceRequest) (response *DescribeAvailableMemResourceResponse, err error) {
	response = CreateDescribeAvailableMemResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvailableMemResourceWithChan invokes the oceanbasepro.DescribeAvailableMemResource API asynchronously
func (client *Client) DescribeAvailableMemResourceWithChan(request *DescribeAvailableMemResourceRequest) (<-chan *DescribeAvailableMemResourceResponse, <-chan error) {
	responseChan := make(chan *DescribeAvailableMemResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvailableMemResource(request)
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

// DescribeAvailableMemResourceWithCallback invokes the oceanbasepro.DescribeAvailableMemResource API asynchronously
func (client *Client) DescribeAvailableMemResourceWithCallback(request *DescribeAvailableMemResourceRequest, callback func(response *DescribeAvailableMemResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvailableMemResourceResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvailableMemResource(request)
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

// DescribeAvailableMemResourceRequest is the request struct for api DescribeAvailableMemResource
type DescribeAvailableMemResourceRequest struct {
	*requests.RpcRequest
	UnitNum    requests.Integer `position:"Body" name:"UnitNum"`
	CpuNum     requests.Integer `position:"Body" name:"CpuNum"`
	InstanceId string           `position:"Body" name:"InstanceId"`
	TenantId   string           `position:"Body" name:"TenantId"`
}

// DescribeAvailableMemResourceResponse is the response struct for api DescribeAvailableMemResource
type DescribeAvailableMemResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeAvailableMemResourceRequest creates a request to invoke DescribeAvailableMemResource API
func CreateDescribeAvailableMemResourceRequest() (request *DescribeAvailableMemResourceRequest) {
	request = &DescribeAvailableMemResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeAvailableMemResource", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAvailableMemResourceResponse creates a response to parse from DescribeAvailableMemResource response
func CreateDescribeAvailableMemResourceResponse() (response *DescribeAvailableMemResourceResponse) {
	response = &DescribeAvailableMemResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}