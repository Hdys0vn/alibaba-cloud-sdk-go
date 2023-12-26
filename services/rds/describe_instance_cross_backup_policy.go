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

// DescribeInstanceCrossBackupPolicy invokes the rds.DescribeInstanceCrossBackupPolicy API synchronously
func (client *Client) DescribeInstanceCrossBackupPolicy(request *DescribeInstanceCrossBackupPolicyRequest) (response *DescribeInstanceCrossBackupPolicyResponse, err error) {
	response = CreateDescribeInstanceCrossBackupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceCrossBackupPolicyWithChan invokes the rds.DescribeInstanceCrossBackupPolicy API asynchronously
func (client *Client) DescribeInstanceCrossBackupPolicyWithChan(request *DescribeInstanceCrossBackupPolicyRequest) (<-chan *DescribeInstanceCrossBackupPolicyResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceCrossBackupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceCrossBackupPolicy(request)
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

// DescribeInstanceCrossBackupPolicyWithCallback invokes the rds.DescribeInstanceCrossBackupPolicy API asynchronously
func (client *Client) DescribeInstanceCrossBackupPolicyWithCallback(request *DescribeInstanceCrossBackupPolicyRequest, callback func(response *DescribeInstanceCrossBackupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceCrossBackupPolicyResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceCrossBackupPolicy(request)
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

// DescribeInstanceCrossBackupPolicyRequest is the request struct for api DescribeInstanceCrossBackupPolicy
type DescribeInstanceCrossBackupPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeInstanceCrossBackupPolicyResponse is the response struct for api DescribeInstanceCrossBackupPolicy
type DescribeInstanceCrossBackupPolicyResponse struct {
	*responses.BaseResponse
	RelServiceId          string `json:"RelServiceId" xml:"RelServiceId"`
	LogBackupEnabledTime  string `json:"LogBackupEnabledTime" xml:"LogBackupEnabledTime"`
	DBInstanceStatusDesc  string `json:"DBInstanceStatusDesc" xml:"DBInstanceStatusDesc"`
	StorageOwner          string `json:"StorageOwner" xml:"StorageOwner"`
	RelService            string `json:"RelService" xml:"RelService"`
	BackupEnabled         string `json:"BackupEnabled" xml:"BackupEnabled"`
	StorageType           string `json:"StorageType" xml:"StorageType"`
	DBInstanceStatus      string `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	LockMode              string `json:"LockMode" xml:"LockMode"`
	EngineVersion         string `json:"EngineVersion" xml:"EngineVersion"`
	RegionId              string `json:"RegionId" xml:"RegionId"`
	LogBackupEnabled      string `json:"LogBackupEnabled" xml:"LogBackupEnabled"`
	RequestId             string `json:"RequestId" xml:"RequestId"`
	RetentType            int    `json:"RetentType" xml:"RetentType"`
	Endpoint              string `json:"Endpoint" xml:"Endpoint"`
	Retention             int    `json:"Retention" xml:"Retention"`
	CrossBackupType       string `json:"CrossBackupType" xml:"CrossBackupType"`
	CrossBackupRegion     string `json:"CrossBackupRegion" xml:"CrossBackupRegion"`
	DBInstanceId          string `json:"DBInstanceId" xml:"DBInstanceId"`
	Engine                string `json:"Engine" xml:"Engine"`
	BackupEnabledTime     string `json:"BackupEnabledTime" xml:"BackupEnabledTime"`
	DBInstanceDescription string `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
}

// CreateDescribeInstanceCrossBackupPolicyRequest creates a request to invoke DescribeInstanceCrossBackupPolicy API
func CreateDescribeInstanceCrossBackupPolicyRequest() (request *DescribeInstanceCrossBackupPolicyRequest) {
	request = &DescribeInstanceCrossBackupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeInstanceCrossBackupPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceCrossBackupPolicyResponse creates a response to parse from DescribeInstanceCrossBackupPolicy response
func CreateDescribeInstanceCrossBackupPolicyResponse() (response *DescribeInstanceCrossBackupPolicyResponse) {
	response = &DescribeInstanceCrossBackupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}