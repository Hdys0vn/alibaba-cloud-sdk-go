package dcdn

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

// DescribeDcdnUserDomains invokes the dcdn.DescribeDcdnUserDomains API synchronously
func (client *Client) DescribeDcdnUserDomains(request *DescribeDcdnUserDomainsRequest) (response *DescribeDcdnUserDomainsResponse, err error) {
	response = CreateDescribeDcdnUserDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnUserDomainsWithChan invokes the dcdn.DescribeDcdnUserDomains API asynchronously
func (client *Client) DescribeDcdnUserDomainsWithChan(request *DescribeDcdnUserDomainsRequest) (<-chan *DescribeDcdnUserDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnUserDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnUserDomains(request)
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

// DescribeDcdnUserDomainsWithCallback invokes the dcdn.DescribeDcdnUserDomains API asynchronously
func (client *Client) DescribeDcdnUserDomainsWithCallback(request *DescribeDcdnUserDomainsRequest, callback func(response *DescribeDcdnUserDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnUserDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnUserDomains(request)
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

// DescribeDcdnUserDomainsRequest is the request struct for api DescribeDcdnUserDomains
type DescribeDcdnUserDomainsRequest struct {
	*requests.RpcRequest
	PageNumber       requests.Integer              `position:"Query" name:"PageNumber"`
	CheckDomainShow  requests.Boolean              `position:"Query" name:"CheckDomainShow"`
	ResourceGroupId  string                        `position:"Query" name:"ResourceGroupId"`
	SecurityToken    string                        `position:"Query" name:"SecurityToken"`
	ChangeEndTime    string                        `position:"Query" name:"ChangeEndTime"`
	PageSize         requests.Integer              `position:"Query" name:"PageSize"`
	Tag              *[]DescribeDcdnUserDomainsTag `position:"Query" name:"Tag"  type:"Repeated"`
	FuncFilter       string                        `position:"Query" name:"FuncFilter"`
	Coverage         string                        `position:"Query" name:"Coverage"`
	DomainName       string                        `position:"Query" name:"DomainName"`
	OwnerId          requests.Integer              `position:"Query" name:"OwnerId"`
	FuncId           string                        `position:"Query" name:"FuncId"`
	DomainStatus     string                        `position:"Query" name:"DomainStatus"`
	DomainSearchType string                        `position:"Query" name:"DomainSearchType"`
	ChangeStartTime  string                        `position:"Query" name:"ChangeStartTime"`
}

// DescribeDcdnUserDomainsTag is a repeated param struct in DescribeDcdnUserDomainsRequest
type DescribeDcdnUserDomainsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeDcdnUserDomainsResponse is the response struct for api DescribeDcdnUserDomains
type DescribeDcdnUserDomainsResponse struct {
	*responses.BaseResponse
	RequestId  string                           `json:"RequestId" xml:"RequestId"`
	PageNumber int64                            `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64                            `json:"PageSize" xml:"PageSize"`
	TotalCount int64                            `json:"TotalCount" xml:"TotalCount"`
	Domains    DomainsInDescribeDcdnUserDomains `json:"Domains" xml:"Domains"`
}

// CreateDescribeDcdnUserDomainsRequest creates a request to invoke DescribeDcdnUserDomains API
func CreateDescribeDcdnUserDomainsRequest() (request *DescribeDcdnUserDomainsRequest) {
	request = &DescribeDcdnUserDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnUserDomains", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnUserDomainsResponse creates a response to parse from DescribeDcdnUserDomains response
func CreateDescribeDcdnUserDomainsResponse() (response *DescribeDcdnUserDomainsResponse) {
	response = &DescribeDcdnUserDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
