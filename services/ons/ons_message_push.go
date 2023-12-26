package ons

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

// OnsMessagePush invokes the ons.OnsMessagePush API synchronously
func (client *Client) OnsMessagePush(request *OnsMessagePushRequest) (response *OnsMessagePushResponse, err error) {
	response = CreateOnsMessagePushResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMessagePushWithChan invokes the ons.OnsMessagePush API asynchronously
func (client *Client) OnsMessagePushWithChan(request *OnsMessagePushRequest) (<-chan *OnsMessagePushResponse, <-chan error) {
	responseChan := make(chan *OnsMessagePushResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMessagePush(request)
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

// OnsMessagePushWithCallback invokes the ons.OnsMessagePush API asynchronously
func (client *Client) OnsMessagePushWithCallback(request *OnsMessagePushRequest, callback func(response *OnsMessagePushResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMessagePushResponse
		var err error
		defer close(result)
		response, err = client.OnsMessagePush(request)
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

// OnsMessagePushRequest is the request struct for api OnsMessagePush
type OnsMessagePushRequest struct {
	*requests.RpcRequest
	ClientId   string `position:"Query" name:"ClientId"`
	GroupId    string `position:"Query" name:"GroupId"`
	MsgId      string `position:"Query" name:"MsgId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
}

// OnsMessagePushResponse is the response struct for api OnsMessagePush
type OnsMessagePushResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

// CreateOnsMessagePushRequest creates a request to invoke OnsMessagePush API
func CreateOnsMessagePushRequest() (request *OnsMessagePushRequest) {
	request = &OnsMessagePushRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMessagePush", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsMessagePushResponse creates a response to parse from OnsMessagePush response
func CreateOnsMessagePushResponse() (response *OnsMessagePushResponse) {
	response = &OnsMessagePushResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}