package cloud_siem

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

// AddUser invokes the cloud_siem.AddUser API synchronously
func (client *Client) AddUser(request *AddUserRequest) (response *AddUserResponse, err error) {
	response = CreateAddUserResponse()
	err = client.DoAction(request, response)
	return
}

// AddUserWithChan invokes the cloud_siem.AddUser API asynchronously
func (client *Client) AddUserWithChan(request *AddUserRequest) (<-chan *AddUserResponse, <-chan error) {
	responseChan := make(chan *AddUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddUser(request)
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

// AddUserWithCallback invokes the cloud_siem.AddUser API asynchronously
func (client *Client) AddUserWithCallback(request *AddUserRequest, callback func(response *AddUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddUserResponse
		var err error
		defer close(result)
		response, err = client.AddUser(request)
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

// AddUserRequest is the request struct for api AddUser
type AddUserRequest struct {
	*requests.RpcRequest
	AddedUserId requests.Integer `position:"Body" name:"AddedUserId"`
}

// AddUserResponse is the response struct for api AddUser
type AddUserResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddUserRequest creates a request to invoke AddUser API
func CreateAddUserRequest() (request *AddUserRequest) {
	request = &AddUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "AddUser", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddUserResponse creates a response to parse from AddUser response
func CreateAddUserResponse() (response *AddUserResponse) {
	response = &AddUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}