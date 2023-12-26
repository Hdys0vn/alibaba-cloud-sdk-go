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

// DescribeRoleZoneInfo invokes the dds.DescribeRoleZoneInfo API synchronously
func (client *Client) DescribeRoleZoneInfo(request *DescribeRoleZoneInfoRequest) (response *DescribeRoleZoneInfoResponse, err error) {
	response = CreateDescribeRoleZoneInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRoleZoneInfoWithChan invokes the dds.DescribeRoleZoneInfo API asynchronously
func (client *Client) DescribeRoleZoneInfoWithChan(request *DescribeRoleZoneInfoRequest) (<-chan *DescribeRoleZoneInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeRoleZoneInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRoleZoneInfo(request)
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

// DescribeRoleZoneInfoWithCallback invokes the dds.DescribeRoleZoneInfo API asynchronously
func (client *Client) DescribeRoleZoneInfoWithCallback(request *DescribeRoleZoneInfoRequest, callback func(response *DescribeRoleZoneInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRoleZoneInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeRoleZoneInfo(request)
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

// DescribeRoleZoneInfoRequest is the request struct for api DescribeRoleZoneInfo
type DescribeRoleZoneInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeRoleZoneInfoResponse is the response struct for api DescribeRoleZoneInfo
type DescribeRoleZoneInfoResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	ZoneInfos ZoneInfos `json:"ZoneInfos" xml:"ZoneInfos"`
}

// CreateDescribeRoleZoneInfoRequest creates a request to invoke DescribeRoleZoneInfo API
func CreateDescribeRoleZoneInfoRequest() (request *DescribeRoleZoneInfoRequest) {
	request = &DescribeRoleZoneInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeRoleZoneInfo", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRoleZoneInfoResponse creates a response to parse from DescribeRoleZoneInfo response
func CreateDescribeRoleZoneInfoResponse() (response *DescribeRoleZoneInfoResponse) {
	response = &DescribeRoleZoneInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
