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

// DescribeVulConfig invokes the aegis.DescribeVulConfig API synchronously
// api document: https://help.aliyun.com/api/aegis/describevulconfig.html
func (client *Client) DescribeVulConfig(request *DescribeVulConfigRequest) (response *DescribeVulConfigResponse, err error) {
	response = CreateDescribeVulConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVulConfigWithChan invokes the aegis.DescribeVulConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describevulconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVulConfigWithChan(request *DescribeVulConfigRequest) (<-chan *DescribeVulConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeVulConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVulConfig(request)
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

// DescribeVulConfigWithCallback invokes the aegis.DescribeVulConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describevulconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVulConfigWithCallback(request *DescribeVulConfigRequest, callback func(response *DescribeVulConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVulConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeVulConfig(request)
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

// DescribeVulConfigRequest is the request struct for api DescribeVulConfig
type DescribeVulConfigRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Type     string `position:"Query" name:"Type"`
}

// DescribeVulConfigResponse is the response struct for api DescribeVulConfig
type DescribeVulConfigResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	TotalCount    int            `json:"TotalCount" xml:"TotalCount"`
	TargetConfigs []TargetConfig `json:"TargetConfigs" xml:"TargetConfigs"`
}

// CreateDescribeVulConfigRequest creates a request to invoke DescribeVulConfig API
func CreateDescribeVulConfigRequest() (request *DescribeVulConfigRequest) {
	request = &DescribeVulConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeVulConfig", "vipaegis", "openAPI")
	return
}

// CreateDescribeVulConfigResponse creates a response to parse from DescribeVulConfig response
func CreateDescribeVulConfigResponse() (response *DescribeVulConfigResponse) {
	response = &DescribeVulConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
