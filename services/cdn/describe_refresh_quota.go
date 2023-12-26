package cdn

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

// DescribeRefreshQuota invokes the cdn.DescribeRefreshQuota API synchronously
func (client *Client) DescribeRefreshQuota(request *DescribeRefreshQuotaRequest) (response *DescribeRefreshQuotaResponse, err error) {
	response = CreateDescribeRefreshQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRefreshQuotaWithChan invokes the cdn.DescribeRefreshQuota API asynchronously
func (client *Client) DescribeRefreshQuotaWithChan(request *DescribeRefreshQuotaRequest) (<-chan *DescribeRefreshQuotaResponse, <-chan error) {
	responseChan := make(chan *DescribeRefreshQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRefreshQuota(request)
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

// DescribeRefreshQuotaWithCallback invokes the cdn.DescribeRefreshQuota API asynchronously
func (client *Client) DescribeRefreshQuotaWithCallback(request *DescribeRefreshQuotaRequest, callback func(response *DescribeRefreshQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRefreshQuotaResponse
		var err error
		defer close(result)
		response, err = client.DescribeRefreshQuota(request)
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

// DescribeRefreshQuotaRequest is the request struct for api DescribeRefreshQuota
type DescribeRefreshQuotaRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeRefreshQuotaResponse is the response struct for api DescribeRefreshQuota
type DescribeRefreshQuotaResponse struct {
	*responses.BaseResponse
	PreloadRemain      string `json:"PreloadRemain" xml:"PreloadRemain"`
	BlockRemain        string `json:"BlockRemain" xml:"BlockRemain"`
	RegexRemain        string `json:"RegexRemain" xml:"RegexRemain"`
	UrlRemain          string `json:"UrlRemain" xml:"UrlRemain"`
	DirRemain          string `json:"DirRemain" xml:"DirRemain"`
	UrlQuota           string `json:"UrlQuota" xml:"UrlQuota"`
	BlockQuota         string `json:"BlockQuota" xml:"BlockQuota"`
	RequestId          string `json:"RequestId" xml:"RequestId"`
	DirQuota           string `json:"DirQuota" xml:"DirQuota"`
	PreloadEdgeQuota   string `json:"PreloadEdgeQuota" xml:"PreloadEdgeQuota"`
	PreloadEdgeRemain  string `json:"PreloadEdgeRemain" xml:"PreloadEdgeRemain"`
	PreloadQuota       string `json:"PreloadQuota" xml:"PreloadQuota"`
	RegexQuota         string `json:"RegexQuota" xml:"RegexQuota"`
	IgnoreParamsQuota  string `json:"IgnoreParamsQuota" xml:"IgnoreParamsQuota"`
	IgnoreParamsRemain string `json:"IgnoreParamsRemain" xml:"IgnoreParamsRemain"`
}

// CreateDescribeRefreshQuotaRequest creates a request to invoke DescribeRefreshQuota API
func CreateDescribeRefreshQuotaRequest() (request *DescribeRefreshQuotaRequest) {
	request = &DescribeRefreshQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeRefreshQuota", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRefreshQuotaResponse creates a response to parse from DescribeRefreshQuota response
func CreateDescribeRefreshQuotaResponse() (response *DescribeRefreshQuotaResponse) {
	response = &DescribeRefreshQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
