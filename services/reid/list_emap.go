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

// ListEmap invokes the reid.ListEmap API synchronously
// api document: https://help.aliyun.com/api/reid/listemap.html
func (client *Client) ListEmap(request *ListEmapRequest) (response *ListEmapResponse, err error) {
	response = CreateListEmapResponse()
	err = client.DoAction(request, response)
	return
}

// ListEmapWithChan invokes the reid.ListEmap API asynchronously
// api document: https://help.aliyun.com/api/reid/listemap.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListEmapWithChan(request *ListEmapRequest) (<-chan *ListEmapResponse, <-chan error) {
	responseChan := make(chan *ListEmapResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEmap(request)
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

// ListEmapWithCallback invokes the reid.ListEmap API asynchronously
// api document: https://help.aliyun.com/api/reid/listemap.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListEmapWithCallback(request *ListEmapRequest, callback func(response *ListEmapResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEmapResponse
		var err error
		defer close(result)
		response, err = client.ListEmap(request)
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

// ListEmapRequest is the request struct for api ListEmap
type ListEmapRequest struct {
	*requests.RpcRequest
	StoreId requests.Integer `position:"Body" name:"StoreId"`
}

// ListEmapResponse is the response struct for api ListEmap
type ListEmapResponse struct {
	*responses.BaseResponse
	ErrorCode    string    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string    `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string    `json:"RequestId" xml:"RequestId"`
	Success      bool      `json:"Success" xml:"Success"`
	OpenEmaps    OpenEmaps `json:"OpenEmaps" xml:"OpenEmaps"`
}

// CreateListEmapRequest creates a request to invoke ListEmap API
func CreateListEmapRequest() (request *ListEmapRequest) {
	request = &ListEmapRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "ListEmap", "1.1.8.2", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListEmapResponse creates a response to parse from ListEmap response
func CreateListEmapResponse() (response *ListEmapResponse) {
	response = &ListEmapResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
