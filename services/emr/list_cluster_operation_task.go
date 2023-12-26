package emr

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

// ListClusterOperationTask invokes the emr.ListClusterOperationTask API synchronously
func (client *Client) ListClusterOperationTask(request *ListClusterOperationTaskRequest) (response *ListClusterOperationTaskResponse, err error) {
	response = CreateListClusterOperationTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterOperationTaskWithChan invokes the emr.ListClusterOperationTask API asynchronously
func (client *Client) ListClusterOperationTaskWithChan(request *ListClusterOperationTaskRequest) (<-chan *ListClusterOperationTaskResponse, <-chan error) {
	responseChan := make(chan *ListClusterOperationTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterOperationTask(request)
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

// ListClusterOperationTaskWithCallback invokes the emr.ListClusterOperationTask API asynchronously
func (client *Client) ListClusterOperationTaskWithCallback(request *ListClusterOperationTaskRequest, callback func(response *ListClusterOperationTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterOperationTaskResponse
		var err error
		defer close(result)
		response, err = client.ListClusterOperationTask(request)
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

// ListClusterOperationTaskRequest is the request struct for api ListClusterOperationTask
type ListClusterOperationTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	OperationId     string           `position:"Query" name:"OperationId"`
	Status          string           `position:"Query" name:"Status"`
}

// ListClusterOperationTaskResponse is the response struct for api ListClusterOperationTask
type ListClusterOperationTaskResponse struct {
	*responses.BaseResponse
	RequestId                string                   `json:"RequestId" xml:"RequestId"`
	TotalCount               int                      `json:"TotalCount" xml:"TotalCount"`
	PageNumber               int                      `json:"PageNumber" xml:"PageNumber"`
	PageSize                 int                      `json:"PageSize" xml:"PageSize"`
	ClusterOperationTaskList ClusterOperationTaskList `json:"ClusterOperationTaskList" xml:"ClusterOperationTaskList"`
}

// CreateListClusterOperationTaskRequest creates a request to invoke ListClusterOperationTask API
func CreateListClusterOperationTaskRequest() (request *ListClusterOperationTaskRequest) {
	request = &ListClusterOperationTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterOperationTask", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListClusterOperationTaskResponse creates a response to parse from ListClusterOperationTask response
func CreateListClusterOperationTaskResponse() (response *ListClusterOperationTaskResponse) {
	response = &ListClusterOperationTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}