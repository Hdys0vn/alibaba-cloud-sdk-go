package cloudauth

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

// DescribeVerifySDK invokes the cloudauth.DescribeVerifySDK API synchronously
func (client *Client) DescribeVerifySDK(request *DescribeVerifySDKRequest) (response *DescribeVerifySDKResponse, err error) {
	response = CreateDescribeVerifySDKResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVerifySDKWithChan invokes the cloudauth.DescribeVerifySDK API asynchronously
func (client *Client) DescribeVerifySDKWithChan(request *DescribeVerifySDKRequest) (<-chan *DescribeVerifySDKResponse, <-chan error) {
	responseChan := make(chan *DescribeVerifySDKResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVerifySDK(request)
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

// DescribeVerifySDKWithCallback invokes the cloudauth.DescribeVerifySDK API asynchronously
func (client *Client) DescribeVerifySDKWithCallback(request *DescribeVerifySDKRequest, callback func(response *DescribeVerifySDKResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVerifySDKResponse
		var err error
		defer close(result)
		response, err = client.DescribeVerifySDK(request)
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

// DescribeVerifySDKRequest is the request struct for api DescribeVerifySDK
type DescribeVerifySDKRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
	TaskId   string `position:"Query" name:"TaskId"`
}

// DescribeVerifySDKResponse is the response struct for api DescribeVerifySDK
type DescribeVerifySDKResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SdkUrl    string `json:"SdkUrl" xml:"SdkUrl"`
}

// CreateDescribeVerifySDKRequest creates a request to invoke DescribeVerifySDK API
func CreateDescribeVerifySDKRequest() (request *DescribeVerifySDKRequest) {
	request = &DescribeVerifySDKRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeVerifySDK", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVerifySDKResponse creates a response to parse from DescribeVerifySDK response
func CreateDescribeVerifySDKResponse() (response *DescribeVerifySDKResponse) {
	response = &DescribeVerifySDKResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
