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

// DescribeRiskWhiteList invokes the aegis.DescribeRiskWhiteList API synchronously
// api document: https://help.aliyun.com/api/aegis/describeriskwhitelist.html
func (client *Client) DescribeRiskWhiteList(request *DescribeRiskWhiteListRequest) (response *DescribeRiskWhiteListResponse, err error) {
	response = CreateDescribeRiskWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRiskWhiteListWithChan invokes the aegis.DescribeRiskWhiteList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeriskwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRiskWhiteListWithChan(request *DescribeRiskWhiteListRequest) (<-chan *DescribeRiskWhiteListResponse, <-chan error) {
	responseChan := make(chan *DescribeRiskWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRiskWhiteList(request)
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

// DescribeRiskWhiteListWithCallback invokes the aegis.DescribeRiskWhiteList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeriskwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRiskWhiteListWithCallback(request *DescribeRiskWhiteListRequest, callback func(response *DescribeRiskWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRiskWhiteListResponse
		var err error
		defer close(result)
		response, err = client.DescribeRiskWhiteList(request)
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

// DescribeRiskWhiteListRequest is the request struct for api DescribeRiskWhiteList
type DescribeRiskWhiteListRequest struct {
	*requests.RpcRequest
	RiskName    string           `position:"Query" name:"RiskName"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
}

// DescribeRiskWhiteListResponse is the response struct for api DescribeRiskWhiteList
type DescribeRiskWhiteListResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Count       int         `json:"Count" xml:"Count"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	CurrentPage int         `json:"CurrentPage" xml:"CurrentPage"`
	WhiteLists  []WhiteList `json:"WhiteLists" xml:"WhiteLists"`
}

// CreateDescribeRiskWhiteListRequest creates a request to invoke DescribeRiskWhiteList API
func CreateDescribeRiskWhiteListRequest() (request *DescribeRiskWhiteListRequest) {
	request = &DescribeRiskWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeRiskWhiteList", "vipaegis", "openAPI")
	return
}

// CreateDescribeRiskWhiteListResponse creates a response to parse from DescribeRiskWhiteList response
func CreateDescribeRiskWhiteListResponse() (response *DescribeRiskWhiteListResponse) {
	response = &DescribeRiskWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
