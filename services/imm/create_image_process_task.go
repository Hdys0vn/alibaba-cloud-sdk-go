package imm

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

// CreateImageProcessTask invokes the imm.CreateImageProcessTask API synchronously
func (client *Client) CreateImageProcessTask(request *CreateImageProcessTaskRequest) (response *CreateImageProcessTaskResponse, err error) {
	response = CreateCreateImageProcessTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateImageProcessTaskWithChan invokes the imm.CreateImageProcessTask API asynchronously
func (client *Client) CreateImageProcessTaskWithChan(request *CreateImageProcessTaskRequest) (<-chan *CreateImageProcessTaskResponse, <-chan error) {
	responseChan := make(chan *CreateImageProcessTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateImageProcessTask(request)
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

// CreateImageProcessTaskWithCallback invokes the imm.CreateImageProcessTask API asynchronously
func (client *Client) CreateImageProcessTaskWithCallback(request *CreateImageProcessTaskRequest, callback func(response *CreateImageProcessTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateImageProcessTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateImageProcessTask(request)
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

// CreateImageProcessTaskRequest is the request struct for api CreateImageProcessTask
type CreateImageProcessTaskRequest struct {
	*requests.RpcRequest
	Project         string `position:"Query" name:"Project"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	TargetList      string `position:"Query" name:"TargetList"`
	ImageUri        string `position:"Query" name:"ImageUri"`
}

// CreateImageProcessTaskResponse is the response struct for api CreateImageProcessTask
type CreateImageProcessTaskResponse struct {
	*responses.BaseResponse
	TaskId    string `json:"TaskId" xml:"TaskId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskType  string `json:"TaskType" xml:"TaskType"`
}

// CreateCreateImageProcessTaskRequest creates a request to invoke CreateImageProcessTask API
func CreateCreateImageProcessTaskRequest() (request *CreateImageProcessTaskRequest) {
	request = &CreateImageProcessTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateImageProcessTask", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateImageProcessTaskResponse creates a response to parse from CreateImageProcessTask response
func CreateCreateImageProcessTaskResponse() (response *CreateImageProcessTaskResponse) {
	response = &CreateImageProcessTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
