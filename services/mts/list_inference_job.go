package mts

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

// ListInferenceJob invokes the mts.ListInferenceJob API synchronously
func (client *Client) ListInferenceJob(request *ListInferenceJobRequest) (response *ListInferenceJobResponse, err error) {
	response = CreateListInferenceJobResponse()
	err = client.DoAction(request, response)
	return
}

// ListInferenceJobWithChan invokes the mts.ListInferenceJob API asynchronously
func (client *Client) ListInferenceJobWithChan(request *ListInferenceJobRequest) (<-chan *ListInferenceJobResponse, <-chan error) {
	responseChan := make(chan *ListInferenceJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInferenceJob(request)
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

// ListInferenceJobWithCallback invokes the mts.ListInferenceJob API asynchronously
func (client *Client) ListInferenceJobWithCallback(request *ListInferenceJobRequest, callback func(response *ListInferenceJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInferenceJobResponse
		var err error
		defer close(result)
		response, err = client.ListInferenceJob(request)
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

// ListInferenceJobRequest is the request struct for api ListInferenceJob
type ListInferenceJobRequest struct {
	*requests.RpcRequest
	MaxPageSize requests.Integer `position:"Query" name:"MaxPageSize"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	ServerName  string           `position:"Query" name:"ServerName"`
}

// ListInferenceJobResponse is the response struct for api ListInferenceJob
type ListInferenceJobResponse struct {
	*responses.BaseResponse
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Code      string     `json:"Code" xml:"Code"`
	TotalSize int64      `json:"TotalSize" xml:"TotalSize"`
	Jobs      []JobsItem `json:"Jobs" xml:"Jobs"`
}

// CreateListInferenceJobRequest creates a request to invoke ListInferenceJob API
func CreateListInferenceJobRequest() (request *ListInferenceJobRequest) {
	request = &ListInferenceJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListInferenceJob", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInferenceJobResponse creates a response to parse from ListInferenceJob response
func CreateListInferenceJobResponse() (response *ListInferenceJobResponse) {
	response = &ListInferenceJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}