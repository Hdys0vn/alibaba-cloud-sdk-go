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

// DescribeSuspiciousUUIDConfig invokes the aegis.DescribeSuspiciousUUIDConfig API synchronously
// api document: https://help.aliyun.com/api/aegis/describesuspiciousuuidconfig.html
func (client *Client) DescribeSuspiciousUUIDConfig(request *DescribeSuspiciousUUIDConfigRequest) (response *DescribeSuspiciousUUIDConfigResponse, err error) {
	response = CreateDescribeSuspiciousUUIDConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSuspiciousUUIDConfigWithChan invokes the aegis.DescribeSuspiciousUUIDConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesuspiciousuuidconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSuspiciousUUIDConfigWithChan(request *DescribeSuspiciousUUIDConfigRequest) (<-chan *DescribeSuspiciousUUIDConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeSuspiciousUUIDConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSuspiciousUUIDConfig(request)
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

// DescribeSuspiciousUUIDConfigWithCallback invokes the aegis.DescribeSuspiciousUUIDConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesuspiciousuuidconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSuspiciousUUIDConfigWithCallback(request *DescribeSuspiciousUUIDConfigRequest, callback func(response *DescribeSuspiciousUUIDConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSuspiciousUUIDConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeSuspiciousUUIDConfig(request)
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

// DescribeSuspiciousUUIDConfigRequest is the request struct for api DescribeSuspiciousUUIDConfig
type DescribeSuspiciousUUIDConfigRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
	Type     string `position:"Query" name:"Type"`
}

// DescribeSuspiciousUUIDConfigResponse is the response struct for api DescribeSuspiciousUUIDConfig
type DescribeSuspiciousUUIDConfigResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Count     int      `json:"Count" xml:"Count"`
	UUIDList  []string `json:"UUIDList" xml:"UUIDList"`
}

// CreateDescribeSuspiciousUUIDConfigRequest creates a request to invoke DescribeSuspiciousUUIDConfig API
func CreateDescribeSuspiciousUUIDConfigRequest() (request *DescribeSuspiciousUUIDConfigRequest) {
	request = &DescribeSuspiciousUUIDConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeSuspiciousUUIDConfig", "vipaegis", "openAPI")
	return
}

// CreateDescribeSuspiciousUUIDConfigResponse creates a response to parse from DescribeSuspiciousUUIDConfig response
func CreateDescribeSuspiciousUUIDConfigResponse() (response *DescribeSuspiciousUUIDConfigResponse) {
	response = &DescribeSuspiciousUUIDConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
