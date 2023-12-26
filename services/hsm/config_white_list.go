package hsm

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

// ConfigWhiteList invokes the hsm.ConfigWhiteList API synchronously
// api document: https://help.aliyun.com/api/hsm/configwhitelist.html
func (client *Client) ConfigWhiteList(request *ConfigWhiteListRequest) (response *ConfigWhiteListResponse, err error) {
	response = CreateConfigWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigWhiteListWithChan invokes the hsm.ConfigWhiteList API asynchronously
// api document: https://help.aliyun.com/api/hsm/configwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigWhiteListWithChan(request *ConfigWhiteListRequest) (<-chan *ConfigWhiteListResponse, <-chan error) {
	responseChan := make(chan *ConfigWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigWhiteList(request)
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

// ConfigWhiteListWithCallback invokes the hsm.ConfigWhiteList API asynchronously
// api document: https://help.aliyun.com/api/hsm/configwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigWhiteListWithCallback(request *ConfigWhiteListRequest, callback func(response *ConfigWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigWhiteListResponse
		var err error
		defer close(result)
		response, err = client.ConfigWhiteList(request)
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

// ConfigWhiteListRequest is the request struct for api ConfigWhiteList
type ConfigWhiteListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	WhiteList       string           `position:"Query" name:"WhiteList"`
}

// ConfigWhiteListResponse is the response struct for api ConfigWhiteList
type ConfigWhiteListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigWhiteListRequest creates a request to invoke ConfigWhiteList API
func CreateConfigWhiteListRequest() (request *ConfigWhiteListRequest) {
	request = &ConfigWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hsm", "2018-01-11", "ConfigWhiteList", "hsm", "openAPI")
	return
}

// CreateConfigWhiteListResponse creates a response to parse from ConfigWhiteList response
func CreateConfigWhiteListResponse() (response *ConfigWhiteListResponse) {
	response = &ConfigWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
