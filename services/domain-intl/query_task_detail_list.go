package domain_intl

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

// QueryTaskDetailList invokes the domain_intl.QueryTaskDetailList API synchronously
// api document: https://help.aliyun.com/api/domain-intl/querytaskdetaillist.html
func (client *Client) QueryTaskDetailList(request *QueryTaskDetailListRequest) (response *QueryTaskDetailListResponse, err error) {
	response = CreateQueryTaskDetailListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTaskDetailListWithChan invokes the domain_intl.QueryTaskDetailList API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/querytaskdetaillist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTaskDetailListWithChan(request *QueryTaskDetailListRequest) (<-chan *QueryTaskDetailListResponse, <-chan error) {
	responseChan := make(chan *QueryTaskDetailListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTaskDetailList(request)
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

// QueryTaskDetailListWithCallback invokes the domain_intl.QueryTaskDetailList API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/querytaskdetaillist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTaskDetailListWithCallback(request *QueryTaskDetailListRequest, callback func(response *QueryTaskDetailListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTaskDetailListResponse
		var err error
		defer close(result)
		response, err = client.QueryTaskDetailList(request)
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

// QueryTaskDetailListRequest is the request struct for api QueryTaskDetailList
type QueryTaskDetailListRequest struct {
	*requests.RpcRequest
	TaskStatus   requests.Integer `position:"Query" name:"TaskStatus"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	TaskNo       string           `position:"Query" name:"TaskNo"`
	DomainName   string           `position:"Query" name:"DomainName"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	PageNum      requests.Integer `position:"Query" name:"PageNum"`
}

// QueryTaskDetailListResponse is the response struct for api QueryTaskDetailList
type QueryTaskDetailListResponse struct {
	*responses.BaseResponse
	RequestId      string                    `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int                       `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int                       `json:"CurrentPageNum" xml:"CurrentPageNum"`
	TotalPageNum   int                       `json:"TotalPageNum" xml:"TotalPageNum"`
	PageSize       int                       `json:"PageSize" xml:"PageSize"`
	PrePage        bool                      `json:"PrePage" xml:"PrePage"`
	NextPage       bool                      `json:"NextPage" xml:"NextPage"`
	Data           DataInQueryTaskDetailList `json:"Data" xml:"Data"`
}

// CreateQueryTaskDetailListRequest creates a request to invoke QueryTaskDetailList API
func CreateQueryTaskDetailListRequest() (request *QueryTaskDetailListRequest) {
	request = &QueryTaskDetailListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryTaskDetailList", "domain", "openAPI")
	return
}

// CreateQueryTaskDetailListResponse creates a response to parse from QueryTaskDetailList response
func CreateQueryTaskDetailListResponse() (response *QueryTaskDetailListResponse) {
	response = &QueryTaskDetailListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
