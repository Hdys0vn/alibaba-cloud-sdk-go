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

// QueryFacerecogJobList invokes the mts.QueryFacerecogJobList API synchronously
func (client *Client) QueryFacerecogJobList(request *QueryFacerecogJobListRequest) (response *QueryFacerecogJobListResponse, err error) {
	response = CreateQueryFacerecogJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryFacerecogJobListWithChan invokes the mts.QueryFacerecogJobList API asynchronously
func (client *Client) QueryFacerecogJobListWithChan(request *QueryFacerecogJobListRequest) (<-chan *QueryFacerecogJobListResponse, <-chan error) {
	responseChan := make(chan *QueryFacerecogJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryFacerecogJobList(request)
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

// QueryFacerecogJobListWithCallback invokes the mts.QueryFacerecogJobList API asynchronously
func (client *Client) QueryFacerecogJobListWithCallback(request *QueryFacerecogJobListRequest, callback func(response *QueryFacerecogJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryFacerecogJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryFacerecogJobList(request)
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

// QueryFacerecogJobListRequest is the request struct for api QueryFacerecogJobList
type QueryFacerecogJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	FacerecogJobIds      string           `position:"Query" name:"FacerecogJobIds"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryFacerecogJobListResponse is the response struct for api QueryFacerecogJobList
type QueryFacerecogJobListResponse struct {
	*responses.BaseResponse
	RequestId        string                             `json:"RequestId" xml:"RequestId"`
	NonExistIds      NonExistIdsInQueryFacerecogJobList `json:"NonExistIds" xml:"NonExistIds"`
	FacerecogJobList FacerecogJobList                   `json:"FacerecogJobList" xml:"FacerecogJobList"`
}

// CreateQueryFacerecogJobListRequest creates a request to invoke QueryFacerecogJobList API
func CreateQueryFacerecogJobListRequest() (request *QueryFacerecogJobListRequest) {
	request = &QueryFacerecogJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryFacerecogJobList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryFacerecogJobListResponse creates a response to parse from QueryFacerecogJobList response
func CreateQueryFacerecogJobListResponse() (response *QueryFacerecogJobListResponse) {
	response = &QueryFacerecogJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}