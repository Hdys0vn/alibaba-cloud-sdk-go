package live

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

// MuteGroupUser invokes the live.MuteGroupUser API synchronously
func (client *Client) MuteGroupUser(request *MuteGroupUserRequest) (response *MuteGroupUserResponse, err error) {
	response = CreateMuteGroupUserResponse()
	err = client.DoAction(request, response)
	return
}

// MuteGroupUserWithChan invokes the live.MuteGroupUser API asynchronously
func (client *Client) MuteGroupUserWithChan(request *MuteGroupUserRequest) (<-chan *MuteGroupUserResponse, <-chan error) {
	responseChan := make(chan *MuteGroupUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MuteGroupUser(request)
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

// MuteGroupUserWithCallback invokes the live.MuteGroupUser API asynchronously
func (client *Client) MuteGroupUserWithCallback(request *MuteGroupUserRequest, callback func(response *MuteGroupUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MuteGroupUserResponse
		var err error
		defer close(result)
		response, err = client.MuteGroupUser(request)
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

// MuteGroupUserRequest is the request struct for api MuteGroupUser
type MuteGroupUserRequest struct {
	*requests.RpcRequest
	MuteUserList   *[]string        `position:"Body" name:"MuteUserList"  type:"Repeated"`
	OperatorUserId string           `position:"Body" name:"OperatorUserId"`
	BroadCastType  requests.Integer `position:"Body" name:"BroadCastType"`
	GroupId        string           `position:"Body" name:"GroupId"`
	MuteTime       requests.Integer `position:"Body" name:"MuteTime"`
	AppId          string           `position:"Body" name:"AppId"`
}

// MuteGroupUserResponse is the response struct for api MuteGroupUser
type MuteGroupUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateMuteGroupUserRequest creates a request to invoke MuteGroupUser API
func CreateMuteGroupUserRequest() (request *MuteGroupUserRequest) {
	request = &MuteGroupUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "MuteGroupUser", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMuteGroupUserResponse creates a response to parse from MuteGroupUser response
func CreateMuteGroupUserResponse() (response *MuteGroupUserResponse) {
	response = &MuteGroupUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
