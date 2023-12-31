package hbase

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

// DescribeInstance invokes the hbase.DescribeInstance API synchronously
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
	response = CreateDescribeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceWithChan invokes the hbase.DescribeInstance API asynchronously
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

// DescribeInstanceWithCallback invokes the hbase.DescribeInstance API asynchronously
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
	ClusterId string `position:"Query" name:"ClusterId"`
}

// DescribeInstanceResponse is the response struct for api DescribeInstance
type DescribeInstanceResponse struct {
	*responses.BaseResponse
	VpcId                string                 `json:"VpcId" xml:"VpcId"`
	Status               string                 `json:"Status" xml:"Status"`
	EncryptionType       string                 `json:"EncryptionType" xml:"EncryptionType"`
	ModuleId             int                    `json:"ModuleId" xml:"ModuleId"`
	VswitchId            string                 `json:"VswitchId" xml:"VswitchId"`
	BackupStatus         string                 `json:"BackupStatus" xml:"BackupStatus"`
	PayType              string                 `json:"PayType" xml:"PayType"`
	CoreDiskType         string                 `json:"CoreDiskType" xml:"CoreDiskType"`
	MasterNodeCount      int                    `json:"MasterNodeCount" xml:"MasterNodeCount"`
	NetworkType          string                 `json:"NetworkType" xml:"NetworkType"`
	CreatedTimeUTC       string                 `json:"CreatedTimeUTC" xml:"CreatedTimeUTC"`
	ColdStorageSize      int                    `json:"ColdStorageSize" xml:"ColdStorageSize"`
	ParentId             string                 `json:"ParentId" xml:"ParentId"`
	IsLatestVersion      bool                   `json:"IsLatestVersion" xml:"IsLatestVersion"`
	ExpireTimeUTC        string                 `json:"ExpireTimeUTC" xml:"ExpireTimeUTC"`
	RequestId            string                 `json:"RequestId" xml:"RequestId"`
	InstanceName         string                 `json:"InstanceName" xml:"InstanceName"`
	MasterInstanceType   string                 `json:"MasterInstanceType" xml:"MasterInstanceType"`
	CoreInstanceType     string                 `json:"CoreInstanceType" xml:"CoreInstanceType"`
	EncryptionKey        string                 `json:"EncryptionKey" xml:"EncryptionKey"`
	CreatedTime          string                 `json:"CreatedTime" xml:"CreatedTime"`
	CoreDiskSize         int                    `json:"CoreDiskSize" xml:"CoreDiskSize"`
	ClusterId            string                 `json:"ClusterId" xml:"ClusterId"`
	ExpireTime           string                 `json:"ExpireTime" xml:"ExpireTime"`
	MaintainStartTime    string                 `json:"MaintainStartTime" xml:"MaintainStartTime"`
	ConfirmMaintainTime  string                 `json:"ConfirmMaintainTime" xml:"ConfirmMaintainTime"`
	IsHa                 bool                   `json:"IsHa" xml:"IsHa"`
	MaintainEndTime      string                 `json:"MaintainEndTime" xml:"MaintainEndTime"`
	InstanceId           string                 `json:"InstanceId" xml:"InstanceId"`
	ColdStorageStatus    string                 `json:"ColdStorageStatus" xml:"ColdStorageStatus"`
	IsDeletionProtection bool                   `json:"IsDeletionProtection" xml:"IsDeletionProtection"`
	MinorVersion         string                 `json:"MinorVersion" xml:"MinorVersion"`
	RegionId             string                 `json:"RegionId" xml:"RegionId"`
	MasterDiskSize       int                    `json:"MasterDiskSize" xml:"MasterDiskSize"`
	MasterDiskType       string                 `json:"MasterDiskType" xml:"MasterDiskType"`
	NeedUpgrade          bool                   `json:"NeedUpgrade" xml:"NeedUpgrade"`
	IsMultiModel         bool                   `json:"IsMultiModel" xml:"IsMultiModel"`
	AutoRenewal          bool                   `json:"AutoRenewal" xml:"AutoRenewal"`
	ClusterType          string                 `json:"ClusterType" xml:"ClusterType"`
	ResourceGroupId      string                 `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ClusterName          string                 `json:"ClusterName" xml:"ClusterName"`
	ZoneId               string                 `json:"ZoneId" xml:"ZoneId"`
	Duration             int                    `json:"Duration" xml:"Duration"`
	ModuleStackVersion   string                 `json:"ModuleStackVersion" xml:"ModuleStackVersion"`
	Engine               string                 `json:"Engine" xml:"Engine"`
	MajorVersion         string                 `json:"MajorVersion" xml:"MajorVersion"`
	CoreDiskCount        string                 `json:"CoreDiskCount" xml:"CoreDiskCount"`
	TaskProgress         string                 `json:"TaskProgress" xml:"TaskProgress"`
	CoreNodeCount        int                    `json:"CoreNodeCount" xml:"CoreNodeCount"`
	NeedUpgradeComps     NeedUpgradeComps       `json:"NeedUpgradeComps" xml:"NeedUpgradeComps"`
	Tags                 TagsInDescribeInstance `json:"Tags" xml:"Tags"`
}

// CreateDescribeInstanceRequest creates a request to invoke DescribeInstance API
func CreateDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DescribeInstance", "hbase", "openAPI")
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
