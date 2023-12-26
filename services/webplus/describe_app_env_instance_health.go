package webplus

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

// DescribeAppEnvInstanceHealth invokes the webplus.DescribeAppEnvInstanceHealth API synchronously
// api document: https://help.aliyun.com/api/webplus/describeappenvinstancehealth.html
func (client *Client) DescribeAppEnvInstanceHealth(request *DescribeAppEnvInstanceHealthRequest) (response *DescribeAppEnvInstanceHealthResponse, err error) {
	response = CreateDescribeAppEnvInstanceHealthResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppEnvInstanceHealthWithChan invokes the webplus.DescribeAppEnvInstanceHealth API asynchronously
// api document: https://help.aliyun.com/api/webplus/describeappenvinstancehealth.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppEnvInstanceHealthWithChan(request *DescribeAppEnvInstanceHealthRequest) (<-chan *DescribeAppEnvInstanceHealthResponse, <-chan error) {
	responseChan := make(chan *DescribeAppEnvInstanceHealthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAppEnvInstanceHealth(request)
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

// DescribeAppEnvInstanceHealthWithCallback invokes the webplus.DescribeAppEnvInstanceHealth API asynchronously
// api document: https://help.aliyun.com/api/webplus/describeappenvinstancehealth.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppEnvInstanceHealthWithCallback(request *DescribeAppEnvInstanceHealthRequest, callback func(response *DescribeAppEnvInstanceHealthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppEnvInstanceHealthResponse
		var err error
		defer close(result)
		response, err = client.DescribeAppEnvInstanceHealth(request)
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

// DescribeAppEnvInstanceHealthRequest is the request struct for api DescribeAppEnvInstanceHealth
type DescribeAppEnvInstanceHealthRequest struct {
	*requests.RoaRequest
	EnvId string `position:"Query" name:"EnvId"`
}

// DescribeAppEnvInstanceHealthResponse is the response struct for api DescribeAppEnvInstanceHealth
type DescribeAppEnvInstanceHealthResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Code              string            `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	EnvInstanceHealth EnvInstanceHealth `json:"EnvInstanceHealth" xml:"EnvInstanceHealth"`
}

// CreateDescribeAppEnvInstanceHealthRequest creates a request to invoke DescribeAppEnvInstanceHealth API
func CreateDescribeAppEnvInstanceHealthRequest() (request *DescribeAppEnvInstanceHealthRequest) {
	request = &DescribeAppEnvInstanceHealthRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DescribeAppEnvInstanceHealth", "/pop/v1/wam/appEnv/instanceHealth", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeAppEnvInstanceHealthResponse creates a response to parse from DescribeAppEnvInstanceHealth response
func CreateDescribeAppEnvInstanceHealthResponse() (response *DescribeAppEnvInstanceHealthResponse) {
	response = &DescribeAppEnvInstanceHealthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}