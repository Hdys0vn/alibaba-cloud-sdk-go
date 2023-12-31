package domain_intl

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

// CheckDomainSunriseClaim invokes the domain_intl.CheckDomainSunriseClaim API synchronously
// api document: https://help.aliyun.com/api/domain-intl/checkdomainsunriseclaim.html
func (client *Client) CheckDomainSunriseClaim(request *CheckDomainSunriseClaimRequest) (response *CheckDomainSunriseClaimResponse, err error) {
	response = CreateCheckDomainSunriseClaimResponse()
	err = client.DoAction(request, response)
	return
}

// CheckDomainSunriseClaimWithChan invokes the domain_intl.CheckDomainSunriseClaim API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/checkdomainsunriseclaim.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckDomainSunriseClaimWithChan(request *CheckDomainSunriseClaimRequest) (<-chan *CheckDomainSunriseClaimResponse, <-chan error) {
	responseChan := make(chan *CheckDomainSunriseClaimResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckDomainSunriseClaim(request)
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

// CheckDomainSunriseClaimWithCallback invokes the domain_intl.CheckDomainSunriseClaim API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/checkdomainsunriseclaim.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckDomainSunriseClaimWithCallback(request *CheckDomainSunriseClaimRequest, callback func(response *CheckDomainSunriseClaimResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckDomainSunriseClaimResponse
		var err error
		defer close(result)
		response, err = client.CheckDomainSunriseClaim(request)
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

// CheckDomainSunriseClaimRequest is the request struct for api CheckDomainSunriseClaim
type CheckDomainSunriseClaimRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// CheckDomainSunriseClaimResponse is the response struct for api CheckDomainSunriseClaim
type CheckDomainSunriseClaimResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    int    `json:"Result" xml:"Result"`
	ClaimKey  string `json:"ClaimKey" xml:"ClaimKey"`
}

// CreateCheckDomainSunriseClaimRequest creates a request to invoke CheckDomainSunriseClaim API
func CreateCheckDomainSunriseClaimRequest() (request *CheckDomainSunriseClaimRequest) {
	request = &CheckDomainSunriseClaimRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "CheckDomainSunriseClaim", "domain", "openAPI")
	return
}

// CreateCheckDomainSunriseClaimResponse creates a response to parse from CheckDomainSunriseClaim response
func CreateCheckDomainSunriseClaimResponse() (response *CheckDomainSunriseClaimResponse) {
	response = &CheckDomainSunriseClaimResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
