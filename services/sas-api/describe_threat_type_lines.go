package sas_api

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

// DescribeThreatTypeLines invokes the sas_api.DescribeThreatTypeLines API synchronously
// api document: https://help.aliyun.com/api/sas-api/describethreattypelines.html
func (client *Client) DescribeThreatTypeLines(request *DescribeThreatTypeLinesRequest) (response *DescribeThreatTypeLinesResponse, err error) {
	response = CreateDescribeThreatTypeLinesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeThreatTypeLinesWithChan invokes the sas_api.DescribeThreatTypeLines API asynchronously
// api document: https://help.aliyun.com/api/sas-api/describethreattypelines.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeThreatTypeLinesWithChan(request *DescribeThreatTypeLinesRequest) (<-chan *DescribeThreatTypeLinesResponse, <-chan error) {
	responseChan := make(chan *DescribeThreatTypeLinesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeThreatTypeLines(request)
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

// DescribeThreatTypeLinesWithCallback invokes the sas_api.DescribeThreatTypeLines API asynchronously
// api document: https://help.aliyun.com/api/sas-api/describethreattypelines.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeThreatTypeLinesWithCallback(request *DescribeThreatTypeLinesRequest, callback func(response *DescribeThreatTypeLinesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeThreatTypeLinesResponse
		var err error
		defer close(result)
		response, err = client.DescribeThreatTypeLines(request)
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

// DescribeThreatTypeLinesRequest is the request struct for api DescribeThreatTypeLines
type DescribeThreatTypeLinesRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	ApiType  requests.Integer `position:"Query" name:"ApiType"`
}

// DescribeThreatTypeLinesResponse is the response struct for api DescribeThreatTypeLines
type DescribeThreatTypeLinesResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	Categories []string `json:"Categories" xml:"Categories"`
	Items      []Item   `json:"Items" xml:"Items"`
}

// CreateDescribeThreatTypeLinesRequest creates a request to invoke DescribeThreatTypeLines API
func CreateDescribeThreatTypeLinesRequest() (request *DescribeThreatTypeLinesRequest) {
	request = &DescribeThreatTypeLinesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas-api", "2017-07-05", "DescribeThreatTypeLines", "sas-api", "openAPI")
	return
}

// CreateDescribeThreatTypeLinesResponse creates a response to parse from DescribeThreatTypeLines response
func CreateDescribeThreatTypeLinesResponse() (response *DescribeThreatTypeLinesResponse) {
	response = &DescribeThreatTypeLinesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
