package resourcecenter

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

// ListMultiAccountTagKeys invokes the resourcecenter.ListMultiAccountTagKeys API synchronously
func (client *Client) ListMultiAccountTagKeys(request *ListMultiAccountTagKeysRequest) (response *ListMultiAccountTagKeysResponse, err error) {
	response = CreateListMultiAccountTagKeysResponse()
	err = client.DoAction(request, response)
	return
}

// ListMultiAccountTagKeysWithChan invokes the resourcecenter.ListMultiAccountTagKeys API asynchronously
func (client *Client) ListMultiAccountTagKeysWithChan(request *ListMultiAccountTagKeysRequest) (<-chan *ListMultiAccountTagKeysResponse, <-chan error) {
	responseChan := make(chan *ListMultiAccountTagKeysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMultiAccountTagKeys(request)
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

// ListMultiAccountTagKeysWithCallback invokes the resourcecenter.ListMultiAccountTagKeys API asynchronously
func (client *Client) ListMultiAccountTagKeysWithCallback(request *ListMultiAccountTagKeysRequest, callback func(response *ListMultiAccountTagKeysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMultiAccountTagKeysResponse
		var err error
		defer close(result)
		response, err = client.ListMultiAccountTagKeys(request)
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

// ListMultiAccountTagKeysRequest is the request struct for api ListMultiAccountTagKeys
type ListMultiAccountTagKeysRequest struct {
	*requests.RpcRequest
	NextToken  string           `position:"Query" name:"NextToken"`
	Scope      string           `position:"Query" name:"Scope"`
	MatchType  string           `position:"Query" name:"MatchType"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
	TagKey     string           `position:"Query" name:"TagKey"`
}

// ListMultiAccountTagKeysResponse is the response struct for api ListMultiAccountTagKeys
type ListMultiAccountTagKeysResponse struct {
	*responses.BaseResponse
	NextToken      string   `json:"NextToken" xml:"NextToken"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	DynamicCode    string   `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string   `json:"DynamicMessage" xml:"DynamicMessage"`
	ErrorCode      string   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string   `json:"ErrorMessage" xml:"ErrorMessage"`
	TagKeys        []string `json:"TagKeys" xml:"TagKeys"`
}

// CreateListMultiAccountTagKeysRequest creates a request to invoke ListMultiAccountTagKeys API
func CreateListMultiAccountTagKeysRequest() (request *ListMultiAccountTagKeysRequest) {
	request = &ListMultiAccountTagKeysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "ListMultiAccountTagKeys", "", "")
	request.Method = requests.POST
	return
}

// CreateListMultiAccountTagKeysResponse creates a response to parse from ListMultiAccountTagKeys response
func CreateListMultiAccountTagKeysResponse() (response *ListMultiAccountTagKeysResponse) {
	response = &ListMultiAccountTagKeysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}