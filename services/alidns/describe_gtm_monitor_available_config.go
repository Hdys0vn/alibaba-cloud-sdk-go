package alidns

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

// DescribeGtmMonitorAvailableConfig invokes the alidns.DescribeGtmMonitorAvailableConfig API synchronously
func (client *Client) DescribeGtmMonitorAvailableConfig(request *DescribeGtmMonitorAvailableConfigRequest) (response *DescribeGtmMonitorAvailableConfigResponse, err error) {
	response = CreateDescribeGtmMonitorAvailableConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGtmMonitorAvailableConfigWithChan invokes the alidns.DescribeGtmMonitorAvailableConfig API asynchronously
func (client *Client) DescribeGtmMonitorAvailableConfigWithChan(request *DescribeGtmMonitorAvailableConfigRequest) (<-chan *DescribeGtmMonitorAvailableConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeGtmMonitorAvailableConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGtmMonitorAvailableConfig(request)
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

// DescribeGtmMonitorAvailableConfigWithCallback invokes the alidns.DescribeGtmMonitorAvailableConfig API asynchronously
func (client *Client) DescribeGtmMonitorAvailableConfigWithCallback(request *DescribeGtmMonitorAvailableConfigRequest, callback func(response *DescribeGtmMonitorAvailableConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGtmMonitorAvailableConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeGtmMonitorAvailableConfig(request)
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

// DescribeGtmMonitorAvailableConfigRequest is the request struct for api DescribeGtmMonitorAvailableConfig
type DescribeGtmMonitorAvailableConfigRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeGtmMonitorAvailableConfigResponse is the response struct for api DescribeGtmMonitorAvailableConfig
type DescribeGtmMonitorAvailableConfigResponse struct {
	*responses.BaseResponse
	RequestId    string                                          `json:"RequestId" xml:"RequestId"`
	IspCityNodes IspCityNodesInDescribeGtmMonitorAvailableConfig `json:"IspCityNodes" xml:"IspCityNodes"`
}

// CreateDescribeGtmMonitorAvailableConfigRequest creates a request to invoke DescribeGtmMonitorAvailableConfig API
func CreateDescribeGtmMonitorAvailableConfigRequest() (request *DescribeGtmMonitorAvailableConfigRequest) {
	request = &DescribeGtmMonitorAvailableConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeGtmMonitorAvailableConfig", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGtmMonitorAvailableConfigResponse creates a response to parse from DescribeGtmMonitorAvailableConfig response
func CreateDescribeGtmMonitorAvailableConfigResponse() (response *DescribeGtmMonitorAvailableConfigResponse) {
	response = &DescribeGtmMonitorAvailableConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
