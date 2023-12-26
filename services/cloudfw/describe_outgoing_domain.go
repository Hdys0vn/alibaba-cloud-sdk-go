package cloudfw

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

// DescribeOutgoingDomain invokes the cloudfw.DescribeOutgoingDomain API synchronously
func (client *Client) DescribeOutgoingDomain(request *DescribeOutgoingDomainRequest) (response *DescribeOutgoingDomainResponse, err error) {
	response = CreateDescribeOutgoingDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOutgoingDomainWithChan invokes the cloudfw.DescribeOutgoingDomain API asynchronously
func (client *Client) DescribeOutgoingDomainWithChan(request *DescribeOutgoingDomainRequest) (<-chan *DescribeOutgoingDomainResponse, <-chan error) {
	responseChan := make(chan *DescribeOutgoingDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOutgoingDomain(request)
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

// DescribeOutgoingDomainWithCallback invokes the cloudfw.DescribeOutgoingDomain API asynchronously
func (client *Client) DescribeOutgoingDomainWithCallback(request *DescribeOutgoingDomainRequest, callback func(response *DescribeOutgoingDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOutgoingDomainResponse
		var err error
		defer close(result)
		response, err = client.DescribeOutgoingDomain(request)
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

// DescribeOutgoingDomainRequest is the request struct for api DescribeOutgoingDomain
type DescribeOutgoingDomainRequest struct {
	*requests.RpcRequest
	SecuritySuggest string `position:"Query" name:"SecuritySuggest"`
	AclCoverage     string `position:"Query" name:"AclCoverage"`
	PublicIP        string `position:"Query" name:"PublicIP"`
	StartTime       string `position:"Query" name:"StartTime"`
	AclStatus       string `position:"Query" name:"AclStatus"`
	PageSize        string `position:"Query" name:"PageSize"`
	Lang            string `position:"Query" name:"Lang"`
	CategoryId      string `position:"Query" name:"CategoryId"`
	Order           string `position:"Query" name:"Order"`
	TagId           string `position:"Query" name:"TagId"`
	PrivateIP       string `position:"Query" name:"PrivateIP"`
	EndTime         string `position:"Query" name:"EndTime"`
	CurrentPage     string `position:"Query" name:"CurrentPage"`
	Sort            string `position:"Query" name:"Sort"`
	TagIdNew        string `position:"Query" name:"TagIdNew"`
	Domain          string `position:"Query" name:"Domain"`
}

// DescribeOutgoingDomainResponse is the response struct for api DescribeOutgoingDomain
type DescribeOutgoingDomainResponse struct {
	*responses.BaseResponse
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainList []Data `json:"DomainList" xml:"DomainList"`
}

// CreateDescribeOutgoingDomainRequest creates a request to invoke DescribeOutgoingDomain API
func CreateDescribeOutgoingDomainRequest() (request *DescribeOutgoingDomainRequest) {
	request = &DescribeOutgoingDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeOutgoingDomain", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeOutgoingDomainResponse creates a response to parse from DescribeOutgoingDomain response
func CreateDescribeOutgoingDomainResponse() (response *DescribeOutgoingDomainResponse) {
	response = &DescribeOutgoingDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}