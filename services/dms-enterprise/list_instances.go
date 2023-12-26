package dms_enterprise

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

// ListInstances invokes the dms_enterprise.ListInstances API synchronously
func (client *Client) ListInstances(request *ListInstancesRequest) (response *ListInstancesResponse, err error) {
	response = CreateListInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstancesWithChan invokes the dms_enterprise.ListInstances API asynchronously
func (client *Client) ListInstancesWithChan(request *ListInstancesRequest) (<-chan *ListInstancesResponse, <-chan error) {
	responseChan := make(chan *ListInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstances(request)
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

// ListInstancesWithCallback invokes the dms_enterprise.ListInstances API asynchronously
func (client *Client) ListInstancesWithCallback(request *ListInstancesRequest, callback func(response *ListInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListInstances(request)
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

// ListInstancesRequest is the request struct for api ListInstances
type ListInstancesRequest struct {
	*requests.RpcRequest
	SearchKey      string           `position:"Query" name:"SearchKey"`
	Tid            requests.Integer `position:"Query" name:"Tid"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	EnvType        string           `position:"Query" name:"EnvType"`
	InstanceSource string           `position:"Query" name:"InstanceSource"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	InstanceState  string           `position:"Query" name:"InstanceState"`
	NetType        string           `position:"Query" name:"NetType"`
	DbType         string           `position:"Query" name:"DbType"`
}

// ListInstancesResponse is the response struct for api ListInstances
type ListInstancesResponse struct {
	*responses.BaseResponse
	TotalCount   int64        `json:"TotalCount" xml:"TotalCount"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ErrorCode    string       `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string       `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool         `json:"Success" xml:"Success"`
	InstanceList InstanceList `json:"InstanceList" xml:"InstanceList"`
}

// CreateListInstancesRequest creates a request to invoke ListInstances API
func CreateListInstancesRequest() (request *ListInstancesRequest) {
	request = &ListInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListInstances", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInstancesResponse creates a response to parse from ListInstances response
func CreateListInstancesResponse() (response *ListInstancesResponse) {
	response = &ListInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}