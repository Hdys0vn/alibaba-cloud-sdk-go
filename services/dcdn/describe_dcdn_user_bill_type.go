package dcdn

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

// DescribeDcdnUserBillType invokes the dcdn.DescribeDcdnUserBillType API synchronously
func (client *Client) DescribeDcdnUserBillType(request *DescribeDcdnUserBillTypeRequest) (response *DescribeDcdnUserBillTypeResponse, err error) {
	response = CreateDescribeDcdnUserBillTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnUserBillTypeWithChan invokes the dcdn.DescribeDcdnUserBillType API asynchronously
func (client *Client) DescribeDcdnUserBillTypeWithChan(request *DescribeDcdnUserBillTypeRequest) (<-chan *DescribeDcdnUserBillTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnUserBillTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnUserBillType(request)
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

// DescribeDcdnUserBillTypeWithCallback invokes the dcdn.DescribeDcdnUserBillType API asynchronously
func (client *Client) DescribeDcdnUserBillTypeWithCallback(request *DescribeDcdnUserBillTypeRequest, callback func(response *DescribeDcdnUserBillTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnUserBillTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnUserBillType(request)
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

// DescribeDcdnUserBillTypeRequest is the request struct for api DescribeDcdnUserBillType
type DescribeDcdnUserBillTypeRequest struct {
	*requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
}

// DescribeDcdnUserBillTypeResponse is the response struct for api DescribeDcdnUserBillType
type DescribeDcdnUserBillTypeResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	BillTypeData BillTypeData `json:"BillTypeData" xml:"BillTypeData"`
}

// CreateDescribeDcdnUserBillTypeRequest creates a request to invoke DescribeDcdnUserBillType API
func CreateDescribeDcdnUserBillTypeRequest() (request *DescribeDcdnUserBillTypeRequest) {
	request = &DescribeDcdnUserBillTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnUserBillType", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnUserBillTypeResponse creates a response to parse from DescribeDcdnUserBillType response
func CreateDescribeDcdnUserBillTypeResponse() (response *DescribeDcdnUserBillTypeResponse) {
	response = &DescribeDcdnUserBillTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}