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

// DescribeUserDomains invokes the cdn.DescribeUserDomains API synchronously
func (client *Client) DescribeUserDomains(request *DescribeUserDomainsRequest) (response *DescribeUserDomainsResponse, err error) {
	response = CreateDescribeUserDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserDomainsWithChan invokes the cdn.DescribeUserDomains API asynchronously
func (client *Client) DescribeUserDomainsWithChan(request *DescribeUserDomainsRequest) (<-chan *DescribeUserDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeUserDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserDomains(request)
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

// DescribeUserDomainsWithCallback invokes the cdn.DescribeUserDomains API asynchronously
func (client *Client) DescribeUserDomainsWithCallback(request *DescribeUserDomainsRequest, callback func(response *DescribeUserDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserDomains(request)
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

// DescribeUserDomainsRequest is the request struct for api DescribeUserDomains
type DescribeUserDomainsRequest struct {
	*requests.RpcRequest
	Source           string                    `position:"Query" name:"Source"`
	PageNumber       requests.Integer          `position:"Query" name:"PageNumber"`
	CheckDomainShow  requests.Boolean          `position:"Query" name:"CheckDomainShow"`
	ResourceGroupId  string                    `position:"Query" name:"ResourceGroupId"`
	SecurityToken    string                    `position:"Query" name:"SecurityToken"`
	CdnType          string                    `position:"Query" name:"CdnType"`
	ChangeEndTime    string                    `position:"Query" name:"ChangeEndTime"`
	PageSize         requests.Integer          `position:"Query" name:"PageSize"`
	Tag              *[]DescribeUserDomainsTag `position:"Query" name:"Tag"  type:"Repeated"`
	FuncFilter       string                    `position:"Query" name:"FuncFilter"`
	Coverage         string                    `position:"Query" name:"Coverage"`
	DomainName       string                    `position:"Query" name:"DomainName"`
	OwnerId          requests.Integer          `position:"Query" name:"OwnerId"`
	FuncId           string                    `position:"Query" name:"FuncId"`
	DomainStatus     string                    `position:"Query" name:"DomainStatus"`
	DomainSearchType string                    `position:"Query" name:"DomainSearchType"`
	ChangeStartTime  string                    `position:"Query" name:"ChangeStartTime"`
}

// DescribeUserDomainsTag is a repeated param struct in DescribeUserDomainsRequest
type DescribeUserDomainsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeUserDomainsResponse is the response struct for api DescribeUserDomains
type DescribeUserDomainsResponse struct {
	*responses.BaseResponse
	RequestId  string                       `json:"RequestId" xml:"RequestId"`
	PageNumber int64                        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64                        `json:"PageSize" xml:"PageSize"`
	TotalCount int64                        `json:"TotalCount" xml:"TotalCount"`
	Domains    DomainsInDescribeUserDomains `json:"Domains" xml:"Domains"`
}

// CreateDescribeUserDomainsRequest creates a request to invoke DescribeUserDomains API
func CreateDescribeUserDomainsRequest() (request *DescribeUserDomainsRequest) {
	request = &DescribeUserDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeUserDomains", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeUserDomainsResponse creates a response to parse from DescribeUserDomains response
func CreateDescribeUserDomainsResponse() (response *DescribeUserDomainsResponse) {
	response = &DescribeUserDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
