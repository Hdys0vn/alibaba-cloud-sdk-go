package cloudcallcenter

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

// AuthorizedVn invokes the cloudcallcenter.AuthorizedVn API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/authorizedvn.html
func (client *Client) AuthorizedVn(request *AuthorizedVnRequest) (response *AuthorizedVnResponse, err error) {
	response = CreateAuthorizedVnResponse()
	err = client.DoAction(request, response)
	return
}

// AuthorizedVnWithChan invokes the cloudcallcenter.AuthorizedVn API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/authorizedvn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AuthorizedVnWithChan(request *AuthorizedVnRequest) (<-chan *AuthorizedVnResponse, <-chan error) {
	responseChan := make(chan *AuthorizedVnResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AuthorizedVn(request)
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

// AuthorizedVnWithCallback invokes the cloudcallcenter.AuthorizedVn API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/authorizedvn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AuthorizedVnWithCallback(request *AuthorizedVnRequest, callback func(response *AuthorizedVnResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AuthorizedVnResponse
		var err error
		defer close(result)
		response, err = client.AuthorizedVn(request)
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

// AuthorizedVnRequest is the request struct for api AuthorizedVn
type AuthorizedVnRequest struct {
	*requests.RpcRequest
	ConversationId string `position:"Query" name:"ConversationId"`
	InitialContext string `position:"Query" name:"InitialContext"`
	InstanceId     string `position:"Query" name:"InstanceId"`
}

// AuthorizedVnResponse is the response struct for api AuthorizedVn
type AuthorizedVnResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TextResponse  string `json:"TextResponse" xml:"TextResponse"`
	Interruptible bool   `json:"Interruptible" xml:"Interruptible"`
	Action        string `json:"Action" xml:"Action"`
	ActionParams  string `json:"ActionParams" xml:"ActionParams"`
}

// CreateAuthorizedVnRequest creates a request to invoke AuthorizedVn API
func CreateAuthorizedVnRequest() (request *AuthorizedVnRequest) {
	request = &AuthorizedVnRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "AuthorizedVn", "", "")
	request.Method = requests.GET
	return
}

// CreateAuthorizedVnResponse creates a response to parse from AuthorizedVn response
func CreateAuthorizedVnResponse() (response *AuthorizedVnResponse) {
	response = &AuthorizedVnResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
