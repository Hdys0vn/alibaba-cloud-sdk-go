package sae

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

// DescribeApplicationConfig invokes the sae.DescribeApplicationConfig API synchronously
func (client *Client) DescribeApplicationConfig(request *DescribeApplicationConfigRequest) (response *DescribeApplicationConfigResponse, err error) {
	response = CreateDescribeApplicationConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApplicationConfigWithChan invokes the sae.DescribeApplicationConfig API asynchronously
func (client *Client) DescribeApplicationConfigWithChan(request *DescribeApplicationConfigRequest) (<-chan *DescribeApplicationConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeApplicationConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApplicationConfig(request)
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

// DescribeApplicationConfigWithCallback invokes the sae.DescribeApplicationConfig API asynchronously
func (client *Client) DescribeApplicationConfigWithCallback(request *DescribeApplicationConfigRequest, callback func(response *DescribeApplicationConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApplicationConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeApplicationConfig(request)
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

// DescribeApplicationConfigRequest is the request struct for api DescribeApplicationConfig
type DescribeApplicationConfigRequest struct {
	*requests.RoaRequest
	VersionId string `position:"Query" name:"VersionId"`
	AppId     string `position:"Query" name:"AppId"`
}

// DescribeApplicationConfigResponse is the response struct for api DescribeApplicationConfig
type DescribeApplicationConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeApplicationConfigRequest creates a request to invoke DescribeApplicationConfig API
func CreateDescribeApplicationConfigRequest() (request *DescribeApplicationConfigRequest) {
	request = &DescribeApplicationConfigRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeApplicationConfig", "/pop/v1/sam/app/describeApplicationConfig", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeApplicationConfigResponse creates a response to parse from DescribeApplicationConfig response
func CreateDescribeApplicationConfigResponse() (response *DescribeApplicationConfigResponse) {
	response = &DescribeApplicationConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
