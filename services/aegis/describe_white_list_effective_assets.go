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

// DescribeWhiteListEffectiveAssets invokes the aegis.DescribeWhiteListEffectiveAssets API synchronously
// api document: https://help.aliyun.com/api/aegis/describewhitelisteffectiveassets.html
func (client *Client) DescribeWhiteListEffectiveAssets(request *DescribeWhiteListEffectiveAssetsRequest) (response *DescribeWhiteListEffectiveAssetsResponse, err error) {
	response = CreateDescribeWhiteListEffectiveAssetsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWhiteListEffectiveAssetsWithChan invokes the aegis.DescribeWhiteListEffectiveAssets API asynchronously
// api document: https://help.aliyun.com/api/aegis/describewhitelisteffectiveassets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWhiteListEffectiveAssetsWithChan(request *DescribeWhiteListEffectiveAssetsRequest) (<-chan *DescribeWhiteListEffectiveAssetsResponse, <-chan error) {
	responseChan := make(chan *DescribeWhiteListEffectiveAssetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWhiteListEffectiveAssets(request)
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

// DescribeWhiteListEffectiveAssetsWithCallback invokes the aegis.DescribeWhiteListEffectiveAssets API asynchronously
// api document: https://help.aliyun.com/api/aegis/describewhitelisteffectiveassets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWhiteListEffectiveAssetsWithCallback(request *DescribeWhiteListEffectiveAssetsRequest, callback func(response *DescribeWhiteListEffectiveAssetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWhiteListEffectiveAssetsResponse
		var err error
		defer close(result)
		response, err = client.DescribeWhiteListEffectiveAssets(request)
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

// DescribeWhiteListEffectiveAssetsRequest is the request struct for api DescribeWhiteListEffectiveAssets
type DescribeWhiteListEffectiveAssetsRequest struct {
	*requests.RpcRequest
	SourceIp       string           `position:"Query" name:"SourceIp"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	Remark         string           `position:"Query" name:"Remark"`
	StrategyId     requests.Integer `position:"Query" name:"StrategyId"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
	Lang           string           `position:"Query" name:"Lang"`
	NeedStatistics requests.Integer `position:"Query" name:"NeedStatistics"`
}

// DescribeWhiteListEffectiveAssetsResponse is the response struct for api DescribeWhiteListEffectiveAssets
type DescribeWhiteListEffectiveAssetsResponse struct {
	*responses.BaseResponse
	RequestId   string  `json:"RequestId" xml:"RequestId"`
	Count       int     `json:"Count" xml:"Count"`
	PageSize    int     `json:"PageSize" xml:"PageSize"`
	TotalCount  int     `json:"TotalCount" xml:"TotalCount"`
	CurrentPage int     `json:"CurrentPage" xml:"CurrentPage"`
	Assets      []Asset `json:"Assets" xml:"Assets"`
}

// CreateDescribeWhiteListEffectiveAssetsRequest creates a request to invoke DescribeWhiteListEffectiveAssets API
func CreateDescribeWhiteListEffectiveAssetsRequest() (request *DescribeWhiteListEffectiveAssetsRequest) {
	request = &DescribeWhiteListEffectiveAssetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeWhiteListEffectiveAssets", "vipaegis", "openAPI")
	return
}

// CreateDescribeWhiteListEffectiveAssetsResponse creates a response to parse from DescribeWhiteListEffectiveAssets response
func CreateDescribeWhiteListEffectiveAssetsResponse() (response *DescribeWhiteListEffectiveAssetsResponse) {
	response = &DescribeWhiteListEffectiveAssetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
