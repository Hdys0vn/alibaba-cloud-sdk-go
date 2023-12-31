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

// PullActionData invokes the reid.PullActionData API synchronously
// api document: https://help.aliyun.com/api/reid/pullactiondata.html
func (client *Client) PullActionData(request *PullActionDataRequest) (response *PullActionDataResponse, err error) {
	response = CreatePullActionDataResponse()
	err = client.DoAction(request, response)
	return
}

// PullActionDataWithChan invokes the reid.PullActionData API asynchronously
// api document: https://help.aliyun.com/api/reid/pullactiondata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PullActionDataWithChan(request *PullActionDataRequest) (<-chan *PullActionDataResponse, <-chan error) {
	responseChan := make(chan *PullActionDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PullActionData(request)
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

// PullActionDataWithCallback invokes the reid.PullActionData API asynchronously
// api document: https://help.aliyun.com/api/reid/pullactiondata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PullActionDataWithCallback(request *PullActionDataRequest, callback func(response *PullActionDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PullActionDataResponse
		var err error
		defer close(result)
		response, err = client.PullActionData(request)
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

// PullActionDataRequest is the request struct for api PullActionData
type PullActionDataRequest struct {
	*requests.RpcRequest
	StoreId        requests.Integer `position:"Body" name:"StoreId"`
	EndMessageId   requests.Integer `position:"Body" name:"EndMessageId"`
	Limit          requests.Integer `position:"Body" name:"Limit"`
	StartMessageId requests.Integer `position:"Body" name:"StartMessageId"`
}

// PullActionDataResponse is the response struct for api PullActionData
type PullActionDataResponse struct {
	*responses.BaseResponse
	ErrorCode     string                  `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage  string                  `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId     string                  `json:"RequestId" xml:"RequestId"`
	Success       bool                    `json:"Success" xml:"Success"`
	NextMessageId int64                   `json:"NextMessageId" xml:"NextMessageId"`
	Actions       ActionsInPullActionData `json:"Actions" xml:"Actions"`
}

// CreatePullActionDataRequest creates a request to invoke PullActionData API
func CreatePullActionDataRequest() (request *PullActionDataRequest) {
	request = &PullActionDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "PullActionData", "1.1.8.2", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePullActionDataResponse creates a response to parse from PullActionData response
func CreatePullActionDataResponse() (response *PullActionDataResponse) {
	response = &PullActionDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
