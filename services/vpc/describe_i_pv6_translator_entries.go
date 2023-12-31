package vpc

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

// DescribeIPv6TranslatorEntries invokes the vpc.DescribeIPv6TranslatorEntries API synchronously
func (client *Client) DescribeIPv6TranslatorEntries(request *DescribeIPv6TranslatorEntriesRequest) (response *DescribeIPv6TranslatorEntriesResponse, err error) {
	response = CreateDescribeIPv6TranslatorEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIPv6TranslatorEntriesWithChan invokes the vpc.DescribeIPv6TranslatorEntries API asynchronously
func (client *Client) DescribeIPv6TranslatorEntriesWithChan(request *DescribeIPv6TranslatorEntriesRequest) (<-chan *DescribeIPv6TranslatorEntriesResponse, <-chan error) {
	responseChan := make(chan *DescribeIPv6TranslatorEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIPv6TranslatorEntries(request)
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

// DescribeIPv6TranslatorEntriesWithCallback invokes the vpc.DescribeIPv6TranslatorEntries API asynchronously
func (client *Client) DescribeIPv6TranslatorEntriesWithCallback(request *DescribeIPv6TranslatorEntriesRequest, callback func(response *DescribeIPv6TranslatorEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIPv6TranslatorEntriesResponse
		var err error
		defer close(result)
		response, err = client.DescribeIPv6TranslatorEntries(request)
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

// DescribeIPv6TranslatorEntriesRequest is the request struct for api DescribeIPv6TranslatorEntries
type DescribeIPv6TranslatorEntriesRequest struct {
	*requests.RpcRequest
	BackendIpv4Port       requests.Integer `position:"Query" name:"BackendIpv4Port"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EntryName             string           `position:"Query" name:"EntryName"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	AclStatus             string           `position:"Query" name:"AclStatus"`
	PageNumber            requests.Integer `position:"Query" name:"PageNumber"`
	AclType               string           `position:"Query" name:"AclType"`
	AllocateIpv6Port      requests.Integer `position:"Query" name:"AllocateIpv6Port"`
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	BackendIpv4Addr       string           `position:"Query" name:"BackendIpv4Addr"`
	AclId                 string           `position:"Query" name:"AclId"`
	Ipv6TranslatorEntryId string           `position:"Query" name:"Ipv6TranslatorEntryId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	AllocateIpv6Addr      string           `position:"Query" name:"AllocateIpv6Addr"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	TransProtocol         string           `position:"Query" name:"TransProtocol"`
	Ipv6TranslatorId      string           `position:"Query" name:"Ipv6TranslatorId"`
}

// DescribeIPv6TranslatorEntriesResponse is the response struct for api DescribeIPv6TranslatorEntries
type DescribeIPv6TranslatorEntriesResponse struct {
	*responses.BaseResponse
	PageSize              int                   `json:"PageSize" xml:"PageSize"`
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	PageNumber            int                   `json:"PageNumber" xml:"PageNumber"`
	TotalCount            int                   `json:"TotalCount" xml:"TotalCount"`
	Ipv6TranslatorEntries Ipv6TranslatorEntries `json:"Ipv6TranslatorEntries" xml:"Ipv6TranslatorEntries"`
}

// CreateDescribeIPv6TranslatorEntriesRequest creates a request to invoke DescribeIPv6TranslatorEntries API
func CreateDescribeIPv6TranslatorEntriesRequest() (request *DescribeIPv6TranslatorEntriesRequest) {
	request = &DescribeIPv6TranslatorEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeIPv6TranslatorEntries", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIPv6TranslatorEntriesResponse creates a response to parse from DescribeIPv6TranslatorEntries response
func CreateDescribeIPv6TranslatorEntriesResponse() (response *DescribeIPv6TranslatorEntriesResponse) {
	response = &DescribeIPv6TranslatorEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
