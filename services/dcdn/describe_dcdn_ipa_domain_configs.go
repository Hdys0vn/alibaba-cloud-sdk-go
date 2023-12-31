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

// DescribeDcdnIpaDomainConfigs invokes the dcdn.DescribeDcdnIpaDomainConfigs API synchronously
func (client *Client) DescribeDcdnIpaDomainConfigs(request *DescribeDcdnIpaDomainConfigsRequest) (response *DescribeDcdnIpaDomainConfigsResponse, err error) {
	response = CreateDescribeDcdnIpaDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnIpaDomainConfigsWithChan invokes the dcdn.DescribeDcdnIpaDomainConfigs API asynchronously
func (client *Client) DescribeDcdnIpaDomainConfigsWithChan(request *DescribeDcdnIpaDomainConfigsRequest) (<-chan *DescribeDcdnIpaDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnIpaDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnIpaDomainConfigs(request)
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

// DescribeDcdnIpaDomainConfigsWithCallback invokes the dcdn.DescribeDcdnIpaDomainConfigs API asynchronously
func (client *Client) DescribeDcdnIpaDomainConfigsWithCallback(request *DescribeDcdnIpaDomainConfigsRequest, callback func(response *DescribeDcdnIpaDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnIpaDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnIpaDomainConfigs(request)
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

// DescribeDcdnIpaDomainConfigsRequest is the request struct for api DescribeDcdnIpaDomainConfigs
type DescribeDcdnIpaDomainConfigsRequest struct {
	*requests.RpcRequest
	FunctionNames string           `position:"Query" name:"FunctionNames"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnIpaDomainConfigsResponse is the response struct for api DescribeDcdnIpaDomainConfigs
type DescribeDcdnIpaDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId     string                                      `json:"RequestId" xml:"RequestId"`
	DomainConfigs DomainConfigsInDescribeDcdnIpaDomainConfigs `json:"DomainConfigs" xml:"DomainConfigs"`
}

// CreateDescribeDcdnIpaDomainConfigsRequest creates a request to invoke DescribeDcdnIpaDomainConfigs API
func CreateDescribeDcdnIpaDomainConfigsRequest() (request *DescribeDcdnIpaDomainConfigsRequest) {
	request = &DescribeDcdnIpaDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnIpaDomainConfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnIpaDomainConfigsResponse creates a response to parse from DescribeDcdnIpaDomainConfigs response
func CreateDescribeDcdnIpaDomainConfigsResponse() (response *DescribeDcdnIpaDomainConfigsResponse) {
	response = &DescribeDcdnIpaDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
