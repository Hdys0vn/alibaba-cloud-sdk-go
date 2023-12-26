package ccc

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

// ModifyUser invokes the ccc.ModifyUser API synchronously
func (client *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
	response = CreateModifyUserResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyUserWithChan invokes the ccc.ModifyUser API asynchronously
func (client *Client) ModifyUserWithChan(request *ModifyUserRequest) (<-chan *ModifyUserResponse, <-chan error) {
	responseChan := make(chan *ModifyUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyUser(request)
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

// ModifyUserWithCallback invokes the ccc.ModifyUser API asynchronously
func (client *Client) ModifyUserWithCallback(request *ModifyUserRequest, callback func(response *ModifyUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyUserResponse
		var err error
		defer close(result)
		response, err = client.ModifyUser(request)
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

// ModifyUserRequest is the request struct for api ModifyUser
type ModifyUserRequest struct {
	*requests.RpcRequest
	RoleId      string           `position:"Query" name:"RoleId"`
	Mobile      string           `position:"Query" name:"Mobile"`
	WorkMode    string           `position:"Query" name:"WorkMode"`
	UserId      string           `position:"Query" name:"UserId"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	DisplayName string           `position:"Query" name:"DisplayName"`
	Force       requests.Boolean `position:"Query" name:"Force"`
	DisplayId   string           `position:"Query" name:"DisplayId"`
}

// ModifyUserResponse is the response struct for api ModifyUser
type ModifyUserResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateModifyUserRequest creates a request to invoke ModifyUser API
func CreateModifyUserRequest() (request *ModifyUserRequest) {
	request = &ModifyUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ModifyUser", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyUserResponse creates a response to parse from ModifyUser response
func CreateModifyUserResponse() (response *ModifyUserResponse) {
	response = &ModifyUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
