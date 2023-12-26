package mpaas

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

// PushReport invokes the mpaas.PushReport API synchronously
func (client *Client) PushReport(request *PushReportRequest) (response *PushReportResponse, err error) {
	response = CreatePushReportResponse()
	err = client.DoAction(request, response)
	return
}

// PushReportWithChan invokes the mpaas.PushReport API asynchronously
func (client *Client) PushReportWithChan(request *PushReportRequest) (<-chan *PushReportResponse, <-chan error) {
	responseChan := make(chan *PushReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushReport(request)
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

// PushReportWithCallback invokes the mpaas.PushReport API asynchronously
func (client *Client) PushReportWithCallback(request *PushReportRequest, callback func(response *PushReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushReportResponse
		var err error
		defer close(result)
		response, err = client.PushReport(request)
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

// PushReportRequest is the request struct for api PushReport
type PushReportRequest struct {
	*requests.RpcRequest
	Channel                 string           `position:"Body" name:"Channel"`
	Imsi                    string           `position:"Body" name:"Imsi"`
	Model                   string           `position:"Body" name:"Model"`
	DeliveryToken           string           `position:"Body" name:"DeliveryToken"`
	ThirdChannelDeviceToken string           `position:"Body" name:"ThirdChannelDeviceToken"`
	AppVersion              string           `position:"Body" name:"AppVersion"`
	OsType                  requests.Integer `position:"Body" name:"OsType"`
	PushVersion             string           `position:"Body" name:"PushVersion"`
	ConnectType             string           `position:"Body" name:"ConnectType"`
	AppId                   string           `position:"Body" name:"AppId"`
	Imei                    string           `position:"Body" name:"Imei"`
	ThirdChannel            requests.Integer `position:"Body" name:"ThirdChannel"`
	WorkspaceId             string           `position:"Body" name:"WorkspaceId"`
}

// PushReportResponse is the response struct for api PushReport
type PushReportResponse struct {
	*responses.BaseResponse
	ResultMessage string     `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string     `json:"ResultCode" xml:"ResultCode"`
	RequestId     string     `json:"RequestId" xml:"RequestId"`
	PushResult    PushResult `json:"PushResult" xml:"PushResult"`
}

// CreatePushReportRequest creates a request to invoke PushReport API
func CreatePushReportRequest() (request *PushReportRequest) {
	request = &PushReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "PushReport", "", "")
	request.Method = requests.POST
	return
}

// CreatePushReportResponse creates a response to parse from PushReport response
func CreatePushReportResponse() (response *PushReportResponse) {
	response = &PushReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
