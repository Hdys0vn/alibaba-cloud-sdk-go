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

// UpdateIspFlushCacheInstanceConfig invokes the alidns.UpdateIspFlushCacheInstanceConfig API synchronously
func (client *Client) UpdateIspFlushCacheInstanceConfig(request *UpdateIspFlushCacheInstanceConfigRequest) (response *UpdateIspFlushCacheInstanceConfigResponse, err error) {
	response = CreateUpdateIspFlushCacheInstanceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateIspFlushCacheInstanceConfigWithChan invokes the alidns.UpdateIspFlushCacheInstanceConfig API asynchronously
func (client *Client) UpdateIspFlushCacheInstanceConfigWithChan(request *UpdateIspFlushCacheInstanceConfigRequest) (<-chan *UpdateIspFlushCacheInstanceConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateIspFlushCacheInstanceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateIspFlushCacheInstanceConfig(request)
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

// UpdateIspFlushCacheInstanceConfigWithCallback invokes the alidns.UpdateIspFlushCacheInstanceConfig API asynchronously
func (client *Client) UpdateIspFlushCacheInstanceConfigWithCallback(request *UpdateIspFlushCacheInstanceConfigRequest, callback func(response *UpdateIspFlushCacheInstanceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateIspFlushCacheInstanceConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateIspFlushCacheInstanceConfig(request)
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

// UpdateIspFlushCacheInstanceConfigRequest is the request struct for api UpdateIspFlushCacheInstanceConfig
type UpdateIspFlushCacheInstanceConfigRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceName string `position:"Query" name:"InstanceName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// UpdateIspFlushCacheInstanceConfigResponse is the response struct for api UpdateIspFlushCacheInstanceConfig
type UpdateIspFlushCacheInstanceConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateIspFlushCacheInstanceConfigRequest creates a request to invoke UpdateIspFlushCacheInstanceConfig API
func CreateUpdateIspFlushCacheInstanceConfigRequest() (request *UpdateIspFlushCacheInstanceConfigRequest) {
	request = &UpdateIspFlushCacheInstanceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateIspFlushCacheInstanceConfig", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateIspFlushCacheInstanceConfigResponse creates a response to parse from UpdateIspFlushCacheInstanceConfig response
func CreateUpdateIspFlushCacheInstanceConfigResponse() (response *UpdateIspFlushCacheInstanceConfigResponse) {
	response = &UpdateIspFlushCacheInstanceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
