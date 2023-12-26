package retailcloud

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

// AllocatePodConfig invokes the retailcloud.AllocatePodConfig API synchronously
func (client *Client) AllocatePodConfig(request *AllocatePodConfigRequest) (response *AllocatePodConfigResponse, err error) {
	response = CreateAllocatePodConfigResponse()
	err = client.DoAction(request, response)
	return
}

// AllocatePodConfigWithChan invokes the retailcloud.AllocatePodConfig API asynchronously
func (client *Client) AllocatePodConfigWithChan(request *AllocatePodConfigRequest) (<-chan *AllocatePodConfigResponse, <-chan error) {
	responseChan := make(chan *AllocatePodConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocatePodConfig(request)
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

// AllocatePodConfigWithCallback invokes the retailcloud.AllocatePodConfig API asynchronously
func (client *Client) AllocatePodConfigWithCallback(request *AllocatePodConfigRequest, callback func(response *AllocatePodConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocatePodConfigResponse
		var err error
		defer close(result)
		response, err = client.AllocatePodConfig(request)
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

// AllocatePodConfigRequest is the request struct for api AllocatePodConfig
type AllocatePodConfigRequest struct {
	*requests.RpcRequest
	RequestId string           `position:"Query" name:"RequestId"`
	AppId     requests.Integer `position:"Query" name:"AppId"`
	EnvId     requests.Integer `position:"Query" name:"EnvId"`
}

// AllocatePodConfigResponse is the response struct for api AllocatePodConfig
type AllocatePodConfigResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateAllocatePodConfigRequest creates a request to invoke AllocatePodConfig API
func CreateAllocatePodConfigRequest() (request *AllocatePodConfigRequest) {
	request = &AllocatePodConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "AllocatePodConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateAllocatePodConfigResponse creates a response to parse from AllocatePodConfig response
func CreateAllocatePodConfigResponse() (response *AllocatePodConfigResponse) {
	response = &AllocatePodConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
