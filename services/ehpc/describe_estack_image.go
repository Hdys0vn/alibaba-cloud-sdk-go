package ehpc

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

// DescribeEstackImage invokes the ehpc.DescribeEstackImage API synchronously
func (client *Client) DescribeEstackImage(request *DescribeEstackImageRequest) (response *DescribeEstackImageResponse, err error) {
	response = CreateDescribeEstackImageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEstackImageWithChan invokes the ehpc.DescribeEstackImage API asynchronously
func (client *Client) DescribeEstackImageWithChan(request *DescribeEstackImageRequest) (<-chan *DescribeEstackImageResponse, <-chan error) {
	responseChan := make(chan *DescribeEstackImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEstackImage(request)
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

// DescribeEstackImageWithCallback invokes the ehpc.DescribeEstackImage API asynchronously
func (client *Client) DescribeEstackImageWithCallback(request *DescribeEstackImageRequest, callback func(response *DescribeEstackImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEstackImageResponse
		var err error
		defer close(result)
		response, err = client.DescribeEstackImage(request)
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

// DescribeEstackImageRequest is the request struct for api DescribeEstackImage
type DescribeEstackImageRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeEstackImageResponse is the response struct for api DescribeEstackImage
type DescribeEstackImageResponse struct {
	*responses.BaseResponse
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	ImageList  ImageList `json:"ImageList" xml:"ImageList"`
}

// CreateDescribeEstackImageRequest creates a request to invoke DescribeEstackImage API
func CreateDescribeEstackImageRequest() (request *DescribeEstackImageRequest) {
	request = &DescribeEstackImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "DescribeEstackImage", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeEstackImageResponse creates a response to parse from DescribeEstackImage response
func CreateDescribeEstackImageResponse() (response *DescribeEstackImageResponse) {
	response = &DescribeEstackImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}