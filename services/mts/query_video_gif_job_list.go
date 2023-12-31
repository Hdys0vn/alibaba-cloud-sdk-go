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

// QueryVideoGifJobList invokes the mts.QueryVideoGifJobList API synchronously
func (client *Client) QueryVideoGifJobList(request *QueryVideoGifJobListRequest) (response *QueryVideoGifJobListResponse, err error) {
	response = CreateQueryVideoGifJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryVideoGifJobListWithChan invokes the mts.QueryVideoGifJobList API asynchronously
func (client *Client) QueryVideoGifJobListWithChan(request *QueryVideoGifJobListRequest) (<-chan *QueryVideoGifJobListResponse, <-chan error) {
	responseChan := make(chan *QueryVideoGifJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryVideoGifJobList(request)
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

// QueryVideoGifJobListWithCallback invokes the mts.QueryVideoGifJobList API asynchronously
func (client *Client) QueryVideoGifJobListWithCallback(request *QueryVideoGifJobListRequest, callback func(response *QueryVideoGifJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryVideoGifJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryVideoGifJobList(request)
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

// QueryVideoGifJobListRequest is the request struct for api QueryVideoGifJobList
type QueryVideoGifJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	JobIds               string           `position:"Query" name:"JobIds"`
}

// QueryVideoGifJobListResponse is the response struct for api QueryVideoGifJobList
type QueryVideoGifJobListResponse struct {
	*responses.BaseResponse
	RequestId   string                            `json:"RequestId" xml:"RequestId"`
	NonExistIds NonExistIdsInQueryVideoGifJobList `json:"NonExistIds" xml:"NonExistIds"`
	JobList     JobListInQueryVideoGifJobList     `json:"JobList" xml:"JobList"`
}

// CreateQueryVideoGifJobListRequest creates a request to invoke QueryVideoGifJobList API
func CreateQueryVideoGifJobListRequest() (request *QueryVideoGifJobListRequest) {
	request = &QueryVideoGifJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryVideoGifJobList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryVideoGifJobListResponse creates a response to parse from QueryVideoGifJobList response
func CreateQueryVideoGifJobListResponse() (response *QueryVideoGifJobListResponse) {
	response = &QueryVideoGifJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
