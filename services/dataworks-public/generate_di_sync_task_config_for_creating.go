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

// GenerateDISyncTaskConfigForCreating invokes the dataworks_public.GenerateDISyncTaskConfigForCreating API synchronously
func (client *Client) GenerateDISyncTaskConfigForCreating(request *GenerateDISyncTaskConfigForCreatingRequest) (response *GenerateDISyncTaskConfigForCreatingResponse, err error) {
	response = CreateGenerateDISyncTaskConfigForCreatingResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateDISyncTaskConfigForCreatingWithChan invokes the dataworks_public.GenerateDISyncTaskConfigForCreating API asynchronously
func (client *Client) GenerateDISyncTaskConfigForCreatingWithChan(request *GenerateDISyncTaskConfigForCreatingRequest) (<-chan *GenerateDISyncTaskConfigForCreatingResponse, <-chan error) {
	responseChan := make(chan *GenerateDISyncTaskConfigForCreatingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateDISyncTaskConfigForCreating(request)
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

// GenerateDISyncTaskConfigForCreatingWithCallback invokes the dataworks_public.GenerateDISyncTaskConfigForCreating API asynchronously
func (client *Client) GenerateDISyncTaskConfigForCreatingWithCallback(request *GenerateDISyncTaskConfigForCreatingRequest, callback func(response *GenerateDISyncTaskConfigForCreatingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateDISyncTaskConfigForCreatingResponse
		var err error
		defer close(result)
		response, err = client.GenerateDISyncTaskConfigForCreating(request)
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

// GenerateDISyncTaskConfigForCreatingRequest is the request struct for api GenerateDISyncTaskConfigForCreating
type GenerateDISyncTaskConfigForCreatingRequest struct {
	*requests.RpcRequest
	TaskType    string           `position:"Query" name:"TaskType"`
	ClientToken string           `position:"Query" name:"ClientToken"`
	TaskParam   string           `position:"Query" name:"TaskParam"`
	ProjectId   requests.Integer `position:"Query" name:"ProjectId"`
}

// GenerateDISyncTaskConfigForCreatingResponse is the response struct for api GenerateDISyncTaskConfigForCreating
type GenerateDISyncTaskConfigForCreatingResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGenerateDISyncTaskConfigForCreatingRequest creates a request to invoke GenerateDISyncTaskConfigForCreating API
func CreateGenerateDISyncTaskConfigForCreatingRequest() (request *GenerateDISyncTaskConfigForCreatingRequest) {
	request = &GenerateDISyncTaskConfigForCreatingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GenerateDISyncTaskConfigForCreating", "", "")
	request.Method = requests.POST
	return
}

// CreateGenerateDISyncTaskConfigForCreatingResponse creates a response to parse from GenerateDISyncTaskConfigForCreating response
func CreateGenerateDISyncTaskConfigForCreatingResponse() (response *GenerateDISyncTaskConfigForCreatingResponse) {
	response = &GenerateDISyncTaskConfigForCreatingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
