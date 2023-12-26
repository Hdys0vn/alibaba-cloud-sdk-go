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

// QueryAnnotationJobList invokes the mts.QueryAnnotationJobList API synchronously
func (client *Client) QueryAnnotationJobList(request *QueryAnnotationJobListRequest) (response *QueryAnnotationJobListResponse, err error) {
	response = CreateQueryAnnotationJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAnnotationJobListWithChan invokes the mts.QueryAnnotationJobList API asynchronously
func (client *Client) QueryAnnotationJobListWithChan(request *QueryAnnotationJobListRequest) (<-chan *QueryAnnotationJobListResponse, <-chan error) {
	responseChan := make(chan *QueryAnnotationJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAnnotationJobList(request)
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

// QueryAnnotationJobListWithCallback invokes the mts.QueryAnnotationJobList API asynchronously
func (client *Client) QueryAnnotationJobListWithCallback(request *QueryAnnotationJobListRequest, callback func(response *QueryAnnotationJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAnnotationJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryAnnotationJobList(request)
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

// QueryAnnotationJobListRequest is the request struct for api QueryAnnotationJobList
type QueryAnnotationJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AnnotationJobIds     string           `position:"Query" name:"AnnotationJobIds"`
}

// QueryAnnotationJobListResponse is the response struct for api QueryAnnotationJobList
type QueryAnnotationJobListResponse struct {
	*responses.BaseResponse
	RequestId         string                              `json:"RequestId" xml:"RequestId"`
	NonExistIds       NonExistIdsInQueryAnnotationJobList `json:"NonExistIds" xml:"NonExistIds"`
	AnnotationJobList AnnotationJobList                   `json:"AnnotationJobList" xml:"AnnotationJobList"`
}

// CreateQueryAnnotationJobListRequest creates a request to invoke QueryAnnotationJobList API
func CreateQueryAnnotationJobListRequest() (request *QueryAnnotationJobListRequest) {
	request = &QueryAnnotationJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryAnnotationJobList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryAnnotationJobListResponse creates a response to parse from QueryAnnotationJobList response
func CreateQueryAnnotationJobListResponse() (response *QueryAnnotationJobListResponse) {
	response = &QueryAnnotationJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}