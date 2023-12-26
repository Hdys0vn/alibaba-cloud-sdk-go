package sgw

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

// DescribeGatewayNFSClients invokes the sgw.DescribeGatewayNFSClients API synchronously
func (client *Client) DescribeGatewayNFSClients(request *DescribeGatewayNFSClientsRequest) (response *DescribeGatewayNFSClientsResponse, err error) {
	response = CreateDescribeGatewayNFSClientsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGatewayNFSClientsWithChan invokes the sgw.DescribeGatewayNFSClients API asynchronously
func (client *Client) DescribeGatewayNFSClientsWithChan(request *DescribeGatewayNFSClientsRequest) (<-chan *DescribeGatewayNFSClientsResponse, <-chan error) {
	responseChan := make(chan *DescribeGatewayNFSClientsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGatewayNFSClients(request)
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

// DescribeGatewayNFSClientsWithCallback invokes the sgw.DescribeGatewayNFSClients API asynchronously
func (client *Client) DescribeGatewayNFSClientsWithCallback(request *DescribeGatewayNFSClientsRequest, callback func(response *DescribeGatewayNFSClientsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGatewayNFSClientsResponse
		var err error
		defer close(result)
		response, err = client.DescribeGatewayNFSClients(request)
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

// DescribeGatewayNFSClientsRequest is the request struct for api DescribeGatewayNFSClients
type DescribeGatewayNFSClientsRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	GatewayId     string           `position:"Query" name:"GatewayId"`
}

// DescribeGatewayNFSClientsResponse is the response struct for api DescribeGatewayNFSClients
type DescribeGatewayNFSClientsResponse struct {
	*responses.BaseResponse
	RequestId        string         `json:"RequestId" xml:"RequestId"`
	Success          bool           `json:"Success" xml:"Success"`
	Code             string         `json:"Code" xml:"Code"`
	Message          string         `json:"Message" xml:"Message"`
	TotalCount       int            `json:"TotalCount" xml:"TotalCount"`
	PageNumber       int            `json:"PageNumber" xml:"PageNumber"`
	PageSize         int            `json:"PageSize" xml:"PageSize"`
	Version3Enabled  bool           `json:"Version3Enabled" xml:"Version3Enabled"`
	Version40Enabled bool           `json:"Version40Enabled" xml:"Version40Enabled"`
	Version41Enabled bool           `json:"Version41Enabled" xml:"Version41Enabled"`
	ClientInfoList   ClientInfoList `json:"ClientInfoList" xml:"ClientInfoList"`
}

// CreateDescribeGatewayNFSClientsRequest creates a request to invoke DescribeGatewayNFSClients API
func CreateDescribeGatewayNFSClientsRequest() (request *DescribeGatewayNFSClientsRequest) {
	request = &DescribeGatewayNFSClientsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeGatewayNFSClients", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGatewayNFSClientsResponse creates a response to parse from DescribeGatewayNFSClients response
func CreateDescribeGatewayNFSClientsResponse() (response *DescribeGatewayNFSClientsResponse) {
	response = &DescribeGatewayNFSClientsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
