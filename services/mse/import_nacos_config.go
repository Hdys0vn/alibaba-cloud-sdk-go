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

// ImportNacosConfig invokes the mse.ImportNacosConfig API synchronously
func (client *Client) ImportNacosConfig(request *ImportNacosConfigRequest) (response *ImportNacosConfigResponse, err error) {
	response = CreateImportNacosConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ImportNacosConfigWithChan invokes the mse.ImportNacosConfig API asynchronously
func (client *Client) ImportNacosConfigWithChan(request *ImportNacosConfigRequest) (<-chan *ImportNacosConfigResponse, <-chan error) {
	responseChan := make(chan *ImportNacosConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportNacosConfig(request)
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

// ImportNacosConfigWithCallback invokes the mse.ImportNacosConfig API asynchronously
func (client *Client) ImportNacosConfigWithCallback(request *ImportNacosConfigRequest, callback func(response *ImportNacosConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportNacosConfigResponse
		var err error
		defer close(result)
		response, err = client.ImportNacosConfig(request)
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

// ImportNacosConfigRequest is the request struct for api ImportNacosConfig
type ImportNacosConfigRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	NamespaceId    string `position:"Query" name:"NamespaceId"`
	Policy         string `position:"Query" name:"Policy"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
	FileUrl        string `position:"Query" name:"FileUrl"`
}

// ImportNacosConfigResponse is the response struct for api ImportNacosConfig
type ImportNacosConfigResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Code           int    `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateImportNacosConfigRequest creates a request to invoke ImportNacosConfig API
func CreateImportNacosConfigRequest() (request *ImportNacosConfigRequest) {
	request = &ImportNacosConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ImportNacosConfig", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateImportNacosConfigResponse creates a response to parse from ImportNacosConfig response
func CreateImportNacosConfigResponse() (response *ImportNacosConfigResponse) {
	response = &ImportNacosConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
