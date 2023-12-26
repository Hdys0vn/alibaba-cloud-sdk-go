package live

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

// UpdateCasterSceneConfig invokes the live.UpdateCasterSceneConfig API synchronously
func (client *Client) UpdateCasterSceneConfig(request *UpdateCasterSceneConfigRequest) (response *UpdateCasterSceneConfigResponse, err error) {
	response = CreateUpdateCasterSceneConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCasterSceneConfigWithChan invokes the live.UpdateCasterSceneConfig API asynchronously
func (client *Client) UpdateCasterSceneConfigWithChan(request *UpdateCasterSceneConfigRequest) (<-chan *UpdateCasterSceneConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateCasterSceneConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCasterSceneConfig(request)
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

// UpdateCasterSceneConfigWithCallback invokes the live.UpdateCasterSceneConfig API asynchronously
func (client *Client) UpdateCasterSceneConfigWithCallback(request *UpdateCasterSceneConfigRequest, callback func(response *UpdateCasterSceneConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCasterSceneConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateCasterSceneConfig(request)
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

// UpdateCasterSceneConfigRequest is the request struct for api UpdateCasterSceneConfig
type UpdateCasterSceneConfigRequest struct {
	*requests.RpcRequest
	LayoutId    string           `position:"Query" name:"LayoutId"`
	ComponentId *[]string        `position:"Query" name:"ComponentId"  type:"Repeated"`
	CasterId    string           `position:"Query" name:"CasterId"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	SceneId     string           `position:"Query" name:"SceneId"`
}

// UpdateCasterSceneConfigResponse is the response struct for api UpdateCasterSceneConfig
type UpdateCasterSceneConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateCasterSceneConfigRequest creates a request to invoke UpdateCasterSceneConfig API
func CreateUpdateCasterSceneConfigRequest() (request *UpdateCasterSceneConfigRequest) {
	request = &UpdateCasterSceneConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateCasterSceneConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateCasterSceneConfigResponse creates a response to parse from UpdateCasterSceneConfig response
func CreateUpdateCasterSceneConfigResponse() (response *UpdateCasterSceneConfigResponse) {
	response = &UpdateCasterSceneConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
