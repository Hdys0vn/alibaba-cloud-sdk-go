package ahas_openapi

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

// GetUserApplicationGroups invokes the ahas_openapi.GetUserApplicationGroups API synchronously
func (client *Client) GetUserApplicationGroups(request *GetUserApplicationGroupsRequest) (response *GetUserApplicationGroupsResponse, err error) {
	response = CreateGetUserApplicationGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserApplicationGroupsWithChan invokes the ahas_openapi.GetUserApplicationGroups API asynchronously
func (client *Client) GetUserApplicationGroupsWithChan(request *GetUserApplicationGroupsRequest) (<-chan *GetUserApplicationGroupsResponse, <-chan error) {
	responseChan := make(chan *GetUserApplicationGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserApplicationGroups(request)
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

// GetUserApplicationGroupsWithCallback invokes the ahas_openapi.GetUserApplicationGroups API asynchronously
func (client *Client) GetUserApplicationGroupsWithCallback(request *GetUserApplicationGroupsRequest, callback func(response *GetUserApplicationGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserApplicationGroupsResponse
		var err error
		defer close(result)
		response, err = client.GetUserApplicationGroups(request)
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

// GetUserApplicationGroupsRequest is the request struct for api GetUserApplicationGroups
type GetUserApplicationGroupsRequest struct {
	*requests.RpcRequest
	AhasRegionId  string `position:"Query" name:"AhasRegionId"`
	NameSpace     string `position:"Query" name:"NameSpace"`
	ApplicationId string `position:"Query" name:"ApplicationId"`
}

// GetUserApplicationGroupsResponse is the response struct for api GetUserApplicationGroups
type GetUserApplicationGroupsResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Success        bool     `json:"Success" xml:"Success"`
	Message        string   `json:"Message" xml:"Message"`
	Code           string   `json:"Code" xml:"Code"`
	AppGroups      []string `json:"AppGroups" xml:"AppGroups"`
}

// CreateGetUserApplicationGroupsRequest creates a request to invoke GetUserApplicationGroups API
func CreateGetUserApplicationGroupsRequest() (request *GetUserApplicationGroupsRequest) {
	request = &GetUserApplicationGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "GetUserApplicationGroups", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUserApplicationGroupsResponse creates a response to parse from GetUserApplicationGroups response
func CreateGetUserApplicationGroupsResponse() (response *GetUserApplicationGroupsResponse) {
	response = &GetUserApplicationGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}