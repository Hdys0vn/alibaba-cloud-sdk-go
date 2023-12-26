package iot

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

// QuerySpeechDevice invokes the iot.QuerySpeechDevice API synchronously
func (client *Client) QuerySpeechDevice(request *QuerySpeechDeviceRequest) (response *QuerySpeechDeviceResponse, err error) {
	response = CreateQuerySpeechDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySpeechDeviceWithChan invokes the iot.QuerySpeechDevice API asynchronously
func (client *Client) QuerySpeechDeviceWithChan(request *QuerySpeechDeviceRequest) (<-chan *QuerySpeechDeviceResponse, <-chan error) {
	responseChan := make(chan *QuerySpeechDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySpeechDevice(request)
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

// QuerySpeechDeviceWithCallback invokes the iot.QuerySpeechDevice API asynchronously
func (client *Client) QuerySpeechDeviceWithCallback(request *QuerySpeechDeviceRequest, callback func(response *QuerySpeechDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySpeechDeviceResponse
		var err error
		defer close(result)
		response, err = client.QuerySpeechDevice(request)
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

// QuerySpeechDeviceRequest is the request struct for api QuerySpeechDevice
type QuerySpeechDeviceRequest struct {
	*requests.RpcRequest
	AvailableSpaceScope string           `position:"Body" name:"AvailableSpaceScope"`
	ProjectCode         string           `position:"Body" name:"ProjectCode"`
	PageId              requests.Integer `position:"Body" name:"PageId"`
	IotInstanceId       string           `position:"Body" name:"IotInstanceId"`
	PageSize            requests.Integer `position:"Body" name:"PageSize"`
	AvailableSpace      string           `position:"Body" name:"AvailableSpace"`
	ApiProduct          string           `position:"Body" name:"ApiProduct"`
	ApiRevision         string           `position:"Body" name:"ApiRevision"`
	DeviceName          string           `position:"Body" name:"DeviceName"`
}

// QuerySpeechDeviceResponse is the response struct for api QuerySpeechDevice
type QuerySpeechDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string                  `json:"RequestId" xml:"RequestId"`
	Success      bool                    `json:"Success" xml:"Success"`
	Code         string                  `json:"Code" xml:"Code"`
	ErrorMessage string                  `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQuerySpeechDevice `json:"Data" xml:"Data"`
}

// CreateQuerySpeechDeviceRequest creates a request to invoke QuerySpeechDevice API
func CreateQuerySpeechDeviceRequest() (request *QuerySpeechDeviceRequest) {
	request = &QuerySpeechDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QuerySpeechDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateQuerySpeechDeviceResponse creates a response to parse from QuerySpeechDevice response
func CreateQuerySpeechDeviceResponse() (response *QuerySpeechDeviceResponse) {
	response = &QuerySpeechDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
