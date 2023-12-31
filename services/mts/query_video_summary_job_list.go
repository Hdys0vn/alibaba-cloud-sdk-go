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

// QueryVideoSummaryJobList invokes the mts.QueryVideoSummaryJobList API synchronously
func (client *Client) QueryVideoSummaryJobList(request *QueryVideoSummaryJobListRequest) (response *QueryVideoSummaryJobListResponse, err error) {
	response = CreateQueryVideoSummaryJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryVideoSummaryJobListWithChan invokes the mts.QueryVideoSummaryJobList API asynchronously
func (client *Client) QueryVideoSummaryJobListWithChan(request *QueryVideoSummaryJobListRequest) (<-chan *QueryVideoSummaryJobListResponse, <-chan error) {
	responseChan := make(chan *QueryVideoSummaryJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryVideoSummaryJobList(request)
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

// QueryVideoSummaryJobListWithCallback invokes the mts.QueryVideoSummaryJobList API asynchronously
func (client *Client) QueryVideoSummaryJobListWithCallback(request *QueryVideoSummaryJobListRequest, callback func(response *QueryVideoSummaryJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryVideoSummaryJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryVideoSummaryJobList(request)
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

// QueryVideoSummaryJobListRequest is the request struct for api QueryVideoSummaryJobList
type QueryVideoSummaryJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	JobIds               string           `position:"Query" name:"JobIds"`
}

// QueryVideoSummaryJobListResponse is the response struct for api QueryVideoSummaryJobList
type QueryVideoSummaryJobListResponse struct {
	*responses.BaseResponse
	RequestId   string                                `json:"RequestId" xml:"RequestId"`
	NonExistIds NonExistIdsInQueryVideoSummaryJobList `json:"NonExistIds" xml:"NonExistIds"`
	JobList     JobListInQueryVideoSummaryJobList     `json:"JobList" xml:"JobList"`
}

// CreateQueryVideoSummaryJobListRequest creates a request to invoke QueryVideoSummaryJobList API
func CreateQueryVideoSummaryJobListRequest() (request *QueryVideoSummaryJobListRequest) {
	request = &QueryVideoSummaryJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryVideoSummaryJobList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryVideoSummaryJobListResponse creates a response to parse from QueryVideoSummaryJobList response
func CreateQueryVideoSummaryJobListResponse() (response *QueryVideoSummaryJobListResponse) {
	response = &QueryVideoSummaryJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
