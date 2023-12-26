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

// DescribeExposedInstanceCriteria invokes the sas.DescribeExposedInstanceCriteria API synchronously
func (client *Client) DescribeExposedInstanceCriteria(request *DescribeExposedInstanceCriteriaRequest) (response *DescribeExposedInstanceCriteriaResponse, err error) {
	response = CreateDescribeExposedInstanceCriteriaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeExposedInstanceCriteriaWithChan invokes the sas.DescribeExposedInstanceCriteria API asynchronously
func (client *Client) DescribeExposedInstanceCriteriaWithChan(request *DescribeExposedInstanceCriteriaRequest) (<-chan *DescribeExposedInstanceCriteriaResponse, <-chan error) {
	responseChan := make(chan *DescribeExposedInstanceCriteriaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeExposedInstanceCriteria(request)
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

// DescribeExposedInstanceCriteriaWithCallback invokes the sas.DescribeExposedInstanceCriteria API asynchronously
func (client *Client) DescribeExposedInstanceCriteriaWithCallback(request *DescribeExposedInstanceCriteriaRequest, callback func(response *DescribeExposedInstanceCriteriaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeExposedInstanceCriteriaResponse
		var err error
		defer close(result)
		response, err = client.DescribeExposedInstanceCriteria(request)
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

// DescribeExposedInstanceCriteriaRequest is the request struct for api DescribeExposedInstanceCriteria
type DescribeExposedInstanceCriteriaRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Value    string `position:"Query" name:"Value"`
}

// DescribeExposedInstanceCriteriaResponse is the response struct for api DescribeExposedInstanceCriteria
type DescribeExposedInstanceCriteriaResponse struct {
	*responses.BaseResponse
	RequestId    string     `json:"RequestId" xml:"RequestId"`
	CriteriaList []Criteria `json:"CriteriaList" xml:"CriteriaList"`
}

// CreateDescribeExposedInstanceCriteriaRequest creates a request to invoke DescribeExposedInstanceCriteria API
func CreateDescribeExposedInstanceCriteriaRequest() (request *DescribeExposedInstanceCriteriaRequest) {
	request = &DescribeExposedInstanceCriteriaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeExposedInstanceCriteria", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeExposedInstanceCriteriaResponse creates a response to parse from DescribeExposedInstanceCriteria response
func CreateDescribeExposedInstanceCriteriaResponse() (response *DescribeExposedInstanceCriteriaResponse) {
	response = &DescribeExposedInstanceCriteriaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}