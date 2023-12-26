package elasticsearch

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

// DeleteDataTask invokes the elasticsearch.DeleteDataTask API synchronously
func (client *Client) DeleteDataTask(request *DeleteDataTaskRequest) (response *DeleteDataTaskResponse, err error) {
	response = CreateDeleteDataTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDataTaskWithChan invokes the elasticsearch.DeleteDataTask API asynchronously
func (client *Client) DeleteDataTaskWithChan(request *DeleteDataTaskRequest) (<-chan *DeleteDataTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteDataTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDataTask(request)
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

// DeleteDataTaskWithCallback invokes the elasticsearch.DeleteDataTask API asynchronously
func (client *Client) DeleteDataTaskWithCallback(request *DeleteDataTaskRequest, callback func(response *DeleteDataTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDataTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteDataTask(request)
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

// DeleteDataTaskRequest is the request struct for api DeleteDataTask
type DeleteDataTaskRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	TaskId      string `position:"Query" name:"taskId"`
}

// DeleteDataTaskResponse is the response struct for api DeleteDataTask
type DeleteDataTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateDeleteDataTaskRequest creates a request to invoke DeleteDataTask API
func CreateDeleteDataTaskRequest() (request *DeleteDataTaskRequest) {
	request = &DeleteDataTaskRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DeleteDataTask", "/openapi/instances/[InstanceId]/data-task", "elasticsearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteDataTaskResponse creates a response to parse from DeleteDataTask response
func CreateDeleteDataTaskResponse() (response *DeleteDataTaskResponse) {
	response = &DeleteDataTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
