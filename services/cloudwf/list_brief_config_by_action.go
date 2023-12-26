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

// ListBriefConfigByAction invokes the cloudwf.ListBriefConfigByAction API synchronously
// api document: https://help.aliyun.com/api/cloudwf/listbriefconfigbyaction.html
func (client *Client) ListBriefConfigByAction(request *ListBriefConfigByActionRequest) (response *ListBriefConfigByActionResponse, err error) {
	response = CreateListBriefConfigByActionResponse()
	err = client.DoAction(request, response)
	return
}

// ListBriefConfigByActionWithChan invokes the cloudwf.ListBriefConfigByAction API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/listbriefconfigbyaction.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListBriefConfigByActionWithChan(request *ListBriefConfigByActionRequest) (<-chan *ListBriefConfigByActionResponse, <-chan error) {
	responseChan := make(chan *ListBriefConfigByActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBriefConfigByAction(request)
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

// ListBriefConfigByActionWithCallback invokes the cloudwf.ListBriefConfigByAction API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/listbriefconfigbyaction.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListBriefConfigByActionWithCallback(request *ListBriefConfigByActionRequest, callback func(response *ListBriefConfigByActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBriefConfigByActionResponse
		var err error
		defer close(result)
		response, err = client.ListBriefConfigByAction(request)
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

// ListBriefConfigByActionRequest is the request struct for api ListBriefConfigByAction
type ListBriefConfigByActionRequest struct {
	*requests.RpcRequest
	AncestorApgroupId requests.Integer `position:"Query" name:"AncestorApgroupId"`
	Limit             requests.Integer `position:"Query" name:"Limit"`
	FuzzySearch       string           `position:"Query" name:"FuzzySearch"`
}

// ListBriefConfigByActionResponse is the response struct for api ListBriefConfigByAction
type ListBriefConfigByActionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateListBriefConfigByActionRequest creates a request to invoke ListBriefConfigByAction API
func CreateListBriefConfigByActionRequest() (request *ListBriefConfigByActionRequest) {
	request = &ListBriefConfigByActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ListBriefConfigByAction", "cloudwf", "openAPI")
	return
}

// CreateListBriefConfigByActionResponse creates a response to parse from ListBriefConfigByAction response
func CreateListBriefConfigByActionResponse() (response *ListBriefConfigByActionResponse) {
	response = &ListBriefConfigByActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
