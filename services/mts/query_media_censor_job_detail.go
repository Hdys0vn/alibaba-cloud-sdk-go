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

// QueryMediaCensorJobDetail invokes the mts.QueryMediaCensorJobDetail API synchronously
func (client *Client) QueryMediaCensorJobDetail(request *QueryMediaCensorJobDetailRequest) (response *QueryMediaCensorJobDetailResponse, err error) {
	response = CreateQueryMediaCensorJobDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMediaCensorJobDetailWithChan invokes the mts.QueryMediaCensorJobDetail API asynchronously
func (client *Client) QueryMediaCensorJobDetailWithChan(request *QueryMediaCensorJobDetailRequest) (<-chan *QueryMediaCensorJobDetailResponse, <-chan error) {
	responseChan := make(chan *QueryMediaCensorJobDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMediaCensorJobDetail(request)
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

// QueryMediaCensorJobDetailWithCallback invokes the mts.QueryMediaCensorJobDetail API asynchronously
func (client *Client) QueryMediaCensorJobDetailWithCallback(request *QueryMediaCensorJobDetailRequest, callback func(response *QueryMediaCensorJobDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMediaCensorJobDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryMediaCensorJobDetail(request)
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

// QueryMediaCensorJobDetailRequest is the request struct for api QueryMediaCensorJobDetail
type QueryMediaCensorJobDetailRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NextPageToken        string           `position:"Query" name:"NextPageToken"`
	JobId                string           `position:"Query" name:"JobId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	MaximumPageSize      requests.Integer `position:"Query" name:"MaximumPageSize"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryMediaCensorJobDetailResponse is the response struct for api QueryMediaCensorJobDetail
type QueryMediaCensorJobDetailResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	MediaCensorJobDetail MediaCensorJobDetail `json:"MediaCensorJobDetail" xml:"MediaCensorJobDetail"`
}

// CreateQueryMediaCensorJobDetailRequest creates a request to invoke QueryMediaCensorJobDetail API
func CreateQueryMediaCensorJobDetailRequest() (request *QueryMediaCensorJobDetailRequest) {
	request = &QueryMediaCensorJobDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaCensorJobDetail", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryMediaCensorJobDetailResponse creates a response to parse from QueryMediaCensorJobDetail response
func CreateQueryMediaCensorJobDetailResponse() (response *QueryMediaCensorJobDetailResponse) {
	response = &QueryMediaCensorJobDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
