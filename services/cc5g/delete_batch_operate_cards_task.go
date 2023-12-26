package cc5g

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

// DeleteBatchOperateCardsTask invokes the cc5g.DeleteBatchOperateCardsTask API synchronously
func (client *Client) DeleteBatchOperateCardsTask(request *DeleteBatchOperateCardsTaskRequest) (response *DeleteBatchOperateCardsTaskResponse, err error) {
	response = CreateDeleteBatchOperateCardsTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBatchOperateCardsTaskWithChan invokes the cc5g.DeleteBatchOperateCardsTask API asynchronously
func (client *Client) DeleteBatchOperateCardsTaskWithChan(request *DeleteBatchOperateCardsTaskRequest) (<-chan *DeleteBatchOperateCardsTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteBatchOperateCardsTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBatchOperateCardsTask(request)
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

// DeleteBatchOperateCardsTaskWithCallback invokes the cc5g.DeleteBatchOperateCardsTask API asynchronously
func (client *Client) DeleteBatchOperateCardsTaskWithCallback(request *DeleteBatchOperateCardsTaskRequest, callback func(response *DeleteBatchOperateCardsTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBatchOperateCardsTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteBatchOperateCardsTask(request)
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

// DeleteBatchOperateCardsTaskRequest is the request struct for api DeleteBatchOperateCardsTask
type DeleteBatchOperateCardsTaskRequest struct {
	*requests.RpcRequest
	DryRun                  requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken             string           `position:"Query" name:"ClientToken"`
	BatchOperateCardsTaskId string           `position:"Query" name:"BatchOperateCardsTaskId"`
}

// DeleteBatchOperateCardsTaskResponse is the response struct for api DeleteBatchOperateCardsTask
type DeleteBatchOperateCardsTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteBatchOperateCardsTaskRequest creates a request to invoke DeleteBatchOperateCardsTask API
func CreateDeleteBatchOperateCardsTaskRequest() (request *DeleteBatchOperateCardsTaskRequest) {
	request = &DeleteBatchOperateCardsTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "DeleteBatchOperateCardsTask", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteBatchOperateCardsTaskResponse creates a response to parse from DeleteBatchOperateCardsTask response
func CreateDeleteBatchOperateCardsTaskResponse() (response *DeleteBatchOperateCardsTaskResponse) {
	response = &DeleteBatchOperateCardsTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}