package r_kvstore

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

// DescribeUserClusterHostInstance invokes the r_kvstore.DescribeUserClusterHostInstance API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeuserclusterhostinstance.html
func (client *Client) DescribeUserClusterHostInstance(request *DescribeUserClusterHostInstanceRequest) (response *DescribeUserClusterHostInstanceResponse, err error) {
	response = CreateDescribeUserClusterHostInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserClusterHostInstanceWithChan invokes the r_kvstore.DescribeUserClusterHostInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeuserclusterhostinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserClusterHostInstanceWithChan(request *DescribeUserClusterHostInstanceRequest) (<-chan *DescribeUserClusterHostInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeUserClusterHostInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserClusterHostInstance(request)
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

// DescribeUserClusterHostInstanceWithCallback invokes the r_kvstore.DescribeUserClusterHostInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeuserclusterhostinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserClusterHostInstanceWithCallback(request *DescribeUserClusterHostInstanceRequest, callback func(response *DescribeUserClusterHostInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserClusterHostInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserClusterHostInstance(request)
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

// DescribeUserClusterHostInstanceRequest is the request struct for api DescribeUserClusterHostInstance
type DescribeUserClusterHostInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           string           `position:"Query" name:"PageNumber"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Engine               string           `position:"Query" name:"Engine"`
	InstanceStatus       string           `position:"Query" name:"InstanceStatus"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	MaxRecordsPerPage    string           `position:"Query" name:"MaxRecordsPerPage"`
	ClusterId            string           `position:"Query" name:"ClusterId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceIds          string           `position:"Query" name:"InstanceIds"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DescribeUserClusterHostInstanceResponse is the response struct for api DescribeUserClusterHostInstance
type DescribeUserClusterHostInstanceResponse struct {
	*responses.BaseResponse
	RequestId         string         `json:"RequestId" xml:"RequestId"`
	MaxRecordsPerPage int            `json:"MaxRecordsPerPage" xml:"MaxRecordsPerPage"`
	PageNumber        int            `json:"PageNumber" xml:"PageNumber"`
	TotalRecords      int            `json:"TotalRecords" xml:"TotalRecords"`
	ItemNumbers       int            `json:"ItemNumbers" xml:"ItemNumbers"`
	InstancesItems    InstancesItems `json:"InstancesItems" xml:"InstancesItems"`
}

// CreateDescribeUserClusterHostInstanceRequest creates a request to invoke DescribeUserClusterHostInstance API
func CreateDescribeUserClusterHostInstanceRequest() (request *DescribeUserClusterHostInstanceRequest) {
	request = &DescribeUserClusterHostInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeUserClusterHostInstance", "redisa", "openAPI")
	return
}

// CreateDescribeUserClusterHostInstanceResponse creates a response to parse from DescribeUserClusterHostInstance response
func CreateDescribeUserClusterHostInstanceResponse() (response *DescribeUserClusterHostInstanceResponse) {
	response = &DescribeUserClusterHostInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
