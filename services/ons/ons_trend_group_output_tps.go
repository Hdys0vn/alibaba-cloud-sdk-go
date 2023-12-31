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

// OnsTrendGroupOutputTps invokes the ons.OnsTrendGroupOutputTps API synchronously
func (client *Client) OnsTrendGroupOutputTps(request *OnsTrendGroupOutputTpsRequest) (response *OnsTrendGroupOutputTpsResponse, err error) {
	response = CreateOnsTrendGroupOutputTpsResponse()
	err = client.DoAction(request, response)
	return
}

// OnsTrendGroupOutputTpsWithChan invokes the ons.OnsTrendGroupOutputTps API asynchronously
func (client *Client) OnsTrendGroupOutputTpsWithChan(request *OnsTrendGroupOutputTpsRequest) (<-chan *OnsTrendGroupOutputTpsResponse, <-chan error) {
	responseChan := make(chan *OnsTrendGroupOutputTpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsTrendGroupOutputTps(request)
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

// OnsTrendGroupOutputTpsWithCallback invokes the ons.OnsTrendGroupOutputTps API asynchronously
func (client *Client) OnsTrendGroupOutputTpsWithCallback(request *OnsTrendGroupOutputTpsRequest, callback func(response *OnsTrendGroupOutputTpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsTrendGroupOutputTpsResponse
		var err error
		defer close(result)
		response, err = client.OnsTrendGroupOutputTps(request)
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

// OnsTrendGroupOutputTpsRequest is the request struct for api OnsTrendGroupOutputTps
type OnsTrendGroupOutputTpsRequest struct {
	*requests.RpcRequest
	Period     requests.Integer `position:"Query" name:"Period"`
	GroupId    string           `position:"Query" name:"GroupId"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	BeginTime  requests.Integer `position:"Query" name:"BeginTime"`
	Type       requests.Integer `position:"Query" name:"Type"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	Topic      string           `position:"Query" name:"Topic"`
}

// OnsTrendGroupOutputTpsResponse is the response struct for api OnsTrendGroupOutputTps
type OnsTrendGroupOutputTpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateOnsTrendGroupOutputTpsRequest creates a request to invoke OnsTrendGroupOutputTps API
func CreateOnsTrendGroupOutputTpsRequest() (request *OnsTrendGroupOutputTpsRequest) {
	request = &OnsTrendGroupOutputTpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsTrendGroupOutputTps", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsTrendGroupOutputTpsResponse creates a response to parse from OnsTrendGroupOutputTps response
func CreateOnsTrendGroupOutputTpsResponse() (response *OnsTrendGroupOutputTpsResponse) {
	response = &OnsTrendGroupOutputTpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
