package ddoscoo

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

// ConfigLayer4Remark invokes the ddoscoo.ConfigLayer4Remark API synchronously
func (client *Client) ConfigLayer4Remark(request *ConfigLayer4RemarkRequest) (response *ConfigLayer4RemarkResponse, err error) {
	response = CreateConfigLayer4RemarkResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigLayer4RemarkWithChan invokes the ddoscoo.ConfigLayer4Remark API asynchronously
func (client *Client) ConfigLayer4RemarkWithChan(request *ConfigLayer4RemarkRequest) (<-chan *ConfigLayer4RemarkResponse, <-chan error) {
	responseChan := make(chan *ConfigLayer4RemarkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigLayer4Remark(request)
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

// ConfigLayer4RemarkWithCallback invokes the ddoscoo.ConfigLayer4Remark API asynchronously
func (client *Client) ConfigLayer4RemarkWithCallback(request *ConfigLayer4RemarkRequest, callback func(response *ConfigLayer4RemarkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigLayer4RemarkResponse
		var err error
		defer close(result)
		response, err = client.ConfigLayer4Remark(request)
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

// ConfigLayer4RemarkRequest is the request struct for api ConfigLayer4Remark
type ConfigLayer4RemarkRequest struct {
	*requests.RpcRequest
	Listeners string `position:"Query" name:"Listeners"`
	SourceIp  string `position:"Query" name:"SourceIp"`
}

// ConfigLayer4RemarkResponse is the response struct for api ConfigLayer4Remark
type ConfigLayer4RemarkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigLayer4RemarkRequest creates a request to invoke ConfigLayer4Remark API
func CreateConfigLayer4RemarkRequest() (request *ConfigLayer4RemarkRequest) {
	request = &ConfigLayer4RemarkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ConfigLayer4Remark", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigLayer4RemarkResponse creates a response to parse from ConfigLayer4Remark response
func CreateConfigLayer4RemarkResponse() (response *ConfigLayer4RemarkResponse) {
	response = &ConfigLayer4RemarkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
