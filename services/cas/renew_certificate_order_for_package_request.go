package cas

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

// RenewCertificateOrderForPackageRequest invokes the cas.RenewCertificateOrderForPackageRequest API synchronously
func (client *Client) RenewCertificateOrderForPackageRequest(request *RenewCertificateOrderForPackageRequestRequest) (response *RenewCertificateOrderForPackageRequestResponse, err error) {
	response = CreateRenewCertificateOrderForPackageRequestResponse()
	err = client.DoAction(request, response)
	return
}

// RenewCertificateOrderForPackageRequestWithChan invokes the cas.RenewCertificateOrderForPackageRequest API asynchronously
func (client *Client) RenewCertificateOrderForPackageRequestWithChan(request *RenewCertificateOrderForPackageRequestRequest) (<-chan *RenewCertificateOrderForPackageRequestResponse, <-chan error) {
	responseChan := make(chan *RenewCertificateOrderForPackageRequestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenewCertificateOrderForPackageRequest(request)
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

// RenewCertificateOrderForPackageRequestWithCallback invokes the cas.RenewCertificateOrderForPackageRequest API asynchronously
func (client *Client) RenewCertificateOrderForPackageRequestWithCallback(request *RenewCertificateOrderForPackageRequestRequest, callback func(response *RenewCertificateOrderForPackageRequestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenewCertificateOrderForPackageRequestResponse
		var err error
		defer close(result)
		response, err = client.RenewCertificateOrderForPackageRequest(request)
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

// RenewCertificateOrderForPackageRequestRequest is the request struct for api RenewCertificateOrderForPackageRequest
type RenewCertificateOrderForPackageRequestRequest struct {
	*requests.RpcRequest
	Csr      string           `position:"Query" name:"Csr"`
	OrderId  requests.Integer `position:"Query" name:"OrderId"`
	SourceIp string           `position:"Query" name:"SourceIp"`
}

// RenewCertificateOrderForPackageRequestResponse is the response struct for api RenewCertificateOrderForPackageRequest
type RenewCertificateOrderForPackageRequestResponse struct {
	*responses.BaseResponse
	OrderId   int64  `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRenewCertificateOrderForPackageRequestRequest creates a request to invoke RenewCertificateOrderForPackageRequest API
func CreateRenewCertificateOrderForPackageRequestRequest() (request *RenewCertificateOrderForPackageRequestRequest) {
	request = &RenewCertificateOrderForPackageRequestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2020-04-07", "RenewCertificateOrderForPackageRequest", "cas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRenewCertificateOrderForPackageRequestResponse creates a response to parse from RenewCertificateOrderForPackageRequest response
func CreateRenewCertificateOrderForPackageRequestResponse() (response *RenewCertificateOrderForPackageRequestResponse) {
	response = &RenewCertificateOrderForPackageRequestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}