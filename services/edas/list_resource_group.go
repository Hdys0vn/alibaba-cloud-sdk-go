package edas

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

// ListResourceGroup invokes the edas.ListResourceGroup API synchronously
func (client *Client) ListResourceGroup(request *ListResourceGroupRequest) (response *ListResourceGroupResponse, err error) {
	response = CreateListResourceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListResourceGroupWithChan invokes the edas.ListResourceGroup API asynchronously
func (client *Client) ListResourceGroupWithChan(request *ListResourceGroupRequest) (<-chan *ListResourceGroupResponse, <-chan error) {
	responseChan := make(chan *ListResourceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListResourceGroup(request)
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

// ListResourceGroupWithCallback invokes the edas.ListResourceGroup API asynchronously
func (client *Client) ListResourceGroupWithCallback(request *ListResourceGroupRequest, callback func(response *ListResourceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListResourceGroupResponse
		var err error
		defer close(result)
		response, err = client.ListResourceGroup(request)
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

// ListResourceGroupRequest is the request struct for api ListResourceGroup
type ListResourceGroupRequest struct {
	*requests.RoaRequest
}

// ListResourceGroupResponse is the response struct for api ListResourceGroup
type ListResourceGroupResponse struct {
	*responses.BaseResponse
	Code              int               `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	ResourceGroupList ResourceGroupList `json:"ResourceGroupList" xml:"ResourceGroupList"`
}

// CreateListResourceGroupRequest creates a request to invoke ListResourceGroup API
func CreateListResourceGroupRequest() (request *ListResourceGroupRequest) {
	request = &ListResourceGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListResourceGroup", "/pop/v5/resource/reg_group_list", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListResourceGroupResponse creates a response to parse from ListResourceGroup response
func CreateListResourceGroupResponse() (response *ListResourceGroupResponse) {
	response = &ListResourceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
