package devops_rdc

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

// ListCommonGroup invokes the devops_rdc.ListCommonGroup API synchronously
func (client *Client) ListCommonGroup(request *ListCommonGroupRequest) (response *ListCommonGroupResponse, err error) {
	response = CreateListCommonGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListCommonGroupWithChan invokes the devops_rdc.ListCommonGroup API asynchronously
func (client *Client) ListCommonGroupWithChan(request *ListCommonGroupRequest) (<-chan *ListCommonGroupResponse, <-chan error) {
	responseChan := make(chan *ListCommonGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCommonGroup(request)
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

// ListCommonGroupWithCallback invokes the devops_rdc.ListCommonGroup API asynchronously
func (client *Client) ListCommonGroupWithCallback(request *ListCommonGroupRequest, callback func(response *ListCommonGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCommonGroupResponse
		var err error
		defer close(result)
		response, err = client.ListCommonGroup(request)
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

// ListCommonGroupRequest is the request struct for api ListCommonGroup
type ListCommonGroupRequest struct {
	*requests.RpcRequest
	All          requests.Boolean `position:"Body" name:"All"`
	SmartGroupId string           `position:"Body" name:"SmartGroupId"`
	ProjectId    string           `position:"Body" name:"ProjectId"`
	OrgId        string           `position:"Body" name:"OrgId"`
}

// ListCommonGroupResponse is the response struct for api ListCommonGroup
type ListCommonGroupResponse struct {
	*responses.BaseResponse
	Successful bool          `json:"Successful" xml:"Successful"`
	ErrorCode  string        `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string        `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string        `json:"RequestId" xml:"RequestId"`
	Object     []CommonGroup `json:"Object" xml:"Object"`
}

// CreateListCommonGroupRequest creates a request to invoke ListCommonGroup API
func CreateListCommonGroupRequest() (request *ListCommonGroupRequest) {
	request = &ListCommonGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "ListCommonGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateListCommonGroupResponse creates a response to parse from ListCommonGroup response
func CreateListCommonGroupResponse() (response *ListCommonGroupResponse) {
	response = &ListCommonGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
