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

// PageQuerySpeechBroadcastHour invokes the iot.PageQuerySpeechBroadcastHour API synchronously
func (client *Client) PageQuerySpeechBroadcastHour(request *PageQuerySpeechBroadcastHourRequest) (response *PageQuerySpeechBroadcastHourResponse, err error) {
	response = CreatePageQuerySpeechBroadcastHourResponse()
	err = client.DoAction(request, response)
	return
}

// PageQuerySpeechBroadcastHourWithChan invokes the iot.PageQuerySpeechBroadcastHour API asynchronously
func (client *Client) PageQuerySpeechBroadcastHourWithChan(request *PageQuerySpeechBroadcastHourRequest) (<-chan *PageQuerySpeechBroadcastHourResponse, <-chan error) {
	responseChan := make(chan *PageQuerySpeechBroadcastHourResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PageQuerySpeechBroadcastHour(request)
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

// PageQuerySpeechBroadcastHourWithCallback invokes the iot.PageQuerySpeechBroadcastHour API asynchronously
func (client *Client) PageQuerySpeechBroadcastHourWithCallback(request *PageQuerySpeechBroadcastHourRequest, callback func(response *PageQuerySpeechBroadcastHourResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PageQuerySpeechBroadcastHourResponse
		var err error
		defer close(result)
		response, err = client.PageQuerySpeechBroadcastHour(request)
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

// PageQuerySpeechBroadcastHourRequest is the request struct for api PageQuerySpeechBroadcastHour
type PageQuerySpeechBroadcastHourRequest struct {
	*requests.RpcRequest
	QueryDateTimeHour string           `position:"Query" name:"QueryDateTimeHour"`
	IotInstanceId     string           `position:"Body" name:"IotInstanceId"`
	PageSize          requests.Integer `position:"Body" name:"PageSize"`
	PageToken         string           `position:"Body" name:"PageToken"`
	ShareTaskCode     string           `position:"Body" name:"ShareTaskCode"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
}

// PageQuerySpeechBroadcastHourResponse is the response struct for api PageQuerySpeechBroadcastHour
type PageQuerySpeechBroadcastHourResponse struct {
	*responses.BaseResponse
	RequestId    string                             `json:"RequestId" xml:"RequestId"`
	Success      bool                               `json:"Success" xml:"Success"`
	Code         string                             `json:"Code" xml:"Code"`
	ErrorMessage string                             `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInPageQuerySpeechBroadcastHour `json:"Data" xml:"Data"`
}

// CreatePageQuerySpeechBroadcastHourRequest creates a request to invoke PageQuerySpeechBroadcastHour API
func CreatePageQuerySpeechBroadcastHourRequest() (request *PageQuerySpeechBroadcastHourRequest) {
	request = &PageQuerySpeechBroadcastHourRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "PageQuerySpeechBroadcastHour", "", "")
	request.Method = requests.POST
	return
}

// CreatePageQuerySpeechBroadcastHourResponse creates a response to parse from PageQuerySpeechBroadcastHour response
func CreatePageQuerySpeechBroadcastHourResponse() (response *PageQuerySpeechBroadcastHourResponse) {
	response = &PageQuerySpeechBroadcastHourResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
