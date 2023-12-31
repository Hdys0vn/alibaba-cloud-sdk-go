package smartag

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

// DescribeSmartAccessGatewayClientUsers invokes the smartag.DescribeSmartAccessGatewayClientUsers API synchronously
func (client *Client) DescribeSmartAccessGatewayClientUsers(request *DescribeSmartAccessGatewayClientUsersRequest) (response *DescribeSmartAccessGatewayClientUsersResponse, err error) {
	response = CreateDescribeSmartAccessGatewayClientUsersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSmartAccessGatewayClientUsersWithChan invokes the smartag.DescribeSmartAccessGatewayClientUsers API asynchronously
func (client *Client) DescribeSmartAccessGatewayClientUsersWithChan(request *DescribeSmartAccessGatewayClientUsersRequest) (<-chan *DescribeSmartAccessGatewayClientUsersResponse, <-chan error) {
	responseChan := make(chan *DescribeSmartAccessGatewayClientUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSmartAccessGatewayClientUsers(request)
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

// DescribeSmartAccessGatewayClientUsersWithCallback invokes the smartag.DescribeSmartAccessGatewayClientUsers API asynchronously
func (client *Client) DescribeSmartAccessGatewayClientUsersWithCallback(request *DescribeSmartAccessGatewayClientUsersRequest, callback func(response *DescribeSmartAccessGatewayClientUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSmartAccessGatewayClientUsersResponse
		var err error
		defer close(result)
		response, err = client.DescribeSmartAccessGatewayClientUsers(request)
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

// DescribeSmartAccessGatewayClientUsersRequest is the request struct for api DescribeSmartAccessGatewayClientUsers
type DescribeSmartAccessGatewayClientUsersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	EndLoginTime         requests.Integer `position:"Query" name:"EndLoginTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UserMail             string           `position:"Query" name:"UserMail"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	BeginLoginTime       requests.Integer `position:"Query" name:"BeginLoginTime"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	UserName             string           `position:"Query" name:"UserName"`
}

// DescribeSmartAccessGatewayClientUsersResponse is the response struct for api DescribeSmartAccessGatewayClientUsers
type DescribeSmartAccessGatewayClientUsersResponse struct {
	*responses.BaseResponse
	TotalCount int                                          `json:"TotalCount" xml:"TotalCount"`
	PageSize   int                                          `json:"PageSize" xml:"PageSize"`
	RequestId  string                                       `json:"RequestId" xml:"RequestId"`
	PageNumber int                                          `json:"PageNumber" xml:"PageNumber"`
	Users      UsersInDescribeSmartAccessGatewayClientUsers `json:"Users" xml:"Users"`
}

// CreateDescribeSmartAccessGatewayClientUsersRequest creates a request to invoke DescribeSmartAccessGatewayClientUsers API
func CreateDescribeSmartAccessGatewayClientUsersRequest() (request *DescribeSmartAccessGatewayClientUsersRequest) {
	request = &DescribeSmartAccessGatewayClientUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSmartAccessGatewayClientUsers", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSmartAccessGatewayClientUsersResponse creates a response to parse from DescribeSmartAccessGatewayClientUsers response
func CreateDescribeSmartAccessGatewayClientUsersResponse() (response *DescribeSmartAccessGatewayClientUsersResponse) {
	response = &DescribeSmartAccessGatewayClientUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
