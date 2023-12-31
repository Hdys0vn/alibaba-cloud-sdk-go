package jarvis

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

// DescribeCpmcPunishList invokes the jarvis.DescribeCpmcPunishList API synchronously
// api document: https://help.aliyun.com/api/jarvis/describecpmcpunishlist.html
func (client *Client) DescribeCpmcPunishList(request *DescribeCpmcPunishListRequest) (response *DescribeCpmcPunishListResponse, err error) {
	response = CreateDescribeCpmcPunishListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCpmcPunishListWithChan invokes the jarvis.DescribeCpmcPunishList API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describecpmcpunishlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCpmcPunishListWithChan(request *DescribeCpmcPunishListRequest) (<-chan *DescribeCpmcPunishListResponse, <-chan error) {
	responseChan := make(chan *DescribeCpmcPunishListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCpmcPunishList(request)
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

// DescribeCpmcPunishListWithCallback invokes the jarvis.DescribeCpmcPunishList API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describecpmcpunishlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCpmcPunishListWithCallback(request *DescribeCpmcPunishListRequest, callback func(response *DescribeCpmcPunishListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCpmcPunishListResponse
		var err error
		defer close(result)
		response, err = client.DescribeCpmcPunishList(request)
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

// DescribeCpmcPunishListRequest is the request struct for api DescribeCpmcPunishList
type DescribeCpmcPunishListRequest struct {
	*requests.RpcRequest
	SrcIP        string           `position:"Query" name:"SrcIP"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	PageSize     requests.Integer `position:"Query" name:"pageSize"`
	CurrentPage  requests.Integer `position:"Query" name:"currentPage"`
	PunishStatus string           `position:"Query" name:"PunishStatus"`
	Lang         string           `position:"Query" name:"Lang"`
	SourceCode   string           `position:"Query" name:"SourceCode"`
}

// DescribeCpmcPunishListResponse is the response struct for api DescribeCpmcPunishList
type DescribeCpmcPunishListResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Module    string   `json:"Module" xml:"Module"`
	PageInfo  PageInfo `json:"PageInfo" xml:"PageInfo"`
	DataList  []Data   `json:"DataList" xml:"DataList"`
}

// CreateDescribeCpmcPunishListRequest creates a request to invoke DescribeCpmcPunishList API
func CreateDescribeCpmcPunishListRequest() (request *DescribeCpmcPunishListRequest) {
	request = &DescribeCpmcPunishListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "DescribeCpmcPunishList", "jarvis", "openAPI")
	return
}

// CreateDescribeCpmcPunishListResponse creates a response to parse from DescribeCpmcPunishList response
func CreateDescribeCpmcPunishListResponse() (response *DescribeCpmcPunishListResponse) {
	response = &DescribeCpmcPunishListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
