package sas

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

// DescribeInstanceAntiBruteForceRules invokes the sas.DescribeInstanceAntiBruteForceRules API synchronously
func (client *Client) DescribeInstanceAntiBruteForceRules(request *DescribeInstanceAntiBruteForceRulesRequest) (response *DescribeInstanceAntiBruteForceRulesResponse, err error) {
	response = CreateDescribeInstanceAntiBruteForceRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceAntiBruteForceRulesWithChan invokes the sas.DescribeInstanceAntiBruteForceRules API asynchronously
func (client *Client) DescribeInstanceAntiBruteForceRulesWithChan(request *DescribeInstanceAntiBruteForceRulesRequest) (<-chan *DescribeInstanceAntiBruteForceRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceAntiBruteForceRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceAntiBruteForceRules(request)
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

// DescribeInstanceAntiBruteForceRulesWithCallback invokes the sas.DescribeInstanceAntiBruteForceRules API asynchronously
func (client *Client) DescribeInstanceAntiBruteForceRulesWithCallback(request *DescribeInstanceAntiBruteForceRulesRequest, callback func(response *DescribeInstanceAntiBruteForceRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceAntiBruteForceRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceAntiBruteForceRules(request)
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

// DescribeInstanceAntiBruteForceRulesRequest is the request struct for api DescribeInstanceAntiBruteForceRules
type DescribeInstanceAntiBruteForceRulesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	UuidList        *[]string        `position:"Query" name:"UuidList"  type:"Repeated"`
}

// DescribeInstanceAntiBruteForceRulesResponse is the response struct for api DescribeInstanceAntiBruteForceRules
type DescribeInstanceAntiBruteForceRulesResponse struct {
	*responses.BaseResponse
	RequestId string                       `json:"RequestId" xml:"RequestId"`
	PageInfo  PageInfo                     `json:"PageInfo" xml:"PageInfo"`
	Rules     []InstanceAntiBruteForceRule `json:"Rules" xml:"Rules"`
}

// CreateDescribeInstanceAntiBruteForceRulesRequest creates a request to invoke DescribeInstanceAntiBruteForceRules API
func CreateDescribeInstanceAntiBruteForceRulesRequest() (request *DescribeInstanceAntiBruteForceRulesRequest) {
	request = &DescribeInstanceAntiBruteForceRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeInstanceAntiBruteForceRules", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceAntiBruteForceRulesResponse creates a response to parse from DescribeInstanceAntiBruteForceRules response
func CreateDescribeInstanceAntiBruteForceRulesResponse() (response *DescribeInstanceAntiBruteForceRulesResponse) {
	response = &DescribeInstanceAntiBruteForceRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
