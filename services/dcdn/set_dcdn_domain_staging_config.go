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

// SetDcdnDomainStagingConfig invokes the dcdn.SetDcdnDomainStagingConfig API synchronously
func (client *Client) SetDcdnDomainStagingConfig(request *SetDcdnDomainStagingConfigRequest) (response *SetDcdnDomainStagingConfigResponse, err error) {
	response = CreateSetDcdnDomainStagingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetDcdnDomainStagingConfigWithChan invokes the dcdn.SetDcdnDomainStagingConfig API asynchronously
func (client *Client) SetDcdnDomainStagingConfigWithChan(request *SetDcdnDomainStagingConfigRequest) (<-chan *SetDcdnDomainStagingConfigResponse, <-chan error) {
	responseChan := make(chan *SetDcdnDomainStagingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDcdnDomainStagingConfig(request)
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

// SetDcdnDomainStagingConfigWithCallback invokes the dcdn.SetDcdnDomainStagingConfig API asynchronously
func (client *Client) SetDcdnDomainStagingConfigWithCallback(request *SetDcdnDomainStagingConfigRequest, callback func(response *SetDcdnDomainStagingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDcdnDomainStagingConfigResponse
		var err error
		defer close(result)
		response, err = client.SetDcdnDomainStagingConfig(request)
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

// SetDcdnDomainStagingConfigRequest is the request struct for api SetDcdnDomainStagingConfig
type SetDcdnDomainStagingConfigRequest struct {
	*requests.RpcRequest
	Functions  string `position:"Query" name:"Functions"`
	DomainName string `position:"Query" name:"DomainName"`
}

// SetDcdnDomainStagingConfigResponse is the response struct for api SetDcdnDomainStagingConfig
type SetDcdnDomainStagingConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDcdnDomainStagingConfigRequest creates a request to invoke SetDcdnDomainStagingConfig API
func CreateSetDcdnDomainStagingConfigRequest() (request *SetDcdnDomainStagingConfigRequest) {
	request = &SetDcdnDomainStagingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "SetDcdnDomainStagingConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateSetDcdnDomainStagingConfigResponse creates a response to parse from SetDcdnDomainStagingConfig response
func CreateSetDcdnDomainStagingConfigResponse() (response *SetDcdnDomainStagingConfigResponse) {
	response = &SetDcdnDomainStagingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
