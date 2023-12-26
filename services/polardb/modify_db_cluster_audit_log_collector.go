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

// ModifyDBClusterAuditLogCollector invokes the polardb.ModifyDBClusterAuditLogCollector API synchronously
func (client *Client) ModifyDBClusterAuditLogCollector(request *ModifyDBClusterAuditLogCollectorRequest) (response *ModifyDBClusterAuditLogCollectorResponse, err error) {
	response = CreateModifyDBClusterAuditLogCollectorResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterAuditLogCollectorWithChan invokes the polardb.ModifyDBClusterAuditLogCollector API asynchronously
func (client *Client) ModifyDBClusterAuditLogCollectorWithChan(request *ModifyDBClusterAuditLogCollectorRequest) (<-chan *ModifyDBClusterAuditLogCollectorResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterAuditLogCollectorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterAuditLogCollector(request)
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

// ModifyDBClusterAuditLogCollectorWithCallback invokes the polardb.ModifyDBClusterAuditLogCollector API asynchronously
func (client *Client) ModifyDBClusterAuditLogCollectorWithCallback(request *ModifyDBClusterAuditLogCollectorRequest, callback func(response *ModifyDBClusterAuditLogCollectorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterAuditLogCollectorResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterAuditLogCollector(request)
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

// ModifyDBClusterAuditLogCollectorRequest is the request struct for api ModifyDBClusterAuditLogCollector
type ModifyDBClusterAuditLogCollectorRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	CollectorStatus      string           `position:"Query" name:"CollectorStatus"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDBClusterAuditLogCollectorResponse is the response struct for api ModifyDBClusterAuditLogCollector
type ModifyDBClusterAuditLogCollectorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterAuditLogCollectorRequest creates a request to invoke ModifyDBClusterAuditLogCollector API
func CreateModifyDBClusterAuditLogCollectorRequest() (request *ModifyDBClusterAuditLogCollectorRequest) {
	request = &ModifyDBClusterAuditLogCollectorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterAuditLogCollector", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterAuditLogCollectorResponse creates a response to parse from ModifyDBClusterAuditLogCollector response
func CreateModifyDBClusterAuditLogCollectorResponse() (response *ModifyDBClusterAuditLogCollectorResponse) {
	response = &ModifyDBClusterAuditLogCollectorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
