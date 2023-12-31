package cloudapi

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

// ReactivateDomain invokes the cloudapi.ReactivateDomain API synchronously
func (client *Client) ReactivateDomain(request *ReactivateDomainRequest) (response *ReactivateDomainResponse, err error) {
	response = CreateReactivateDomainResponse()
	err = client.DoAction(request, response)
	return
}

// ReactivateDomainWithChan invokes the cloudapi.ReactivateDomain API asynchronously
func (client *Client) ReactivateDomainWithChan(request *ReactivateDomainRequest) (<-chan *ReactivateDomainResponse, <-chan error) {
	responseChan := make(chan *ReactivateDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReactivateDomain(request)
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

// ReactivateDomainWithCallback invokes the cloudapi.ReactivateDomain API asynchronously
func (client *Client) ReactivateDomainWithCallback(request *ReactivateDomainRequest, callback func(response *ReactivateDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReactivateDomainResponse
		var err error
		defer close(result)
		response, err = client.ReactivateDomain(request)
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

// ReactivateDomainRequest is the request struct for api ReactivateDomain
type ReactivateDomainRequest struct {
	*requests.RpcRequest
	GroupId       string `position:"Query" name:"GroupId"`
	DomainName    string `position:"Query" name:"DomainName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// ReactivateDomainResponse is the response struct for api ReactivateDomain
type ReactivateDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReactivateDomainRequest creates a request to invoke ReactivateDomain API
func CreateReactivateDomainRequest() (request *ReactivateDomainRequest) {
	request = &ReactivateDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ReactivateDomain", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReactivateDomainResponse creates a response to parse from ReactivateDomain response
func CreateReactivateDomainResponse() (response *ReactivateDomainResponse) {
	response = &ReactivateDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
