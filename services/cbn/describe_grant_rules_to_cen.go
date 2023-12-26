package cbn

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

// DescribeGrantRulesToCen invokes the cbn.DescribeGrantRulesToCen API synchronously
func (client *Client) DescribeGrantRulesToCen(request *DescribeGrantRulesToCenRequest) (response *DescribeGrantRulesToCenResponse, err error) {
	response = CreateDescribeGrantRulesToCenResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGrantRulesToCenWithChan invokes the cbn.DescribeGrantRulesToCen API asynchronously
func (client *Client) DescribeGrantRulesToCenWithChan(request *DescribeGrantRulesToCenRequest) (<-chan *DescribeGrantRulesToCenResponse, <-chan error) {
	responseChan := make(chan *DescribeGrantRulesToCenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGrantRulesToCen(request)
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

// DescribeGrantRulesToCenWithCallback invokes the cbn.DescribeGrantRulesToCen API asynchronously
func (client *Client) DescribeGrantRulesToCenWithCallback(request *DescribeGrantRulesToCenRequest, callback func(response *DescribeGrantRulesToCenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGrantRulesToCenResponse
		var err error
		defer close(result)
		response, err = client.DescribeGrantRulesToCen(request)
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

// DescribeGrantRulesToCenRequest is the request struct for api DescribeGrantRulesToCen
type DescribeGrantRulesToCenRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	ProductType          string           `position:"Query" name:"ProductType"`
	NextToken            string           `position:"Query" name:"NextToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Version              string           `position:"Query" name:"Version"`
	ChildInstanceOwnerId requests.Integer `position:"Query" name:"ChildInstanceOwnerId"`
	ChildInstanceId      string           `position:"Query" name:"ChildInstanceId"`
	MaxResults           requests.Integer `position:"Query" name:"MaxResults"`
}

// DescribeGrantRulesToCenResponse is the response struct for api DescribeGrantRulesToCen
type DescribeGrantRulesToCenResponse struct {
	*responses.BaseResponse
	RequestId  string                              `json:"RequestId" xml:"RequestId"`
	TotalCount int64                               `json:"TotalCount" xml:"TotalCount"`
	MaxResults int64                               `json:"MaxResults" xml:"MaxResults"`
	NextToken  string                              `json:"NextToken" xml:"NextToken"`
	GrantRules GrantRulesInDescribeGrantRulesToCen `json:"GrantRules" xml:"GrantRules"`
}

// CreateDescribeGrantRulesToCenRequest creates a request to invoke DescribeGrantRulesToCen API
func CreateDescribeGrantRulesToCenRequest() (request *DescribeGrantRulesToCenRequest) {
	request = &DescribeGrantRulesToCenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DescribeGrantRulesToCen", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeGrantRulesToCenResponse creates a response to parse from DescribeGrantRulesToCen response
func CreateDescribeGrantRulesToCenResponse() (response *DescribeGrantRulesToCenResponse) {
	response = &DescribeGrantRulesToCenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
