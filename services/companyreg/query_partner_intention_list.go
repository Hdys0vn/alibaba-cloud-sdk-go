package companyreg

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

// QueryPartnerIntentionList invokes the companyreg.QueryPartnerIntentionList API synchronously
func (client *Client) QueryPartnerIntentionList(request *QueryPartnerIntentionListRequest) (response *QueryPartnerIntentionListResponse, err error) {
	response = CreateQueryPartnerIntentionListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPartnerIntentionListWithChan invokes the companyreg.QueryPartnerIntentionList API asynchronously
func (client *Client) QueryPartnerIntentionListWithChan(request *QueryPartnerIntentionListRequest) (<-chan *QueryPartnerIntentionListResponse, <-chan error) {
	responseChan := make(chan *QueryPartnerIntentionListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPartnerIntentionList(request)
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

// QueryPartnerIntentionListWithCallback invokes the companyreg.QueryPartnerIntentionList API asynchronously
func (client *Client) QueryPartnerIntentionListWithCallback(request *QueryPartnerIntentionListRequest, callback func(response *QueryPartnerIntentionListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPartnerIntentionListResponse
		var err error
		defer close(result)
		response, err = client.QueryPartnerIntentionList(request)
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

// QueryPartnerIntentionListRequest is the request struct for api QueryPartnerIntentionList
type QueryPartnerIntentionListRequest struct {
	*requests.RpcRequest
	BizType  string           `position:"Query" name:"BizType"`
	BizId    string           `position:"Query" name:"BizId"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	PageNum  requests.Integer `position:"Query" name:"PageNum"`
}

// QueryPartnerIntentionListResponse is the response struct for api QueryPartnerIntentionList
type QueryPartnerIntentionListResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	CurrentPageNum int64       `json:"CurrentPageNum" xml:"CurrentPageNum"`
	TotalPageNum   int64       `json:"TotalPageNum" xml:"TotalPageNum"`
	PageSize       int64       `json:"PageSize" xml:"PageSize"`
	TotalItemNum   int64       `json:"TotalItemNum" xml:"TotalItemNum"`
	Data           []Intention `json:"Data" xml:"Data"`
}

// CreateQueryPartnerIntentionListRequest creates a request to invoke QueryPartnerIntentionList API
func CreateQueryPartnerIntentionListRequest() (request *QueryPartnerIntentionListRequest) {
	request = &QueryPartnerIntentionListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "QueryPartnerIntentionList", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryPartnerIntentionListResponse creates a response to parse from QueryPartnerIntentionList response
func CreateQueryPartnerIntentionListResponse() (response *QueryPartnerIntentionListResponse) {
	response = &QueryPartnerIntentionListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
