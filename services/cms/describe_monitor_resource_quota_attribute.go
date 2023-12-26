package cms

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

// DescribeMonitorResourceQuotaAttribute invokes the cms.DescribeMonitorResourceQuotaAttribute API synchronously
func (client *Client) DescribeMonitorResourceQuotaAttribute(request *DescribeMonitorResourceQuotaAttributeRequest) (response *DescribeMonitorResourceQuotaAttributeResponse, err error) {
	response = CreateDescribeMonitorResourceQuotaAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMonitorResourceQuotaAttributeWithChan invokes the cms.DescribeMonitorResourceQuotaAttribute API asynchronously
func (client *Client) DescribeMonitorResourceQuotaAttributeWithChan(request *DescribeMonitorResourceQuotaAttributeRequest) (<-chan *DescribeMonitorResourceQuotaAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeMonitorResourceQuotaAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMonitorResourceQuotaAttribute(request)
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

// DescribeMonitorResourceQuotaAttributeWithCallback invokes the cms.DescribeMonitorResourceQuotaAttribute API asynchronously
func (client *Client) DescribeMonitorResourceQuotaAttributeWithCallback(request *DescribeMonitorResourceQuotaAttributeRequest, callback func(response *DescribeMonitorResourceQuotaAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMonitorResourceQuotaAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeMonitorResourceQuotaAttribute(request)
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

// DescribeMonitorResourceQuotaAttributeRequest is the request struct for api DescribeMonitorResourceQuotaAttribute
type DescribeMonitorResourceQuotaAttributeRequest struct {
	*requests.RpcRequest
	ShowUsed requests.Boolean `position:"Query" name:"ShowUsed"`
}

// DescribeMonitorResourceQuotaAttributeResponse is the response struct for api DescribeMonitorResourceQuotaAttribute
type DescribeMonitorResourceQuotaAttributeResponse struct {
	*responses.BaseResponse
	Code          string        `json:"Code" xml:"Code"`
	Message       string        `json:"Message" xml:"Message"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ResourceQuota ResourceQuota `json:"ResourceQuota" xml:"ResourceQuota"`
}

// CreateDescribeMonitorResourceQuotaAttributeRequest creates a request to invoke DescribeMonitorResourceQuotaAttribute API
func CreateDescribeMonitorResourceQuotaAttributeRequest() (request *DescribeMonitorResourceQuotaAttributeRequest) {
	request = &DescribeMonitorResourceQuotaAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMonitorResourceQuotaAttribute", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMonitorResourceQuotaAttributeResponse creates a response to parse from DescribeMonitorResourceQuotaAttribute response
func CreateDescribeMonitorResourceQuotaAttributeResponse() (response *DescribeMonitorResourceQuotaAttributeResponse) {
	response = &DescribeMonitorResourceQuotaAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
