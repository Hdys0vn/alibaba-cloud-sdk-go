package alidns

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

// DescribeDnsCacheDomains invokes the alidns.DescribeDnsCacheDomains API synchronously
func (client *Client) DescribeDnsCacheDomains(request *DescribeDnsCacheDomainsRequest) (response *DescribeDnsCacheDomainsResponse, err error) {
	response = CreateDescribeDnsCacheDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDnsCacheDomainsWithChan invokes the alidns.DescribeDnsCacheDomains API asynchronously
func (client *Client) DescribeDnsCacheDomainsWithChan(request *DescribeDnsCacheDomainsRequest) (<-chan *DescribeDnsCacheDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeDnsCacheDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDnsCacheDomains(request)
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

// DescribeDnsCacheDomainsWithCallback invokes the alidns.DescribeDnsCacheDomains API asynchronously
func (client *Client) DescribeDnsCacheDomainsWithCallback(request *DescribeDnsCacheDomainsRequest, callback func(response *DescribeDnsCacheDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDnsCacheDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDnsCacheDomains(request)
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

// DescribeDnsCacheDomainsRequest is the request struct for api DescribeDnsCacheDomains
type DescribeDnsCacheDomainsRequest struct {
	*requests.RpcRequest
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	Keyword      string           `position:"Query" name:"Keyword"`
}

// DescribeDnsCacheDomainsResponse is the response struct for api DescribeDnsCacheDomains
type DescribeDnsCacheDomainsResponse struct {
	*responses.BaseResponse
	TotalCount int64    `json:"TotalCount" xml:"TotalCount"`
	PageSize   int64    `json:"PageSize" xml:"PageSize"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	PageNumber int64    `json:"PageNumber" xml:"PageNumber"`
	Domains    []Domain `json:"Domains" xml:"Domains"`
}

// CreateDescribeDnsCacheDomainsRequest creates a request to invoke DescribeDnsCacheDomains API
func CreateDescribeDnsCacheDomainsRequest() (request *DescribeDnsCacheDomainsRequest) {
	request = &DescribeDnsCacheDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsCacheDomains", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDnsCacheDomainsResponse creates a response to parse from DescribeDnsCacheDomains response
func CreateDescribeDnsCacheDomainsResponse() (response *DescribeDnsCacheDomainsResponse) {
	response = &DescribeDnsCacheDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
