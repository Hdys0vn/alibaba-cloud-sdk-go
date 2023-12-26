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

// GetUserGroupInfo invokes the quickbi_public.GetUserGroupInfo API synchronously
func (client *Client) GetUserGroupInfo(request *GetUserGroupInfoRequest) (response *GetUserGroupInfoResponse, err error) {
	response = CreateGetUserGroupInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserGroupInfoWithChan invokes the quickbi_public.GetUserGroupInfo API asynchronously
func (client *Client) GetUserGroupInfoWithChan(request *GetUserGroupInfoRequest) (<-chan *GetUserGroupInfoResponse, <-chan error) {
	responseChan := make(chan *GetUserGroupInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserGroupInfo(request)
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

// GetUserGroupInfoWithCallback invokes the quickbi_public.GetUserGroupInfo API asynchronously
func (client *Client) GetUserGroupInfoWithCallback(request *GetUserGroupInfoRequest, callback func(response *GetUserGroupInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserGroupInfoResponse
		var err error
		defer close(result)
		response, err = client.GetUserGroupInfo(request)
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

// GetUserGroupInfoRequest is the request struct for api GetUserGroupInfo
type GetUserGroupInfoRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	Keyword     string `position:"Query" name:"Keyword"`
}

// GetUserGroupInfoResponse is the response struct for api GetUserGroupInfo
type GetUserGroupInfoResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    []Data `json:"Result" xml:"Result"`
}

// CreateGetUserGroupInfoRequest creates a request to invoke GetUserGroupInfo API
func CreateGetUserGroupInfoRequest() (request *GetUserGroupInfoRequest) {
	request = &GetUserGroupInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "GetUserGroupInfo", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUserGroupInfoResponse creates a response to parse from GetUserGroupInfo response
func CreateGetUserGroupInfoResponse() (response *GetUserGroupInfoResponse) {
	response = &GetUserGroupInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
