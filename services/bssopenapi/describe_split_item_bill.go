package bssopenapi

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

// DescribeSplitItemBill invokes the bssopenapi.DescribeSplitItemBill API synchronously
func (client *Client) DescribeSplitItemBill(request *DescribeSplitItemBillRequest) (response *DescribeSplitItemBillResponse, err error) {
	response = CreateDescribeSplitItemBillResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSplitItemBillWithChan invokes the bssopenapi.DescribeSplitItemBill API asynchronously
func (client *Client) DescribeSplitItemBillWithChan(request *DescribeSplitItemBillRequest) (<-chan *DescribeSplitItemBillResponse, <-chan error) {
	responseChan := make(chan *DescribeSplitItemBillResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSplitItemBill(request)
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

// DescribeSplitItemBillWithCallback invokes the bssopenapi.DescribeSplitItemBill API asynchronously
func (client *Client) DescribeSplitItemBillWithCallback(request *DescribeSplitItemBillRequest, callback func(response *DescribeSplitItemBillResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSplitItemBillResponse
		var err error
		defer close(result)
		response, err = client.DescribeSplitItemBill(request)
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

// DescribeSplitItemBillRequest is the request struct for api DescribeSplitItemBill
type DescribeSplitItemBillRequest struct {
	*requests.RpcRequest
	ProductCode      string                            `position:"Query" name:"ProductCode"`
	SubscriptionType string                            `position:"Query" name:"SubscriptionType"`
	BillOwnerId      requests.Integer                  `position:"Query" name:"BillOwnerId"`
	ProductType      string                            `position:"Query" name:"ProductType"`
	NextToken        string                            `position:"Query" name:"NextToken"`
	SplitItemID      string                            `position:"Query" name:"SplitItemID"`
	BillingCycle     string                            `position:"Query" name:"BillingCycle"`
	OwnerId          requests.Integer                  `position:"Query" name:"OwnerId"`
	TagFilter        *[]DescribeSplitItemBillTagFilter `position:"Query" name:"TagFilter"  type:"Repeated"`
	BillingDate      string                            `position:"Query" name:"BillingDate"`
	InstanceID       string                            `position:"Query" name:"InstanceID"`
	Granularity      string                            `position:"Query" name:"Granularity"`
	MaxResults       requests.Integer                  `position:"Query" name:"MaxResults"`
}

// DescribeSplitItemBillTagFilter is a repeated param struct in DescribeSplitItemBillRequest
type DescribeSplitItemBillTagFilter struct {
	TagValues *[]string `name:"TagValues" type:"Repeated"`
	TagKey    string    `name:"TagKey"`
}

// DescribeSplitItemBillResponse is the response struct for api DescribeSplitItemBill
type DescribeSplitItemBillResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeSplitItemBillRequest creates a request to invoke DescribeSplitItemBill API
func CreateDescribeSplitItemBillRequest() (request *DescribeSplitItemBillRequest) {
	request = &DescribeSplitItemBillRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "DescribeSplitItemBill", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSplitItemBillResponse creates a response to parse from DescribeSplitItemBill response
func CreateDescribeSplitItemBillResponse() (response *DescribeSplitItemBillResponse) {
	response = &DescribeSplitItemBillResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
