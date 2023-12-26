package sgw

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

// DescribePayAsYouGoPrice invokes the sgw.DescribePayAsYouGoPrice API synchronously
func (client *Client) DescribePayAsYouGoPrice(request *DescribePayAsYouGoPriceRequest) (response *DescribePayAsYouGoPriceResponse, err error) {
	response = CreateDescribePayAsYouGoPriceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePayAsYouGoPriceWithChan invokes the sgw.DescribePayAsYouGoPrice API asynchronously
func (client *Client) DescribePayAsYouGoPriceWithChan(request *DescribePayAsYouGoPriceRequest) (<-chan *DescribePayAsYouGoPriceResponse, <-chan error) {
	responseChan := make(chan *DescribePayAsYouGoPriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePayAsYouGoPrice(request)
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

// DescribePayAsYouGoPriceWithCallback invokes the sgw.DescribePayAsYouGoPrice API asynchronously
func (client *Client) DescribePayAsYouGoPriceWithCallback(request *DescribePayAsYouGoPriceRequest, callback func(response *DescribePayAsYouGoPriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePayAsYouGoPriceResponse
		var err error
		defer close(result)
		response, err = client.DescribePayAsYouGoPrice(request)
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

// DescribePayAsYouGoPriceRequest is the request struct for api DescribePayAsYouGoPrice
type DescribePayAsYouGoPriceRequest struct {
	*requests.RpcRequest
	GatewayClass  string `position:"Query" name:"GatewayClass"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DescribePayAsYouGoPriceResponse is the response struct for api DescribePayAsYouGoPrice
type DescribePayAsYouGoPriceResponse struct {
	*responses.BaseResponse
	RequestId                     string  `json:"RequestId" xml:"RequestId"`
	Success                       bool    `json:"Success" xml:"Success"`
	Code                          string  `json:"Code" xml:"Code"`
	Message                       string  `json:"Message" xml:"Message"`
	Currency                      string  `json:"Currency" xml:"Currency"`
	GatewayClassPrice             float64 `json:"GatewayClassPrice" xml:"GatewayClassPrice"`
	CacheCloudEfficiencySizePrice float64 `json:"CacheCloudEfficiencySizePrice" xml:"CacheCloudEfficiencySizePrice"`
	CacheCloudSSDSizePrice        float64 `json:"CacheCloudSSDSizePrice" xml:"CacheCloudSSDSizePrice"`
}

// CreateDescribePayAsYouGoPriceRequest creates a request to invoke DescribePayAsYouGoPrice API
func CreateDescribePayAsYouGoPriceRequest() (request *DescribePayAsYouGoPriceRequest) {
	request = &DescribePayAsYouGoPriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribePayAsYouGoPrice", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePayAsYouGoPriceResponse creates a response to parse from DescribePayAsYouGoPrice response
func CreateDescribePayAsYouGoPriceResponse() (response *DescribePayAsYouGoPriceResponse) {
	response = &DescribePayAsYouGoPriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
