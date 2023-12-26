package drds

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

// SubmitSmoothExpandPreCheckTask invokes the drds.SubmitSmoothExpandPreCheckTask API synchronously
func (client *Client) SubmitSmoothExpandPreCheckTask(request *SubmitSmoothExpandPreCheckTaskRequest) (response *SubmitSmoothExpandPreCheckTaskResponse, err error) {
	response = CreateSubmitSmoothExpandPreCheckTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitSmoothExpandPreCheckTaskWithChan invokes the drds.SubmitSmoothExpandPreCheckTask API asynchronously
func (client *Client) SubmitSmoothExpandPreCheckTaskWithChan(request *SubmitSmoothExpandPreCheckTaskRequest) (<-chan *SubmitSmoothExpandPreCheckTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitSmoothExpandPreCheckTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitSmoothExpandPreCheckTask(request)
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

// SubmitSmoothExpandPreCheckTaskWithCallback invokes the drds.SubmitSmoothExpandPreCheckTask API asynchronously
func (client *Client) SubmitSmoothExpandPreCheckTaskWithCallback(request *SubmitSmoothExpandPreCheckTaskRequest, callback func(response *SubmitSmoothExpandPreCheckTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitSmoothExpandPreCheckTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitSmoothExpandPreCheckTask(request)
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

// SubmitSmoothExpandPreCheckTaskRequest is the request struct for api SubmitSmoothExpandPreCheckTask
type SubmitSmoothExpandPreCheckTaskRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// SubmitSmoothExpandPreCheckTaskResponse is the response struct for api SubmitSmoothExpandPreCheckTask
type SubmitSmoothExpandPreCheckTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Msg       string `json:"Msg" xml:"Msg"`
	TaskId    int64  `json:"TaskId" xml:"TaskId"`
}

// CreateSubmitSmoothExpandPreCheckTaskRequest creates a request to invoke SubmitSmoothExpandPreCheckTask API
func CreateSubmitSmoothExpandPreCheckTaskRequest() (request *SubmitSmoothExpandPreCheckTaskRequest) {
	request = &SubmitSmoothExpandPreCheckTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "SubmitSmoothExpandPreCheckTask", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitSmoothExpandPreCheckTaskResponse creates a response to parse from SubmitSmoothExpandPreCheckTask response
func CreateSubmitSmoothExpandPreCheckTaskResponse() (response *SubmitSmoothExpandPreCheckTaskResponse) {
	response = &SubmitSmoothExpandPreCheckTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
