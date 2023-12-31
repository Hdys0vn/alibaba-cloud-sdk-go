package cms

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

// DescribeTagKeyList invokes the cms.DescribeTagKeyList API synchronously
func (client *Client) DescribeTagKeyList(request *DescribeTagKeyListRequest) (response *DescribeTagKeyListResponse, err error) {
	response = CreateDescribeTagKeyListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTagKeyListWithChan invokes the cms.DescribeTagKeyList API asynchronously
func (client *Client) DescribeTagKeyListWithChan(request *DescribeTagKeyListRequest) (<-chan *DescribeTagKeyListResponse, <-chan error) {
	responseChan := make(chan *DescribeTagKeyListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTagKeyList(request)
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

// DescribeTagKeyListWithCallback invokes the cms.DescribeTagKeyList API asynchronously
func (client *Client) DescribeTagKeyListWithCallback(request *DescribeTagKeyListRequest, callback func(response *DescribeTagKeyListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTagKeyListResponse
		var err error
		defer close(result)
		response, err = client.DescribeTagKeyList(request)
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

// DescribeTagKeyListRequest is the request struct for api DescribeTagKeyList
type DescribeTagKeyListRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeTagKeyListResponse is the response struct for api DescribeTagKeyList
type DescribeTagKeyListResponse struct {
	*responses.BaseResponse
	Code      string                      `json:"Code" xml:"Code"`
	Message   string                      `json:"Message" xml:"Message"`
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	Success   bool                        `json:"Success" xml:"Success"`
	TagKeys   TagKeysInDescribeTagKeyList `json:"TagKeys" xml:"TagKeys"`
}

// CreateDescribeTagKeyListRequest creates a request to invoke DescribeTagKeyList API
func CreateDescribeTagKeyListRequest() (request *DescribeTagKeyListRequest) {
	request = &DescribeTagKeyListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeTagKeyList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTagKeyListResponse creates a response to parse from DescribeTagKeyList response
func CreateDescribeTagKeyListResponse() (response *DescribeTagKeyListResponse) {
	response = &DescribeTagKeyListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
