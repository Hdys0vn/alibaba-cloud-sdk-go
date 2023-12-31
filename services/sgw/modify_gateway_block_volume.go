package sgw

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

// ModifyGatewayBlockVolume invokes the sgw.ModifyGatewayBlockVolume API synchronously
func (client *Client) ModifyGatewayBlockVolume(request *ModifyGatewayBlockVolumeRequest) (response *ModifyGatewayBlockVolumeResponse, err error) {
	response = CreateModifyGatewayBlockVolumeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyGatewayBlockVolumeWithChan invokes the sgw.ModifyGatewayBlockVolume API asynchronously
func (client *Client) ModifyGatewayBlockVolumeWithChan(request *ModifyGatewayBlockVolumeRequest) (<-chan *ModifyGatewayBlockVolumeResponse, <-chan error) {
	responseChan := make(chan *ModifyGatewayBlockVolumeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyGatewayBlockVolume(request)
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

// ModifyGatewayBlockVolumeWithCallback invokes the sgw.ModifyGatewayBlockVolume API asynchronously
func (client *Client) ModifyGatewayBlockVolumeWithCallback(request *ModifyGatewayBlockVolumeRequest, callback func(response *ModifyGatewayBlockVolumeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyGatewayBlockVolumeResponse
		var err error
		defer close(result)
		response, err = client.ModifyGatewayBlockVolume(request)
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

// ModifyGatewayBlockVolumeRequest is the request struct for api ModifyGatewayBlockVolume
type ModifyGatewayBlockVolumeRequest struct {
	*requests.RpcRequest
	CacheConfig   string `position:"Query" name:"CacheConfig"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	IndexId       string `position:"Query" name:"IndexId"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// ModifyGatewayBlockVolumeResponse is the response struct for api ModifyGatewayBlockVolume
type ModifyGatewayBlockVolumeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateModifyGatewayBlockVolumeRequest creates a request to invoke ModifyGatewayBlockVolume API
func CreateModifyGatewayBlockVolumeRequest() (request *ModifyGatewayBlockVolumeRequest) {
	request = &ModifyGatewayBlockVolumeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "ModifyGatewayBlockVolume", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyGatewayBlockVolumeResponse creates a response to parse from ModifyGatewayBlockVolume response
func CreateModifyGatewayBlockVolumeResponse() (response *ModifyGatewayBlockVolumeResponse) {
	response = &ModifyGatewayBlockVolumeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
