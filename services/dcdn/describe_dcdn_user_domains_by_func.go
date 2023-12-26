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

// DescribeDcdnUserDomainsByFunc invokes the dcdn.DescribeDcdnUserDomainsByFunc API synchronously
func (client *Client) DescribeDcdnUserDomainsByFunc(request *DescribeDcdnUserDomainsByFuncRequest) (response *DescribeDcdnUserDomainsByFuncResponse, err error) {
	response = CreateDescribeDcdnUserDomainsByFuncResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnUserDomainsByFuncWithChan invokes the dcdn.DescribeDcdnUserDomainsByFunc API asynchronously
func (client *Client) DescribeDcdnUserDomainsByFuncWithChan(request *DescribeDcdnUserDomainsByFuncRequest) (<-chan *DescribeDcdnUserDomainsByFuncResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnUserDomainsByFuncResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnUserDomainsByFunc(request)
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

// DescribeDcdnUserDomainsByFuncWithCallback invokes the dcdn.DescribeDcdnUserDomainsByFunc API asynchronously
func (client *Client) DescribeDcdnUserDomainsByFuncWithCallback(request *DescribeDcdnUserDomainsByFuncRequest, callback func(response *DescribeDcdnUserDomainsByFuncResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnUserDomainsByFuncResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnUserDomainsByFunc(request)
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

// DescribeDcdnUserDomainsByFuncRequest is the request struct for api DescribeDcdnUserDomainsByFunc
type DescribeDcdnUserDomainsByFuncRequest struct {
	*requests.RpcRequest
	FuncFilter      string           `position:"Query" name:"FuncFilter"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	DomainName      string           `position:"Query" name:"DomainName"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	FuncId          requests.Integer `position:"Query" name:"FuncId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeDcdnUserDomainsByFuncResponse is the response struct for api DescribeDcdnUserDomainsByFunc
type DescribeDcdnUserDomainsByFuncResponse struct {
	*responses.BaseResponse
	RequestId  string                                 `json:"RequestId" xml:"RequestId"`
	PageNumber int64                                  `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64                                  `json:"PageSize" xml:"PageSize"`
	TotalCount int64                                  `json:"TotalCount" xml:"TotalCount"`
	Domains    DomainsInDescribeDcdnUserDomainsByFunc `json:"Domains" xml:"Domains"`
}

// CreateDescribeDcdnUserDomainsByFuncRequest creates a request to invoke DescribeDcdnUserDomainsByFunc API
func CreateDescribeDcdnUserDomainsByFuncRequest() (request *DescribeDcdnUserDomainsByFuncRequest) {
	request = &DescribeDcdnUserDomainsByFuncRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnUserDomainsByFunc", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnUserDomainsByFuncResponse creates a response to parse from DescribeDcdnUserDomainsByFunc response
func CreateDescribeDcdnUserDomainsByFuncResponse() (response *DescribeDcdnUserDomainsByFuncResponse) {
	response = &DescribeDcdnUserDomainsByFuncResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
