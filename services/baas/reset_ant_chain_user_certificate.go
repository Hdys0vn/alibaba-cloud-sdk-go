package baas

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

// ResetAntChainUserCertificate invokes the baas.ResetAntChainUserCertificate API synchronously
// api document: https://help.aliyun.com/api/baas/resetantchainusercertificate.html
func (client *Client) ResetAntChainUserCertificate(request *ResetAntChainUserCertificateRequest) (response *ResetAntChainUserCertificateResponse, err error) {
	response = CreateResetAntChainUserCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// ResetAntChainUserCertificateWithChan invokes the baas.ResetAntChainUserCertificate API asynchronously
// api document: https://help.aliyun.com/api/baas/resetantchainusercertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetAntChainUserCertificateWithChan(request *ResetAntChainUserCertificateRequest) (<-chan *ResetAntChainUserCertificateResponse, <-chan error) {
	responseChan := make(chan *ResetAntChainUserCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetAntChainUserCertificate(request)
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

// ResetAntChainUserCertificateWithCallback invokes the baas.ResetAntChainUserCertificate API asynchronously
// api document: https://help.aliyun.com/api/baas/resetantchainusercertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetAntChainUserCertificateWithCallback(request *ResetAntChainUserCertificateRequest, callback func(response *ResetAntChainUserCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetAntChainUserCertificateResponse
		var err error
		defer close(result)
		response, err = client.ResetAntChainUserCertificate(request)
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

// ResetAntChainUserCertificateRequest is the request struct for api ResetAntChainUserCertificate
type ResetAntChainUserCertificateRequest struct {
	*requests.RpcRequest
	AntChainId string `position:"Body" name:"AntChainId"`
	Username   string `position:"Body" name:"Username"`
}

// ResetAntChainUserCertificateResponse is the response struct for api ResetAntChainUserCertificate
type ResetAntChainUserCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateResetAntChainUserCertificateRequest creates a request to invoke ResetAntChainUserCertificate API
func CreateResetAntChainUserCertificateRequest() (request *ResetAntChainUserCertificateRequest) {
	request = &ResetAntChainUserCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "ResetAntChainUserCertificate", "baas", "openAPI")
	return
}

// CreateResetAntChainUserCertificateResponse creates a response to parse from ResetAntChainUserCertificate response
func CreateResetAntChainUserCertificateResponse() (response *ResetAntChainUserCertificateResponse) {
	response = &ResetAntChainUserCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
