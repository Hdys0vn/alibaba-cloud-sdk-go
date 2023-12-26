package ddosbgp

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

// ListOpenedAccessLogInstances invokes the ddosbgp.ListOpenedAccessLogInstances API synchronously
func (client *Client) ListOpenedAccessLogInstances(request *ListOpenedAccessLogInstancesRequest) (response *ListOpenedAccessLogInstancesResponse, err error) {
	response = CreateListOpenedAccessLogInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListOpenedAccessLogInstancesWithChan invokes the ddosbgp.ListOpenedAccessLogInstances API asynchronously
func (client *Client) ListOpenedAccessLogInstancesWithChan(request *ListOpenedAccessLogInstancesRequest) (<-chan *ListOpenedAccessLogInstancesResponse, <-chan error) {
	responseChan := make(chan *ListOpenedAccessLogInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOpenedAccessLogInstances(request)
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

// ListOpenedAccessLogInstancesWithCallback invokes the ddosbgp.ListOpenedAccessLogInstances API asynchronously
func (client *Client) ListOpenedAccessLogInstancesWithCallback(request *ListOpenedAccessLogInstancesRequest, callback func(response *ListOpenedAccessLogInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOpenedAccessLogInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListOpenedAccessLogInstances(request)
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

// ListOpenedAccessLogInstancesRequest is the request struct for api ListOpenedAccessLogInstances
type ListOpenedAccessLogInstancesRequest struct {
	*requests.RpcRequest
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// ListOpenedAccessLogInstancesResponse is the response struct for api ListOpenedAccessLogInstances
type ListOpenedAccessLogInstancesResponse struct {
	*responses.BaseResponse
	TotalCount      int              `json:"TotalCount" xml:"TotalCount"`
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	SlsConfigStatus []OpenedInstance `json:"SlsConfigStatus" xml:"SlsConfigStatus"`
}

// CreateListOpenedAccessLogInstancesRequest creates a request to invoke ListOpenedAccessLogInstances API
func CreateListOpenedAccessLogInstancesRequest() (request *ListOpenedAccessLogInstancesRequest) {
	request = &ListOpenedAccessLogInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "ListOpenedAccessLogInstances", "", "")
	request.Method = requests.POST
	return
}

// CreateListOpenedAccessLogInstancesResponse creates a response to parse from ListOpenedAccessLogInstances response
func CreateListOpenedAccessLogInstancesResponse() (response *ListOpenedAccessLogInstancesResponse) {
	response = &ListOpenedAccessLogInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}