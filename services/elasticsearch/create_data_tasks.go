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

// CreateDataTasks invokes the elasticsearch.CreateDataTasks API synchronously
func (client *Client) CreateDataTasks(request *CreateDataTasksRequest) (response *CreateDataTasksResponse, err error) {
	response = CreateCreateDataTasksResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataTasksWithChan invokes the elasticsearch.CreateDataTasks API asynchronously
func (client *Client) CreateDataTasksWithChan(request *CreateDataTasksRequest) (<-chan *CreateDataTasksResponse, <-chan error) {
	responseChan := make(chan *CreateDataTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataTasks(request)
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

// CreateDataTasksWithCallback invokes the elasticsearch.CreateDataTasks API asynchronously
func (client *Client) CreateDataTasksWithCallback(request *CreateDataTasksRequest, callback func(response *CreateDataTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataTasksResponse
		var err error
		defer close(result)
		response, err = client.CreateDataTasks(request)
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

// CreateDataTasksRequest is the request struct for api CreateDataTasks
type CreateDataTasksRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

// CreateDataTasksResponse is the response struct for api CreateDataTasks
type CreateDataTasksResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateCreateDataTasksRequest creates a request to invoke CreateDataTasks API
func CreateCreateDataTasksRequest() (request *CreateDataTasksRequest) {
	request = &CreateDataTasksRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "CreateDataTasks", "/openapi/instances/[InstanceId]/data-task", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDataTasksResponse creates a response to parse from CreateDataTasks response
func CreateCreateDataTasksResponse() (response *CreateDataTasksResponse) {
	response = &CreateDataTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
