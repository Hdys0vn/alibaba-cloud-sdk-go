package mse

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

// ListClusterHealthCheckTask invokes the mse.ListClusterHealthCheckTask API synchronously
func (client *Client) ListClusterHealthCheckTask(request *ListClusterHealthCheckTaskRequest) (response *ListClusterHealthCheckTaskResponse, err error) {
	response = CreateListClusterHealthCheckTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterHealthCheckTaskWithChan invokes the mse.ListClusterHealthCheckTask API asynchronously
func (client *Client) ListClusterHealthCheckTaskWithChan(request *ListClusterHealthCheckTaskRequest) (<-chan *ListClusterHealthCheckTaskResponse, <-chan error) {
	responseChan := make(chan *ListClusterHealthCheckTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterHealthCheckTask(request)
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

// ListClusterHealthCheckTaskWithCallback invokes the mse.ListClusterHealthCheckTask API asynchronously
func (client *Client) ListClusterHealthCheckTaskWithCallback(request *ListClusterHealthCheckTaskRequest, callback func(response *ListClusterHealthCheckTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterHealthCheckTaskResponse
		var err error
		defer close(result)
		response, err = client.ListClusterHealthCheckTask(request)
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

// ListClusterHealthCheckTaskRequest is the request struct for api ListClusterHealthCheckTask
type ListClusterHealthCheckTaskRequest struct {
	*requests.RpcRequest
	MseSessionId   string           `position:"Query" name:"MseSessionId"`
	PageNum        requests.Integer `position:"Query" name:"PageNum"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	RequestPars    string           `position:"Query" name:"RequestPars"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
}

// ListClusterHealthCheckTaskResponse is the response struct for api ListClusterHealthCheckTask
type ListClusterHealthCheckTaskResponse struct {
	*responses.BaseResponse
	RequestId      string                           `json:"RequestId" xml:"RequestId"`
	Success        bool                             `json:"Success" xml:"Success"`
	Code           int                              `json:"Code" xml:"Code"`
	ErrorCode      string                           `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode int                              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                           `json:"Message" xml:"Message"`
	DynamicCode    string                           `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string                           `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           DataInListClusterHealthCheckTask `json:"Data" xml:"Data"`
}

// CreateListClusterHealthCheckTaskRequest creates a request to invoke ListClusterHealthCheckTask API
func CreateListClusterHealthCheckTaskRequest() (request *ListClusterHealthCheckTaskRequest) {
	request = &ListClusterHealthCheckTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListClusterHealthCheckTask", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListClusterHealthCheckTaskResponse creates a response to parse from ListClusterHealthCheckTask response
func CreateListClusterHealthCheckTaskResponse() (response *ListClusterHealthCheckTaskResponse) {
	response = &ListClusterHealthCheckTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
