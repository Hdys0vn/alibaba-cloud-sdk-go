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

// DescribeDBClusterAccessWhitelist invokes the polardb.DescribeDBClusterAccessWhitelist API synchronously
func (client *Client) DescribeDBClusterAccessWhitelist(request *DescribeDBClusterAccessWhitelistRequest) (response *DescribeDBClusterAccessWhitelistResponse, err error) {
	response = CreateDescribeDBClusterAccessWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterAccessWhitelistWithChan invokes the polardb.DescribeDBClusterAccessWhitelist API asynchronously
func (client *Client) DescribeDBClusterAccessWhitelistWithChan(request *DescribeDBClusterAccessWhitelistRequest) (<-chan *DescribeDBClusterAccessWhitelistResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterAccessWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterAccessWhitelist(request)
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

// DescribeDBClusterAccessWhitelistWithCallback invokes the polardb.DescribeDBClusterAccessWhitelist API asynchronously
func (client *Client) DescribeDBClusterAccessWhitelistWithCallback(request *DescribeDBClusterAccessWhitelistRequest, callback func(response *DescribeDBClusterAccessWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterAccessWhitelistResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterAccessWhitelist(request)
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

// DescribeDBClusterAccessWhitelistRequest is the request struct for api DescribeDBClusterAccessWhitelist
type DescribeDBClusterAccessWhitelistRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterAccessWhitelistResponse is the response struct for api DescribeDBClusterAccessWhitelist
type DescribeDBClusterAccessWhitelistResponse struct {
	*responses.BaseResponse
	RequestId               string                                  `json:"RequestId" xml:"RequestId"`
	Items                   ItemsInDescribeDBClusterAccessWhitelist `json:"Items" xml:"Items"`
	DBClusterSecurityGroups DBClusterSecurityGroups                 `json:"DBClusterSecurityGroups" xml:"DBClusterSecurityGroups"`
}

// CreateDescribeDBClusterAccessWhitelistRequest creates a request to invoke DescribeDBClusterAccessWhitelist API
func CreateDescribeDBClusterAccessWhitelistRequest() (request *DescribeDBClusterAccessWhitelistRequest) {
	request = &DescribeDBClusterAccessWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterAccessWhitelist", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterAccessWhitelistResponse creates a response to parse from DescribeDBClusterAccessWhitelist response
func CreateDescribeDBClusterAccessWhitelistResponse() (response *DescribeDBClusterAccessWhitelistResponse) {
	response = &DescribeDBClusterAccessWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
