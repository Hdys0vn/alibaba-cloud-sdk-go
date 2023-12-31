package ivpd

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

// ListUserBuckets invokes the ivpd.ListUserBuckets API synchronously
// api document: https://help.aliyun.com/api/ivpd/listuserbuckets.html
func (client *Client) ListUserBuckets(request *ListUserBucketsRequest) (response *ListUserBucketsResponse, err error) {
	response = CreateListUserBucketsResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserBucketsWithChan invokes the ivpd.ListUserBuckets API asynchronously
// api document: https://help.aliyun.com/api/ivpd/listuserbuckets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUserBucketsWithChan(request *ListUserBucketsRequest) (<-chan *ListUserBucketsResponse, <-chan error) {
	responseChan := make(chan *ListUserBucketsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserBuckets(request)
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

// ListUserBucketsWithCallback invokes the ivpd.ListUserBuckets API asynchronously
// api document: https://help.aliyun.com/api/ivpd/listuserbuckets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUserBucketsWithCallback(request *ListUserBucketsRequest, callback func(response *ListUserBucketsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserBucketsResponse
		var err error
		defer close(result)
		response, err = client.ListUserBuckets(request)
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

// ListUserBucketsRequest is the request struct for api ListUserBuckets
type ListUserBucketsRequest struct {
	*requests.RpcRequest
	Data *[]ListUserBucketsData `position:"Body" name:"Data"  type:"Repeated"`
}

// ListUserBucketsData is a repeated param struct in ListUserBucketsRequest
type ListUserBucketsData struct {
	RegionId string `name:"RegionId"`
}

// ListUserBucketsResponse is the response struct for api ListUserBuckets
type ListUserBucketsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	Data      []string `json:"Data" xml:"Data"`
}

// CreateListUserBucketsRequest creates a request to invoke ListUserBuckets API
func CreateListUserBucketsRequest() (request *ListUserBucketsRequest) {
	request = &ListUserBucketsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivpd", "2019-06-25", "ListUserBuckets", "ivpd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListUserBucketsResponse creates a response to parse from ListUserBuckets response
func CreateListUserBucketsResponse() (response *ListUserBucketsResponse) {
	response = &ListUserBucketsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
