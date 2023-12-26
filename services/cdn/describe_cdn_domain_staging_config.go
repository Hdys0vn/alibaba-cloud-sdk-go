package cdn

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

// DescribeCdnDomainStagingConfig invokes the cdn.DescribeCdnDomainStagingConfig API synchronously
func (client *Client) DescribeCdnDomainStagingConfig(request *DescribeCdnDomainStagingConfigRequest) (response *DescribeCdnDomainStagingConfigResponse, err error) {
	response = CreateDescribeCdnDomainStagingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnDomainStagingConfigWithChan invokes the cdn.DescribeCdnDomainStagingConfig API asynchronously
func (client *Client) DescribeCdnDomainStagingConfigWithChan(request *DescribeCdnDomainStagingConfigRequest) (<-chan *DescribeCdnDomainStagingConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnDomainStagingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnDomainStagingConfig(request)
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

// DescribeCdnDomainStagingConfigWithCallback invokes the cdn.DescribeCdnDomainStagingConfig API asynchronously
func (client *Client) DescribeCdnDomainStagingConfigWithCallback(request *DescribeCdnDomainStagingConfigRequest, callback func(response *DescribeCdnDomainStagingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnDomainStagingConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnDomainStagingConfig(request)
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

// DescribeCdnDomainStagingConfigRequest is the request struct for api DescribeCdnDomainStagingConfig
type DescribeCdnDomainStagingConfigRequest struct {
	*requests.RpcRequest
	FunctionNames string `position:"Query" name:"FunctionNames"`
	DomainName    string `position:"Query" name:"DomainName"`
}

// DescribeCdnDomainStagingConfigResponse is the response struct for api DescribeCdnDomainStagingConfig
type DescribeCdnDomainStagingConfigResponse struct {
	*responses.BaseResponse
	DomainName    string         `json:"DomainName" xml:"DomainName"`
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	DomainConfigs []DomainConfig `json:"DomainConfigs" xml:"DomainConfigs"`
}

// CreateDescribeCdnDomainStagingConfigRequest creates a request to invoke DescribeCdnDomainStagingConfig API
func CreateDescribeCdnDomainStagingConfigRequest() (request *DescribeCdnDomainStagingConfigRequest) {
	request = &DescribeCdnDomainStagingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCdnDomainStagingConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCdnDomainStagingConfigResponse creates a response to parse from DescribeCdnDomainStagingConfig response
func CreateDescribeCdnDomainStagingConfigResponse() (response *DescribeCdnDomainStagingConfigResponse) {
	response = &DescribeCdnDomainStagingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
