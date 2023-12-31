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

// UpdateDnsCacheDomain invokes the alidns.UpdateDnsCacheDomain API synchronously
func (client *Client) UpdateDnsCacheDomain(request *UpdateDnsCacheDomainRequest) (response *UpdateDnsCacheDomainResponse, err error) {
	response = CreateUpdateDnsCacheDomainResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDnsCacheDomainWithChan invokes the alidns.UpdateDnsCacheDomain API asynchronously
func (client *Client) UpdateDnsCacheDomainWithChan(request *UpdateDnsCacheDomainRequest) (<-chan *UpdateDnsCacheDomainResponse, <-chan error) {
	responseChan := make(chan *UpdateDnsCacheDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDnsCacheDomain(request)
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

// UpdateDnsCacheDomainWithCallback invokes the alidns.UpdateDnsCacheDomain API asynchronously
func (client *Client) UpdateDnsCacheDomainWithCallback(request *UpdateDnsCacheDomainRequest, callback func(response *UpdateDnsCacheDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDnsCacheDomainResponse
		var err error
		defer close(result)
		response, err = client.UpdateDnsCacheDomain(request)
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

// UpdateDnsCacheDomainRequest is the request struct for api UpdateDnsCacheDomain
type UpdateDnsCacheDomainRequest struct {
	*requests.RpcRequest
	SourceProtocol  string                                 `position:"Query" name:"SourceProtocol"`
	Lang            string                                 `position:"Query" name:"Lang"`
	DomainName      string                                 `position:"Query" name:"DomainName"`
	CacheTtlMax     requests.Integer                       `position:"Query" name:"CacheTtlMax"`
	InstanceId      string                                 `position:"Query" name:"InstanceId"`
	SourceEdns      string                                 `position:"Query" name:"SourceEdns"`
	UserClientIp    string                                 `position:"Query" name:"UserClientIp"`
	CacheTtlMin     requests.Integer                       `position:"Query" name:"CacheTtlMin"`
	SourceDnsServer *[]UpdateDnsCacheDomainSourceDnsServer `position:"Query" name:"SourceDnsServer"  type:"Repeated"`
}

// UpdateDnsCacheDomainSourceDnsServer is a repeated param struct in UpdateDnsCacheDomainRequest
type UpdateDnsCacheDomainSourceDnsServer struct {
	Port string `name:"Port"`
	Host string `name:"Host"`
}

// UpdateDnsCacheDomainResponse is the response struct for api UpdateDnsCacheDomain
type UpdateDnsCacheDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDnsCacheDomainRequest creates a request to invoke UpdateDnsCacheDomain API
func CreateUpdateDnsCacheDomainRequest() (request *UpdateDnsCacheDomainRequest) {
	request = &UpdateDnsCacheDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDnsCacheDomain", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDnsCacheDomainResponse creates a response to parse from UpdateDnsCacheDomain response
func CreateUpdateDnsCacheDomainResponse() (response *UpdateDnsCacheDomainResponse) {
	response = &UpdateDnsCacheDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
