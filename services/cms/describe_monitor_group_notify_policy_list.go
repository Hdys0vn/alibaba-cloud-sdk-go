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

// DescribeMonitorGroupNotifyPolicyList invokes the cms.DescribeMonitorGroupNotifyPolicyList API synchronously
func (client *Client) DescribeMonitorGroupNotifyPolicyList(request *DescribeMonitorGroupNotifyPolicyListRequest) (response *DescribeMonitorGroupNotifyPolicyListResponse, err error) {
	response = CreateDescribeMonitorGroupNotifyPolicyListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMonitorGroupNotifyPolicyListWithChan invokes the cms.DescribeMonitorGroupNotifyPolicyList API asynchronously
func (client *Client) DescribeMonitorGroupNotifyPolicyListWithChan(request *DescribeMonitorGroupNotifyPolicyListRequest) (<-chan *DescribeMonitorGroupNotifyPolicyListResponse, <-chan error) {
	responseChan := make(chan *DescribeMonitorGroupNotifyPolicyListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMonitorGroupNotifyPolicyList(request)
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

// DescribeMonitorGroupNotifyPolicyListWithCallback invokes the cms.DescribeMonitorGroupNotifyPolicyList API asynchronously
func (client *Client) DescribeMonitorGroupNotifyPolicyListWithCallback(request *DescribeMonitorGroupNotifyPolicyListRequest, callback func(response *DescribeMonitorGroupNotifyPolicyListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMonitorGroupNotifyPolicyListResponse
		var err error
		defer close(result)
		response, err = client.DescribeMonitorGroupNotifyPolicyList(request)
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

// DescribeMonitorGroupNotifyPolicyListRequest is the request struct for api DescribeMonitorGroupNotifyPolicyList
type DescribeMonitorGroupNotifyPolicyListRequest struct {
	*requests.RpcRequest
	PolicyType string           `position:"Query" name:"PolicyType"`
	GroupId    string           `position:"Query" name:"GroupId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeMonitorGroupNotifyPolicyListResponse is the response struct for api DescribeMonitorGroupNotifyPolicyList
type DescribeMonitorGroupNotifyPolicyListResponse struct {
	*responses.BaseResponse
	Code             string           `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Total            int              `json:"Total" xml:"Total"`
	Success          string           `json:"Success" xml:"Success"`
	NotifyPolicyList NotifyPolicyList `json:"NotifyPolicyList" xml:"NotifyPolicyList"`
}

// CreateDescribeMonitorGroupNotifyPolicyListRequest creates a request to invoke DescribeMonitorGroupNotifyPolicyList API
func CreateDescribeMonitorGroupNotifyPolicyListRequest() (request *DescribeMonitorGroupNotifyPolicyListRequest) {
	request = &DescribeMonitorGroupNotifyPolicyListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMonitorGroupNotifyPolicyList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMonitorGroupNotifyPolicyListResponse creates a response to parse from DescribeMonitorGroupNotifyPolicyList response
func CreateDescribeMonitorGroupNotifyPolicyListResponse() (response *DescribeMonitorGroupNotifyPolicyListResponse) {
	response = &DescribeMonitorGroupNotifyPolicyListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
