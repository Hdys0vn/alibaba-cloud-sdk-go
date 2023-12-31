package domain

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

// QueryDomainGroupList invokes the domain.QueryDomainGroupList API synchronously
func (client *Client) QueryDomainGroupList(request *QueryDomainGroupListRequest) (response *QueryDomainGroupListResponse, err error) {
	response = CreateQueryDomainGroupListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDomainGroupListWithChan invokes the domain.QueryDomainGroupList API asynchronously
func (client *Client) QueryDomainGroupListWithChan(request *QueryDomainGroupListRequest) (<-chan *QueryDomainGroupListResponse, <-chan error) {
	responseChan := make(chan *QueryDomainGroupListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDomainGroupList(request)
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

// QueryDomainGroupListWithCallback invokes the domain.QueryDomainGroupList API asynchronously
func (client *Client) QueryDomainGroupListWithCallback(request *QueryDomainGroupListRequest, callback func(response *QueryDomainGroupListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDomainGroupListResponse
		var err error
		defer close(result)
		response, err = client.QueryDomainGroupList(request)
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

// QueryDomainGroupListRequest is the request struct for api QueryDomainGroupList
type QueryDomainGroupListRequest struct {
	*requests.RpcRequest
	ShowDeletingGroup requests.Boolean `position:"Query" name:"ShowDeletingGroup"`
	UserClientIp      string           `position:"Query" name:"UserClientIp"`
	DomainGroupName   string           `position:"Query" name:"DomainGroupName"`
	Lang              string           `position:"Query" name:"Lang"`
}

// QueryDomainGroupListResponse is the response struct for api QueryDomainGroupList
type QueryDomainGroupListResponse struct {
	*responses.BaseResponse
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Data      DataInQueryDomainGroupList `json:"Data" xml:"Data"`
}

// CreateQueryDomainGroupListRequest creates a request to invoke QueryDomainGroupList API
func CreateQueryDomainGroupListRequest() (request *QueryDomainGroupListRequest) {
	request = &QueryDomainGroupListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryDomainGroupList", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDomainGroupListResponse creates a response to parse from QueryDomainGroupList response
func CreateQueryDomainGroupListResponse() (response *QueryDomainGroupListResponse) {
	response = &QueryDomainGroupListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
