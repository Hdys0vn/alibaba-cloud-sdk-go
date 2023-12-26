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

// GetMonoRecording invokes the ccc.GetMonoRecording API synchronously
func (client *Client) GetMonoRecording(request *GetMonoRecordingRequest) (response *GetMonoRecordingResponse, err error) {
	response = CreateGetMonoRecordingResponse()
	err = client.DoAction(request, response)
	return
}

// GetMonoRecordingWithChan invokes the ccc.GetMonoRecording API asynchronously
func (client *Client) GetMonoRecordingWithChan(request *GetMonoRecordingRequest) (<-chan *GetMonoRecordingResponse, <-chan error) {
	responseChan := make(chan *GetMonoRecordingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMonoRecording(request)
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

// GetMonoRecordingWithCallback invokes the ccc.GetMonoRecording API asynchronously
func (client *Client) GetMonoRecordingWithCallback(request *GetMonoRecordingRequest, callback func(response *GetMonoRecordingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMonoRecordingResponse
		var err error
		defer close(result)
		response, err = client.GetMonoRecording(request)
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

// GetMonoRecordingRequest is the request struct for api GetMonoRecording
type GetMonoRecordingRequest struct {
	*requests.RpcRequest
	ContactId  string `position:"Query" name:"ContactId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetMonoRecordingResponse is the response struct for api GetMonoRecording
type GetMonoRecordingResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetMonoRecordingRequest creates a request to invoke GetMonoRecording API
func CreateGetMonoRecordingRequest() (request *GetMonoRecordingRequest) {
	request = &GetMonoRecordingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "GetMonoRecording", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMonoRecordingResponse creates a response to parse from GetMonoRecording response
func CreateGetMonoRecordingResponse() (response *GetMonoRecordingResponse) {
	response = &GetMonoRecordingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
