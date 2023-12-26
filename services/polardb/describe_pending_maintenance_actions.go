package polardb

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

// DescribePendingMaintenanceActions invokes the polardb.DescribePendingMaintenanceActions API synchronously
func (client *Client) DescribePendingMaintenanceActions(request *DescribePendingMaintenanceActionsRequest) (response *DescribePendingMaintenanceActionsResponse, err error) {
	response = CreateDescribePendingMaintenanceActionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePendingMaintenanceActionsWithChan invokes the polardb.DescribePendingMaintenanceActions API asynchronously
func (client *Client) DescribePendingMaintenanceActionsWithChan(request *DescribePendingMaintenanceActionsRequest) (<-chan *DescribePendingMaintenanceActionsResponse, <-chan error) {
	responseChan := make(chan *DescribePendingMaintenanceActionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePendingMaintenanceActions(request)
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

// DescribePendingMaintenanceActionsWithCallback invokes the polardb.DescribePendingMaintenanceActions API asynchronously
func (client *Client) DescribePendingMaintenanceActionsWithCallback(request *DescribePendingMaintenanceActionsRequest, callback func(response *DescribePendingMaintenanceActionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePendingMaintenanceActionsResponse
		var err error
		defer close(result)
		response, err = client.DescribePendingMaintenanceActions(request)
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

// DescribePendingMaintenanceActionsRequest is the request struct for api DescribePendingMaintenanceActions
type DescribePendingMaintenanceActionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	IsHistory            requests.Integer `position:"Query" name:"IsHistory"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribePendingMaintenanceActionsResponse is the response struct for api DescribePendingMaintenanceActions
type DescribePendingMaintenanceActionsResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	TypeList  []Items `json:"TypeList" xml:"TypeList"`
}

// CreateDescribePendingMaintenanceActionsRequest creates a request to invoke DescribePendingMaintenanceActions API
func CreateDescribePendingMaintenanceActionsRequest() (request *DescribePendingMaintenanceActionsRequest) {
	request = &DescribePendingMaintenanceActionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribePendingMaintenanceActions", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePendingMaintenanceActionsResponse creates a response to parse from DescribePendingMaintenanceActions response
func CreateDescribePendingMaintenanceActionsResponse() (response *DescribePendingMaintenanceActionsResponse) {
	response = &DescribePendingMaintenanceActionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}