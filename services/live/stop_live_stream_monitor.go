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

// StopLiveStreamMonitor invokes the live.StopLiveStreamMonitor API synchronously
func (client *Client) StopLiveStreamMonitor(request *StopLiveStreamMonitorRequest) (response *StopLiveStreamMonitorResponse, err error) {
	response = CreateStopLiveStreamMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// StopLiveStreamMonitorWithChan invokes the live.StopLiveStreamMonitor API asynchronously
func (client *Client) StopLiveStreamMonitorWithChan(request *StopLiveStreamMonitorRequest) (<-chan *StopLiveStreamMonitorResponse, <-chan error) {
	responseChan := make(chan *StopLiveStreamMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopLiveStreamMonitor(request)
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

// StopLiveStreamMonitorWithCallback invokes the live.StopLiveStreamMonitor API asynchronously
func (client *Client) StopLiveStreamMonitorWithCallback(request *StopLiveStreamMonitorRequest, callback func(response *StopLiveStreamMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopLiveStreamMonitorResponse
		var err error
		defer close(result)
		response, err = client.StopLiveStreamMonitor(request)
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

// StopLiveStreamMonitorRequest is the request struct for api StopLiveStreamMonitor
type StopLiveStreamMonitorRequest struct {
	*requests.RpcRequest
	MonitorId string           `position:"Query" name:"MonitorId"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// StopLiveStreamMonitorResponse is the response struct for api StopLiveStreamMonitor
type StopLiveStreamMonitorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopLiveStreamMonitorRequest creates a request to invoke StopLiveStreamMonitor API
func CreateStopLiveStreamMonitorRequest() (request *StopLiveStreamMonitorRequest) {
	request = &StopLiveStreamMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "StopLiveStreamMonitor", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopLiveStreamMonitorResponse creates a response to parse from StopLiveStreamMonitor response
func CreateStopLiveStreamMonitorResponse() (response *StopLiveStreamMonitorResponse) {
	response = &StopLiveStreamMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
