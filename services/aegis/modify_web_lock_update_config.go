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

// ModifyWebLockUpdateConfig invokes the aegis.ModifyWebLockUpdateConfig API synchronously
// api document: https://help.aliyun.com/api/aegis/modifyweblockupdateconfig.html
func (client *Client) ModifyWebLockUpdateConfig(request *ModifyWebLockUpdateConfigRequest) (response *ModifyWebLockUpdateConfigResponse, err error) {
	response = CreateModifyWebLockUpdateConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyWebLockUpdateConfigWithChan invokes the aegis.ModifyWebLockUpdateConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifyweblockupdateconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWebLockUpdateConfigWithChan(request *ModifyWebLockUpdateConfigRequest) (<-chan *ModifyWebLockUpdateConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyWebLockUpdateConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyWebLockUpdateConfig(request)
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

// ModifyWebLockUpdateConfigWithCallback invokes the aegis.ModifyWebLockUpdateConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifyweblockupdateconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWebLockUpdateConfigWithCallback(request *ModifyWebLockUpdateConfigRequest, callback func(response *ModifyWebLockUpdateConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyWebLockUpdateConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyWebLockUpdateConfig(request)
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

// ModifyWebLockUpdateConfigRequest is the request struct for api ModifyWebLockUpdateConfig
type ModifyWebLockUpdateConfigRequest struct {
	*requests.RpcRequest
	LocalBackupDir    string           `position:"Query" name:"LocalBackupDir"`
	SourceIp          string           `position:"Query" name:"SourceIp"`
	ExclusiveFileType string           `position:"Query" name:"ExclusiveFileType"`
	Id                requests.Integer `position:"Query" name:"Id"`
	Lang              string           `position:"Query" name:"Lang"`
	Dir               string           `position:"Query" name:"Dir"`
	Uuid              string           `position:"Query" name:"Uuid"`
	ExclusiveDir      string           `position:"Query" name:"ExclusiveDir"`
}

// ModifyWebLockUpdateConfigResponse is the response struct for api ModifyWebLockUpdateConfig
type ModifyWebLockUpdateConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyWebLockUpdateConfigRequest creates a request to invoke ModifyWebLockUpdateConfig API
func CreateModifyWebLockUpdateConfigRequest() (request *ModifyWebLockUpdateConfigRequest) {
	request = &ModifyWebLockUpdateConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifyWebLockUpdateConfig", "vipaegis", "openAPI")
	return
}

// CreateModifyWebLockUpdateConfigResponse creates a response to parse from ModifyWebLockUpdateConfig response
func CreateModifyWebLockUpdateConfigResponse() (response *ModifyWebLockUpdateConfigResponse) {
	response = &ModifyWebLockUpdateConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
