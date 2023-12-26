package hbase

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

// DescribeSubDomain invokes the hbase.DescribeSubDomain API synchronously
func (client *Client) DescribeSubDomain(request *DescribeSubDomainRequest) (response *DescribeSubDomainResponse, err error) {
	response = CreateDescribeSubDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSubDomainWithChan invokes the hbase.DescribeSubDomain API asynchronously
func (client *Client) DescribeSubDomainWithChan(request *DescribeSubDomainRequest) (<-chan *DescribeSubDomainResponse, <-chan error) {
	responseChan := make(chan *DescribeSubDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSubDomain(request)
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

// DescribeSubDomainWithCallback invokes the hbase.DescribeSubDomain API asynchronously
func (client *Client) DescribeSubDomainWithCallback(request *DescribeSubDomainRequest, callback func(response *DescribeSubDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSubDomainResponse
		var err error
		defer close(result)
		response, err = client.DescribeSubDomain(request)
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

// DescribeSubDomainRequest is the request struct for api DescribeSubDomain
type DescribeSubDomainRequest struct {
	*requests.RpcRequest
	ZoneId string `position:"Query" name:"ZoneId"`
}

// DescribeSubDomainResponse is the response struct for api DescribeSubDomain
type DescribeSubDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SubDomain string `json:"SubDomain" xml:"SubDomain"`
}

// CreateDescribeSubDomainRequest creates a request to invoke DescribeSubDomain API
func CreateDescribeSubDomainRequest() (request *DescribeSubDomainRequest) {
	request = &DescribeSubDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DescribeSubDomain", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSubDomainResponse creates a response to parse from DescribeSubDomain response
func CreateDescribeSubDomainResponse() (response *DescribeSubDomainResponse) {
	response = &DescribeSubDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}