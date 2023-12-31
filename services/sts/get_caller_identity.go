package sts

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

// GetCallerIdentity invokes the sts.GetCallerIdentity API synchronously
func (client *Client) GetCallerIdentity(request *GetCallerIdentityRequest) (response *GetCallerIdentityResponse, err error) {
	response = CreateGetCallerIdentityResponse()
	err = client.DoAction(request, response)
	return
}

// GetCallerIdentityWithChan invokes the sts.GetCallerIdentity API asynchronously
func (client *Client) GetCallerIdentityWithChan(request *GetCallerIdentityRequest) (<-chan *GetCallerIdentityResponse, <-chan error) {
	responseChan := make(chan *GetCallerIdentityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCallerIdentity(request)
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

// GetCallerIdentityWithCallback invokes the sts.GetCallerIdentity API asynchronously
func (client *Client) GetCallerIdentityWithCallback(request *GetCallerIdentityRequest, callback func(response *GetCallerIdentityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCallerIdentityResponse
		var err error
		defer close(result)
		response, err = client.GetCallerIdentity(request)
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

// GetCallerIdentityRequest is the request struct for api GetCallerIdentity
type GetCallerIdentityRequest struct {
	*requests.RpcRequest
}

// GetCallerIdentityResponse is the response struct for api GetCallerIdentity
type GetCallerIdentityResponse struct {
	*responses.BaseResponse
	IdentityType string `json:"IdentityType" xml:"IdentityType"`
	AccountId    string `json:"AccountId" xml:"AccountId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	PrincipalId  string `json:"PrincipalId" xml:"PrincipalId"`
	UserId       string `json:"UserId" xml:"UserId"`
	Arn          string `json:"Arn" xml:"Arn"`
	RoleId       string `json:"RoleId" xml:"RoleId"`
}

// CreateGetCallerIdentityRequest creates a request to invoke GetCallerIdentity API
func CreateGetCallerIdentityRequest() (request *GetCallerIdentityRequest) {
	request = &GetCallerIdentityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sts", "2015-04-01", "GetCallerIdentity", "", "")
	request.Method = requests.POST
	return
}

// CreateGetCallerIdentityResponse creates a response to parse from GetCallerIdentity response
func CreateGetCallerIdentityResponse() (response *GetCallerIdentityResponse) {
	response = &GetCallerIdentityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
