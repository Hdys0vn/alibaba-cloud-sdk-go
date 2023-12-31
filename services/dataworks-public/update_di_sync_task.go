package dataworks_public

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

// UpdateDISyncTask invokes the dataworks_public.UpdateDISyncTask API synchronously
func (client *Client) UpdateDISyncTask(request *UpdateDISyncTaskRequest) (response *UpdateDISyncTaskResponse, err error) {
	response = CreateUpdateDISyncTaskResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDISyncTaskWithChan invokes the dataworks_public.UpdateDISyncTask API asynchronously
func (client *Client) UpdateDISyncTaskWithChan(request *UpdateDISyncTaskRequest) (<-chan *UpdateDISyncTaskResponse, <-chan error) {
	responseChan := make(chan *UpdateDISyncTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDISyncTask(request)
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

// UpdateDISyncTaskWithCallback invokes the dataworks_public.UpdateDISyncTask API asynchronously
func (client *Client) UpdateDISyncTaskWithCallback(request *UpdateDISyncTaskRequest, callback func(response *UpdateDISyncTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDISyncTaskResponse
		var err error
		defer close(result)
		response, err = client.UpdateDISyncTask(request)
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

// UpdateDISyncTaskRequest is the request struct for api UpdateDISyncTask
type UpdateDISyncTaskRequest struct {
	*requests.RpcRequest
	TaskType    string           `position:"Query" name:"TaskType"`
	TaskParam   string           `position:"Query" name:"TaskParam"`
	TaskContent string           `position:"Query" name:"TaskContent"`
	ProjectId   requests.Integer `position:"Query" name:"ProjectId"`
	FileId      requests.Integer `position:"Query" name:"FileId"`
}

// UpdateDISyncTaskResponse is the response struct for api UpdateDISyncTask
type UpdateDISyncTaskResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdateDISyncTaskRequest creates a request to invoke UpdateDISyncTask API
func CreateUpdateDISyncTaskRequest() (request *UpdateDISyncTaskRequest) {
	request = &UpdateDISyncTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateDISyncTask", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateDISyncTaskResponse creates a response to parse from UpdateDISyncTask response
func CreateUpdateDISyncTaskResponse() (response *UpdateDISyncTaskResponse) {
	response = &UpdateDISyncTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
