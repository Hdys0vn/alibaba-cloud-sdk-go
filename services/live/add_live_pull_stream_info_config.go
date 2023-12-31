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

// AddLivePullStreamInfoConfig invokes the live.AddLivePullStreamInfoConfig API synchronously
func (client *Client) AddLivePullStreamInfoConfig(request *AddLivePullStreamInfoConfigRequest) (response *AddLivePullStreamInfoConfigResponse, err error) {
	response = CreateAddLivePullStreamInfoConfigResponse()
	err = client.DoAction(request, response)
	return
}

// AddLivePullStreamInfoConfigWithChan invokes the live.AddLivePullStreamInfoConfig API asynchronously
func (client *Client) AddLivePullStreamInfoConfigWithChan(request *AddLivePullStreamInfoConfigRequest) (<-chan *AddLivePullStreamInfoConfigResponse, <-chan error) {
	responseChan := make(chan *AddLivePullStreamInfoConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLivePullStreamInfoConfig(request)
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

// AddLivePullStreamInfoConfigWithCallback invokes the live.AddLivePullStreamInfoConfig API asynchronously
func (client *Client) AddLivePullStreamInfoConfigWithCallback(request *AddLivePullStreamInfoConfigRequest, callback func(response *AddLivePullStreamInfoConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLivePullStreamInfoConfigResponse
		var err error
		defer close(result)
		response, err = client.AddLivePullStreamInfoConfig(request)
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

// AddLivePullStreamInfoConfigRequest is the request struct for api AddLivePullStreamInfoConfig
type AddLivePullStreamInfoConfigRequest struct {
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

// AddLivePullStreamInfoConfigResponse is the response struct for api AddLivePullStreamInfoConfig
type AddLivePullStreamInfoConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLivePullStreamInfoConfigRequest creates a request to invoke AddLivePullStreamInfoConfig API
func CreateAddLivePullStreamInfoConfigRequest() (request *AddLivePullStreamInfoConfigRequest) {
	request = &AddLivePullStreamInfoConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddLivePullStreamInfoConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddLivePullStreamInfoConfigResponse creates a response to parse from AddLivePullStreamInfoConfig response
func CreateAddLivePullStreamInfoConfigResponse() (response *AddLivePullStreamInfoConfigResponse) {
	response = &AddLivePullStreamInfoConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
