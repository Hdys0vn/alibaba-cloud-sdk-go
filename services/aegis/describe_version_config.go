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

// DescribeVersionConfig invokes the aegis.DescribeVersionConfig API synchronously
// api document: https://help.aliyun.com/api/aegis/describeversionconfig.html
func (client *Client) DescribeVersionConfig(request *DescribeVersionConfigRequest) (response *DescribeVersionConfigResponse, err error) {
	response = CreateDescribeVersionConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVersionConfigWithChan invokes the aegis.DescribeVersionConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeversionconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVersionConfigWithChan(request *DescribeVersionConfigRequest) (<-chan *DescribeVersionConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeVersionConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVersionConfig(request)
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

// DescribeVersionConfigWithCallback invokes the aegis.DescribeVersionConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeversionconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVersionConfigWithCallback(request *DescribeVersionConfigRequest, callback func(response *DescribeVersionConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVersionConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeVersionConfig(request)
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

// DescribeVersionConfigRequest is the request struct for api DescribeVersionConfig
type DescribeVersionConfigRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeVersionConfigResponse is the response struct for api DescribeVersionConfig
type DescribeVersionConfigResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	AssetLevel            int    `json:"AssetLevel" xml:"AssetLevel"`
	AvdsFlag              int    `json:"AvdsFlag" xml:"AvdsFlag"`
	CreateTime            int    `json:"CreateTime" xml:"CreateTime"`
	Flag                  int    `json:"Flag" xml:"Flag"`
	InstanceId            string `json:"InstanceId" xml:"InstanceId"`
	IsSasOpening          bool   `json:"IsSasOpening" xml:"IsSasOpening"`
	IsTrialVersion        int    `json:"IsTrialVersion" xml:"IsTrialVersion"`
	LogCapacity           int    `json:"LogCapacity" xml:"LogCapacity"`
	LogTime               int    `json:"LogTime" xml:"LogTime"`
	ReleaseTime           int    `json:"ReleaseTime" xml:"ReleaseTime"`
	SasLog                int    `json:"SasLog" xml:"SasLog"`
	SasScreen             int    `json:"SasScreen" xml:"SasScreen"`
	Version               int    `json:"Version" xml:"Version"`
	UserDefinedAlarms     int    `json:"UserDefinedAlarms" xml:"UserDefinedAlarms"`
	WebLock               int    `json:"WebLock" xml:"WebLock"`
	WebLockAuthCount      int    `json:"WebLockAuthCount" xml:"WebLockAuthCount"`
	AppWhiteListAuthCount int    `json:"AppWhiteListAuthCount" xml:"AppWhiteListAuthCount"`
	AppWhiteList          int    `json:"AppWhiteList" xml:"AppWhiteList"`
	SlsCapacity           int    `json:"SlsCapacity" xml:"SlsCapacity"`
}

// CreateDescribeVersionConfigRequest creates a request to invoke DescribeVersionConfig API
func CreateDescribeVersionConfigRequest() (request *DescribeVersionConfigRequest) {
	request = &DescribeVersionConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeVersionConfig", "vipaegis", "openAPI")
	return
}

// CreateDescribeVersionConfigResponse creates a response to parse from DescribeVersionConfig response
func CreateDescribeVersionConfigResponse() (response *DescribeVersionConfigResponse) {
	response = &DescribeVersionConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
