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

// QueryEnsAssociation invokes the domain_intl.QueryEnsAssociation API synchronously
// api document: https://help.aliyun.com/api/domain-intl/queryensassociation.html
func (client *Client) QueryEnsAssociation(request *QueryEnsAssociationRequest) (response *QueryEnsAssociationResponse, err error) {
	response = CreateQueryEnsAssociationResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEnsAssociationWithChan invokes the domain_intl.QueryEnsAssociation API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/queryensassociation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEnsAssociationWithChan(request *QueryEnsAssociationRequest) (<-chan *QueryEnsAssociationResponse, <-chan error) {
	responseChan := make(chan *QueryEnsAssociationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEnsAssociation(request)
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

// QueryEnsAssociationWithCallback invokes the domain_intl.QueryEnsAssociation API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/queryensassociation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEnsAssociationWithCallback(request *QueryEnsAssociationRequest, callback func(response *QueryEnsAssociationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEnsAssociationResponse
		var err error
		defer close(result)
		response, err = client.QueryEnsAssociation(request)
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

// QueryEnsAssociationRequest is the request struct for api QueryEnsAssociation
type QueryEnsAssociationRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryEnsAssociationResponse is the response struct for api QueryEnsAssociation
type QueryEnsAssociationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Address   string `json:"Address" xml:"Address"`
}

// CreateQueryEnsAssociationRequest creates a request to invoke QueryEnsAssociation API
func CreateQueryEnsAssociationRequest() (request *QueryEnsAssociationRequest) {
	request = &QueryEnsAssociationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryEnsAssociation", "domain", "openAPI")
	return
}

// CreateQueryEnsAssociationResponse creates a response to parse from QueryEnsAssociation response
func CreateQueryEnsAssociationResponse() (response *QueryEnsAssociationResponse) {
	response = &QueryEnsAssociationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
