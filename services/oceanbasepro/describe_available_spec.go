package oceanbasepro

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

// DescribeAvailableSpec invokes the oceanbasepro.DescribeAvailableSpec API synchronously
func (client *Client) DescribeAvailableSpec(request *DescribeAvailableSpecRequest) (response *DescribeAvailableSpecResponse, err error) {
	response = CreateDescribeAvailableSpecResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvailableSpecWithChan invokes the oceanbasepro.DescribeAvailableSpec API asynchronously
func (client *Client) DescribeAvailableSpecWithChan(request *DescribeAvailableSpecRequest) (<-chan *DescribeAvailableSpecResponse, <-chan error) {
	responseChan := make(chan *DescribeAvailableSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvailableSpec(request)
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

// DescribeAvailableSpecWithCallback invokes the oceanbasepro.DescribeAvailableSpec API asynchronously
func (client *Client) DescribeAvailableSpecWithCallback(request *DescribeAvailableSpecRequest, callback func(response *DescribeAvailableSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvailableSpecResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvailableSpec(request)
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

// DescribeAvailableSpecRequest is the request struct for api DescribeAvailableSpec
type DescribeAvailableSpecRequest struct {
	*requests.RpcRequest
	UpgradeType string `position:"Body" name:"UpgradeType"`
	Spec        string `position:"Body" name:"Spec"`
	InstanceId  string `position:"Body" name:"InstanceId"`
}

// DescribeAvailableSpecResponse is the response struct for api DescribeAvailableSpec
type DescribeAvailableSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeAvailableSpecRequest creates a request to invoke DescribeAvailableSpec API
func CreateDescribeAvailableSpecRequest() (request *DescribeAvailableSpecRequest) {
	request = &DescribeAvailableSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeAvailableSpec", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAvailableSpecResponse creates a response to parse from DescribeAvailableSpec response
func CreateDescribeAvailableSpecResponse() (response *DescribeAvailableSpecResponse) {
	response = &DescribeAvailableSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}