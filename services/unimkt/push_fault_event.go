package unimkt

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

// PushFaultEvent invokes the unimkt.PushFaultEvent API synchronously
func (client *Client) PushFaultEvent(request *PushFaultEventRequest) (response *PushFaultEventResponse, err error) {
	response = CreatePushFaultEventResponse()
	err = client.DoAction(request, response)
	return
}

// PushFaultEventWithChan invokes the unimkt.PushFaultEvent API asynchronously
func (client *Client) PushFaultEventWithChan(request *PushFaultEventRequest) (<-chan *PushFaultEventResponse, <-chan error) {
	responseChan := make(chan *PushFaultEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushFaultEvent(request)
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

// PushFaultEventWithCallback invokes the unimkt.PushFaultEvent API asynchronously
func (client *Client) PushFaultEventWithCallback(request *PushFaultEventRequest, callback func(response *PushFaultEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushFaultEventResponse
		var err error
		defer close(result)
		response, err = client.PushFaultEvent(request)
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

// PushFaultEventRequest is the request struct for api PushFaultEvent
type PushFaultEventRequest struct {
	*requests.RpcRequest
	FaultComment string           `position:"Body" name:"FaultComment"`
	Time         requests.Integer `position:"Body" name:"Time"`
	Type         requests.Integer `position:"Body" name:"Type"`
	DeviceSn     string           `position:"Body" name:"DeviceSn"`
	ChannelId    string           `position:"Body" name:"ChannelId"`
	FaultType    string           `position:"Body" name:"FaultType"`
}

// PushFaultEventResponse is the response struct for api PushFaultEvent
type PushFaultEventResponse struct {
	*responses.BaseResponse
	Status    bool   `json:"Status" xml:"Status"`
	Msg       string `json:"Msg" xml:"Msg"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePushFaultEventRequest creates a request to invoke PushFaultEvent API
func CreatePushFaultEventRequest() (request *PushFaultEventRequest) {
	request = &PushFaultEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "PushFaultEvent", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePushFaultEventResponse creates a response to parse from PushFaultEvent response
func CreatePushFaultEventResponse() (response *PushFaultEventResponse) {
	response = &PushFaultEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
