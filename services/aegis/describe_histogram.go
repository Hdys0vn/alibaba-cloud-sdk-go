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

// DescribeHistogram invokes the aegis.DescribeHistogram API synchronously
// api document: https://help.aliyun.com/api/aegis/describehistogram.html
func (client *Client) DescribeHistogram(request *DescribeHistogramRequest) (response *DescribeHistogramResponse, err error) {
	response = CreateDescribeHistogramResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHistogramWithChan invokes the aegis.DescribeHistogram API asynchronously
// api document: https://help.aliyun.com/api/aegis/describehistogram.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHistogramWithChan(request *DescribeHistogramRequest) (<-chan *DescribeHistogramResponse, <-chan error) {
	responseChan := make(chan *DescribeHistogramResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHistogram(request)
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

// DescribeHistogramWithCallback invokes the aegis.DescribeHistogram API asynchronously
// api document: https://help.aliyun.com/api/aegis/describehistogram.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHistogramWithCallback(request *DescribeHistogramRequest, callback func(response *DescribeHistogramResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHistogramResponse
		var err error
		defer close(result)
		response, err = client.DescribeHistogram(request)
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

// DescribeHistogramRequest is the request struct for api DescribeHistogram
type DescribeHistogramRequest struct {
	*requests.RpcRequest
	SourceIp  string `position:"Query" name:"SourceIp"`
	Query     string `position:"Query" name:"Query"`
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
}

// DescribeHistogramResponse is the response struct for api DescribeHistogram
type DescribeHistogramResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	Categories []string    `json:"Categories" xml:"Categories"`
	Items      []ItemsItem `json:"Items" xml:"Items"`
}

// CreateDescribeHistogramRequest creates a request to invoke DescribeHistogram API
func CreateDescribeHistogramRequest() (request *DescribeHistogramRequest) {
	request = &DescribeHistogramRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeHistogram", "vipaegis", "openAPI")
	return
}

// CreateDescribeHistogramResponse creates a response to parse from DescribeHistogram response
func CreateDescribeHistogramResponse() (response *DescribeHistogramResponse) {
	response = &DescribeHistogramResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
