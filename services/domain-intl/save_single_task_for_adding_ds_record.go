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

// SaveSingleTaskForAddingDSRecord invokes the domain_intl.SaveSingleTaskForAddingDSRecord API synchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskforaddingdsrecord.html
func (client *Client) SaveSingleTaskForAddingDSRecord(request *SaveSingleTaskForAddingDSRecordRequest) (response *SaveSingleTaskForAddingDSRecordResponse, err error) {
	response = CreateSaveSingleTaskForAddingDSRecordResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForAddingDSRecordWithChan invokes the domain_intl.SaveSingleTaskForAddingDSRecord API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskforaddingdsrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForAddingDSRecordWithChan(request *SaveSingleTaskForAddingDSRecordRequest) (<-chan *SaveSingleTaskForAddingDSRecordResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForAddingDSRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForAddingDSRecord(request)
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

// SaveSingleTaskForAddingDSRecordWithCallback invokes the domain_intl.SaveSingleTaskForAddingDSRecord API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskforaddingdsrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForAddingDSRecordWithCallback(request *SaveSingleTaskForAddingDSRecordRequest, callback func(response *SaveSingleTaskForAddingDSRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForAddingDSRecordResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForAddingDSRecord(request)
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

// SaveSingleTaskForAddingDSRecordRequest is the request struct for api SaveSingleTaskForAddingDSRecord
type SaveSingleTaskForAddingDSRecordRequest struct {
	*requests.RpcRequest
	KeyTag       requests.Integer `position:"Query" name:"KeyTag"`
	DomainName   string           `position:"Query" name:"DomainName"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	DigestType   requests.Integer `position:"Query" name:"DigestType"`
	Digest       string           `position:"Query" name:"Digest"`
	Lang         string           `position:"Query" name:"Lang"`
	Algorithm    requests.Integer `position:"Query" name:"Algorithm"`
}

// SaveSingleTaskForAddingDSRecordResponse is the response struct for api SaveSingleTaskForAddingDSRecord
type SaveSingleTaskForAddingDSRecordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForAddingDSRecordRequest creates a request to invoke SaveSingleTaskForAddingDSRecord API
func CreateSaveSingleTaskForAddingDSRecordRequest() (request *SaveSingleTaskForAddingDSRecordRequest) {
	request = &SaveSingleTaskForAddingDSRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForAddingDSRecord", "domain", "openAPI")
	return
}

// CreateSaveSingleTaskForAddingDSRecordResponse creates a response to parse from SaveSingleTaskForAddingDSRecord response
func CreateSaveSingleTaskForAddingDSRecordResponse() (response *SaveSingleTaskForAddingDSRecordResponse) {
	response = &SaveSingleTaskForAddingDSRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
