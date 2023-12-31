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

// DescribeNsasSuspEventType invokes the aegis.DescribeNsasSuspEventType API synchronously
// api document: https://help.aliyun.com/api/aegis/describensassuspeventtype.html
func (client *Client) DescribeNsasSuspEventType(request *DescribeNsasSuspEventTypeRequest) (response *DescribeNsasSuspEventTypeResponse, err error) {
	response = CreateDescribeNsasSuspEventTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNsasSuspEventTypeWithChan invokes the aegis.DescribeNsasSuspEventType API asynchronously
// api document: https://help.aliyun.com/api/aegis/describensassuspeventtype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNsasSuspEventTypeWithChan(request *DescribeNsasSuspEventTypeRequest) (<-chan *DescribeNsasSuspEventTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeNsasSuspEventTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNsasSuspEventType(request)
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

// DescribeNsasSuspEventTypeWithCallback invokes the aegis.DescribeNsasSuspEventType API asynchronously
// api document: https://help.aliyun.com/api/aegis/describensassuspeventtype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNsasSuspEventTypeWithCallback(request *DescribeNsasSuspEventTypeRequest, callback func(response *DescribeNsasSuspEventTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNsasSuspEventTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeNsasSuspEventType(request)
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

// DescribeNsasSuspEventTypeRequest is the request struct for api DescribeNsasSuspEventType
type DescribeNsasSuspEventTypeRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Name     string `position:"Query" name:"Name"`
	Remark   string `position:"Query" name:"Remark"`
	From     string `position:"Query" name:"From"`
	Lang     string `position:"Query" name:"Lang"`
}

// DescribeNsasSuspEventTypeResponse is the response struct for api DescribeNsasSuspEventType
type DescribeNsasSuspEventTypeResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	EventTypes []Data `json:"EventTypes" xml:"EventTypes"`
}

// CreateDescribeNsasSuspEventTypeRequest creates a request to invoke DescribeNsasSuspEventType API
func CreateDescribeNsasSuspEventTypeRequest() (request *DescribeNsasSuspEventTypeRequest) {
	request = &DescribeNsasSuspEventTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeNsasSuspEventType", "vipaegis", "openAPI")
	return
}

// CreateDescribeNsasSuspEventTypeResponse creates a response to parse from DescribeNsasSuspEventType response
func CreateDescribeNsasSuspEventTypeResponse() (response *DescribeNsasSuspEventTypeResponse) {
	response = &DescribeNsasSuspEventTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
