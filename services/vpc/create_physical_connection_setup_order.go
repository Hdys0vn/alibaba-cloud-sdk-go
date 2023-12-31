package vpc

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

// CreatePhysicalConnectionSetupOrder invokes the vpc.CreatePhysicalConnectionSetupOrder API synchronously
func (client *Client) CreatePhysicalConnectionSetupOrder(request *CreatePhysicalConnectionSetupOrderRequest) (response *CreatePhysicalConnectionSetupOrderResponse, err error) {
	response = CreateCreatePhysicalConnectionSetupOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePhysicalConnectionSetupOrderWithChan invokes the vpc.CreatePhysicalConnectionSetupOrder API asynchronously
func (client *Client) CreatePhysicalConnectionSetupOrderWithChan(request *CreatePhysicalConnectionSetupOrderRequest) (<-chan *CreatePhysicalConnectionSetupOrderResponse, <-chan error) {
	responseChan := make(chan *CreatePhysicalConnectionSetupOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePhysicalConnectionSetupOrder(request)
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

// CreatePhysicalConnectionSetupOrderWithCallback invokes the vpc.CreatePhysicalConnectionSetupOrder API asynchronously
func (client *Client) CreatePhysicalConnectionSetupOrderWithCallback(request *CreatePhysicalConnectionSetupOrderRequest, callback func(response *CreatePhysicalConnectionSetupOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePhysicalConnectionSetupOrderResponse
		var err error
		defer close(result)
		response, err = client.CreatePhysicalConnectionSetupOrder(request)
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

// CreatePhysicalConnectionSetupOrderRequest is the request struct for api CreatePhysicalConnectionSetupOrder
type CreatePhysicalConnectionSetupOrderRequest struct {
	*requests.RpcRequest
	AccessPointId                 string           `position:"Query" name:"AccessPointId"`
	ResourceOwnerId               requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PortType                      string           `position:"Query" name:"PortType"`
	ClientToken                   string           `position:"Query" name:"ClientToken"`
	RedundantPhysicalConnectionId string           `position:"Query" name:"RedundantPhysicalConnectionId"`
	AutoPay                       requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount          string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string           `position:"Query" name:"OwnerAccount"`
	OwnerId                       requests.Integer `position:"Query" name:"OwnerId"`
	LineOperator                  string           `position:"Query" name:"LineOperator"`
}

// CreatePhysicalConnectionSetupOrderResponse is the response struct for api CreatePhysicalConnectionSetupOrder
type CreatePhysicalConnectionSetupOrderResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	OrderId              string `json:"OrderId" xml:"OrderId"`
	PhysicalConnectionId string `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
}

// CreateCreatePhysicalConnectionSetupOrderRequest creates a request to invoke CreatePhysicalConnectionSetupOrder API
func CreateCreatePhysicalConnectionSetupOrderRequest() (request *CreatePhysicalConnectionSetupOrderRequest) {
	request = &CreatePhysicalConnectionSetupOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreatePhysicalConnectionSetupOrder", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatePhysicalConnectionSetupOrderResponse creates a response to parse from CreatePhysicalConnectionSetupOrder response
func CreateCreatePhysicalConnectionSetupOrderResponse() (response *CreatePhysicalConnectionSetupOrderResponse) {
	response = &CreatePhysicalConnectionSetupOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
