package scdn

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

// DescribeScdnDomainConfigs invokes the scdn.DescribeScdnDomainConfigs API synchronously
func (client *Client) DescribeScdnDomainConfigs(request *DescribeScdnDomainConfigsRequest) (response *DescribeScdnDomainConfigsResponse, err error) {
	response = CreateDescribeScdnDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainConfigsWithChan invokes the scdn.DescribeScdnDomainConfigs API asynchronously
func (client *Client) DescribeScdnDomainConfigsWithChan(request *DescribeScdnDomainConfigsRequest) (<-chan *DescribeScdnDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainConfigs(request)
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

// DescribeScdnDomainConfigsWithCallback invokes the scdn.DescribeScdnDomainConfigs API asynchronously
func (client *Client) DescribeScdnDomainConfigsWithCallback(request *DescribeScdnDomainConfigsRequest, callback func(response *DescribeScdnDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainConfigs(request)
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

// DescribeScdnDomainConfigsRequest is the request struct for api DescribeScdnDomainConfigs
type DescribeScdnDomainConfigsRequest struct {
	*requests.RpcRequest
	FunctionNames string           `position:"Query" name:"FunctionNames"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	ConfigId      string           `position:"Query" name:"ConfigId"`
}

// DescribeScdnDomainConfigsResponse is the response struct for api DescribeScdnDomainConfigs
type DescribeScdnDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	DomainConfigs DomainConfigs `json:"DomainConfigs" xml:"DomainConfigs"`
}

// CreateDescribeScdnDomainConfigsRequest creates a request to invoke DescribeScdnDomainConfigs API
func CreateDescribeScdnDomainConfigsRequest() (request *DescribeScdnDomainConfigsRequest) {
	request = &DescribeScdnDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainConfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnDomainConfigsResponse creates a response to parse from DescribeScdnDomainConfigs response
func CreateDescribeScdnDomainConfigsResponse() (response *DescribeScdnDomainConfigsResponse) {
	response = &DescribeScdnDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
