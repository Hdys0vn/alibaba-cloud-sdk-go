package oms

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

// GetReadyFlagAlertConfig invokes the oms.GetReadyFlagAlertConfig API synchronously
func (client *Client) GetReadyFlagAlertConfig(request *GetReadyFlagAlertConfigRequest) (response *GetReadyFlagAlertConfigResponse, err error) {
	response = CreateGetReadyFlagAlertConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetReadyFlagAlertConfigWithChan invokes the oms.GetReadyFlagAlertConfig API asynchronously
func (client *Client) GetReadyFlagAlertConfigWithChan(request *GetReadyFlagAlertConfigRequest) (<-chan *GetReadyFlagAlertConfigResponse, <-chan error) {
	responseChan := make(chan *GetReadyFlagAlertConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetReadyFlagAlertConfig(request)
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

// GetReadyFlagAlertConfigWithCallback invokes the oms.GetReadyFlagAlertConfig API asynchronously
func (client *Client) GetReadyFlagAlertConfigWithCallback(request *GetReadyFlagAlertConfigRequest, callback func(response *GetReadyFlagAlertConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetReadyFlagAlertConfigResponse
		var err error
		defer close(result)
		response, err = client.GetReadyFlagAlertConfig(request)
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

// GetReadyFlagAlertConfigRequest is the request struct for api GetReadyFlagAlertConfig
type GetReadyFlagAlertConfigRequest struct {
	*requests.RpcRequest
	DomainCode     string           `position:"Query" name:"DomainCode"`
	DataType       string           `position:"Query" name:"DataType"`
	CompressEnable requests.Boolean `position:"Query" name:"CompressEnable"`
}

// GetReadyFlagAlertConfigResponse is the response struct for api GetReadyFlagAlertConfig
type GetReadyFlagAlertConfigResponse struct {
	*responses.BaseResponse
	Compressed bool   `json:"Compressed" xml:"Compressed"`
	Data       string `json:"Data" xml:"Data"`
	DataType   string `json:"DataType" xml:"DataType"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainCode string `json:"DomainCode" xml:"DomainCode"`
}

// CreateGetReadyFlagAlertConfigRequest creates a request to invoke GetReadyFlagAlertConfig API
func CreateGetReadyFlagAlertConfigRequest() (request *GetReadyFlagAlertConfigRequest) {
	request = &GetReadyFlagAlertConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oms", "2016-06-15", "GetReadyFlagAlertConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateGetReadyFlagAlertConfigResponse creates a response to parse from GetReadyFlagAlertConfig response
func CreateGetReadyFlagAlertConfigResponse() (response *GetReadyFlagAlertConfigResponse) {
	response = &GetReadyFlagAlertConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}