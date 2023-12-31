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

// MarkDayuVoiceService invokes the cloudcallcenter.MarkDayuVoiceService API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/markdayuvoiceservice.html
func (client *Client) MarkDayuVoiceService(request *MarkDayuVoiceServiceRequest) (response *MarkDayuVoiceServiceResponse, err error) {
	response = CreateMarkDayuVoiceServiceResponse()
	err = client.DoAction(request, response)
	return
}

// MarkDayuVoiceServiceWithChan invokes the cloudcallcenter.MarkDayuVoiceService API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/markdayuvoiceservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MarkDayuVoiceServiceWithChan(request *MarkDayuVoiceServiceRequest) (<-chan *MarkDayuVoiceServiceResponse, <-chan error) {
	responseChan := make(chan *MarkDayuVoiceServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MarkDayuVoiceService(request)
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

// MarkDayuVoiceServiceWithCallback invokes the cloudcallcenter.MarkDayuVoiceService API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/markdayuvoiceservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MarkDayuVoiceServiceWithCallback(request *MarkDayuVoiceServiceRequest, callback func(response *MarkDayuVoiceServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MarkDayuVoiceServiceResponse
		var err error
		defer close(result)
		response, err = client.MarkDayuVoiceService(request)
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

// MarkDayuVoiceServiceRequest is the request struct for api MarkDayuVoiceService
type MarkDayuVoiceServiceRequest struct {
	*requests.RpcRequest
}

// MarkDayuVoiceServiceResponse is the response struct for api MarkDayuVoiceService
type MarkDayuVoiceServiceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateMarkDayuVoiceServiceRequest creates a request to invoke MarkDayuVoiceService API
func CreateMarkDayuVoiceServiceRequest() (request *MarkDayuVoiceServiceRequest) {
	request = &MarkDayuVoiceServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "MarkDayuVoiceService", "", "")
	request.Method = requests.POST
	return
}

// CreateMarkDayuVoiceServiceResponse creates a response to parse from MarkDayuVoiceService response
func CreateMarkDayuVoiceServiceResponse() (response *MarkDayuVoiceServiceResponse) {
	response = &MarkDayuVoiceServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
