package mse

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

// GetPluginConfig invokes the mse.GetPluginConfig API synchronously
func (client *Client) GetPluginConfig(request *GetPluginConfigRequest) (response *GetPluginConfigResponse, err error) {
	response = CreateGetPluginConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetPluginConfigWithChan invokes the mse.GetPluginConfig API asynchronously
func (client *Client) GetPluginConfigWithChan(request *GetPluginConfigRequest) (<-chan *GetPluginConfigResponse, <-chan error) {
	responseChan := make(chan *GetPluginConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPluginConfig(request)
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

// GetPluginConfigWithCallback invokes the mse.GetPluginConfig API asynchronously
func (client *Client) GetPluginConfigWithCallback(request *GetPluginConfigRequest, callback func(response *GetPluginConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPluginConfigResponse
		var err error
		defer close(result)
		response, err = client.GetPluginConfig(request)
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

// GetPluginConfigRequest is the request struct for api GetPluginConfig
type GetPluginConfigRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	PluginId        requests.Integer `position:"Query" name:"PluginId"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
}

// GetPluginConfigResponse is the response struct for api GetPluginConfig
type GetPluginConfigResponse struct {
	*responses.BaseResponse
	RequestId      string                `json:"RequestId" xml:"RequestId"`
	Success        bool                  `json:"Success" xml:"Success"`
	Code           int                   `json:"Code" xml:"Code"`
	ErrorCode      string                `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode int                   `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                `json:"Message" xml:"Message"`
	DynamicCode    string                `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string                `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           DataInGetPluginConfig `json:"Data" xml:"Data"`
}

// CreateGetPluginConfigRequest creates a request to invoke GetPluginConfig API
func CreateGetPluginConfigRequest() (request *GetPluginConfigRequest) {
	request = &GetPluginConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetPluginConfig", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetPluginConfigResponse creates a response to parse from GetPluginConfig response
func CreateGetPluginConfigResponse() (response *GetPluginConfigResponse) {
	response = &GetPluginConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
