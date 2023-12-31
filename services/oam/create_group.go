package oam

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

// CreateGroup invokes the oam.CreateGroup API synchronously
// api document: https://help.aliyun.com/api/oam/creategroup.html
func (client *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
	response = CreateCreateGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGroupWithChan invokes the oam.CreateGroup API asynchronously
// api document: https://help.aliyun.com/api/oam/creategroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGroupWithChan(request *CreateGroupRequest) (<-chan *CreateGroupResponse, <-chan error) {
	responseChan := make(chan *CreateGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGroup(request)
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

// CreateGroupWithCallback invokes the oam.CreateGroup API asynchronously
// api document: https://help.aliyun.com/api/oam/creategroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGroupWithCallback(request *CreateGroupRequest, callback func(response *CreateGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateGroup(request)
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

// CreateGroupRequest is the request struct for api CreateGroup
type CreateGroupRequest struct {
	*requests.RpcRequest
	Description string `position:"Query" name:"Description"`
	GroupName   string `position:"Query" name:"GroupName"`
}

// CreateGroupResponse is the response struct for api CreateGroup
type CreateGroupResponse struct {
	*responses.BaseResponse
	Data    string `json:"Data" xml:"Data"`
	Code    string `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
}

// CreateCreateGroupRequest creates a request to invoke CreateGroup API
func CreateCreateGroupRequest() (request *CreateGroupRequest) {
	request = &CreateGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "CreateGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateGroupResponse creates a response to parse from CreateGroup response
func CreateCreateGroupResponse() (response *CreateGroupResponse) {
	response = &CreateGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
