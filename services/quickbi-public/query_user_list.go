package quickbi_public

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

// QueryUserList invokes the quickbi_public.QueryUserList API synchronously
func (client *Client) QueryUserList(request *QueryUserListRequest) (response *QueryUserListResponse, err error) {
	response = CreateQueryUserListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryUserListWithChan invokes the quickbi_public.QueryUserList API asynchronously
func (client *Client) QueryUserListWithChan(request *QueryUserListRequest) (<-chan *QueryUserListResponse, <-chan error) {
	responseChan := make(chan *QueryUserListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryUserList(request)
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

// QueryUserListWithCallback invokes the quickbi_public.QueryUserList API asynchronously
func (client *Client) QueryUserListWithCallback(request *QueryUserListRequest, callback func(response *QueryUserListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryUserListResponse
		var err error
		defer close(result)
		response, err = client.QueryUserList(request)
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

// QueryUserListRequest is the request struct for api QueryUserList
type QueryUserListRequest struct {
	*requests.RpcRequest
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	AccountType requests.Integer `position:"Query" name:"AccountType"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	SignType    string           `position:"Query" name:"SignType"`
	Keyword     string           `position:"Query" name:"Keyword"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
}

// QueryUserListResponse is the response struct for api QueryUserList
type QueryUserListResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryUserListRequest creates a request to invoke QueryUserList API
func CreateQueryUserListRequest() (request *QueryUserListRequest) {
	request = &QueryUserListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryUserList", "2.2.0", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryUserListResponse creates a response to parse from QueryUserList response
func CreateQueryUserListResponse() (response *QueryUserListResponse) {
	response = &QueryUserListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}