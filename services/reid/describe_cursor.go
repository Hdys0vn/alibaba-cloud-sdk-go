package reid

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

// DescribeCursor invokes the reid.DescribeCursor API synchronously
// api document: https://help.aliyun.com/api/reid/describecursor.html
func (client *Client) DescribeCursor(request *DescribeCursorRequest) (response *DescribeCursorResponse, err error) {
	response = CreateDescribeCursorResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCursorWithChan invokes the reid.DescribeCursor API asynchronously
// api document: https://help.aliyun.com/api/reid/describecursor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCursorWithChan(request *DescribeCursorRequest) (<-chan *DescribeCursorResponse, <-chan error) {
	responseChan := make(chan *DescribeCursorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCursor(request)
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

// DescribeCursorWithCallback invokes the reid.DescribeCursor API asynchronously
// api document: https://help.aliyun.com/api/reid/describecursor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCursorWithCallback(request *DescribeCursorRequest, callback func(response *DescribeCursorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCursorResponse
		var err error
		defer close(result)
		response, err = client.DescribeCursor(request)
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

// DescribeCursorRequest is the request struct for api DescribeCursor
type DescribeCursorRequest struct {
	*requests.RpcRequest
	StoreId requests.Integer `position:"Body" name:"StoreId"`
	Time    string           `position:"Body" name:"Time"`
}

// DescribeCursorResponse is the response struct for api DescribeCursor
type DescribeCursorResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Cursor       string `json:"Cursor" xml:"Cursor"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateDescribeCursorRequest creates a request to invoke DescribeCursor API
func CreateDescribeCursorRequest() (request *DescribeCursorRequest) {
	request = &DescribeCursorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "DescribeCursor", "1.1.8.2", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCursorResponse creates a response to parse from DescribeCursor response
func CreateDescribeCursorResponse() (response *DescribeCursorResponse) {
	response = &DescribeCursorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
