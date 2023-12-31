package jarvis

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

// DescribeCdnVendor invokes the jarvis.DescribeCdnVendor API synchronously
// api document: https://help.aliyun.com/api/jarvis/describecdnvendor.html
func (client *Client) DescribeCdnVendor(request *DescribeCdnVendorRequest) (response *DescribeCdnVendorResponse, err error) {
	response = CreateDescribeCdnVendorResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnVendorWithChan invokes the jarvis.DescribeCdnVendor API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describecdnvendor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnVendorWithChan(request *DescribeCdnVendorRequest) (<-chan *DescribeCdnVendorResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnVendorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnVendor(request)
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

// DescribeCdnVendorWithCallback invokes the jarvis.DescribeCdnVendor API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describecdnvendor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnVendorWithCallback(request *DescribeCdnVendorRequest, callback func(response *DescribeCdnVendorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnVendorResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnVendor(request)
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

// DescribeCdnVendorRequest is the request struct for api DescribeCdnVendor
type DescribeCdnVendorRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Lang        string           `position:"Query" name:"Lang"`
	SourceCode  string           `position:"Query" name:"SourceCode"`
}

// DescribeCdnVendorResponse is the response struct for api DescribeCdnVendor
type DescribeCdnVendorResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Module    string   `json:"Module" xml:"Module"`
	PageInfo  PageInfo `json:"PageInfo" xml:"PageInfo"`
	DataList  []Data   `json:"DataList" xml:"DataList"`
}

// CreateDescribeCdnVendorRequest creates a request to invoke DescribeCdnVendor API
func CreateDescribeCdnVendorRequest() (request *DescribeCdnVendorRequest) {
	request = &DescribeCdnVendorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "DescribeCdnVendor", "jarvis", "openAPI")
	return
}

// CreateDescribeCdnVendorResponse creates a response to parse from DescribeCdnVendor response
func CreateDescribeCdnVendorResponse() (response *DescribeCdnVendorResponse) {
	response = &DescribeCdnVendorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
