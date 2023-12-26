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

// HeadquartersOverview invokes the cloudwf.HeadquartersOverview API synchronously
// api document: https://help.aliyun.com/api/cloudwf/headquartersoverview.html
func (client *Client) HeadquartersOverview(request *HeadquartersOverviewRequest) (response *HeadquartersOverviewResponse, err error) {
	response = CreateHeadquartersOverviewResponse()
	err = client.DoAction(request, response)
	return
}

// HeadquartersOverviewWithChan invokes the cloudwf.HeadquartersOverview API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/headquartersoverview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) HeadquartersOverviewWithChan(request *HeadquartersOverviewRequest) (<-chan *HeadquartersOverviewResponse, <-chan error) {
	responseChan := make(chan *HeadquartersOverviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.HeadquartersOverview(request)
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

// HeadquartersOverviewWithCallback invokes the cloudwf.HeadquartersOverview API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/headquartersoverview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) HeadquartersOverviewWithCallback(request *HeadquartersOverviewRequest, callback func(response *HeadquartersOverviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *HeadquartersOverviewResponse
		var err error
		defer close(result)
		response, err = client.HeadquartersOverview(request)
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

// HeadquartersOverviewRequest is the request struct for api HeadquartersOverview
type HeadquartersOverviewRequest struct {
	*requests.RpcRequest
	Bid requests.Integer `position:"Query" name:"Bid"`
}

// HeadquartersOverviewResponse is the response struct for api HeadquartersOverview
type HeadquartersOverviewResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateHeadquartersOverviewRequest creates a request to invoke HeadquartersOverview API
func CreateHeadquartersOverviewRequest() (request *HeadquartersOverviewRequest) {
	request = &HeadquartersOverviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersOverview", "cloudwf", "openAPI")
	return
}

// CreateHeadquartersOverviewResponse creates a response to parse from HeadquartersOverview response
func CreateHeadquartersOverviewResponse() (response *HeadquartersOverviewResponse) {
	response = &HeadquartersOverviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
