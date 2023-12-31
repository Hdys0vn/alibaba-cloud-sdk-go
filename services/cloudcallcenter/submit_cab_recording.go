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

// SubmitCabRecording invokes the cloudcallcenter.SubmitCabRecording API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/submitcabrecording.html
func (client *Client) SubmitCabRecording(request *SubmitCabRecordingRequest) (response *SubmitCabRecordingResponse, err error) {
	response = CreateSubmitCabRecordingResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitCabRecordingWithChan invokes the cloudcallcenter.SubmitCabRecording API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/submitcabrecording.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitCabRecordingWithChan(request *SubmitCabRecordingRequest) (<-chan *SubmitCabRecordingResponse, <-chan error) {
	responseChan := make(chan *SubmitCabRecordingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitCabRecording(request)
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

// SubmitCabRecordingWithCallback invokes the cloudcallcenter.SubmitCabRecording API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/submitcabrecording.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitCabRecordingWithCallback(request *SubmitCabRecordingRequest, callback func(response *SubmitCabRecordingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitCabRecordingResponse
		var err error
		defer close(result)
		response, err = client.SubmitCabRecording(request)
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

// SubmitCabRecordingRequest is the request struct for api SubmitCabRecording
type SubmitCabRecordingRequest struct {
	*requests.RpcRequest
	MergedRecording   string           `position:"Query" name:"MergedRecording"`
	ResourceRecording string           `position:"Query" name:"ResourceRecording"`
	InstanceId        string           `position:"Query" name:"InstanceId"`
	InstanceOwnerId   requests.Integer `position:"Query" name:"InstanceOwnerId"`
	TaskId            string           `position:"Query" name:"TaskId"`
}

// SubmitCabRecordingResponse is the response struct for api SubmitCabRecording
type SubmitCabRecordingResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateSubmitCabRecordingRequest creates a request to invoke SubmitCabRecording API
func CreateSubmitCabRecordingRequest() (request *SubmitCabRecordingRequest) {
	request = &SubmitCabRecordingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "SubmitCabRecording", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitCabRecordingResponse creates a response to parse from SubmitCabRecording response
func CreateSubmitCabRecordingResponse() (response *SubmitCabRecordingResponse) {
	response = &SubmitCabRecordingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
