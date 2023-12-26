package dyvmsapi

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

// CreateCallTask invokes the dyvmsapi.CreateCallTask API synchronously
func (client *Client) CreateCallTask(request *CreateCallTaskRequest) (response *CreateCallTaskResponse, err error) {
	response = CreateCreateCallTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCallTaskWithChan invokes the dyvmsapi.CreateCallTask API asynchronously
func (client *Client) CreateCallTaskWithChan(request *CreateCallTaskRequest) (<-chan *CreateCallTaskResponse, <-chan error) {
	responseChan := make(chan *CreateCallTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCallTask(request)
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

// CreateCallTaskWithCallback invokes the dyvmsapi.CreateCallTask API asynchronously
func (client *Client) CreateCallTaskWithCallback(request *CreateCallTaskRequest, callback func(response *CreateCallTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCallTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateCallTask(request)
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

// CreateCallTaskRequest is the request struct for api CreateCallTask
type CreateCallTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ScheduleType         string           `position:"Query" name:"ScheduleType"`
	Data                 string           `position:"Query" name:"Data"`
	TaskName             string           `position:"Query" name:"TaskName"`
	StopTime             string           `position:"Query" name:"StopTime"`
	DataType             string           `position:"Query" name:"DataType"`
	TemplateName         string           `position:"Query" name:"TemplateName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Resource             string           `position:"Query" name:"Resource"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
	BizType              string           `position:"Query" name:"BizType"`
	FireTime             string           `position:"Query" name:"FireTime"`
	TemplateCode         string           `position:"Query" name:"TemplateCode"`
}

// CreateCallTaskResponse is the response struct for api CreateCallTask
type CreateCallTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      int64  `json:"Data" xml:"Data"`
}

// CreateCreateCallTaskRequest creates a request to invoke CreateCallTask API
func CreateCreateCallTaskRequest() (request *CreateCallTaskRequest) {
	request = &CreateCallTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "CreateCallTask", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCallTaskResponse creates a response to parse from CreateCallTask response
func CreateCreateCallTaskResponse() (response *CreateCallTaskResponse) {
	response = &CreateCallTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
