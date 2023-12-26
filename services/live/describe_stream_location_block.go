package live

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

// DescribeStreamLocationBlock invokes the live.DescribeStreamLocationBlock API synchronously
func (client *Client) DescribeStreamLocationBlock(request *DescribeStreamLocationBlockRequest) (response *DescribeStreamLocationBlockResponse, err error) {
	response = CreateDescribeStreamLocationBlockResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStreamLocationBlockWithChan invokes the live.DescribeStreamLocationBlock API asynchronously
func (client *Client) DescribeStreamLocationBlockWithChan(request *DescribeStreamLocationBlockRequest) (<-chan *DescribeStreamLocationBlockResponse, <-chan error) {
	responseChan := make(chan *DescribeStreamLocationBlockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStreamLocationBlock(request)
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

// DescribeStreamLocationBlockWithCallback invokes the live.DescribeStreamLocationBlock API asynchronously
func (client *Client) DescribeStreamLocationBlockWithCallback(request *DescribeStreamLocationBlockRequest, callback func(response *DescribeStreamLocationBlockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStreamLocationBlockResponse
		var err error
		defer close(result)
		response, err = client.DescribeStreamLocationBlock(request)
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

// DescribeStreamLocationBlockRequest is the request struct for api DescribeStreamLocationBlock
type DescribeStreamLocationBlockRequest struct {
	*requests.RpcRequest
	BlockType  string           `position:"Query" name:"BlockType"`
	PageNum    requests.Integer `position:"Query" name:"PageNum"`
	AppName    string           `position:"Query" name:"AppName"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	StreamName string           `position:"Query" name:"StreamName"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeStreamLocationBlockResponse is the response struct for api DescribeStreamLocationBlock
type DescribeStreamLocationBlockResponse struct {
	*responses.BaseResponse
	TotalPage       int             `json:"TotalPage" xml:"TotalPage"`
	PageNum         int             `json:"PageNum" xml:"PageNum"`
	PageSize        int             `json:"PageSize" xml:"PageSize"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Count           int             `json:"Count" xml:"Count"`
	StreamBlockList StreamBlockList `json:"StreamBlockList" xml:"StreamBlockList"`
}

// CreateDescribeStreamLocationBlockRequest creates a request to invoke DescribeStreamLocationBlock API
func CreateDescribeStreamLocationBlockRequest() (request *DescribeStreamLocationBlockRequest) {
	request = &DescribeStreamLocationBlockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeStreamLocationBlock", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeStreamLocationBlockResponse creates a response to parse from DescribeStreamLocationBlock response
func CreateDescribeStreamLocationBlockResponse() (response *DescribeStreamLocationBlockResponse) {
	response = &DescribeStreamLocationBlockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
