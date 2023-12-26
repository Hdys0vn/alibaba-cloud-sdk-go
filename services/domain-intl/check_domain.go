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

// CheckDomain invokes the domain_intl.CheckDomain API synchronously
// api document: https://help.aliyun.com/api/domain-intl/checkdomain.html
func (client *Client) CheckDomain(request *CheckDomainRequest) (response *CheckDomainResponse, err error) {
	response = CreateCheckDomainResponse()
	err = client.DoAction(request, response)
	return
}

// CheckDomainWithChan invokes the domain_intl.CheckDomain API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/checkdomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckDomainWithChan(request *CheckDomainRequest) (<-chan *CheckDomainResponse, <-chan error) {
	responseChan := make(chan *CheckDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckDomain(request)
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

// CheckDomainWithCallback invokes the domain_intl.CheckDomain API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/checkdomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckDomainWithCallback(request *CheckDomainRequest, callback func(response *CheckDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckDomainResponse
		var err error
		defer close(result)
		response, err = client.CheckDomain(request)
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

// CheckDomainRequest is the request struct for api CheckDomain
type CheckDomainRequest struct {
	*requests.RpcRequest
	FeeCurrency  string           `position:"Query" name:"FeeCurrency"`
	FeePeriod    requests.Integer `position:"Query" name:"FeePeriod"`
	DomainName   string           `position:"Query" name:"DomainName"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	FeeCommand   string           `position:"Query" name:"FeeCommand"`
	Lang         string           `position:"Query" name:"Lang"`
}

// CheckDomainResponse is the response struct for api CheckDomain
type CheckDomainResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DomainName   string `json:"DomainName" xml:"DomainName"`
	Avail        string `json:"Avail" xml:"Avail"`
	Premium      string `json:"Premium" xml:"Premium"`
	Reason       string `json:"Reason" xml:"Reason"`
	Price        int    `json:"Price" xml:"Price"`
	DynamicCheck bool   `json:"DynamicCheck" xml:"DynamicCheck"`
}

// CreateCheckDomainRequest creates a request to invoke CheckDomain API
func CreateCheckDomainRequest() (request *CheckDomainRequest) {
	request = &CheckDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "CheckDomain", "domain", "openAPI")
	return
}

// CreateCheckDomainResponse creates a response to parse from CheckDomain response
func CreateCheckDomainResponse() (response *CheckDomainResponse) {
	response = &CheckDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
