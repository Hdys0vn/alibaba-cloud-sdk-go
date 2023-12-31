package servicemesh

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

// DescribeServiceMeshDetail invokes the servicemesh.DescribeServiceMeshDetail API synchronously
func (client *Client) DescribeServiceMeshDetail(request *DescribeServiceMeshDetailRequest) (response *DescribeServiceMeshDetailResponse, err error) {
	response = CreateDescribeServiceMeshDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeServiceMeshDetailWithChan invokes the servicemesh.DescribeServiceMeshDetail API asynchronously
func (client *Client) DescribeServiceMeshDetailWithChan(request *DescribeServiceMeshDetailRequest) (<-chan *DescribeServiceMeshDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeServiceMeshDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeServiceMeshDetail(request)
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

// DescribeServiceMeshDetailWithCallback invokes the servicemesh.DescribeServiceMeshDetail API asynchronously
func (client *Client) DescribeServiceMeshDetailWithCallback(request *DescribeServiceMeshDetailRequest, callback func(response *DescribeServiceMeshDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeServiceMeshDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeServiceMeshDetail(request)
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

// DescribeServiceMeshDetailRequest is the request struct for api DescribeServiceMeshDetail
type DescribeServiceMeshDetailRequest struct {
	*requests.RpcRequest
	ServiceMeshId string `position:"Body" name:"ServiceMeshId"`
}

// DescribeServiceMeshDetailResponse is the response struct for api DescribeServiceMeshDetail
type DescribeServiceMeshDetailResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ServiceMesh ServiceMesh `json:"ServiceMesh" xml:"ServiceMesh"`
}

// CreateDescribeServiceMeshDetailRequest creates a request to invoke DescribeServiceMeshDetail API
func CreateDescribeServiceMeshDetailRequest() (request *DescribeServiceMeshDetailRequest) {
	request = &DescribeServiceMeshDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("servicemesh", "2020-01-11", "DescribeServiceMeshDetail", "servicemesh", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeServiceMeshDetailResponse creates a response to parse from DescribeServiceMeshDetail response
func CreateDescribeServiceMeshDetailResponse() (response *DescribeServiceMeshDetailResponse) {
	response = &DescribeServiceMeshDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
