package sae

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

// DescribeConfigMap invokes the sae.DescribeConfigMap API synchronously
func (client *Client) DescribeConfigMap(request *DescribeConfigMapRequest) (response *DescribeConfigMapResponse, err error) {
	response = CreateDescribeConfigMapResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConfigMapWithChan invokes the sae.DescribeConfigMap API asynchronously
func (client *Client) DescribeConfigMapWithChan(request *DescribeConfigMapRequest) (<-chan *DescribeConfigMapResponse, <-chan error) {
	responseChan := make(chan *DescribeConfigMapResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConfigMap(request)
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

// DescribeConfigMapWithCallback invokes the sae.DescribeConfigMap API asynchronously
func (client *Client) DescribeConfigMapWithCallback(request *DescribeConfigMapRequest, callback func(response *DescribeConfigMapResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConfigMapResponse
		var err error
		defer close(result)
		response, err = client.DescribeConfigMap(request)
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

// DescribeConfigMapRequest is the request struct for api DescribeConfigMap
type DescribeConfigMapRequest struct {
	*requests.RoaRequest
	ConfigMapId requests.Integer `position:"Query" name:"ConfigMapId"`
}

// DescribeConfigMapResponse is the response struct for api DescribeConfigMap
type DescribeConfigMapResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeConfigMapRequest creates a request to invoke DescribeConfigMap API
func CreateDescribeConfigMapRequest() (request *DescribeConfigMapRequest) {
	request = &DescribeConfigMapRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeConfigMap", "/pop/v1/sam/configmap/configMap", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeConfigMapResponse creates a response to parse from DescribeConfigMap response
func CreateDescribeConfigMapResponse() (response *DescribeConfigMapResponse) {
	response = &DescribeConfigMapResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}