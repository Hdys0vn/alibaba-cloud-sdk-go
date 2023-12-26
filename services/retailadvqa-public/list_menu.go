package retailadvqa_public

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

// ListMenu invokes the retailadvqa_public.ListMenu API synchronously
func (client *Client) ListMenu(request *ListMenuRequest) (response *ListMenuResponse, err error) {
	response = CreateListMenuResponse()
	err = client.DoAction(request, response)
	return
}

// ListMenuWithChan invokes the retailadvqa_public.ListMenu API asynchronously
func (client *Client) ListMenuWithChan(request *ListMenuRequest) (<-chan *ListMenuResponse, <-chan error) {
	responseChan := make(chan *ListMenuResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMenu(request)
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

// ListMenuWithCallback invokes the retailadvqa_public.ListMenu API asynchronously
func (client *Client) ListMenuWithCallback(request *ListMenuRequest, callback func(response *ListMenuResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMenuResponse
		var err error
		defer close(result)
		response, err = client.ListMenu(request)
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

// ListMenuRequest is the request struct for api ListMenu
type ListMenuRequest struct {
	*requests.RpcRequest
	AccessId    string `position:"Query" name:"AccessId"`
	TenantId    string `position:"Query" name:"TenantId"`
	RoleSign    string `position:"Query" name:"RoleSign"`
	WorkspaceId string `position:"Query" name:"WorkspaceId"`
}

// ListMenuResponse is the response struct for api ListMenu
type ListMenuResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	ErrorDesc string     `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string     `json:"TraceId" xml:"TraceId"`
	ErrorCode string     `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool       `json:"Success" xml:"Success"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListMenuRequest creates a request to invoke ListMenu API
func CreateListMenuRequest() (request *ListMenuRequest) {
	request = &ListMenuRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "ListMenu", "", "")
	request.Method = requests.POST
	return
}

// CreateListMenuResponse creates a response to parse from ListMenu response
func CreateListMenuResponse() (response *ListMenuResponse) {
	response = &ListMenuResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
