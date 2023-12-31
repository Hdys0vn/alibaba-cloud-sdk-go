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

// UnFreezeMember invokes the lto.UnFreezeMember API synchronously
func (client *Client) UnFreezeMember(request *UnFreezeMemberRequest) (response *UnFreezeMemberResponse, err error) {
	response = CreateUnFreezeMemberResponse()
	err = client.DoAction(request, response)
	return
}

// UnFreezeMemberWithChan invokes the lto.UnFreezeMember API asynchronously
func (client *Client) UnFreezeMemberWithChan(request *UnFreezeMemberRequest) (<-chan *UnFreezeMemberResponse, <-chan error) {
	responseChan := make(chan *UnFreezeMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnFreezeMember(request)
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

// UnFreezeMemberWithCallback invokes the lto.UnFreezeMember API asynchronously
func (client *Client) UnFreezeMemberWithCallback(request *UnFreezeMemberRequest, callback func(response *UnFreezeMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnFreezeMemberResponse
		var err error
		defer close(result)
		response, err = client.UnFreezeMember(request)
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

// UnFreezeMemberRequest is the request struct for api UnFreezeMember
type UnFreezeMemberRequest struct {
	*requests.RpcRequest
	MemberId string `position:"Query" name:"MemberId"`
}

// UnFreezeMemberResponse is the response struct for api UnFreezeMember
type UnFreezeMemberResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateUnFreezeMemberRequest creates a request to invoke UnFreezeMember API
func CreateUnFreezeMemberRequest() (request *UnFreezeMemberRequest) {
	request = &UnFreezeMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "UnFreezeMember", "", "")
	request.Method = requests.POST
	return
}

// CreateUnFreezeMemberResponse creates a response to parse from UnFreezeMember response
func CreateUnFreezeMemberResponse() (response *UnFreezeMemberResponse) {
	response = &UnFreezeMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
