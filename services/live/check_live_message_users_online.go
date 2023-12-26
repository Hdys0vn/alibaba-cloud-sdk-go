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

// CheckLiveMessageUsersOnline invokes the live.CheckLiveMessageUsersOnline API synchronously
func (client *Client) CheckLiveMessageUsersOnline(request *CheckLiveMessageUsersOnlineRequest) (response *CheckLiveMessageUsersOnlineResponse, err error) {
	response = CreateCheckLiveMessageUsersOnlineResponse()
	err = client.DoAction(request, response)
	return
}

// CheckLiveMessageUsersOnlineWithChan invokes the live.CheckLiveMessageUsersOnline API asynchronously
func (client *Client) CheckLiveMessageUsersOnlineWithChan(request *CheckLiveMessageUsersOnlineRequest) (<-chan *CheckLiveMessageUsersOnlineResponse, <-chan error) {
	responseChan := make(chan *CheckLiveMessageUsersOnlineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckLiveMessageUsersOnline(request)
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

// CheckLiveMessageUsersOnlineWithCallback invokes the live.CheckLiveMessageUsersOnline API asynchronously
func (client *Client) CheckLiveMessageUsersOnlineWithCallback(request *CheckLiveMessageUsersOnlineRequest, callback func(response *CheckLiveMessageUsersOnlineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckLiveMessageUsersOnlineResponse
		var err error
		defer close(result)
		response, err = client.CheckLiveMessageUsersOnline(request)
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

// CheckLiveMessageUsersOnlineRequest is the request struct for api CheckLiveMessageUsersOnline
type CheckLiveMessageUsersOnlineRequest struct {
	*requests.RpcRequest
	DataCenter string    `position:"Query" name:"DataCenter"`
	UserIds    *[]string `position:"Query" name:"UserIds"  type:"Repeated"`
	AppId      string    `position:"Query" name:"AppId"`
}

// CheckLiveMessageUsersOnlineResponse is the response struct for api CheckLiveMessageUsersOnline
type CheckLiveMessageUsersOnlineResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	UserList  []Users `json:"UserList" xml:"UserList"`
}

// CreateCheckLiveMessageUsersOnlineRequest creates a request to invoke CheckLiveMessageUsersOnline API
func CreateCheckLiveMessageUsersOnlineRequest() (request *CheckLiveMessageUsersOnlineRequest) {
	request = &CheckLiveMessageUsersOnlineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CheckLiveMessageUsersOnline", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckLiveMessageUsersOnlineResponse creates a response to parse from CheckLiveMessageUsersOnline response
func CreateCheckLiveMessageUsersOnlineResponse() (response *CheckLiveMessageUsersOnlineResponse) {
	response = &CheckLiveMessageUsersOnlineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}