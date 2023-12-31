package cloudcallcenter

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

// ModifyVnTTSConfig invokes the cloudcallcenter.ModifyVnTTSConfig API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnttsconfig.html
func (client *Client) ModifyVnTTSConfig(request *ModifyVnTTSConfigRequest) (response *ModifyVnTTSConfigResponse, err error) {
	response = CreateModifyVnTTSConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVnTTSConfigWithChan invokes the cloudcallcenter.ModifyVnTTSConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnttsconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnTTSConfigWithChan(request *ModifyVnTTSConfigRequest) (<-chan *ModifyVnTTSConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyVnTTSConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVnTTSConfig(request)
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

// ModifyVnTTSConfigWithCallback invokes the cloudcallcenter.ModifyVnTTSConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnttsconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnTTSConfigWithCallback(request *ModifyVnTTSConfigRequest, callback func(response *ModifyVnTTSConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVnTTSConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyVnTTSConfig(request)
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

// ModifyVnTTSConfigRequest is the request struct for api ModifyVnTTSConfig
type ModifyVnTTSConfigRequest struct {
	*requests.RpcRequest
	Voice      string `position:"Query" name:"Voice"`
	Volume     string `position:"Query" name:"Volume"`
	InstanceId string `position:"Query" name:"InstanceId"`
	SpeechRate string `position:"Query" name:"SpeechRate"`
}

// ModifyVnTTSConfigResponse is the response struct for api ModifyVnTTSConfig
type ModifyVnTTSConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVnTTSConfigRequest creates a request to invoke ModifyVnTTSConfig API
func CreateModifyVnTTSConfigRequest() (request *ModifyVnTTSConfigRequest) {
	request = &ModifyVnTTSConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyVnTTSConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateModifyVnTTSConfigResponse creates a response to parse from ModifyVnTTSConfig response
func CreateModifyVnTTSConfigResponse() (response *ModifyVnTTSConfigResponse) {
	response = &ModifyVnTTSConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
