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

// QueryFpDBDeleteJobList invokes the mts.QueryFpDBDeleteJobList API synchronously
func (client *Client) QueryFpDBDeleteJobList(request *QueryFpDBDeleteJobListRequest) (response *QueryFpDBDeleteJobListResponse, err error) {
	response = CreateQueryFpDBDeleteJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryFpDBDeleteJobListWithChan invokes the mts.QueryFpDBDeleteJobList API asynchronously
func (client *Client) QueryFpDBDeleteJobListWithChan(request *QueryFpDBDeleteJobListRequest) (<-chan *QueryFpDBDeleteJobListResponse, <-chan error) {
	responseChan := make(chan *QueryFpDBDeleteJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryFpDBDeleteJobList(request)
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

// QueryFpDBDeleteJobListWithCallback invokes the mts.QueryFpDBDeleteJobList API asynchronously
func (client *Client) QueryFpDBDeleteJobListWithCallback(request *QueryFpDBDeleteJobListRequest, callback func(response *QueryFpDBDeleteJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryFpDBDeleteJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryFpDBDeleteJobList(request)
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

// QueryFpDBDeleteJobListRequest is the request struct for api QueryFpDBDeleteJobList
type QueryFpDBDeleteJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	JobIds               string           `position:"Query" name:"JobIds"`
}

// QueryFpDBDeleteJobListResponse is the response struct for api QueryFpDBDeleteJobList
type QueryFpDBDeleteJobListResponse struct {
	*responses.BaseResponse
	RequestId         string                              `json:"RequestId" xml:"RequestId"`
	NonExistIds       NonExistIdsInQueryFpDBDeleteJobList `json:"NonExistIds" xml:"NonExistIds"`
	FpDBDeleteJobList FpDBDeleteJobList                   `json:"FpDBDeleteJobList" xml:"FpDBDeleteJobList"`
}

// CreateQueryFpDBDeleteJobListRequest creates a request to invoke QueryFpDBDeleteJobList API
func CreateQueryFpDBDeleteJobListRequest() (request *QueryFpDBDeleteJobListRequest) {
	request = &QueryFpDBDeleteJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryFpDBDeleteJobList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryFpDBDeleteJobListResponse creates a response to parse from QueryFpDBDeleteJobList response
func CreateQueryFpDBDeleteJobListResponse() (response *QueryFpDBDeleteJobListResponse) {
	response = &QueryFpDBDeleteJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
