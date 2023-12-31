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

// UpdateLiveAudioAuditConfig invokes the live.UpdateLiveAudioAuditConfig API synchronously
func (client *Client) UpdateLiveAudioAuditConfig(request *UpdateLiveAudioAuditConfigRequest) (response *UpdateLiveAudioAuditConfigResponse, err error) {
	response = CreateUpdateLiveAudioAuditConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLiveAudioAuditConfigWithChan invokes the live.UpdateLiveAudioAuditConfig API asynchronously
func (client *Client) UpdateLiveAudioAuditConfigWithChan(request *UpdateLiveAudioAuditConfigRequest) (<-chan *UpdateLiveAudioAuditConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateLiveAudioAuditConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLiveAudioAuditConfig(request)
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

// UpdateLiveAudioAuditConfigWithCallback invokes the live.UpdateLiveAudioAuditConfig API asynchronously
func (client *Client) UpdateLiveAudioAuditConfigWithCallback(request *UpdateLiveAudioAuditConfigRequest, callback func(response *UpdateLiveAudioAuditConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLiveAudioAuditConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateLiveAudioAuditConfig(request)
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

// UpdateLiveAudioAuditConfigRequest is the request struct for api UpdateLiveAudioAuditConfig
type UpdateLiveAudioAuditConfigRequest struct {
	*requests.RpcRequest
	OssEndpoint string           `position:"Query" name:"OssEndpoint"`
	OssObject   string           `position:"Query" name:"OssObject"`
	AppName     string           `position:"Query" name:"AppName"`
	StreamName  string           `position:"Query" name:"StreamName"`
	OssBucket   string           `position:"Query" name:"OssBucket"`
	DomainName  string           `position:"Query" name:"DomainName"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	BizType     string           `position:"Query" name:"BizType"`
}

// UpdateLiveAudioAuditConfigResponse is the response struct for api UpdateLiveAudioAuditConfig
type UpdateLiveAudioAuditConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLiveAudioAuditConfigRequest creates a request to invoke UpdateLiveAudioAuditConfig API
func CreateUpdateLiveAudioAuditConfigRequest() (request *UpdateLiveAudioAuditConfigRequest) {
	request = &UpdateLiveAudioAuditConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateLiveAudioAuditConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateLiveAudioAuditConfigResponse creates a response to parse from UpdateLiveAudioAuditConfig response
func CreateUpdateLiveAudioAuditConfigResponse() (response *UpdateLiveAudioAuditConfigResponse) {
	response = &UpdateLiveAudioAuditConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
