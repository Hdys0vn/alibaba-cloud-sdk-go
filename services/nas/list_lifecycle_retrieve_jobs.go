package nas

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

// ListLifecycleRetrieveJobs invokes the nas.ListLifecycleRetrieveJobs API synchronously
func (client *Client) ListLifecycleRetrieveJobs(request *ListLifecycleRetrieveJobsRequest) (response *ListLifecycleRetrieveJobsResponse, err error) {
	response = CreateListLifecycleRetrieveJobsResponse()
	err = client.DoAction(request, response)
	return
}

// ListLifecycleRetrieveJobsWithChan invokes the nas.ListLifecycleRetrieveJobs API asynchronously
func (client *Client) ListLifecycleRetrieveJobsWithChan(request *ListLifecycleRetrieveJobsRequest) (<-chan *ListLifecycleRetrieveJobsResponse, <-chan error) {
	responseChan := make(chan *ListLifecycleRetrieveJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLifecycleRetrieveJobs(request)
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

// ListLifecycleRetrieveJobsWithCallback invokes the nas.ListLifecycleRetrieveJobs API asynchronously
func (client *Client) ListLifecycleRetrieveJobsWithCallback(request *ListLifecycleRetrieveJobsRequest, callback func(response *ListLifecycleRetrieveJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLifecycleRetrieveJobsResponse
		var err error
		defer close(result)
		response, err = client.ListLifecycleRetrieveJobs(request)
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

// ListLifecycleRetrieveJobsRequest is the request struct for api ListLifecycleRetrieveJobs
type ListLifecycleRetrieveJobsRequest struct {
	*requests.RpcRequest
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	FileSystemId string           `position:"Query" name:"FileSystemId"`
	Status       string           `position:"Query" name:"Status"`
}

// ListLifecycleRetrieveJobsResponse is the response struct for api ListLifecycleRetrieveJobs
type ListLifecycleRetrieveJobsResponse struct {
	*responses.BaseResponse
	TotalCount            int                    `json:"TotalCount" xml:"TotalCount"`
	PageSize              int                    `json:"PageSize" xml:"PageSize"`
	RequestId             string                 `json:"RequestId" xml:"RequestId"`
	PageNumber            int                    `json:"PageNumber" xml:"PageNumber"`
	LifecycleRetrieveJobs []LifecycleRetrieveJob `json:"LifecycleRetrieveJobs" xml:"LifecycleRetrieveJobs"`
}

// CreateListLifecycleRetrieveJobsRequest creates a request to invoke ListLifecycleRetrieveJobs API
func CreateListLifecycleRetrieveJobsRequest() (request *ListLifecycleRetrieveJobsRequest) {
	request = &ListLifecycleRetrieveJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "ListLifecycleRetrieveJobs", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListLifecycleRetrieveJobsResponse creates a response to parse from ListLifecycleRetrieveJobs response
func CreateListLifecycleRetrieveJobsResponse() (response *ListLifecycleRetrieveJobsResponse) {
	response = &ListLifecycleRetrieveJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
