package market

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

// DescribeApiMetering invokes the market.DescribeApiMetering API synchronously
func (client *Client) DescribeApiMetering(request *DescribeApiMeteringRequest) (response *DescribeApiMeteringResponse, err error) {
	response = CreateDescribeApiMeteringResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiMeteringWithChan invokes the market.DescribeApiMetering API asynchronously
func (client *Client) DescribeApiMeteringWithChan(request *DescribeApiMeteringRequest) (<-chan *DescribeApiMeteringResponse, <-chan error) {
	responseChan := make(chan *DescribeApiMeteringResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiMetering(request)
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

// DescribeApiMeteringWithCallback invokes the market.DescribeApiMetering API asynchronously
func (client *Client) DescribeApiMeteringWithCallback(request *DescribeApiMeteringRequest, callback func(response *DescribeApiMeteringResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiMeteringResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiMetering(request)
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

// DescribeApiMeteringRequest is the request struct for api DescribeApiMetering
type DescribeApiMeteringRequest struct {
	*requests.RpcRequest
	ProductCode string           `position:"Query" name:"productCode"`
	Type        requests.Integer `position:"Query" name:"type"`
	PageNum     requests.Integer `position:"Query" name:"pageNum"`
}

// DescribeApiMeteringResponse is the response struct for api DescribeApiMetering
type DescribeApiMeteringResponse struct {
	*responses.BaseResponse
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	Code       string       `json:"Code" xml:"Code"`
	PageNumber int64        `json:"PageNumber" xml:"PageNumber"`
	Success    bool         `json:"Success" xml:"Success"`
	Count      int64        `json:"Count" xml:"Count"`
	PageSize   int64        `json:"PageSize" xml:"PageSize"`
	Message    string       `json:"Message" xml:"Message"`
	Version    string       `json:"Version" xml:"Version"`
	Fatal      bool         `json:"Fatal" xml:"Fatal"`
	Result     []ResultItem `json:"Result" xml:"Result"`
}

// CreateDescribeApiMeteringRequest creates a request to invoke DescribeApiMetering API
func CreateDescribeApiMeteringRequest() (request *DescribeApiMeteringRequest) {
	request = &DescribeApiMeteringRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Market", "2015-11-01", "DescribeApiMetering", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeApiMeteringResponse creates a response to parse from DescribeApiMetering response
func CreateDescribeApiMeteringResponse() (response *DescribeApiMeteringResponse) {
	response = &DescribeApiMeteringResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
