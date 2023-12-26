package live

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

// PublishLiveStagingConfigToProduction invokes the live.PublishLiveStagingConfigToProduction API synchronously
func (client *Client) PublishLiveStagingConfigToProduction(request *PublishLiveStagingConfigToProductionRequest) (response *PublishLiveStagingConfigToProductionResponse, err error) {
	response = CreatePublishLiveStagingConfigToProductionResponse()
	err = client.DoAction(request, response)
	return
}

// PublishLiveStagingConfigToProductionWithChan invokes the live.PublishLiveStagingConfigToProduction API asynchronously
func (client *Client) PublishLiveStagingConfigToProductionWithChan(request *PublishLiveStagingConfigToProductionRequest) (<-chan *PublishLiveStagingConfigToProductionResponse, <-chan error) {
	responseChan := make(chan *PublishLiveStagingConfigToProductionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishLiveStagingConfigToProduction(request)
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

// PublishLiveStagingConfigToProductionWithCallback invokes the live.PublishLiveStagingConfigToProduction API asynchronously
func (client *Client) PublishLiveStagingConfigToProductionWithCallback(request *PublishLiveStagingConfigToProductionRequest, callback func(response *PublishLiveStagingConfigToProductionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishLiveStagingConfigToProductionResponse
		var err error
		defer close(result)
		response, err = client.PublishLiveStagingConfigToProduction(request)
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

// PublishLiveStagingConfigToProductionRequest is the request struct for api PublishLiveStagingConfigToProduction
type PublishLiveStagingConfigToProductionRequest struct {
	*requests.RpcRequest
	FunctionName string           `position:"Query" name:"FunctionName"`
	DomainName   string           `position:"Query" name:"DomainName"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
}

// PublishLiveStagingConfigToProductionResponse is the response struct for api PublishLiveStagingConfigToProduction
type PublishLiveStagingConfigToProductionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePublishLiveStagingConfigToProductionRequest creates a request to invoke PublishLiveStagingConfigToProduction API
func CreatePublishLiveStagingConfigToProductionRequest() (request *PublishLiveStagingConfigToProductionRequest) {
	request = &PublishLiveStagingConfigToProductionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "PublishLiveStagingConfigToProduction", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePublishLiveStagingConfigToProductionResponse creates a response to parse from PublishLiveStagingConfigToProduction response
func CreatePublishLiveStagingConfigToProductionResponse() (response *PublishLiveStagingConfigToProductionResponse) {
	response = &PublishLiveStagingConfigToProductionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
