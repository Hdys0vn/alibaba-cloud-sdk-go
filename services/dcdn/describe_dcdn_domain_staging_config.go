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

// DescribeDcdnDomainStagingConfig invokes the dcdn.DescribeDcdnDomainStagingConfig API synchronously
func (client *Client) DescribeDcdnDomainStagingConfig(request *DescribeDcdnDomainStagingConfigRequest) (response *DescribeDcdnDomainStagingConfigResponse, err error) {
	response = CreateDescribeDcdnDomainStagingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainStagingConfigWithChan invokes the dcdn.DescribeDcdnDomainStagingConfig API asynchronously
func (client *Client) DescribeDcdnDomainStagingConfigWithChan(request *DescribeDcdnDomainStagingConfigRequest) (<-chan *DescribeDcdnDomainStagingConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainStagingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainStagingConfig(request)
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

// DescribeDcdnDomainStagingConfigWithCallback invokes the dcdn.DescribeDcdnDomainStagingConfig API asynchronously
func (client *Client) DescribeDcdnDomainStagingConfigWithCallback(request *DescribeDcdnDomainStagingConfigRequest, callback func(response *DescribeDcdnDomainStagingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainStagingConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainStagingConfig(request)
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

// DescribeDcdnDomainStagingConfigRequest is the request struct for api DescribeDcdnDomainStagingConfig
type DescribeDcdnDomainStagingConfigRequest struct {
	*requests.RpcRequest
	FunctionNames string `position:"Query" name:"FunctionNames"`
	DomainName    string `position:"Query" name:"DomainName"`
}

// DescribeDcdnDomainStagingConfigResponse is the response struct for api DescribeDcdnDomainStagingConfig
type DescribeDcdnDomainStagingConfigResponse struct {
	*responses.BaseResponse
	RequestId     string                                          `json:"RequestId" xml:"RequestId"`
	DomainConfigs []DomainConfigInDescribeDcdnDomainStagingConfig `json:"DomainConfigs" xml:"DomainConfigs"`
}

// CreateDescribeDcdnDomainStagingConfigRequest creates a request to invoke DescribeDcdnDomainStagingConfig API
func CreateDescribeDcdnDomainStagingConfigRequest() (request *DescribeDcdnDomainStagingConfigRequest) {
	request = &DescribeDcdnDomainStagingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainStagingConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainStagingConfigResponse creates a response to parse from DescribeDcdnDomainStagingConfig response
func CreateDescribeDcdnDomainStagingConfigResponse() (response *DescribeDcdnDomainStagingConfigResponse) {
	response = &DescribeDcdnDomainStagingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
