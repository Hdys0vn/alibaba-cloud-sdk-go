package lto

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

// ListAllMember invokes the lto.ListAllMember API synchronously
func (client *Client) ListAllMember(request *ListAllMemberRequest) (response *ListAllMemberResponse, err error) {
	response = CreateListAllMemberResponse()
	err = client.DoAction(request, response)
	return
}

// ListAllMemberWithChan invokes the lto.ListAllMember API asynchronously
func (client *Client) ListAllMemberWithChan(request *ListAllMemberRequest) (<-chan *ListAllMemberResponse, <-chan error) {
	responseChan := make(chan *ListAllMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAllMember(request)
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

// ListAllMemberWithCallback invokes the lto.ListAllMember API asynchronously
func (client *Client) ListAllMemberWithCallback(request *ListAllMemberRequest, callback func(response *ListAllMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAllMemberResponse
		var err error
		defer close(result)
		response, err = client.ListAllMember(request)
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

// ListAllMemberRequest is the request struct for api ListAllMember
type ListAllMemberRequest struct {
	*requests.RpcRequest
}

// ListAllMemberResponse is the response struct for api ListAllMember
type ListAllMemberResponse struct {
	*responses.BaseResponse
	Code           string       `json:"Code" xml:"Code"`
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string       `json:"Message" xml:"Message"`
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	Data           []MemberInfo `json:"Data" xml:"Data"`
}

// CreateListAllMemberRequest creates a request to invoke ListAllMember API
func CreateListAllMemberRequest() (request *ListAllMemberRequest) {
	request = &ListAllMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "ListAllMember", "", "")
	request.Method = requests.POST
	return
}

// CreateListAllMemberResponse creates a response to parse from ListAllMember response
func CreateListAllMemberResponse() (response *ListAllMemberResponse) {
	response = &ListAllMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
