package trademark

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

// QueryTradeIntentionUserList invokes the trademark.QueryTradeIntentionUserList API synchronously
// api document: https://help.aliyun.com/api/trademark/querytradeintentionuserlist.html
func (client *Client) QueryTradeIntentionUserList(request *QueryTradeIntentionUserListRequest) (response *QueryTradeIntentionUserListResponse, err error) {
	response = CreateQueryTradeIntentionUserListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTradeIntentionUserListWithChan invokes the trademark.QueryTradeIntentionUserList API asynchronously
// api document: https://help.aliyun.com/api/trademark/querytradeintentionuserlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTradeIntentionUserListWithChan(request *QueryTradeIntentionUserListRequest) (<-chan *QueryTradeIntentionUserListResponse, <-chan error) {
	responseChan := make(chan *QueryTradeIntentionUserListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTradeIntentionUserList(request)
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

// QueryTradeIntentionUserListWithCallback invokes the trademark.QueryTradeIntentionUserList API asynchronously
// api document: https://help.aliyun.com/api/trademark/querytradeintentionuserlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTradeIntentionUserListWithCallback(request *QueryTradeIntentionUserListRequest, callback func(response *QueryTradeIntentionUserListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTradeIntentionUserListResponse
		var err error
		defer close(result)
		response, err = client.QueryTradeIntentionUserList(request)
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

// QueryTradeIntentionUserListRequest is the request struct for api QueryTradeIntentionUserList
type QueryTradeIntentionUserListRequest struct {
	*requests.RpcRequest
	End   requests.Integer `position:"Query" name:"End"`
	Begin requests.Integer `position:"Query" name:"Begin"`
}

// QueryTradeIntentionUserListResponse is the response struct for api QueryTradeIntentionUserList
type QueryTradeIntentionUserListResponse struct {
	*responses.BaseResponse
	RequestId      string                            `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int                               `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int                               `json:"CurrentPageNum" xml:"CurrentPageNum"`
	PageSize       int                               `json:"PageSize" xml:"PageSize"`
	TotalPageNum   int                               `json:"TotalPageNum" xml:"TotalPageNum"`
	Data           DataInQueryTradeIntentionUserList `json:"Data" xml:"Data"`
}

// CreateQueryTradeIntentionUserListRequest creates a request to invoke QueryTradeIntentionUserList API
func CreateQueryTradeIntentionUserListRequest() (request *QueryTradeIntentionUserListRequest) {
	request = &QueryTradeIntentionUserListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "QueryTradeIntentionUserList", "trademark", "openAPI")
	return
}

// CreateQueryTradeIntentionUserListResponse creates a response to parse from QueryTradeIntentionUserList response
func CreateQueryTradeIntentionUserListResponse() (response *QueryTradeIntentionUserListResponse) {
	response = &QueryTradeIntentionUserListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
