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

// DescribeEventLevelCount invokes the aegis.DescribeEventLevelCount API synchronously
// api document: https://help.aliyun.com/api/aegis/describeeventlevelcount.html
func (client *Client) DescribeEventLevelCount(request *DescribeEventLevelCountRequest) (response *DescribeEventLevelCountResponse, err error) {
	response = CreateDescribeEventLevelCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEventLevelCountWithChan invokes the aegis.DescribeEventLevelCount API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeeventlevelcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEventLevelCountWithChan(request *DescribeEventLevelCountRequest) (<-chan *DescribeEventLevelCountResponse, <-chan error) {
	responseChan := make(chan *DescribeEventLevelCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEventLevelCount(request)
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

// DescribeEventLevelCountWithCallback invokes the aegis.DescribeEventLevelCount API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeeventlevelcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEventLevelCountWithCallback(request *DescribeEventLevelCountRequest, callback func(response *DescribeEventLevelCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEventLevelCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeEventLevelCount(request)
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

// DescribeEventLevelCountRequest is the request struct for api DescribeEventLevelCount
type DescribeEventLevelCountRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	From     string `position:"Query" name:"From"`
	Type     string `position:"Query" name:"Type"`
}

// DescribeEventLevelCountResponse is the response struct for api DescribeEventLevelCount
type DescribeEventLevelCountResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	EventLevels EventLevels `json:"EventLevels" xml:"EventLevels"`
}

// CreateDescribeEventLevelCountRequest creates a request to invoke DescribeEventLevelCount API
func CreateDescribeEventLevelCountRequest() (request *DescribeEventLevelCountRequest) {
	request = &DescribeEventLevelCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeEventLevelCount", "vipaegis", "openAPI")
	return
}

// CreateDescribeEventLevelCountResponse creates a response to parse from DescribeEventLevelCount response
func CreateDescribeEventLevelCountResponse() (response *DescribeEventLevelCountResponse) {
	response = &DescribeEventLevelCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
