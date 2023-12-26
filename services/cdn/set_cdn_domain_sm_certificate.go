package cdn

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

// SetCdnDomainSMCertificate invokes the cdn.SetCdnDomainSMCertificate API synchronously
func (client *Client) SetCdnDomainSMCertificate(request *SetCdnDomainSMCertificateRequest) (response *SetCdnDomainSMCertificateResponse, err error) {
	response = CreateSetCdnDomainSMCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// SetCdnDomainSMCertificateWithChan invokes the cdn.SetCdnDomainSMCertificate API asynchronously
func (client *Client) SetCdnDomainSMCertificateWithChan(request *SetCdnDomainSMCertificateRequest) (<-chan *SetCdnDomainSMCertificateResponse, <-chan error) {
	responseChan := make(chan *SetCdnDomainSMCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetCdnDomainSMCertificate(request)
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

// SetCdnDomainSMCertificateWithCallback invokes the cdn.SetCdnDomainSMCertificate API asynchronously
func (client *Client) SetCdnDomainSMCertificateWithCallback(request *SetCdnDomainSMCertificateRequest, callback func(response *SetCdnDomainSMCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetCdnDomainSMCertificateResponse
		var err error
		defer close(result)
		response, err = client.SetCdnDomainSMCertificate(request)
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

// SetCdnDomainSMCertificateRequest is the request struct for api SetCdnDomainSMCertificate
type SetCdnDomainSMCertificateRequest struct {
	*requests.RpcRequest
	SSLProtocol    string           `position:"Query" name:"SSLProtocol"`
	DomainName     string           `position:"Query" name:"DomainName"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken  string           `position:"Query" name:"SecurityToken"`
	CertIdentifier string           `position:"Query" name:"CertIdentifier"`
}

// SetCdnDomainSMCertificateResponse is the response struct for api SetCdnDomainSMCertificate
type SetCdnDomainSMCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetCdnDomainSMCertificateRequest creates a request to invoke SetCdnDomainSMCertificate API
func CreateSetCdnDomainSMCertificateRequest() (request *SetCdnDomainSMCertificateRequest) {
	request = &SetCdnDomainSMCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "SetCdnDomainSMCertificate", "", "")
	request.Method = requests.POST
	return
}

// CreateSetCdnDomainSMCertificateResponse creates a response to parse from SetCdnDomainSMCertificate response
func CreateSetCdnDomainSMCertificateResponse() (response *SetCdnDomainSMCertificateResponse) {
	response = &SetCdnDomainSMCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
