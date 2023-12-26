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

// GetSidsAndGids4Bid invokes the cloudwf.GetSidsAndGids4Bid API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getsidsandgids4bid.html
func (client *Client) GetSidsAndGids4Bid(request *GetSidsAndGids4BidRequest) (response *GetSidsAndGids4BidResponse, err error) {
	response = CreateGetSidsAndGids4BidResponse()
	err = client.DoAction(request, response)
	return
}

// GetSidsAndGids4BidWithChan invokes the cloudwf.GetSidsAndGids4Bid API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getsidsandgids4bid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSidsAndGids4BidWithChan(request *GetSidsAndGids4BidRequest) (<-chan *GetSidsAndGids4BidResponse, <-chan error) {
	responseChan := make(chan *GetSidsAndGids4BidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSidsAndGids4Bid(request)
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

// GetSidsAndGids4BidWithCallback invokes the cloudwf.GetSidsAndGids4Bid API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getsidsandgids4bid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSidsAndGids4BidWithCallback(request *GetSidsAndGids4BidRequest, callback func(response *GetSidsAndGids4BidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSidsAndGids4BidResponse
		var err error
		defer close(result)
		response, err = client.GetSidsAndGids4Bid(request)
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

// GetSidsAndGids4BidRequest is the request struct for api GetSidsAndGids4Bid
type GetSidsAndGids4BidRequest struct {
	*requests.RpcRequest
	QueryType requests.Integer `position:"Query" name:"QueryType"`
	QueryId   requests.Integer `position:"Query" name:"QueryId"`
}

// GetSidsAndGids4BidResponse is the response struct for api GetSidsAndGids4Bid
type GetSidsAndGids4BidResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetSidsAndGids4BidRequest creates a request to invoke GetSidsAndGids4Bid API
func CreateGetSidsAndGids4BidRequest() (request *GetSidsAndGids4BidRequest) {
	request = &GetSidsAndGids4BidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetSidsAndGids4Bid", "cloudwf", "openAPI")
	return
}

// CreateGetSidsAndGids4BidResponse creates a response to parse from GetSidsAndGids4Bid response
func CreateGetSidsAndGids4BidResponse() (response *GetSidsAndGids4BidResponse) {
	response = &GetSidsAndGids4BidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
