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

// DescribeScreenSetting invokes the aegis.DescribeScreenSetting API synchronously
// api document: https://help.aliyun.com/api/aegis/describescreensetting.html
func (client *Client) DescribeScreenSetting(request *DescribeScreenSettingRequest) (response *DescribeScreenSettingResponse, err error) {
	response = CreateDescribeScreenSettingResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScreenSettingWithChan invokes the aegis.DescribeScreenSetting API asynchronously
// api document: https://help.aliyun.com/api/aegis/describescreensetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScreenSettingWithChan(request *DescribeScreenSettingRequest) (<-chan *DescribeScreenSettingResponse, <-chan error) {
	responseChan := make(chan *DescribeScreenSettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScreenSetting(request)
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

// DescribeScreenSettingWithCallback invokes the aegis.DescribeScreenSetting API asynchronously
// api document: https://help.aliyun.com/api/aegis/describescreensetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScreenSettingWithCallback(request *DescribeScreenSettingRequest, callback func(response *DescribeScreenSettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScreenSettingResponse
		var err error
		defer close(result)
		response, err = client.DescribeScreenSetting(request)
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

// DescribeScreenSettingRequest is the request struct for api DescribeScreenSetting
type DescribeScreenSettingRequest struct {
	*requests.RpcRequest
	SourceIp    string `position:"Query" name:"SourceIp"`
	ScreenTitle string `position:"Query" name:"ScreenTitle"`
}

// DescribeScreenSettingResponse is the response struct for api DescribeScreenSetting
type DescribeScreenSettingResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	SasScreenSetting SasScreenSetting `json:"SasScreenSetting" xml:"SasScreenSetting"`
}

// CreateDescribeScreenSettingRequest creates a request to invoke DescribeScreenSetting API
func CreateDescribeScreenSettingRequest() (request *DescribeScreenSettingRequest) {
	request = &DescribeScreenSettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeScreenSetting", "vipaegis", "openAPI")
	return
}

// CreateDescribeScreenSettingResponse creates a response to parse from DescribeScreenSetting response
func CreateDescribeScreenSettingResponse() (response *DescribeScreenSettingResponse) {
	response = &DescribeScreenSettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
