package kms

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

// GetCertificate invokes the kms.GetCertificate API synchronously
func (client *Client) GetCertificate(request *GetCertificateRequest) (response *GetCertificateResponse, err error) {
	response = CreateGetCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// GetCertificateWithChan invokes the kms.GetCertificate API asynchronously
func (client *Client) GetCertificateWithChan(request *GetCertificateRequest) (<-chan *GetCertificateResponse, <-chan error) {
	responseChan := make(chan *GetCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCertificate(request)
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

// GetCertificateWithCallback invokes the kms.GetCertificate API asynchronously
func (client *Client) GetCertificateWithCallback(request *GetCertificateRequest, callback func(response *GetCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCertificateResponse
		var err error
		defer close(result)
		response, err = client.GetCertificate(request)
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

// GetCertificateRequest is the request struct for api GetCertificate
type GetCertificateRequest struct {
	*requests.RpcRequest
	CertificateId string `position:"Query" name:"CertificateId"`
}

// GetCertificateResponse is the response struct for api GetCertificate
type GetCertificateResponse struct {
	*responses.BaseResponse
	CertificateChain string `json:"CertificateChain" xml:"CertificateChain"`
	Certificate      string `json:"Certificate" xml:"Certificate"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	CertificateId    string `json:"CertificateId" xml:"CertificateId"`
	Csr              string `json:"Csr" xml:"Csr"`
}

// CreateGetCertificateRequest creates a request to invoke GetCertificate API
func CreateGetCertificateRequest() (request *GetCertificateRequest) {
	request = &GetCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "GetCertificate", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetCertificateResponse creates a response to parse from GetCertificate response
func CreateGetCertificateResponse() (response *GetCertificateResponse) {
	response = &GetCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
