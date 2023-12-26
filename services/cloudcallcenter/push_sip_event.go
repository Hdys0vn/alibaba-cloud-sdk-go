package cloudcallcenter

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

// PushSIPEvent invokes the cloudcallcenter.PushSIPEvent API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/pushsipevent.html
func (client *Client) PushSIPEvent(request *PushSIPEventRequest) (response *PushSIPEventResponse, err error) {
	response = CreatePushSIPEventResponse()
	err = client.DoAction(request, response)
	return
}

// PushSIPEventWithChan invokes the cloudcallcenter.PushSIPEvent API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/pushsipevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PushSIPEventWithChan(request *PushSIPEventRequest) (<-chan *PushSIPEventResponse, <-chan error) {
	responseChan := make(chan *PushSIPEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushSIPEvent(request)
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

// PushSIPEventWithCallback invokes the cloudcallcenter.PushSIPEvent API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/pushsipevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PushSIPEventWithCallback(request *PushSIPEventRequest, callback func(response *PushSIPEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushSIPEventResponse
		var err error
		defer close(result)
		response, err = client.PushSIPEvent(request)
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

// PushSIPEventRequest is the request struct for api PushSIPEvent
type PushSIPEventRequest struct {
	*requests.RpcRequest
	AgentId       string           `position:"Query" name:"AgentId"`
	ContactId     string           `position:"Query" name:"ContactId"`
	Callee        string           `position:"Query" name:"Callee"`
	Initiator     string           `position:"Query" name:"Initiator"`
	EventTime     requests.Integer `position:"Query" name:"EventTime"`
	Broker        string           `position:"Query" name:"Broker"`
	Caller        string           `position:"Query" name:"Caller"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	Event         string           `position:"Query" name:"Event"`
	CallType      string           `position:"Query" name:"CallType"`
	ReleaseReason string           `position:"Query" name:"ReleaseReason"`
}

// PushSIPEventResponse is the response struct for api PushSIPEvent
type PushSIPEventResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	StatusCode     string `json:"StatusCode" xml:"StatusCode"`
	StatusDesc     string `json:"StatusDesc" xml:"StatusDesc"`
}

// CreatePushSIPEventRequest creates a request to invoke PushSIPEvent API
func CreatePushSIPEventRequest() (request *PushSIPEventRequest) {
	request = &PushSIPEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "PushSIPEvent", "", "")
	request.Method = requests.POST
	return
}

// CreatePushSIPEventResponse creates a response to parse from PushSIPEvent response
func CreatePushSIPEventResponse() (response *PushSIPEventResponse) {
	response = &PushSIPEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
