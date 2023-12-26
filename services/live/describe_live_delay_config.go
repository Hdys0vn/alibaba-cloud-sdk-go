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

// DescribeLiveDelayConfig invokes the live.DescribeLiveDelayConfig API synchronously
func (client *Client) DescribeLiveDelayConfig(request *DescribeLiveDelayConfigRequest) (response *DescribeLiveDelayConfigResponse, err error) {
	response = CreateDescribeLiveDelayConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDelayConfigWithChan invokes the live.DescribeLiveDelayConfig API asynchronously
func (client *Client) DescribeLiveDelayConfigWithChan(request *DescribeLiveDelayConfigRequest) (<-chan *DescribeLiveDelayConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDelayConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDelayConfig(request)
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

// DescribeLiveDelayConfigWithCallback invokes the live.DescribeLiveDelayConfig API asynchronously
func (client *Client) DescribeLiveDelayConfigWithCallback(request *DescribeLiveDelayConfigRequest, callback func(response *DescribeLiveDelayConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDelayConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDelayConfig(request)
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

// DescribeLiveDelayConfigRequest is the request struct for api DescribeLiveDelayConfig
type DescribeLiveDelayConfigRequest struct {
	*requests.RpcRequest
	Stream  string           `position:"Query" name:"Stream"`
	App     string           `position:"Query" name:"App"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	Domain  string           `position:"Query" name:"Domain"`
}

// DescribeLiveDelayConfigResponse is the response struct for api DescribeLiveDelayConfig
type DescribeLiveDelayConfigResponse struct {
	*responses.BaseResponse
	Domain          string `json:"Domain" xml:"Domain"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	TaskTriggerMode string `json:"TaskTriggerMode" xml:"TaskTriggerMode"`
	App             string `json:"App" xml:"App"`
	DelayTime       string `json:"DelayTime" xml:"DelayTime"`
	Stream          string `json:"Stream" xml:"Stream"`
}

// CreateDescribeLiveDelayConfigRequest creates a request to invoke DescribeLiveDelayConfig API
func CreateDescribeLiveDelayConfigRequest() (request *DescribeLiveDelayConfigRequest) {
	request = &DescribeLiveDelayConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDelayConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDelayConfigResponse creates a response to parse from DescribeLiveDelayConfig response
func CreateDescribeLiveDelayConfigResponse() (response *DescribeLiveDelayConfigResponse) {
	response = &DescribeLiveDelayConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}