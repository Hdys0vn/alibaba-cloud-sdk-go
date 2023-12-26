package yundun_dbaudit

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

// ConfigInstanceWhiteList invokes the yundun_dbaudit.ConfigInstanceWhiteList API synchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/configinstancewhitelist.html
func (client *Client) ConfigInstanceWhiteList(request *ConfigInstanceWhiteListRequest) (response *ConfigInstanceWhiteListResponse, err error) {
	response = CreateConfigInstanceWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigInstanceWhiteListWithChan invokes the yundun_dbaudit.ConfigInstanceWhiteList API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/configinstancewhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigInstanceWhiteListWithChan(request *ConfigInstanceWhiteListRequest) (<-chan *ConfigInstanceWhiteListResponse, <-chan error) {
	responseChan := make(chan *ConfigInstanceWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigInstanceWhiteList(request)
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

// ConfigInstanceWhiteListWithCallback invokes the yundun_dbaudit.ConfigInstanceWhiteList API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/configinstancewhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigInstanceWhiteListWithCallback(request *ConfigInstanceWhiteListRequest, callback func(response *ConfigInstanceWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigInstanceWhiteListResponse
		var err error
		defer close(result)
		response, err = client.ConfigInstanceWhiteList(request)
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

// ConfigInstanceWhiteListRequest is the request struct for api ConfigInstanceWhiteList
type ConfigInstanceWhiteListRequest struct {
	*requests.RpcRequest
	WhiteList  *[]string `position:"Query" name:"WhiteList"  type:"Repeated"`
	InstanceId string    `position:"Query" name:"InstanceId"`
	SourceIp   string    `position:"Query" name:"SourceIp"`
	IpVersion  string    `position:"Query" name:"IpVersion"`
	Lang       string    `position:"Query" name:"Lang"`
}

// ConfigInstanceWhiteListResponse is the response struct for api ConfigInstanceWhiteList
type ConfigInstanceWhiteListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigInstanceWhiteListRequest creates a request to invoke ConfigInstanceWhiteList API
func CreateConfigInstanceWhiteListRequest() (request *ConfigInstanceWhiteListRequest) {
	request = &ConfigInstanceWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-dbaudit", "2018-10-29", "ConfigInstanceWhiteList", "dbaudit", "openAPI")
	return
}

// CreateConfigInstanceWhiteListResponse creates a response to parse from ConfigInstanceWhiteList response
func CreateConfigInstanceWhiteListResponse() (response *ConfigInstanceWhiteListResponse) {
	response = &ConfigInstanceWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
