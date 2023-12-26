package rds

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

// TransformDBInstancePayType invokes the rds.TransformDBInstancePayType API synchronously
func (client *Client) TransformDBInstancePayType(request *TransformDBInstancePayTypeRequest) (response *TransformDBInstancePayTypeResponse, err error) {
	response = CreateTransformDBInstancePayTypeResponse()
	err = client.DoAction(request, response)
	return
}

// TransformDBInstancePayTypeWithChan invokes the rds.TransformDBInstancePayType API asynchronously
func (client *Client) TransformDBInstancePayTypeWithChan(request *TransformDBInstancePayTypeRequest) (<-chan *TransformDBInstancePayTypeResponse, <-chan error) {
	responseChan := make(chan *TransformDBInstancePayTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TransformDBInstancePayType(request)
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

// TransformDBInstancePayTypeWithCallback invokes the rds.TransformDBInstancePayType API asynchronously
func (client *Client) TransformDBInstancePayTypeWithCallback(request *TransformDBInstancePayTypeRequest, callback func(response *TransformDBInstancePayTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TransformDBInstancePayTypeResponse
		var err error
		defer close(result)
		response, err = client.TransformDBInstancePayType(request)
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

// TransformDBInstancePayTypeRequest is the request struct for api TransformDBInstancePayType
type TransformDBInstancePayTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	Period               string           `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Resource             string           `position:"Query" name:"Resource"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	UsedTime             requests.Integer `position:"Query" name:"UsedTime"`
	AutoRenew            string           `position:"Query" name:"AutoRenew"`
	PayType              string           `position:"Query" name:"PayType"`
}

// TransformDBInstancePayTypeResponse is the response struct for api TransformDBInstancePayType
type TransformDBInstancePayTypeResponse struct {
	*responses.BaseResponse
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ExpiredTime  string `json:"ExpiredTime" xml:"ExpiredTime"`
	OrderId      int64  `json:"OrderId" xml:"OrderId"`
	ChargeType   string `json:"ChargeType" xml:"ChargeType"`
}

// CreateTransformDBInstancePayTypeRequest creates a request to invoke TransformDBInstancePayType API
func CreateTransformDBInstancePayTypeRequest() (request *TransformDBInstancePayTypeRequest) {
	request = &TransformDBInstancePayTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "TransformDBInstancePayType", "", "")
	request.Method = requests.POST
	return
}

// CreateTransformDBInstancePayTypeResponse creates a response to parse from TransformDBInstancePayType response
func CreateTransformDBInstancePayTypeResponse() (response *TransformDBInstancePayTypeResponse) {
	response = &TransformDBInstancePayTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
