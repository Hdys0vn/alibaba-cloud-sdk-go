package partnerbill

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

// GetPartnerInstancesMonthBill invokes the partnerbill.GetPartnerInstancesMonthBill API synchronously
// api document: https://help.aliyun.com/api/partnerbill/getpartnerinstancesmonthbill.html
func (client *Client) GetPartnerInstancesMonthBill(request *GetPartnerInstancesMonthBillRequest) (response *GetPartnerInstancesMonthBillResponse, err error) {
	response = CreateGetPartnerInstancesMonthBillResponse()
	err = client.DoAction(request, response)
	return
}

// GetPartnerInstancesMonthBillWithChan invokes the partnerbill.GetPartnerInstancesMonthBill API asynchronously
// api document: https://help.aliyun.com/api/partnerbill/getpartnerinstancesmonthbill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPartnerInstancesMonthBillWithChan(request *GetPartnerInstancesMonthBillRequest) (<-chan *GetPartnerInstancesMonthBillResponse, <-chan error) {
	responseChan := make(chan *GetPartnerInstancesMonthBillResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPartnerInstancesMonthBill(request)
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

// GetPartnerInstancesMonthBillWithCallback invokes the partnerbill.GetPartnerInstancesMonthBill API asynchronously
// api document: https://help.aliyun.com/api/partnerbill/getpartnerinstancesmonthbill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPartnerInstancesMonthBillWithCallback(request *GetPartnerInstancesMonthBillRequest, callback func(response *GetPartnerInstancesMonthBillResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPartnerInstancesMonthBillResponse
		var err error
		defer close(result)
		response, err = client.GetPartnerInstancesMonthBill(request)
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

// GetPartnerInstancesMonthBillRequest is the request struct for api GetPartnerInstancesMonthBill
type GetPartnerInstancesMonthBillRequest struct {
	*requests.RpcRequest
	AliyunPk            requests.Integer `position:"Query" name:"AliyunPk"`
	ContainRelatedOrder requests.Boolean `position:"Query" name:"ContainRelatedOrder"`
	Test                requests.Boolean `position:"Query" name:"Test"`
	BillPeriod          string           `position:"Query" name:"BillPeriod"`
}

// GetPartnerInstancesMonthBillResponse is the response struct for api GetPartnerInstancesMonthBill
type GetPartnerInstancesMonthBillResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetPartnerInstancesMonthBillRequest creates a request to invoke GetPartnerInstancesMonthBill API
func CreateGetPartnerInstancesMonthBillRequest() (request *GetPartnerInstancesMonthBillRequest) {
	request = &GetPartnerInstancesMonthBillRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PartnerBill", "2020-01-08", "GetPartnerInstancesMonthBill", "GetPartnerInstancesMonthBill", "openAPI")
	return
}

// CreateGetPartnerInstancesMonthBillResponse creates a response to parse from GetPartnerInstancesMonthBill response
func CreateGetPartnerInstancesMonthBillResponse() (response *GetPartnerInstancesMonthBillResponse) {
	response = &GetPartnerInstancesMonthBillResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
