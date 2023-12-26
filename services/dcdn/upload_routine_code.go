package dcdn

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

// UploadRoutineCode invokes the dcdn.UploadRoutineCode API synchronously
func (client *Client) UploadRoutineCode(request *UploadRoutineCodeRequest) (response *UploadRoutineCodeResponse, err error) {
	response = CreateUploadRoutineCodeResponse()
	err = client.DoAction(request, response)
	return
}

// UploadRoutineCodeWithChan invokes the dcdn.UploadRoutineCode API asynchronously
func (client *Client) UploadRoutineCodeWithChan(request *UploadRoutineCodeRequest) (<-chan *UploadRoutineCodeResponse, <-chan error) {
	responseChan := make(chan *UploadRoutineCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadRoutineCode(request)
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

// UploadRoutineCodeWithCallback invokes the dcdn.UploadRoutineCode API asynchronously
func (client *Client) UploadRoutineCodeWithCallback(request *UploadRoutineCodeRequest, callback func(response *UploadRoutineCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadRoutineCodeResponse
		var err error
		defer close(result)
		response, err = client.UploadRoutineCode(request)
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

// UploadRoutineCodeRequest is the request struct for api UploadRoutineCode
type UploadRoutineCodeRequest struct {
	*requests.RpcRequest
	CodeDescription string `position:"Body" name:"CodeDescription"`
	Name            string `position:"Body" name:"Name"`
}

// UploadRoutineCodeResponse is the response struct for api UploadRoutineCode
type UploadRoutineCodeResponse struct {
	*responses.BaseResponse
	Content   map[string]interface{} `json:"Content" xml:"Content"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateUploadRoutineCodeRequest creates a request to invoke UploadRoutineCode API
func CreateUploadRoutineCodeRequest() (request *UploadRoutineCodeRequest) {
	request = &UploadRoutineCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "UploadRoutineCode", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadRoutineCodeResponse creates a response to parse from UploadRoutineCode response
func CreateUploadRoutineCodeResponse() (response *UploadRoutineCodeResponse) {
	response = &UploadRoutineCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}