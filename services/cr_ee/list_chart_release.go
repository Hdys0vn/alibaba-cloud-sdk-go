package cr_ee

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

// ListChartRelease invokes the cr.ListChartRelease API synchronously
// api document: https://help.aliyun.com/api/cr/listchartrelease.html
func (client *Client) ListChartRelease(request *ListChartReleaseRequest) (response *ListChartReleaseResponse, err error) {
	response = CreateListChartReleaseResponse()
	err = client.DoAction(request, response)
	return
}

// ListChartReleaseWithChan invokes the cr.ListChartRelease API asynchronously
// api document: https://help.aliyun.com/api/cr/listchartrelease.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListChartReleaseWithChan(request *ListChartReleaseRequest) (<-chan *ListChartReleaseResponse, <-chan error) {
	responseChan := make(chan *ListChartReleaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListChartRelease(request)
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

// ListChartReleaseWithCallback invokes the cr.ListChartRelease API asynchronously
// api document: https://help.aliyun.com/api/cr/listchartrelease.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListChartReleaseWithCallback(request *ListChartReleaseRequest, callback func(response *ListChartReleaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListChartReleaseResponse
		var err error
		defer close(result)
		response, err = client.ListChartRelease(request)
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

// ListChartReleaseRequest is the request struct for api ListChartRelease
type ListChartReleaseRequest struct {
	*requests.RpcRequest
	InstanceId        string           `position:"Query" name:"InstanceId"`
	RepoName          string           `position:"Query" name:"RepoName"`
	RepoNamespaceName string           `position:"Query" name:"RepoNamespaceName"`
	PageNo            requests.Integer `position:"Query" name:"PageNo"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	Chart             string           `position:"Query" name:"Chart"`
}

// ListChartReleaseResponse is the response struct for api ListChartRelease
type ListChartReleaseResponse struct {
	*responses.BaseResponse
	ListChartReleaseIsSuccess bool                `json:"IsSuccess" xml:"IsSuccess"`
	Code                      string              `json:"Code" xml:"Code"`
	RequestId                 string              `json:"RequestId" xml:"RequestId"`
	PageNo                    int                 `json:"PageNo" xml:"PageNo"`
	PageSize                  int                 `json:"PageSize" xml:"PageSize"`
	TotalCount                string              `json:"TotalCount" xml:"TotalCount"`
	ChartReleases             []ChartReleasesItem `json:"ChartReleases" xml:"ChartReleases"`
}

// CreateListChartReleaseRequest creates a request to invoke ListChartRelease API
func CreateListChartReleaseRequest() (request *ListChartReleaseRequest) {
	request = &ListChartReleaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "ListChartRelease", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListChartReleaseResponse creates a response to parse from ListChartRelease response
func CreateListChartReleaseResponse() (response *ListChartReleaseResponse) {
	response = &ListChartReleaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
