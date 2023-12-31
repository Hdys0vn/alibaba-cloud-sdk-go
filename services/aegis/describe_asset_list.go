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

// DescribeAssetList invokes the aegis.DescribeAssetList API synchronously
// api document: https://help.aliyun.com/api/aegis/describeassetlist.html
func (client *Client) DescribeAssetList(request *DescribeAssetListRequest) (response *DescribeAssetListResponse, err error) {
	response = CreateDescribeAssetListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAssetListWithChan invokes the aegis.DescribeAssetList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeassetlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAssetListWithChan(request *DescribeAssetListRequest) (<-chan *DescribeAssetListResponse, <-chan error) {
	responseChan := make(chan *DescribeAssetListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAssetList(request)
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

// DescribeAssetListWithCallback invokes the aegis.DescribeAssetList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeassetlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAssetListWithCallback(request *DescribeAssetListRequest, callback func(response *DescribeAssetListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAssetListResponse
		var err error
		defer close(result)
		response, err = client.DescribeAssetList(request)
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

// DescribeAssetListRequest is the request struct for api DescribeAssetList
type DescribeAssetListRequest struct {
	*requests.RpcRequest
	SourceIp         string           `position:"Query" name:"SourceIp"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage      requests.Integer `position:"Query" name:"CurrentPage"`
	Lang             string           `position:"Query" name:"Lang"`
	FilterConditions string           `position:"Query" name:"FilterConditions"`
}

// DescribeAssetListResponse is the response struct for api DescribeAssetList
type DescribeAssetListResponse struct {
	*responses.BaseResponse
	RequestId   string  `json:"RequestId" xml:"RequestId"`
	PageSize    int     `json:"PageSize" xml:"PageSize"`
	CurrentPage int     `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int     `json:"TotalCount" xml:"TotalCount"`
	AssetList   []Asset `json:"AssetList" xml:"AssetList"`
}

// CreateDescribeAssetListRequest creates a request to invoke DescribeAssetList API
func CreateDescribeAssetListRequest() (request *DescribeAssetListRequest) {
	request = &DescribeAssetListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeAssetList", "vipaegis", "openAPI")
	return
}

// CreateDescribeAssetListResponse creates a response to parse from DescribeAssetList response
func CreateDescribeAssetListResponse() (response *DescribeAssetListResponse) {
	response = &DescribeAssetListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
