package bssopenapi

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

// QueryRelationList invokes the bssopenapi.QueryRelationList API synchronously
func (client *Client) QueryRelationList(request *QueryRelationListRequest) (response *QueryRelationListResponse, err error) {
	response = CreateQueryRelationListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRelationListWithChan invokes the bssopenapi.QueryRelationList API asynchronously
func (client *Client) QueryRelationListWithChan(request *QueryRelationListRequest) (<-chan *QueryRelationListResponse, <-chan error) {
	responseChan := make(chan *QueryRelationListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRelationList(request)
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

// QueryRelationListWithCallback invokes the bssopenapi.QueryRelationList API asynchronously
func (client *Client) QueryRelationListWithCallback(request *QueryRelationListRequest, callback func(response *QueryRelationListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRelationListResponse
		var err error
		defer close(result)
		response, err = client.QueryRelationList(request)
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

// QueryRelationListRequest is the request struct for api QueryRelationList
type QueryRelationListRequest struct {
	*requests.RpcRequest
	StatusList *[]string        `position:"Query" name:"StatusList"  type:"Repeated"`
	PageNum    requests.Integer `position:"Query" name:"PageNum"`
	UserId     requests.Integer `position:"Query" name:"UserId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// QueryRelationListResponse is the response struct for api QueryRelationList
type QueryRelationListResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryRelationListRequest creates a request to invoke QueryRelationList API
func CreateQueryRelationListRequest() (request *QueryRelationListRequest) {
	request = &QueryRelationListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryRelationList", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryRelationListResponse creates a response to parse from QueryRelationList response
func CreateQueryRelationListResponse() (response *QueryRelationListResponse) {
	response = &QueryRelationListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
