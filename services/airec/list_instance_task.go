package airec

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

// ListInstanceTask invokes the airec.ListInstanceTask API synchronously
func (client *Client) ListInstanceTask(request *ListInstanceTaskRequest) (response *ListInstanceTaskResponse, err error) {
	response = CreateListInstanceTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstanceTaskWithChan invokes the airec.ListInstanceTask API asynchronously
func (client *Client) ListInstanceTaskWithChan(request *ListInstanceTaskRequest) (<-chan *ListInstanceTaskResponse, <-chan error) {
	responseChan := make(chan *ListInstanceTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstanceTask(request)
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

// ListInstanceTaskWithCallback invokes the airec.ListInstanceTask API asynchronously
func (client *Client) ListInstanceTaskWithCallback(request *ListInstanceTaskRequest, callback func(response *ListInstanceTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstanceTaskResponse
		var err error
		defer close(result)
		response, err = client.ListInstanceTask(request)
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

// ListInstanceTaskRequest is the request struct for api ListInstanceTask
type ListInstanceTaskRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"instanceId"`
}

// ListInstanceTaskResponse is the response struct for api ListInstanceTask
type ListInstanceTaskResponse struct {
	*responses.BaseResponse
	Code      string       `json:"code" xml:"code"`
	Message   string       `json:"message" xml:"message"`
	RequestId string       `json:"requestId" xml:"requestId"`
	Result    []ResultItem `json:"result" xml:"result"`
}

// CreateListInstanceTaskRequest creates a request to invoke ListInstanceTask API
func CreateListInstanceTaskRequest() (request *ListInstanceTaskRequest) {
	request = &ListInstanceTaskRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "ListInstanceTask", "/v2/openapi/instances/[instanceId]/tasks", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListInstanceTaskResponse creates a response to parse from ListInstanceTask response
func CreateListInstanceTaskResponse() (response *ListInstanceTaskResponse) {
	response = &ListInstanceTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
