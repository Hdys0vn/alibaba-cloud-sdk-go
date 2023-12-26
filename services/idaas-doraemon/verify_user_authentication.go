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

// VerifyUserAuthentication invokes the idaas_doraemon.VerifyUserAuthentication API synchronously
func (client *Client) VerifyUserAuthentication(request *VerifyUserAuthenticationRequest) (response *VerifyUserAuthenticationResponse, err error) {
	response = CreateVerifyUserAuthenticationResponse()
	err = client.DoAction(request, response)
	return
}

// VerifyUserAuthenticationWithChan invokes the idaas_doraemon.VerifyUserAuthentication API asynchronously
func (client *Client) VerifyUserAuthenticationWithChan(request *VerifyUserAuthenticationRequest) (<-chan *VerifyUserAuthenticationResponse, <-chan error) {
	responseChan := make(chan *VerifyUserAuthenticationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VerifyUserAuthentication(request)
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

// VerifyUserAuthenticationWithCallback invokes the idaas_doraemon.VerifyUserAuthentication API asynchronously
func (client *Client) VerifyUserAuthenticationWithCallback(request *VerifyUserAuthenticationRequest, callback func(response *VerifyUserAuthenticationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VerifyUserAuthenticationResponse
		var err error
		defer close(result)
		response, err = client.VerifyUserAuthentication(request)
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

// VerifyUserAuthenticationRequest is the request struct for api VerifyUserAuthentication
type VerifyUserAuthenticationRequest struct {
	*requests.RpcRequest
	LogParams                  string `position:"Query" name:"LogParams"`
	ClientExtendParamsJson     string `position:"Query" name:"ClientExtendParamsJson"`
	UserId                     string `position:"Query" name:"UserId"`
	LogTag                     string `position:"Query" name:"LogTag"`
	ServerExtendParamsJson     string `position:"Query" name:"ServerExtendParamsJson"`
	RequireBindHashBase64      string `position:"Query" name:"RequireBindHashBase64"`
	AuthenticationContext      string `position:"Query" name:"AuthenticationContext"`
	RequireChallengeBase64     string `position:"Query" name:"RequireChallengeBase64"`
	AuthenticatorType          string `position:"Query" name:"AuthenticatorType"`
	ClientExtendParamsJsonSign string `position:"Query" name:"ClientExtendParamsJsonSign"`
	UserSourceIp               string `position:"Query" name:"UserSourceIp"`
	ApplicationExternalId      string `position:"Query" name:"ApplicationExternalId"`
}

// VerifyUserAuthenticationResponse is the response struct for api VerifyUserAuthentication
type VerifyUserAuthenticationResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	VerifyResult           bool                   `json:"VerifyResult" xml:"VerifyResult"`
	EtasSDKString          string                 `json:"EtasSDKString" xml:"EtasSDKString"`
	IdToken                string                 `json:"IdToken" xml:"IdToken"`
	AuthenticateResultInfo AuthenticateResultInfo `json:"AuthenticateResultInfo" xml:"AuthenticateResultInfo"`
}

// CreateVerifyUserAuthenticationRequest creates a request to invoke VerifyUserAuthentication API
func CreateVerifyUserAuthenticationRequest() (request *VerifyUserAuthenticationRequest) {
	request = &VerifyUserAuthenticationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idaas-doraemon", "2021-05-20", "VerifyUserAuthentication", "", "")
	request.Method = requests.POST
	return
}

// CreateVerifyUserAuthenticationResponse creates a response to parse from VerifyUserAuthentication response
func CreateVerifyUserAuthenticationResponse() (response *VerifyUserAuthenticationResponse) {
	response = &VerifyUserAuthenticationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}