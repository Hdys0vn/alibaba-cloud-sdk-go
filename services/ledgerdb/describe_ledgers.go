package ledgerdb

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

// DescribeLedgers invokes the ledgerdb.DescribeLedgers API synchronously
// api document: https://help.aliyun.com/api/ledgerdb/describeledgers.html
func (client *Client) DescribeLedgers(request *DescribeLedgersRequest) (response *DescribeLedgersResponse, err error) {
	response = CreateDescribeLedgersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLedgersWithChan invokes the ledgerdb.DescribeLedgers API asynchronously
// api document: https://help.aliyun.com/api/ledgerdb/describeledgers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLedgersWithChan(request *DescribeLedgersRequest) (<-chan *DescribeLedgersResponse, <-chan error) {
	responseChan := make(chan *DescribeLedgersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLedgers(request)
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

// DescribeLedgersWithCallback invokes the ledgerdb.DescribeLedgers API asynchronously
// api document: https://help.aliyun.com/api/ledgerdb/describeledgers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLedgersWithCallback(request *DescribeLedgersRequest, callback func(response *DescribeLedgersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLedgersResponse
		var err error
		defer close(result)
		response, err = client.DescribeLedgers(request)
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

// DescribeLedgersRequest is the request struct for api DescribeLedgers
type DescribeLedgersRequest struct {
	*requests.RpcRequest
	NextToken  string           `position:"Query" name:"NextToken"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
}

// DescribeLedgersResponse is the response struct for api DescribeLedgers
type DescribeLedgersResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	NextToken  string   `json:"NextToken" xml:"NextToken"`
	MaxResults int      `json:"MaxResults" xml:"MaxResults"`
	Ledgers    []Ledger `json:"Ledgers" xml:"Ledgers"`
}

// CreateDescribeLedgersRequest creates a request to invoke DescribeLedgers API
func CreateDescribeLedgersRequest() (request *DescribeLedgersRequest) {
	request = &DescribeLedgersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ledgerdb", "2019-11-22", "DescribeLedgers", "ledgerdb", "openAPI")
	return
}

// CreateDescribeLedgersResponse creates a response to parse from DescribeLedgers response
func CreateDescribeLedgersResponse() (response *DescribeLedgersResponse) {
	response = &DescribeLedgersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
