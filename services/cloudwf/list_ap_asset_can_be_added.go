package cloudwf

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

// ListApAssetCanBeAdded invokes the cloudwf.ListApAssetCanBeAdded API synchronously
// api document: https://help.aliyun.com/api/cloudwf/listapassetcanbeadded.html
func (client *Client) ListApAssetCanBeAdded(request *ListApAssetCanBeAddedRequest) (response *ListApAssetCanBeAddedResponse, err error) {
	response = CreateListApAssetCanBeAddedResponse()
	err = client.DoAction(request, response)
	return
}

// ListApAssetCanBeAddedWithChan invokes the cloudwf.ListApAssetCanBeAdded API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/listapassetcanbeadded.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListApAssetCanBeAddedWithChan(request *ListApAssetCanBeAddedRequest) (<-chan *ListApAssetCanBeAddedResponse, <-chan error) {
	responseChan := make(chan *ListApAssetCanBeAddedResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListApAssetCanBeAdded(request)
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

// ListApAssetCanBeAddedWithCallback invokes the cloudwf.ListApAssetCanBeAdded API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/listapassetcanbeadded.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListApAssetCanBeAddedWithCallback(request *ListApAssetCanBeAddedRequest, callback func(response *ListApAssetCanBeAddedResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListApAssetCanBeAddedResponse
		var err error
		defer close(result)
		response, err = client.ListApAssetCanBeAdded(request)
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

// ListApAssetCanBeAddedRequest is the request struct for api ListApAssetCanBeAdded
type ListApAssetCanBeAddedRequest struct {
	*requests.RpcRequest
	SearchName  string           `position:"Query" name:"SearchName"`
	ApgroupId   requests.Integer `position:"Query" name:"ApgroupId"`
	Length      requests.Integer `position:"Query" name:"Length"`
	PageIndex   requests.Integer `position:"Query" name:"PageIndex"`
	SearchMac   string           `position:"Query" name:"SearchMac"`
	SearchModel string           `position:"Query" name:"SearchModel"`
}

// ListApAssetCanBeAddedResponse is the response struct for api ListApAssetCanBeAdded
type ListApAssetCanBeAddedResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateListApAssetCanBeAddedRequest creates a request to invoke ListApAssetCanBeAdded API
func CreateListApAssetCanBeAddedRequest() (request *ListApAssetCanBeAddedRequest) {
	request = &ListApAssetCanBeAddedRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ListApAssetCanBeAdded", "cloudwf", "openAPI")
	return
}

// CreateListApAssetCanBeAddedResponse creates a response to parse from ListApAssetCanBeAdded response
func CreateListApAssetCanBeAddedResponse() (response *ListApAssetCanBeAddedResponse) {
	response = &ListApAssetCanBeAddedResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
