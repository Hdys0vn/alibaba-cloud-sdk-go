package live

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

// DescribeLiveDomainCertificateInfo invokes the live.DescribeLiveDomainCertificateInfo API synchronously
func (client *Client) DescribeLiveDomainCertificateInfo(request *DescribeLiveDomainCertificateInfoRequest) (response *DescribeLiveDomainCertificateInfoResponse, err error) {
	response = CreateDescribeLiveDomainCertificateInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainCertificateInfoWithChan invokes the live.DescribeLiveDomainCertificateInfo API asynchronously
func (client *Client) DescribeLiveDomainCertificateInfoWithChan(request *DescribeLiveDomainCertificateInfoRequest) (<-chan *DescribeLiveDomainCertificateInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainCertificateInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainCertificateInfo(request)
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

// DescribeLiveDomainCertificateInfoWithCallback invokes the live.DescribeLiveDomainCertificateInfo API asynchronously
func (client *Client) DescribeLiveDomainCertificateInfoWithCallback(request *DescribeLiveDomainCertificateInfoRequest, callback func(response *DescribeLiveDomainCertificateInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainCertificateInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainCertificateInfo(request)
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

// DescribeLiveDomainCertificateInfoRequest is the request struct for api DescribeLiveDomainCertificateInfo
type DescribeLiveDomainCertificateInfoRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveDomainCertificateInfoResponse is the response struct for api DescribeLiveDomainCertificateInfo
type DescribeLiveDomainCertificateInfoResponse struct {
	*responses.BaseResponse
	RequestId string                                       `json:"RequestId" xml:"RequestId"`
	CertInfos CertInfosInDescribeLiveDomainCertificateInfo `json:"CertInfos" xml:"CertInfos"`
}

// CreateDescribeLiveDomainCertificateInfoRequest creates a request to invoke DescribeLiveDomainCertificateInfo API
func CreateDescribeLiveDomainCertificateInfoRequest() (request *DescribeLiveDomainCertificateInfoRequest) {
	request = &DescribeLiveDomainCertificateInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainCertificateInfo", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainCertificateInfoResponse creates a response to parse from DescribeLiveDomainCertificateInfo response
func CreateDescribeLiveDomainCertificateInfoResponse() (response *DescribeLiveDomainCertificateInfoResponse) {
	response = &DescribeLiveDomainCertificateInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}