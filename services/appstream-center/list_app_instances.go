package appstream_center

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

// ListAppInstances invokes the appstream_center.ListAppInstances API synchronously
func (client *Client) ListAppInstances(request *ListAppInstancesRequest) (response *ListAppInstancesResponse, err error) {
	response = CreateListAppInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppInstancesWithChan invokes the appstream_center.ListAppInstances API asynchronously
func (client *Client) ListAppInstancesWithChan(request *ListAppInstancesRequest) (<-chan *ListAppInstancesResponse, <-chan error) {
	responseChan := make(chan *ListAppInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppInstances(request)
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

// ListAppInstancesWithCallback invokes the appstream_center.ListAppInstances API asynchronously
func (client *Client) ListAppInstancesWithCallback(request *ListAppInstancesRequest, callback func(response *ListAppInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListAppInstances(request)
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

// ListAppInstancesRequest is the request struct for api ListAppInstances
type ListAppInstancesRequest struct {
	*requests.RpcRequest
	PageNumber         requests.Integer `position:"Query" name:"PageNumber"`
	AppInstanceGroupId string           `position:"Query" name:"AppInstanceGroupId"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
	AppInstanceId      string           `position:"Query" name:"AppInstanceId"`
	Status             *[]string        `position:"Body" name:"Status"  type:"Repeated"`
}

// ListAppInstancesResponse is the response struct for api ListAppInstances
type ListAppInstancesResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	TotalCount        int    `json:"TotalCount" xml:"TotalCount"`
	PageSize          int    `json:"PageSize" xml:"PageSize"`
	PageNumber        int    `json:"PageNumber" xml:"PageNumber"`
	AppInstanceModels []Data `json:"AppInstanceModels" xml:"AppInstanceModels"`
}

// CreateListAppInstancesRequest creates a request to invoke ListAppInstances API
func CreateListAppInstancesRequest() (request *ListAppInstancesRequest) {
	request = &ListAppInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("appstream-center", "2021-09-01", "ListAppInstances", "", "")
	request.Method = requests.POST
	return
}

// CreateListAppInstancesResponse creates a response to parse from ListAppInstances response
func CreateListAppInstancesResponse() (response *ListAppInstancesResponse) {
	response = &ListAppInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}