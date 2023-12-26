package vs

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

// ListBuckets invokes the vs.ListBuckets API synchronously
func (client *Client) ListBuckets(request *ListBucketsRequest) (response *ListBucketsResponse, err error) {
	response = CreateListBucketsResponse()
	err = client.DoAction(request, response)
	return
}

// ListBucketsWithChan invokes the vs.ListBuckets API asynchronously
func (client *Client) ListBucketsWithChan(request *ListBucketsRequest) (<-chan *ListBucketsResponse, <-chan error) {
	responseChan := make(chan *ListBucketsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBuckets(request)
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

// ListBucketsWithCallback invokes the vs.ListBuckets API asynchronously
func (client *Client) ListBucketsWithCallback(request *ListBucketsRequest, callback func(response *ListBucketsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBucketsResponse
		var err error
		defer close(result)
		response, err = client.ListBuckets(request)
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

// ListBucketsRequest is the request struct for api ListBuckets
type ListBucketsRequest struct {
	*requests.RpcRequest
	Prefix     string           `position:"Query" name:"Prefix"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Keyword    string           `position:"Query" name:"Keyword"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Marker     string           `position:"Query" name:"Marker"`
}

// ListBucketsResponse is the response struct for api ListBuckets
type ListBucketsResponse struct {
	*responses.BaseResponse
	TotalCount  int64        `json:"TotalCount" xml:"TotalCount"`
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	BucketInfos []BucketInfo `json:"BucketInfos" xml:"BucketInfos"`
}

// CreateListBucketsRequest creates a request to invoke ListBuckets API
func CreateListBucketsRequest() (request *ListBucketsRequest) {
	request = &ListBucketsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "ListBuckets", "", "")
	request.Method = requests.POST
	return
}

// CreateListBucketsResponse creates a response to parse from ListBuckets response
func CreateListBucketsResponse() (response *ListBucketsResponse) {
	response = &ListBucketsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
