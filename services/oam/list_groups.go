package oam

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

// ListGroups invokes the oam.ListGroups API synchronously
// api document: https://help.aliyun.com/api/oam/listgroups.html
func (client *Client) ListGroups(request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
	response = CreateListGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListGroupsWithChan invokes the oam.ListGroups API asynchronously
// api document: https://help.aliyun.com/api/oam/listgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListGroupsWithChan(request *ListGroupsRequest) (<-chan *ListGroupsResponse, <-chan error) {
	responseChan := make(chan *ListGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGroups(request)
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

// ListGroupsWithCallback invokes the oam.ListGroups API asynchronously
// api document: https://help.aliyun.com/api/oam/listgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListGroupsWithCallback(request *ListGroupsRequest, callback func(response *ListGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListGroups(request)
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

// ListGroupsRequest is the request struct for api ListGroups
type ListGroupsRequest struct {
	*requests.RpcRequest
	OwnerName string           `position:"Query" name:"OwnerName"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	PageIndex requests.Integer `position:"Query" name:"PageIndex"`
}

// ListGroupsResponse is the response struct for api ListGroups
type ListGroupsResponse struct {
	*responses.BaseResponse
	Code     string           `json:"Code" xml:"Code"`
	Message  string           `json:"Message" xml:"Message"`
	PageInfo PageInfo         `json:"PageInfo" xml:"PageInfo"`
	Data     DataInListGroups `json:"Data" xml:"Data"`
}

// CreateListGroupsRequest creates a request to invoke ListGroups API
func CreateListGroupsRequest() (request *ListGroupsRequest) {
	request = &ListGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "ListGroups", "", "")
	request.Method = requests.POST
	return
}

// CreateListGroupsResponse creates a response to parse from ListGroups response
func CreateListGroupsResponse() (response *ListGroupsResponse) {
	response = &ListGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
