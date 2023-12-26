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

// DescribeBlockedRegions invokes the cdn.DescribeBlockedRegions API synchronously
func (client *Client) DescribeBlockedRegions(request *DescribeBlockedRegionsRequest) (response *DescribeBlockedRegionsResponse, err error) {
	response = CreateDescribeBlockedRegionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBlockedRegionsWithChan invokes the cdn.DescribeBlockedRegions API asynchronously
func (client *Client) DescribeBlockedRegionsWithChan(request *DescribeBlockedRegionsRequest) (<-chan *DescribeBlockedRegionsResponse, <-chan error) {
	responseChan := make(chan *DescribeBlockedRegionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBlockedRegions(request)
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

// DescribeBlockedRegionsWithCallback invokes the cdn.DescribeBlockedRegions API asynchronously
func (client *Client) DescribeBlockedRegionsWithCallback(request *DescribeBlockedRegionsRequest, callback func(response *DescribeBlockedRegionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBlockedRegionsResponse
		var err error
		defer close(result)
		response, err = client.DescribeBlockedRegions(request)
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

// DescribeBlockedRegionsRequest is the request struct for api DescribeBlockedRegions
type DescribeBlockedRegionsRequest struct {
	*requests.RpcRequest
	Language string `position:"Query" name:"Language"`
}

// DescribeBlockedRegionsResponse is the response struct for api DescribeBlockedRegions
type DescribeBlockedRegionsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	InfoList  InfoList `json:"InfoList" xml:"InfoList"`
}

// CreateDescribeBlockedRegionsRequest creates a request to invoke DescribeBlockedRegions API
func CreateDescribeBlockedRegionsRequest() (request *DescribeBlockedRegionsRequest) {
	request = &DescribeBlockedRegionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeBlockedRegions", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeBlockedRegionsResponse creates a response to parse from DescribeBlockedRegions response
func CreateDescribeBlockedRegionsResponse() (response *DescribeBlockedRegionsResponse) {
	response = &DescribeBlockedRegionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
