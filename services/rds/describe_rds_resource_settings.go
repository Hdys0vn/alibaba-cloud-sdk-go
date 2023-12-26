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

// DescribeRdsResourceSettings invokes the rds.DescribeRdsResourceSettings API synchronously
func (client *Client) DescribeRdsResourceSettings(request *DescribeRdsResourceSettingsRequest) (response *DescribeRdsResourceSettingsResponse, err error) {
	response = CreateDescribeRdsResourceSettingsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRdsResourceSettingsWithChan invokes the rds.DescribeRdsResourceSettings API asynchronously
func (client *Client) DescribeRdsResourceSettingsWithChan(request *DescribeRdsResourceSettingsRequest) (<-chan *DescribeRdsResourceSettingsResponse, <-chan error) {
	responseChan := make(chan *DescribeRdsResourceSettingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRdsResourceSettings(request)
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

// DescribeRdsResourceSettingsWithCallback invokes the rds.DescribeRdsResourceSettings API asynchronously
func (client *Client) DescribeRdsResourceSettingsWithCallback(request *DescribeRdsResourceSettingsRequest, callback func(response *DescribeRdsResourceSettingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRdsResourceSettingsResponse
		var err error
		defer close(result)
		response, err = client.DescribeRdsResourceSettings(request)
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

// DescribeRdsResourceSettingsRequest is the request struct for api DescribeRdsResourceSettings
type DescribeRdsResourceSettingsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceNiche        string           `position:"Query" name:"ResourceNiche"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeRdsResourceSettingsResponse is the response struct for api DescribeRdsResourceSettings
type DescribeRdsResourceSettingsResponse struct {
	*responses.BaseResponse
	RequestId                   string                      `json:"RequestId" xml:"RequestId"`
	RdsInstanceResourceSettings RdsInstanceResourceSettings `json:"RdsInstanceResourceSettings" xml:"RdsInstanceResourceSettings"`
}

// CreateDescribeRdsResourceSettingsRequest creates a request to invoke DescribeRdsResourceSettings API
func CreateDescribeRdsResourceSettingsRequest() (request *DescribeRdsResourceSettingsRequest) {
	request = &DescribeRdsResourceSettingsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeRdsResourceSettings", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRdsResourceSettingsResponse creates a response to parse from DescribeRdsResourceSettings response
func CreateDescribeRdsResourceSettingsResponse() (response *DescribeRdsResourceSettingsResponse) {
	response = &DescribeRdsResourceSettingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
