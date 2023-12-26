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

// DescribeInstanceLinkedWhitelistTemplate invokes the rds.DescribeInstanceLinkedWhitelistTemplate API synchronously
func (client *Client) DescribeInstanceLinkedWhitelistTemplate(request *DescribeInstanceLinkedWhitelistTemplateRequest) (response *DescribeInstanceLinkedWhitelistTemplateResponse, err error) {
	response = CreateDescribeInstanceLinkedWhitelistTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceLinkedWhitelistTemplateWithChan invokes the rds.DescribeInstanceLinkedWhitelistTemplate API asynchronously
func (client *Client) DescribeInstanceLinkedWhitelistTemplateWithChan(request *DescribeInstanceLinkedWhitelistTemplateRequest) (<-chan *DescribeInstanceLinkedWhitelistTemplateResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceLinkedWhitelistTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceLinkedWhitelistTemplate(request)
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

// DescribeInstanceLinkedWhitelistTemplateWithCallback invokes the rds.DescribeInstanceLinkedWhitelistTemplate API asynchronously
func (client *Client) DescribeInstanceLinkedWhitelistTemplateWithCallback(request *DescribeInstanceLinkedWhitelistTemplateRequest, callback func(response *DescribeInstanceLinkedWhitelistTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceLinkedWhitelistTemplateResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceLinkedWhitelistTemplate(request)
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

// DescribeInstanceLinkedWhitelistTemplateRequest is the request struct for api DescribeInstanceLinkedWhitelistTemplate
type DescribeInstanceLinkedWhitelistTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	InsName              string           `position:"Query" name:"InsName"`
}

// DescribeInstanceLinkedWhitelistTemplateResponse is the response struct for api DescribeInstanceLinkedWhitelistTemplate
type DescribeInstanceLinkedWhitelistTemplateResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateDescribeInstanceLinkedWhitelistTemplateRequest creates a request to invoke DescribeInstanceLinkedWhitelistTemplate API
func CreateDescribeInstanceLinkedWhitelistTemplateRequest() (request *DescribeInstanceLinkedWhitelistTemplateRequest) {
	request = &DescribeInstanceLinkedWhitelistTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeInstanceLinkedWhitelistTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceLinkedWhitelistTemplateResponse creates a response to parse from DescribeInstanceLinkedWhitelistTemplate response
func CreateDescribeInstanceLinkedWhitelistTemplateResponse() (response *DescribeInstanceLinkedWhitelistTemplateResponse) {
	response = &DescribeInstanceLinkedWhitelistTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
