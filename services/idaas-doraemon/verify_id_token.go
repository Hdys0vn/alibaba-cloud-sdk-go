package idaas_doraemon

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

// VerifyIdToken invokes the idaas_doraemon.VerifyIdToken API synchronously
func (client *Client) VerifyIdToken(request *VerifyIdTokenRequest) (response *VerifyIdTokenResponse, err error) {
	response = CreateVerifyIdTokenResponse()
	err = client.DoAction(request, response)
	return
}

// VerifyIdTokenWithChan invokes the idaas_doraemon.VerifyIdToken API asynchronously
func (client *Client) VerifyIdTokenWithChan(request *VerifyIdTokenRequest) (<-chan *VerifyIdTokenResponse, <-chan error) {
	responseChan := make(chan *VerifyIdTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VerifyIdToken(request)
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

// VerifyIdTokenWithCallback invokes the idaas_doraemon.VerifyIdToken API asynchronously
func (client *Client) VerifyIdTokenWithCallback(request *VerifyIdTokenRequest, callback func(response *VerifyIdTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VerifyIdTokenResponse
		var err error
		defer close(result)
		response, err = client.VerifyIdToken(request)
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

// VerifyIdTokenRequest is the request struct for api VerifyIdToken
type VerifyIdTokenRequest struct {
	*requests.RpcRequest
	JwtIdToken            string `position:"Query" name:"JwtIdToken"`
	ApplicationExternalId string `position:"Query" name:"ApplicationExternalId"`
}

// VerifyIdTokenResponse is the response struct for api VerifyIdToken
type VerifyIdTokenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	UserId    string `json:"UserId" xml:"UserId"`
}

// CreateVerifyIdTokenRequest creates a request to invoke VerifyIdToken API
func CreateVerifyIdTokenRequest() (request *VerifyIdTokenRequest) {
	request = &VerifyIdTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idaas-doraemon", "2021-05-20", "VerifyIdToken", "", "")
	request.Method = requests.POST
	return
}

// CreateVerifyIdTokenResponse creates a response to parse from VerifyIdToken response
func CreateVerifyIdTokenResponse() (response *VerifyIdTokenResponse) {
	response = &VerifyIdTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
