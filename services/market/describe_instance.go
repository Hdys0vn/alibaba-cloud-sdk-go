package market

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

// DescribeInstance invokes the market.DescribeInstance API synchronously
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
	response = CreateDescribeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceWithChan invokes the market.DescribeInstance API asynchronously
func (client *Client) DescribeInstanceWithChan(request *DescribeInstanceRequest) (<-chan *DescribeInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstance(request)
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

// DescribeInstanceWithCallback invokes the market.DescribeInstance API asynchronously
func (client *Client) DescribeInstanceWithCallback(request *DescribeInstanceRequest, callback func(response *DescribeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstance(request)
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

// DescribeInstanceRequest is the request struct for api DescribeInstance
type DescribeInstanceRequest struct {
	*requests.RpcRequest
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	OrderType  string           `position:"Query" name:"OrderType"`
}

// DescribeInstanceResponse is the response struct for api DescribeInstance
type DescribeInstanceResponse struct {
	*responses.BaseResponse
	Status         string                    `json:"Status" xml:"Status"`
	AppJson        string                    `json:"AppJson" xml:"AppJson"`
	ProductName    string                    `json:"ProductName" xml:"ProductName"`
	InstanceId     int64                     `json:"InstanceId" xml:"InstanceId"`
	ExtendJson     string                    `json:"ExtendJson" xml:"ExtendJson"`
	IsTrial        bool                      `json:"IsTrial" xml:"IsTrial"`
	BeganOn        int64                     `json:"BeganOn" xml:"BeganOn"`
	ComponentJson  string                    `json:"ComponentJson" xml:"ComponentJson"`
	Constraints    string                    `json:"Constraints" xml:"Constraints"`
	ProductType    string                    `json:"ProductType" xml:"ProductType"`
	HostJson       string                    `json:"HostJson" xml:"HostJson"`
	ProductSkuCode string                    `json:"ProductSkuCode" xml:"ProductSkuCode"`
	CreatedOn      int64                     `json:"CreatedOn" xml:"CreatedOn"`
	EndOn          int64                     `json:"EndOn" xml:"EndOn"`
	OrderId        int64                     `json:"OrderId" xml:"OrderId"`
	ProductCode    string                    `json:"ProductCode" xml:"ProductCode"`
	SupplierName   string                    `json:"SupplierName" xml:"SupplierName"`
	AutoRenewal    string                    `json:"AutoRenewal" xml:"AutoRenewal"`
	RelationalData RelationalData            `json:"RelationalData" xml:"RelationalData"`
	Modules        ModulesInDescribeInstance `json:"Modules" xml:"Modules"`
}

// CreateDescribeInstanceRequest creates a request to invoke DescribeInstance API
func CreateDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Market", "2015-11-01", "DescribeInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceResponse creates a response to parse from DescribeInstance response
func CreateDescribeInstanceResponse() (response *DescribeInstanceResponse) {
	response = &DescribeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
