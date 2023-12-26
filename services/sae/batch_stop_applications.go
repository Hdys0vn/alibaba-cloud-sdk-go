package sae

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

// BatchStopApplications invokes the sae.BatchStopApplications API synchronously
func (client *Client) BatchStopApplications(request *BatchStopApplicationsRequest) (response *BatchStopApplicationsResponse, err error) {
	response = CreateBatchStopApplicationsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchStopApplicationsWithChan invokes the sae.BatchStopApplications API asynchronously
func (client *Client) BatchStopApplicationsWithChan(request *BatchStopApplicationsRequest) (<-chan *BatchStopApplicationsResponse, <-chan error) {
	responseChan := make(chan *BatchStopApplicationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchStopApplications(request)
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

// BatchStopApplicationsWithCallback invokes the sae.BatchStopApplications API asynchronously
func (client *Client) BatchStopApplicationsWithCallback(request *BatchStopApplicationsRequest, callback func(response *BatchStopApplicationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchStopApplicationsResponse
		var err error
		defer close(result)
		response, err = client.BatchStopApplications(request)
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

// BatchStopApplicationsRequest is the request struct for api BatchStopApplications
type BatchStopApplicationsRequest struct {
	*requests.RoaRequest
	AppIds      string `position:"Query" name:"AppIds"`
	NamespaceId string `position:"Query" name:"NamespaceId"`
}

// BatchStopApplicationsResponse is the response struct for api BatchStopApplications
type BatchStopApplicationsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateBatchStopApplicationsRequest creates a request to invoke BatchStopApplications API
func CreateBatchStopApplicationsRequest() (request *BatchStopApplicationsRequest) {
	request = &BatchStopApplicationsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "BatchStopApplications", "/pop/v1/sam/app/batchStopApplications", "serverless", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateBatchStopApplicationsResponse creates a response to parse from BatchStopApplications response
func CreateBatchStopApplicationsResponse() (response *BatchStopApplicationsResponse) {
	response = &BatchStopApplicationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
