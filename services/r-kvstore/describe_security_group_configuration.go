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

// DescribeSecurityGroupConfiguration invokes the r_kvstore.DescribeSecurityGroupConfiguration API synchronously
func (client *Client) DescribeSecurityGroupConfiguration(request *DescribeSecurityGroupConfigurationRequest) (response *DescribeSecurityGroupConfigurationResponse, err error) {
	response = CreateDescribeSecurityGroupConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecurityGroupConfigurationWithChan invokes the r_kvstore.DescribeSecurityGroupConfiguration API asynchronously
func (client *Client) DescribeSecurityGroupConfigurationWithChan(request *DescribeSecurityGroupConfigurationRequest) (<-chan *DescribeSecurityGroupConfigurationResponse, <-chan error) {
	responseChan := make(chan *DescribeSecurityGroupConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecurityGroupConfiguration(request)
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

// DescribeSecurityGroupConfigurationWithCallback invokes the r_kvstore.DescribeSecurityGroupConfiguration API asynchronously
func (client *Client) DescribeSecurityGroupConfigurationWithCallback(request *DescribeSecurityGroupConfigurationRequest, callback func(response *DescribeSecurityGroupConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecurityGroupConfigurationResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecurityGroupConfiguration(request)
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

// DescribeSecurityGroupConfigurationRequest is the request struct for api DescribeSecurityGroupConfiguration
type DescribeSecurityGroupConfigurationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeSecurityGroupConfigurationResponse is the response struct for api DescribeSecurityGroupConfiguration
type DescribeSecurityGroupConfigurationResponse struct {
	*responses.BaseResponse
	RequestId string                                    `json:"RequestId" xml:"RequestId"`
	Items     ItemsInDescribeSecurityGroupConfiguration `json:"Items" xml:"Items"`
}

// CreateDescribeSecurityGroupConfigurationRequest creates a request to invoke DescribeSecurityGroupConfiguration API
func CreateDescribeSecurityGroupConfigurationRequest() (request *DescribeSecurityGroupConfigurationRequest) {
	request = &DescribeSecurityGroupConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeSecurityGroupConfiguration", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSecurityGroupConfigurationResponse creates a response to parse from DescribeSecurityGroupConfiguration response
func CreateDescribeSecurityGroupConfigurationResponse() (response *DescribeSecurityGroupConfigurationResponse) {
	response = &DescribeSecurityGroupConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
