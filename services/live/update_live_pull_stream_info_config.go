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

// UpdateLivePullStreamInfoConfig invokes the live.UpdateLivePullStreamInfoConfig API synchronously
func (client *Client) UpdateLivePullStreamInfoConfig(request *UpdateLivePullStreamInfoConfigRequest) (response *UpdateLivePullStreamInfoConfigResponse, err error) {
	response = CreateUpdateLivePullStreamInfoConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLivePullStreamInfoConfigWithChan invokes the live.UpdateLivePullStreamInfoConfig API asynchronously
func (client *Client) UpdateLivePullStreamInfoConfigWithChan(request *UpdateLivePullStreamInfoConfigRequest) (<-chan *UpdateLivePullStreamInfoConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateLivePullStreamInfoConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLivePullStreamInfoConfig(request)
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

// UpdateLivePullStreamInfoConfigWithCallback invokes the live.UpdateLivePullStreamInfoConfig API asynchronously
func (client *Client) UpdateLivePullStreamInfoConfigWithCallback(request *UpdateLivePullStreamInfoConfigRequest, callback func(response *UpdateLivePullStreamInfoConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLivePullStreamInfoConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateLivePullStreamInfoConfig(request)
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

// UpdateLivePullStreamInfoConfigRequest is the request struct for api UpdateLivePullStreamInfoConfig
type UpdateLivePullStreamInfoConfigRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	PullAlways string           `position:"Query" name:"PullAlways"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	SourceUrl  string           `position:"Query" name:"SourceUrl"`
}

// UpdateLivePullStreamInfoConfigResponse is the response struct for api UpdateLivePullStreamInfoConfig
type UpdateLivePullStreamInfoConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLivePullStreamInfoConfigRequest creates a request to invoke UpdateLivePullStreamInfoConfig API
func CreateUpdateLivePullStreamInfoConfigRequest() (request *UpdateLivePullStreamInfoConfigRequest) {
	request = &UpdateLivePullStreamInfoConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateLivePullStreamInfoConfig", "live", "openAPI")
	request.Method = requests.GET
	return
}

// CreateUpdateLivePullStreamInfoConfigResponse creates a response to parse from UpdateLivePullStreamInfoConfig response
func CreateUpdateLivePullStreamInfoConfigResponse() (response *UpdateLivePullStreamInfoConfigResponse) {
	response = &UpdateLivePullStreamInfoConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
