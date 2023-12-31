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

// SetCdnDomainCSRCertificate invokes the cdn.SetCdnDomainCSRCertificate API synchronously
func (client *Client) SetCdnDomainCSRCertificate(request *SetCdnDomainCSRCertificateRequest) (response *SetCdnDomainCSRCertificateResponse, err error) {
	response = CreateSetCdnDomainCSRCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// SetCdnDomainCSRCertificateWithChan invokes the cdn.SetCdnDomainCSRCertificate API asynchronously
func (client *Client) SetCdnDomainCSRCertificateWithChan(request *SetCdnDomainCSRCertificateRequest) (<-chan *SetCdnDomainCSRCertificateResponse, <-chan error) {
	responseChan := make(chan *SetCdnDomainCSRCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetCdnDomainCSRCertificate(request)
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

// SetCdnDomainCSRCertificateWithCallback invokes the cdn.SetCdnDomainCSRCertificate API asynchronously
func (client *Client) SetCdnDomainCSRCertificateWithCallback(request *SetCdnDomainCSRCertificateRequest, callback func(response *SetCdnDomainCSRCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetCdnDomainCSRCertificateResponse
		var err error
		defer close(result)
		response, err = client.SetCdnDomainCSRCertificate(request)
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

// SetCdnDomainCSRCertificateRequest is the request struct for api SetCdnDomainCSRCertificate
type SetCdnDomainCSRCertificateRequest struct {
	*requests.RpcRequest
	ServerCertificate string `position:"Query" name:"ServerCertificate"`
	DomainName        string `position:"Query" name:"DomainName"`
}

// SetCdnDomainCSRCertificateResponse is the response struct for api SetCdnDomainCSRCertificate
type SetCdnDomainCSRCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetCdnDomainCSRCertificateRequest creates a request to invoke SetCdnDomainCSRCertificate API
func CreateSetCdnDomainCSRCertificateRequest() (request *SetCdnDomainCSRCertificateRequest) {
	request = &SetCdnDomainCSRCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "SetCdnDomainCSRCertificate", "", "")
	request.Method = requests.POST
	return
}

// CreateSetCdnDomainCSRCertificateResponse creates a response to parse from SetCdnDomainCSRCertificate response
func CreateSetCdnDomainCSRCertificateResponse() (response *SetCdnDomainCSRCertificateResponse) {
	response = &SetCdnDomainCSRCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
