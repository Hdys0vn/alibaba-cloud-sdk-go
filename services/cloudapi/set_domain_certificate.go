package cloudapi

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

// SetDomainCertificate invokes the cloudapi.SetDomainCertificate API synchronously
func (client *Client) SetDomainCertificate(request *SetDomainCertificateRequest) (response *SetDomainCertificateResponse, err error) {
	response = CreateSetDomainCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// SetDomainCertificateWithChan invokes the cloudapi.SetDomainCertificate API asynchronously
func (client *Client) SetDomainCertificateWithChan(request *SetDomainCertificateRequest) (<-chan *SetDomainCertificateResponse, <-chan error) {
	responseChan := make(chan *SetDomainCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDomainCertificate(request)
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

// SetDomainCertificateWithCallback invokes the cloudapi.SetDomainCertificate API asynchronously
func (client *Client) SetDomainCertificateWithCallback(request *SetDomainCertificateRequest, callback func(response *SetDomainCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDomainCertificateResponse
		var err error
		defer close(result)
		response, err = client.SetDomainCertificate(request)
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

// SetDomainCertificateRequest is the request struct for api SetDomainCertificate
type SetDomainCertificateRequest struct {
	*requests.RpcRequest
	CertificatePrivateKey string `position:"Query" name:"CertificatePrivateKey"`
	GroupId               string `position:"Query" name:"GroupId"`
	DomainName            string `position:"Query" name:"DomainName"`
	CertificateBody       string `position:"Query" name:"CertificateBody"`
	SslVerifyDepth        string `position:"Query" name:"SslVerifyDepth"`
	CaCertificateBody     string `position:"Query" name:"CaCertificateBody"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	CertificateName       string `position:"Query" name:"CertificateName"`
}

// SetDomainCertificateResponse is the response struct for api SetDomainCertificate
type SetDomainCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDomainCertificateRequest creates a request to invoke SetDomainCertificate API
func CreateSetDomainCertificateRequest() (request *SetDomainCertificateRequest) {
	request = &SetDomainCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "SetDomainCertificate", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetDomainCertificateResponse creates a response to parse from SetDomainCertificate response
func CreateSetDomainCertificateResponse() (response *SetDomainCertificateResponse) {
	response = &SetDomainCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}