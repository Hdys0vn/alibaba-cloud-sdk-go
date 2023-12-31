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

// BridgeRtcCall invokes the ccc.BridgeRtcCall API synchronously
func (client *Client) BridgeRtcCall(request *BridgeRtcCallRequest) (response *BridgeRtcCallResponse, err error) {
	response = CreateBridgeRtcCallResponse()
	err = client.DoAction(request, response)
	return
}

// BridgeRtcCallWithChan invokes the ccc.BridgeRtcCall API asynchronously
func (client *Client) BridgeRtcCallWithChan(request *BridgeRtcCallRequest) (<-chan *BridgeRtcCallResponse, <-chan error) {
	responseChan := make(chan *BridgeRtcCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BridgeRtcCall(request)
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

// BridgeRtcCallWithCallback invokes the ccc.BridgeRtcCall API asynchronously
func (client *Client) BridgeRtcCallWithCallback(request *BridgeRtcCallRequest, callback func(response *BridgeRtcCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BridgeRtcCallResponse
		var err error
		defer close(result)
		response, err = client.BridgeRtcCall(request)
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

// BridgeRtcCallRequest is the request struct for api BridgeRtcCall
type BridgeRtcCallRequest struct {
	*requests.RpcRequest
	ServiceProvider string           `position:"Query" name:"ServiceProvider"`
	Callee          string           `position:"Query" name:"Callee"`
	UserId          string           `position:"Query" name:"UserId"`
	DeviceId        string           `position:"Query" name:"DeviceId"`
	Tags            string           `position:"Query" name:"Tags"`
	TimeoutSeconds  requests.Integer `position:"Query" name:"TimeoutSeconds"`
	Caller          string           `position:"Query" name:"Caller"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	VideoEnabled    requests.Boolean `position:"Query" name:"VideoEnabled"`
}

// BridgeRtcCallResponse is the response struct for api BridgeRtcCall
type BridgeRtcCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateBridgeRtcCallRequest creates a request to invoke BridgeRtcCall API
func CreateBridgeRtcCallRequest() (request *BridgeRtcCallRequest) {
	request = &BridgeRtcCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "BridgeRtcCall", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBridgeRtcCallResponse creates a response to parse from BridgeRtcCall response
func CreateBridgeRtcCallResponse() (response *BridgeRtcCallResponse) {
	response = &BridgeRtcCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
