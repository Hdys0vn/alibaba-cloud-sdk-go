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

// OperateBatchDomain invokes the alidns.OperateBatchDomain API synchronously
func (client *Client) OperateBatchDomain(request *OperateBatchDomainRequest) (response *OperateBatchDomainResponse, err error) {
	response = CreateOperateBatchDomainResponse()
	err = client.DoAction(request, response)
	return
}

// OperateBatchDomainWithChan invokes the alidns.OperateBatchDomain API asynchronously
func (client *Client) OperateBatchDomainWithChan(request *OperateBatchDomainRequest) (<-chan *OperateBatchDomainResponse, <-chan error) {
	responseChan := make(chan *OperateBatchDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OperateBatchDomain(request)
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

// OperateBatchDomainWithCallback invokes the alidns.OperateBatchDomain API asynchronously
func (client *Client) OperateBatchDomainWithCallback(request *OperateBatchDomainRequest, callback func(response *OperateBatchDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OperateBatchDomainResponse
		var err error
		defer close(result)
		response, err = client.OperateBatchDomain(request)
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

// OperateBatchDomainRequest is the request struct for api OperateBatchDomain
type OperateBatchDomainRequest struct {
	*requests.RpcRequest
	DomainRecordInfo *[]OperateBatchDomainDomainRecordInfo `position:"Query" name:"DomainRecordInfo"  type:"Repeated"`
	Type             string                                `position:"Query" name:"Type"`
	UserClientIp     string                                `position:"Query" name:"UserClientIp"`
	Lang             string                                `position:"Query" name:"Lang"`
}

// OperateBatchDomainDomainRecordInfo is a repeated param struct in OperateBatchDomainRequest
type OperateBatchDomainDomainRecordInfo struct {
	Rr       string `name:"Rr"`
	NewType  string `name:"NewType"`
	NewValue string `name:"NewValue"`
	Line     string `name:"Line"`
	Domain   string `name:"Domain"`
	Type     string `name:"Type"`
	Priority string `name:"Priority"`
	Value    string `name:"Value"`
	Ttl      string `name:"Ttl"`
	NewRr    string `name:"NewRr"`
}

// OperateBatchDomainResponse is the response struct for api OperateBatchDomain
type OperateBatchDomainResponse struct {
	*responses.BaseResponse
	TaskId    int64  `json:"TaskId" xml:"TaskId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOperateBatchDomainRequest creates a request to invoke OperateBatchDomain API
func CreateOperateBatchDomainRequest() (request *OperateBatchDomainRequest) {
	request = &OperateBatchDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "OperateBatchDomain", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOperateBatchDomainResponse creates a response to parse from OperateBatchDomain response
func CreateOperateBatchDomainResponse() (response *OperateBatchDomainResponse) {
	response = &OperateBatchDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
