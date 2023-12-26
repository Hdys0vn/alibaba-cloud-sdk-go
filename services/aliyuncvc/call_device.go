package aliyuncvc

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

// CallDevice invokes the aliyuncvc.CallDevice API synchronously
func (client *Client) CallDevice(request *CallDeviceRequest) (response *CallDeviceResponse, err error) {
	response = CreateCallDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// CallDeviceWithChan invokes the aliyuncvc.CallDevice API asynchronously
func (client *Client) CallDeviceWithChan(request *CallDeviceRequest) (<-chan *CallDeviceResponse, <-chan error) {
	responseChan := make(chan *CallDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CallDevice(request)
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

// CallDeviceWithCallback invokes the aliyuncvc.CallDevice API asynchronously
func (client *Client) CallDeviceWithCallback(request *CallDeviceRequest, callback func(response *CallDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CallDeviceResponse
		var err error
		defer close(result)
		response, err = client.CallDevice(request)
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

// CallDeviceRequest is the request struct for api CallDevice
type CallDeviceRequest struct {
	*requests.RpcRequest
	InviteName      string           `position:"Body" name:"InviteName"`
	OperateUserId   string           `position:"Body" name:"OperateUserId"`
	JoinMeetingFlag requests.Boolean `position:"Query" name:"JoinMeetingFlag"`
	MeetingCode     string           `position:"Body" name:"MeetingCode"`
	SN              string           `position:"Body" name:"SN"`
}

// CallDeviceResponse is the response struct for api CallDevice
type CallDeviceResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	MessageId string `json:"MessageId" xml:"MessageId"`
}

// CreateCallDeviceRequest creates a request to invoke CallDevice API
func CreateCallDeviceRequest() (request *CallDeviceRequest) {
	request = &CallDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "CallDevice", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCallDeviceResponse creates a response to parse from CallDevice response
func CreateCallDeviceResponse() (response *CallDeviceResponse) {
	response = &CallDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}