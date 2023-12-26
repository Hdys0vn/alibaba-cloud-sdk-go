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

// DescribeVsStreamsPublishList invokes the vs.DescribeVsStreamsPublishList API synchronously
func (client *Client) DescribeVsStreamsPublishList(request *DescribeVsStreamsPublishListRequest) (response *DescribeVsStreamsPublishListResponse, err error) {
	response = CreateDescribeVsStreamsPublishListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVsStreamsPublishListWithChan invokes the vs.DescribeVsStreamsPublishList API asynchronously
func (client *Client) DescribeVsStreamsPublishListWithChan(request *DescribeVsStreamsPublishListRequest) (<-chan *DescribeVsStreamsPublishListResponse, <-chan error) {
	responseChan := make(chan *DescribeVsStreamsPublishListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVsStreamsPublishList(request)
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

// DescribeVsStreamsPublishListWithCallback invokes the vs.DescribeVsStreamsPublishList API asynchronously
func (client *Client) DescribeVsStreamsPublishListWithCallback(request *DescribeVsStreamsPublishListRequest, callback func(response *DescribeVsStreamsPublishListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVsStreamsPublishListResponse
		var err error
		defer close(result)
		response, err = client.DescribeVsStreamsPublishList(request)
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

// DescribeVsStreamsPublishListRequest is the request struct for api DescribeVsStreamsPublishList
type DescribeVsStreamsPublishListRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	AppName    string           `position:"Query" name:"AppName"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	StreamName string           `position:"Query" name:"StreamName"`
	QueryType  string           `position:"Query" name:"QueryType"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	StreamType string           `position:"Query" name:"StreamType"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OrderBy    string           `position:"Query" name:"OrderBy"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVsStreamsPublishListResponse is the response struct for api DescribeVsStreamsPublishList
type DescribeVsStreamsPublishListResponse struct {
	*responses.BaseResponse
	TotalPage   int         `json:"TotalPage" xml:"TotalPage"`
	PageNum     int         `json:"PageNum" xml:"PageNum"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	TotalNum    int         `json:"TotalNum" xml:"TotalNum"`
	PublishInfo PublishInfo `json:"PublishInfo" xml:"PublishInfo"`
}

// CreateDescribeVsStreamsPublishListRequest creates a request to invoke DescribeVsStreamsPublishList API
func CreateDescribeVsStreamsPublishListRequest() (request *DescribeVsStreamsPublishListRequest) {
	request = &DescribeVsStreamsPublishListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeVsStreamsPublishList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeVsStreamsPublishListResponse creates a response to parse from DescribeVsStreamsPublishList response
func CreateDescribeVsStreamsPublishListResponse() (response *DescribeVsStreamsPublishListResponse) {
	response = &DescribeVsStreamsPublishListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
