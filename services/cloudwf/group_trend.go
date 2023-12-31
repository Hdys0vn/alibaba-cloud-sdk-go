package cloudwf

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

// GroupTrend invokes the cloudwf.GroupTrend API synchronously
// api document: https://help.aliyun.com/api/cloudwf/grouptrend.html
func (client *Client) GroupTrend(request *GroupTrendRequest) (response *GroupTrendResponse, err error) {
	response = CreateGroupTrendResponse()
	err = client.DoAction(request, response)
	return
}

// GroupTrendWithChan invokes the cloudwf.GroupTrend API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/grouptrend.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GroupTrendWithChan(request *GroupTrendRequest) (<-chan *GroupTrendResponse, <-chan error) {
	responseChan := make(chan *GroupTrendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GroupTrend(request)
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

// GroupTrendWithCallback invokes the cloudwf.GroupTrend API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/grouptrend.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GroupTrendWithCallback(request *GroupTrendRequest, callback func(response *GroupTrendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GroupTrendResponse
		var err error
		defer close(result)
		response, err = client.GroupTrend(request)
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

// GroupTrendRequest is the request struct for api GroupTrend
type GroupTrendRequest struct {
	*requests.RpcRequest
	Gsid requests.Integer `position:"Query" name:"Gsid"`
}

// GroupTrendResponse is the response struct for api GroupTrend
type GroupTrendResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGroupTrendRequest creates a request to invoke GroupTrend API
func CreateGroupTrendRequest() (request *GroupTrendRequest) {
	request = &GroupTrendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GroupTrend", "cloudwf", "openAPI")
	return
}

// CreateGroupTrendResponse creates a response to parse from GroupTrend response
func CreateGroupTrendResponse() (response *GroupTrendResponse) {
	response = &GroupTrendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
