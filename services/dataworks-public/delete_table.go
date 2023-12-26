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

// DeleteTable invokes the dataworks_public.DeleteTable API synchronously
func (client *Client) DeleteTable(request *DeleteTableRequest) (response *DeleteTableResponse, err error) {
	response = CreateDeleteTableResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTableWithChan invokes the dataworks_public.DeleteTable API asynchronously
func (client *Client) DeleteTableWithChan(request *DeleteTableRequest) (<-chan *DeleteTableResponse, <-chan error) {
	responseChan := make(chan *DeleteTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTable(request)
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

// DeleteTableWithCallback invokes the dataworks_public.DeleteTable API asynchronously
func (client *Client) DeleteTableWithCallback(request *DeleteTableRequest, callback func(response *DeleteTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTableResponse
		var err error
		defer close(result)
		response, err = client.DeleteTable(request)
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

// DeleteTableRequest is the request struct for api DeleteTable
type DeleteTableRequest struct {
	*requests.RpcRequest
	Schema    string           `position:"Query" name:"Schema"`
	EnvType   requests.Integer `position:"Query" name:"EnvType"`
	TableName string           `position:"Query" name:"TableName"`
	AppGuid   string           `position:"Query" name:"AppGuid"`
	ProjectId requests.Integer `position:"Query" name:"ProjectId"`
}

// DeleteTableResponse is the response struct for api DeleteTable
type DeleteTableResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	TaskInfo  TaskInfo `json:"TaskInfo" xml:"TaskInfo"`
}

// CreateDeleteTableRequest creates a request to invoke DeleteTable API
func CreateDeleteTableRequest() (request *DeleteTableRequest) {
	request = &DeleteTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeleteTable", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteTableResponse creates a response to parse from DeleteTable response
func CreateDeleteTableResponse() (response *DeleteTableResponse) {
	response = &DeleteTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
