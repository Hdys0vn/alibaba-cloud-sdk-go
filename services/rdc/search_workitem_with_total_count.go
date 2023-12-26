package rdc

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

// SearchWorkitemWithTotalCount invokes the rdc.SearchWorkitemWithTotalCount API synchronously
// api document: https://help.aliyun.com/api/rdc/searchworkitemwithtotalcount.html
func (client *Client) SearchWorkitemWithTotalCount(request *SearchWorkitemWithTotalCountRequest) (response *SearchWorkitemWithTotalCountResponse, err error) {
	response = CreateSearchWorkitemWithTotalCountResponse()
	err = client.DoAction(request, response)
	return
}

// SearchWorkitemWithTotalCountWithChan invokes the rdc.SearchWorkitemWithTotalCount API asynchronously
// api document: https://help.aliyun.com/api/rdc/searchworkitemwithtotalcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchWorkitemWithTotalCountWithChan(request *SearchWorkitemWithTotalCountRequest) (<-chan *SearchWorkitemWithTotalCountResponse, <-chan error) {
	responseChan := make(chan *SearchWorkitemWithTotalCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchWorkitemWithTotalCount(request)
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

// SearchWorkitemWithTotalCountWithCallback invokes the rdc.SearchWorkitemWithTotalCount API asynchronously
// api document: https://help.aliyun.com/api/rdc/searchworkitemwithtotalcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchWorkitemWithTotalCountWithCallback(request *SearchWorkitemWithTotalCountRequest, callback func(response *SearchWorkitemWithTotalCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchWorkitemWithTotalCountResponse
		var err error
		defer close(result)
		response, err = client.SearchWorkitemWithTotalCount(request)
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

// SearchWorkitemWithTotalCountRequest is the request struct for api SearchWorkitemWithTotalCount
type SearchWorkitemWithTotalCountRequest struct {
	*requests.RpcRequest
	SprintId       requests.Integer `position:"Body" name:"SprintId"`
	CorpIdentifier string           `position:"Query" name:"CorpIdentifier"`
	ToPage         requests.Integer `position:"Body" name:"ToPage"`
	PageSize       requests.Integer `position:"Body" name:"PageSize"`
	Stamp          string           `position:"Body" name:"Stamp"`
	AKProjectId    requests.Integer `position:"Body" name:"AKProjectId"`
}

// SearchWorkitemWithTotalCountResponse is the response struct for api SearchWorkitemWithTotalCount
type SearchWorkitemWithTotalCountResponse struct {
	*responses.BaseResponse
	Code       int        `json:"Code" xml:"Code"`
	Success    bool       `json:"Success" xml:"Success"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	TotalPages int        `json:"TotalPages" xml:"TotalPages"`
	ToPage     int        `json:"ToPage" xml:"ToPage"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	Data       []DataItem `json:"Data" xml:"Data"`
}

// CreateSearchWorkitemWithTotalCountRequest creates a request to invoke SearchWorkitemWithTotalCount API
func CreateSearchWorkitemWithTotalCountRequest() (request *SearchWorkitemWithTotalCountRequest) {
	request = &SearchWorkitemWithTotalCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rdc", "2018-08-21", "SearchWorkitemWithTotalCount", "rdc", "openAPI")
	return
}

// CreateSearchWorkitemWithTotalCountResponse creates a response to parse from SearchWorkitemWithTotalCount response
func CreateSearchWorkitemWithTotalCountResponse() (response *SearchWorkitemWithTotalCountResponse) {
	response = &SearchWorkitemWithTotalCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
