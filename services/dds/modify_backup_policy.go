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

// ModifyBackupPolicy invokes the dds.ModifyBackupPolicy API synchronously
func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (response *ModifyBackupPolicyResponse, err error) {
	response = CreateModifyBackupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBackupPolicyWithChan invokes the dds.ModifyBackupPolicy API asynchronously
func (client *Client) ModifyBackupPolicyWithChan(request *ModifyBackupPolicyRequest) (<-chan *ModifyBackupPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyBackupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBackupPolicy(request)
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

// ModifyBackupPolicyWithCallback invokes the dds.ModifyBackupPolicy API asynchronously
func (client *Client) ModifyBackupPolicyWithCallback(request *ModifyBackupPolicyRequest, callback func(response *ModifyBackupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBackupPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyBackupPolicy(request)
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

// ModifyBackupPolicyRequest is the request struct for api ModifyBackupPolicy
type ModifyBackupPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BackupInterval           string           `position:"Query" name:"BackupInterval"`
	SecurityToken            string           `position:"Query" name:"SecurityToken"`
	DBInstanceId             string           `position:"Query" name:"DBInstanceId"`
	EnableBackupLog          requests.Integer `position:"Query" name:"EnableBackupLog"`
	PreferredBackupPeriod    string           `position:"Query" name:"PreferredBackupPeriod"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	SnapshotBackupType       string           `position:"Query" name:"SnapshotBackupType"`
	PreferredBackupTime      string           `position:"Query" name:"PreferredBackupTime"`
	BackupRetentionPeriod    requests.Integer `position:"Query" name:"BackupRetentionPeriod"`
	LogBackupRetentionPeriod requests.Integer `position:"Query" name:"LogBackupRetentionPeriod"`
}

// ModifyBackupPolicyResponse is the response struct for api ModifyBackupPolicy
type ModifyBackupPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyBackupPolicyRequest creates a request to invoke ModifyBackupPolicy API
func CreateModifyBackupPolicyRequest() (request *ModifyBackupPolicyRequest) {
	request = &ModifyBackupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyBackupPolicy", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyBackupPolicyResponse creates a response to parse from ModifyBackupPolicy response
func CreateModifyBackupPolicyResponse() (response *ModifyBackupPolicyResponse) {
	response = &ModifyBackupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
