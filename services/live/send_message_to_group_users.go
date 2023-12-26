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

// SendMessageToGroupUsers invokes the live.SendMessageToGroupUsers API synchronously
func (client *Client) SendMessageToGroupUsers(request *SendMessageToGroupUsersRequest) (response *SendMessageToGroupUsersResponse, err error) {
	response = CreateSendMessageToGroupUsersResponse()
	err = client.DoAction(request, response)
	return
}

// SendMessageToGroupUsersWithChan invokes the live.SendMessageToGroupUsers API asynchronously
func (client *Client) SendMessageToGroupUsersWithChan(request *SendMessageToGroupUsersRequest) (<-chan *SendMessageToGroupUsersResponse, <-chan error) {
	responseChan := make(chan *SendMessageToGroupUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendMessageToGroupUsers(request)
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

// SendMessageToGroupUsersWithCallback invokes the live.SendMessageToGroupUsers API asynchronously
func (client *Client) SendMessageToGroupUsersWithCallback(request *SendMessageToGroupUsersRequest, callback func(response *SendMessageToGroupUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendMessageToGroupUsersResponse
		var err error
		defer close(result)
		response, err = client.SendMessageToGroupUsers(request)
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

// SendMessageToGroupUsersRequest is the request struct for api SendMessageToGroupUsers
type SendMessageToGroupUsersRequest struct {
	*requests.RpcRequest
	Data           string           `position:"Body" name:"Data"`
	SkipAudit      requests.Boolean `position:"Query" name:"SkipAudit"`
	Type           requests.Integer `position:"Body" name:"Type"`
	OperatorUserId string           `position:"Body" name:"OperatorUserId"`
	ReceiverIdList *[]string        `position:"Body" name:"ReceiverIdList"  type:"Repeated"`
	GroupId        string           `position:"Body" name:"GroupId"`
	AppId          string           `position:"Body" name:"AppId"`
}

// SendMessageToGroupUsersResponse is the response struct for api SendMessageToGroupUsers
type SendMessageToGroupUsersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateSendMessageToGroupUsersRequest creates a request to invoke SendMessageToGroupUsers API
func CreateSendMessageToGroupUsersRequest() (request *SendMessageToGroupUsersRequest) {
	request = &SendMessageToGroupUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "SendMessageToGroupUsers", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSendMessageToGroupUsersResponse creates a response to parse from SendMessageToGroupUsers response
func CreateSendMessageToGroupUsersResponse() (response *SendMessageToGroupUsersResponse) {
	response = &SendMessageToGroupUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
