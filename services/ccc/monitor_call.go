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

// MonitorCall invokes the ccc.MonitorCall API synchronously
func (client *Client) MonitorCall(request *MonitorCallRequest) (response *MonitorCallResponse, err error) {
	response = CreateMonitorCallResponse()
	err = client.DoAction(request, response)
	return
}

// MonitorCallWithChan invokes the ccc.MonitorCall API asynchronously
func (client *Client) MonitorCallWithChan(request *MonitorCallRequest) (<-chan *MonitorCallResponse, <-chan error) {
	responseChan := make(chan *MonitorCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MonitorCall(request)
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

// MonitorCallWithCallback invokes the ccc.MonitorCall API asynchronously
func (client *Client) MonitorCallWithCallback(request *MonitorCallRequest, callback func(response *MonitorCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MonitorCallResponse
		var err error
		defer close(result)
		response, err = client.MonitorCall(request)
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

// MonitorCallRequest is the request struct for api MonitorCall
type MonitorCallRequest struct {
	*requests.RpcRequest
	UserId          string           `position:"Query" name:"UserId"`
	DeviceId        string           `position:"Query" name:"DeviceId"`
	TimeoutSeconds  requests.Integer `position:"Query" name:"TimeoutSeconds"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	MonitoredUserId string           `position:"Query" name:"MonitoredUserId"`
}

// MonitorCallResponse is the response struct for api MonitorCall
type MonitorCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateMonitorCallRequest creates a request to invoke MonitorCall API
func CreateMonitorCallRequest() (request *MonitorCallRequest) {
	request = &MonitorCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "MonitorCall", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMonitorCallResponse creates a response to parse from MonitorCall response
func CreateMonitorCallResponse() (response *MonitorCallResponse) {
	response = &MonitorCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
