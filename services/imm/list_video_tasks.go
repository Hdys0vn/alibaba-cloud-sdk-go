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

// ListVideoTasks invokes the imm.ListVideoTasks API synchronously
func (client *Client) ListVideoTasks(request *ListVideoTasksRequest) (response *ListVideoTasksResponse, err error) {
	response = CreateListVideoTasksResponse()
	err = client.DoAction(request, response)
	return
}

// ListVideoTasksWithChan invokes the imm.ListVideoTasks API asynchronously
func (client *Client) ListVideoTasksWithChan(request *ListVideoTasksRequest) (<-chan *ListVideoTasksResponse, <-chan error) {
	responseChan := make(chan *ListVideoTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVideoTasks(request)
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

// ListVideoTasksWithCallback invokes the imm.ListVideoTasks API asynchronously
func (client *Client) ListVideoTasksWithCallback(request *ListVideoTasksRequest, callback func(response *ListVideoTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVideoTasksResponse
		var err error
		defer close(result)
		response, err = client.ListVideoTasks(request)
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

// ListVideoTasksRequest is the request struct for api ListVideoTasks
type ListVideoTasksRequest struct {
	*requests.RpcRequest
	MaxKeys  requests.Integer `position:"Query" name:"MaxKeys"`
	Project  string           `position:"Query" name:"Project"`
	TaskType string           `position:"Query" name:"TaskType"`
	Marker   string           `position:"Query" name:"Marker"`
}

// ListVideoTasksResponse is the response struct for api ListVideoTasks
type ListVideoTasksResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	NextMarker string      `json:"NextMarker" xml:"NextMarker"`
	Tasks      []TasksItem `json:"Tasks" xml:"Tasks"`
}

// CreateListVideoTasksRequest creates a request to invoke ListVideoTasks API
func CreateListVideoTasksRequest() (request *ListVideoTasksRequest) {
	request = &ListVideoTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListVideoTasks", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListVideoTasksResponse creates a response to parse from ListVideoTasks response
func CreateListVideoTasksResponse() (response *ListVideoTasksResponse) {
	response = &ListVideoTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
