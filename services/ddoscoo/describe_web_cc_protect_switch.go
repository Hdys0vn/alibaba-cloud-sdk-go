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

// DescribeWebCcProtectSwitch invokes the ddoscoo.DescribeWebCcProtectSwitch API synchronously
func (client *Client) DescribeWebCcProtectSwitch(request *DescribeWebCcProtectSwitchRequest) (response *DescribeWebCcProtectSwitchResponse, err error) {
	response = CreateDescribeWebCcProtectSwitchResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebCcProtectSwitchWithChan invokes the ddoscoo.DescribeWebCcProtectSwitch API asynchronously
func (client *Client) DescribeWebCcProtectSwitchWithChan(request *DescribeWebCcProtectSwitchRequest) (<-chan *DescribeWebCcProtectSwitchResponse, <-chan error) {
	responseChan := make(chan *DescribeWebCcProtectSwitchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebCcProtectSwitch(request)
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

// DescribeWebCcProtectSwitchWithCallback invokes the ddoscoo.DescribeWebCcProtectSwitch API asynchronously
func (client *Client) DescribeWebCcProtectSwitchWithCallback(request *DescribeWebCcProtectSwitchRequest, callback func(response *DescribeWebCcProtectSwitchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebCcProtectSwitchResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebCcProtectSwitch(request)
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

// DescribeWebCcProtectSwitchRequest is the request struct for api DescribeWebCcProtectSwitch
type DescribeWebCcProtectSwitchRequest struct {
	*requests.RpcRequest
	Domains         *[]string `position:"Query" name:"Domains"  type:"Repeated"`
	ResourceGroupId string    `position:"Query" name:"ResourceGroupId"`
	SourceIp        string    `position:"Query" name:"SourceIp"`
}

// DescribeWebCcProtectSwitchResponse is the response struct for api DescribeWebCcProtectSwitch
type DescribeWebCcProtectSwitchResponse struct {
	*responses.BaseResponse
	RequestId         string          `json:"RequestId" xml:"RequestId"`
	ProtectSwitchList []ProtectSwitch `json:"ProtectSwitchList" xml:"ProtectSwitchList"`
}

// CreateDescribeWebCcProtectSwitchRequest creates a request to invoke DescribeWebCcProtectSwitch API
func CreateDescribeWebCcProtectSwitchRequest() (request *DescribeWebCcProtectSwitchRequest) {
	request = &DescribeWebCcProtectSwitchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeWebCcProtectSwitch", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeWebCcProtectSwitchResponse creates a response to parse from DescribeWebCcProtectSwitch response
func CreateDescribeWebCcProtectSwitchResponse() (response *DescribeWebCcProtectSwitchResponse) {
	response = &DescribeWebCcProtectSwitchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
