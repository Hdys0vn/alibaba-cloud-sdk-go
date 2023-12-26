package dds

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

// UpgradeDBInstanceKernelVersion invokes the dds.UpgradeDBInstanceKernelVersion API synchronously
func (client *Client) UpgradeDBInstanceKernelVersion(request *UpgradeDBInstanceKernelVersionRequest) (response *UpgradeDBInstanceKernelVersionResponse, err error) {
	response = CreateUpgradeDBInstanceKernelVersionResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeDBInstanceKernelVersionWithChan invokes the dds.UpgradeDBInstanceKernelVersion API asynchronously
func (client *Client) UpgradeDBInstanceKernelVersionWithChan(request *UpgradeDBInstanceKernelVersionRequest) (<-chan *UpgradeDBInstanceKernelVersionResponse, <-chan error) {
	responseChan := make(chan *UpgradeDBInstanceKernelVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeDBInstanceKernelVersion(request)
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

// UpgradeDBInstanceKernelVersionWithCallback invokes the dds.UpgradeDBInstanceKernelVersion API asynchronously
func (client *Client) UpgradeDBInstanceKernelVersionWithCallback(request *UpgradeDBInstanceKernelVersionRequest, callback func(response *UpgradeDBInstanceKernelVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeDBInstanceKernelVersionResponse
		var err error
		defer close(result)
		response, err = client.UpgradeDBInstanceKernelVersion(request)
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

// UpgradeDBInstanceKernelVersionRequest is the request struct for api UpgradeDBInstanceKernelVersion
type UpgradeDBInstanceKernelVersionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// UpgradeDBInstanceKernelVersionResponse is the response struct for api UpgradeDBInstanceKernelVersion
type UpgradeDBInstanceKernelVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpgradeDBInstanceKernelVersionRequest creates a request to invoke UpgradeDBInstanceKernelVersion API
func CreateUpgradeDBInstanceKernelVersionRequest() (request *UpgradeDBInstanceKernelVersionRequest) {
	request = &UpgradeDBInstanceKernelVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "UpgradeDBInstanceKernelVersion", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpgradeDBInstanceKernelVersionResponse creates a response to parse from UpgradeDBInstanceKernelVersion response
func CreateUpgradeDBInstanceKernelVersionResponse() (response *UpgradeDBInstanceKernelVersionResponse) {
	response = &UpgradeDBInstanceKernelVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
