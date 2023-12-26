package aegis

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

// ModifyWebLockCreateConfig invokes the aegis.ModifyWebLockCreateConfig API synchronously
// api document: https://help.aliyun.com/api/aegis/modifyweblockcreateconfig.html
func (client *Client) ModifyWebLockCreateConfig(request *ModifyWebLockCreateConfigRequest) (response *ModifyWebLockCreateConfigResponse, err error) {
	response = CreateModifyWebLockCreateConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyWebLockCreateConfigWithChan invokes the aegis.ModifyWebLockCreateConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifyweblockcreateconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWebLockCreateConfigWithChan(request *ModifyWebLockCreateConfigRequest) (<-chan *ModifyWebLockCreateConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyWebLockCreateConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyWebLockCreateConfig(request)
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

// ModifyWebLockCreateConfigWithCallback invokes the aegis.ModifyWebLockCreateConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifyweblockcreateconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWebLockCreateConfigWithCallback(request *ModifyWebLockCreateConfigRequest, callback func(response *ModifyWebLockCreateConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyWebLockCreateConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyWebLockCreateConfig(request)
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

// ModifyWebLockCreateConfigRequest is the request struct for api ModifyWebLockCreateConfig
type ModifyWebLockCreateConfigRequest struct {
	*requests.RpcRequest
	LocalBackupDir    string `position:"Query" name:"LocalBackupDir"`
	SourceIp          string `position:"Query" name:"SourceIp"`
	ExclusiveFileType string `position:"Query" name:"ExclusiveFileType"`
	Lang              string `position:"Query" name:"Lang"`
	Dir               string `position:"Query" name:"Dir"`
	Uuid              string `position:"Query" name:"Uuid"`
	ExclusiveDir      string `position:"Query" name:"ExclusiveDir"`
}

// ModifyWebLockCreateConfigResponse is the response struct for api ModifyWebLockCreateConfig
type ModifyWebLockCreateConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyWebLockCreateConfigRequest creates a request to invoke ModifyWebLockCreateConfig API
func CreateModifyWebLockCreateConfigRequest() (request *ModifyWebLockCreateConfigRequest) {
	request = &ModifyWebLockCreateConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifyWebLockCreateConfig", "vipaegis", "openAPI")
	return
}

// CreateModifyWebLockCreateConfigResponse creates a response to parse from ModifyWebLockCreateConfig response
func CreateModifyWebLockCreateConfigResponse() (response *ModifyWebLockCreateConfigResponse) {
	response = &ModifyWebLockCreateConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
