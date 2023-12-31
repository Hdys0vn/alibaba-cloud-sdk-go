package cusanalytic_sc_online

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

// DescribeActionData invokes the cusanalytic_sc_online.DescribeActionData API synchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/describeactiondata.html
func (client *Client) DescribeActionData(request *DescribeActionDataRequest) (response *DescribeActionDataResponse, err error) {
	response = CreateDescribeActionDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeActionDataWithChan invokes the cusanalytic_sc_online.DescribeActionData API asynchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/describeactiondata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeActionDataWithChan(request *DescribeActionDataRequest) (<-chan *DescribeActionDataResponse, <-chan error) {
	responseChan := make(chan *DescribeActionDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeActionData(request)
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

// DescribeActionDataWithCallback invokes the cusanalytic_sc_online.DescribeActionData API asynchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/describeactiondata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeActionDataWithCallback(request *DescribeActionDataRequest, callback func(response *DescribeActionDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeActionDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeActionData(request)
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

// DescribeActionDataRequest is the request struct for api DescribeActionData
type DescribeActionDataRequest struct {
	*requests.RpcRequest
	TsEnd     requests.Integer `position:"Body" name:"TsEnd"`
	StoreId   string           `position:"Body" name:"StoreId"`
	PageLimit requests.Integer `position:"Body" name:"PageLimit"`
	PageNo    requests.Integer `position:"Body" name:"PageNo"`
	TsStart   requests.Integer `position:"Body" name:"TsStart"`
}

// DescribeActionDataResponse is the response struct for api DescribeActionData
type DescribeActionDataResponse struct {
	*responses.BaseResponse
	PageNo                      int             `json:"PageNo" xml:"PageNo"`
	TsStart                     int64           `json:"TsStart" xml:"TsStart"`
	PageLimit                   int             `json:"PageLimit" xml:"PageLimit"`
	PageCount                   int             `json:"PageCount" xml:"PageCount"`
	DescribeActionDataIsSuccess bool            `json:"IsSuccess" xml:"IsSuccess"`
	ErrorMsg                    string          `json:"ErrorMsg" xml:"ErrorMsg"`
	StoreId                     string          `json:"StoreId" xml:"StoreId"`
	TsEnd                       int64           `json:"TsEnd" xml:"TsEnd"`
	ActionsMsgItems             ActionsMsgItems `json:"ActionsMsgItems" xml:"ActionsMsgItems"`
}

// CreateDescribeActionDataRequest creates a request to invoke DescribeActionData API
func CreateDescribeActionDataRequest() (request *DescribeActionDataRequest) {
	request = &DescribeActionDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cusanalytic_sc_online", "2019-05-24", "DescribeActionData", "", "")
	return
}

// CreateDescribeActionDataResponse creates a response to parse from DescribeActionData response
func CreateDescribeActionDataResponse() (response *DescribeActionDataResponse) {
	response = &DescribeActionDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
