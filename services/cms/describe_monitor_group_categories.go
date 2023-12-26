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

// DescribeMonitorGroupCategories invokes the cms.DescribeMonitorGroupCategories API synchronously
func (client *Client) DescribeMonitorGroupCategories(request *DescribeMonitorGroupCategoriesRequest) (response *DescribeMonitorGroupCategoriesResponse, err error) {
	response = CreateDescribeMonitorGroupCategoriesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMonitorGroupCategoriesWithChan invokes the cms.DescribeMonitorGroupCategories API asynchronously
func (client *Client) DescribeMonitorGroupCategoriesWithChan(request *DescribeMonitorGroupCategoriesRequest) (<-chan *DescribeMonitorGroupCategoriesResponse, <-chan error) {
	responseChan := make(chan *DescribeMonitorGroupCategoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMonitorGroupCategories(request)
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

// DescribeMonitorGroupCategoriesWithCallback invokes the cms.DescribeMonitorGroupCategories API asynchronously
func (client *Client) DescribeMonitorGroupCategoriesWithCallback(request *DescribeMonitorGroupCategoriesRequest, callback func(response *DescribeMonitorGroupCategoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMonitorGroupCategoriesResponse
		var err error
		defer close(result)
		response, err = client.DescribeMonitorGroupCategories(request)
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

// DescribeMonitorGroupCategoriesRequest is the request struct for api DescribeMonitorGroupCategories
type DescribeMonitorGroupCategoriesRequest struct {
	*requests.RpcRequest
	GroupId requests.Integer `position:"Query" name:"GroupId"`
}

// DescribeMonitorGroupCategoriesResponse is the response struct for api DescribeMonitorGroupCategories
type DescribeMonitorGroupCategoriesResponse struct {
	*responses.BaseResponse
	Code                   int                    `json:"Code" xml:"Code"`
	Message                string                 `json:"Message" xml:"Message"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	Success                bool                   `json:"Success" xml:"Success"`
	MonitorGroupCategories MonitorGroupCategories `json:"MonitorGroupCategories" xml:"MonitorGroupCategories"`
}

// CreateDescribeMonitorGroupCategoriesRequest creates a request to invoke DescribeMonitorGroupCategories API
func CreateDescribeMonitorGroupCategoriesRequest() (request *DescribeMonitorGroupCategoriesRequest) {
	request = &DescribeMonitorGroupCategoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMonitorGroupCategories", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMonitorGroupCategoriesResponse creates a response to parse from DescribeMonitorGroupCategories response
func CreateDescribeMonitorGroupCategoriesResponse() (response *DescribeMonitorGroupCategoriesResponse) {
	response = &DescribeMonitorGroupCategoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
