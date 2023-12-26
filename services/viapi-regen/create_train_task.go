package viapi_regen

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

// CreateTrainTask invokes the viapi_regen.CreateTrainTask API synchronously
func (client *Client) CreateTrainTask(request *CreateTrainTaskRequest) (response *CreateTrainTaskResponse, err error) {
	response = CreateCreateTrainTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTrainTaskWithChan invokes the viapi_regen.CreateTrainTask API asynchronously
func (client *Client) CreateTrainTaskWithChan(request *CreateTrainTaskRequest) (<-chan *CreateTrainTaskResponse, <-chan error) {
	responseChan := make(chan *CreateTrainTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTrainTask(request)
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

// CreateTrainTaskWithCallback invokes the viapi_regen.CreateTrainTask API asynchronously
func (client *Client) CreateTrainTaskWithCallback(request *CreateTrainTaskRequest, callback func(response *CreateTrainTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTrainTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateTrainTask(request)
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

// CreateTrainTaskRequest is the request struct for api CreateTrainTask
type CreateTrainTaskRequest struct {
	*requests.RpcRequest
	Description        string           `position:"Body" name:"Description"`
	TrainMode          string           `position:"Body" name:"TrainMode"`
	DatasetIds         string           `position:"Body" name:"DatasetIds"`
	PreTrainTaskId     requests.Integer `position:"Body" name:"PreTrainTaskId"`
	AdvancedParameters string           `position:"Body" name:"AdvancedParameters"`
	LabelId            requests.Integer `position:"Body" name:"LabelId"`
	Name               string           `position:"Body" name:"Name"`
	DatasetId          requests.Integer `position:"Body" name:"DatasetId"`
	LabelIds           string           `position:"Body" name:"LabelIds"`
	WorkspaceId        requests.Integer `position:"Body" name:"WorkspaceId"`
}

// CreateTrainTaskResponse is the response struct for api CreateTrainTask
type CreateTrainTaskResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateTrainTaskRequest creates a request to invoke CreateTrainTask API
func CreateCreateTrainTaskRequest() (request *CreateTrainTaskRequest) {
	request = &CreateTrainTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "CreateTrainTask", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTrainTaskResponse creates a response to parse from CreateTrainTask response
func CreateCreateTrainTaskResponse() (response *CreateTrainTaskResponse) {
	response = &CreateTrainTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
