package alimt

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

// CreateDocTranslateTask invokes the alimt.CreateDocTranslateTask API synchronously
func (client *Client) CreateDocTranslateTask(request *CreateDocTranslateTaskRequest) (response *CreateDocTranslateTaskResponse, err error) {
	response = CreateCreateDocTranslateTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDocTranslateTaskWithChan invokes the alimt.CreateDocTranslateTask API asynchronously
func (client *Client) CreateDocTranslateTaskWithChan(request *CreateDocTranslateTaskRequest) (<-chan *CreateDocTranslateTaskResponse, <-chan error) {
	responseChan := make(chan *CreateDocTranslateTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDocTranslateTask(request)
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

// CreateDocTranslateTaskWithCallback invokes the alimt.CreateDocTranslateTask API asynchronously
func (client *Client) CreateDocTranslateTaskWithCallback(request *CreateDocTranslateTaskRequest, callback func(response *CreateDocTranslateTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDocTranslateTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateDocTranslateTask(request)
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

// CreateDocTranslateTaskRequest is the request struct for api CreateDocTranslateTask
type CreateDocTranslateTaskRequest struct {
	*requests.RpcRequest
	SourceLanguage string `position:"Body" name:"SourceLanguage"`
	ClientToken    string `position:"Body" name:"ClientToken"`
	Scene          string `position:"Body" name:"Scene"`
	FileUrl        string `position:"Body" name:"FileUrl"`
	TargetLanguage string `position:"Body" name:"TargetLanguage"`
	CallbackUrl    string `position:"Body" name:"CallbackUrl"`
}

// CreateDocTranslateTaskResponse is the response struct for api CreateDocTranslateTask
type CreateDocTranslateTaskResponse struct {
	*responses.BaseResponse
	Status    string `json:"Status" xml:"Status"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateDocTranslateTaskRequest creates a request to invoke CreateDocTranslateTask API
func CreateCreateDocTranslateTaskRequest() (request *CreateDocTranslateTaskRequest) {
	request = &CreateDocTranslateTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "CreateDocTranslateTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDocTranslateTaskResponse creates a response to parse from CreateDocTranslateTask response
func CreateCreateDocTranslateTaskResponse() (response *CreateDocTranslateTaskResponse) {
	response = &CreateDocTranslateTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
