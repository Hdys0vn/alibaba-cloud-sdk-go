package cdn

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

// DescribeCdnOrderCommodityCode invokes the cdn.DescribeCdnOrderCommodityCode API synchronously
func (client *Client) DescribeCdnOrderCommodityCode(request *DescribeCdnOrderCommodityCodeRequest) (response *DescribeCdnOrderCommodityCodeResponse, err error) {
	response = CreateDescribeCdnOrderCommodityCodeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnOrderCommodityCodeWithChan invokes the cdn.DescribeCdnOrderCommodityCode API asynchronously
func (client *Client) DescribeCdnOrderCommodityCodeWithChan(request *DescribeCdnOrderCommodityCodeRequest) (<-chan *DescribeCdnOrderCommodityCodeResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnOrderCommodityCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnOrderCommodityCode(request)
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

// DescribeCdnOrderCommodityCodeWithCallback invokes the cdn.DescribeCdnOrderCommodityCode API asynchronously
func (client *Client) DescribeCdnOrderCommodityCodeWithCallback(request *DescribeCdnOrderCommodityCodeRequest, callback func(response *DescribeCdnOrderCommodityCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnOrderCommodityCodeResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnOrderCommodityCode(request)
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

// DescribeCdnOrderCommodityCodeRequest is the request struct for api DescribeCdnOrderCommodityCode
type DescribeCdnOrderCommodityCodeRequest struct {
	*requests.RpcRequest
	CommodityCode string           `position:"Query" name:"CommodityCode"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeCdnOrderCommodityCodeResponse is the response struct for api DescribeCdnOrderCommodityCode
type DescribeCdnOrderCommodityCodeResponse struct {
	*responses.BaseResponse
	OrderCommodityCode string `json:"OrderCommodityCode" xml:"OrderCommodityCode"`
	RequestId          string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeCdnOrderCommodityCodeRequest creates a request to invoke DescribeCdnOrderCommodityCode API
func CreateDescribeCdnOrderCommodityCodeRequest() (request *DescribeCdnOrderCommodityCodeRequest) {
	request = &DescribeCdnOrderCommodityCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCdnOrderCommodityCode", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCdnOrderCommodityCodeResponse creates a response to parse from DescribeCdnOrderCommodityCode response
func CreateDescribeCdnOrderCommodityCodeResponse() (response *DescribeCdnOrderCommodityCodeResponse) {
	response = &DescribeCdnOrderCommodityCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
