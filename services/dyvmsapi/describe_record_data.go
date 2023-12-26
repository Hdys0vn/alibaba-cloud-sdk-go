package dyvmsapi

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

// DescribeRecordData invokes the dyvmsapi.DescribeRecordData API synchronously
func (client *Client) DescribeRecordData(request *DescribeRecordDataRequest) (response *DescribeRecordDataResponse, err error) {
	response = CreateDescribeRecordDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecordDataWithChan invokes the dyvmsapi.DescribeRecordData API asynchronously
func (client *Client) DescribeRecordDataWithChan(request *DescribeRecordDataRequest) (<-chan *DescribeRecordDataResponse, <-chan error) {
	responseChan := make(chan *DescribeRecordDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecordData(request)
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

// DescribeRecordDataWithCallback invokes the dyvmsapi.DescribeRecordData API asynchronously
func (client *Client) DescribeRecordDataWithCallback(request *DescribeRecordDataRequest, callback func(response *DescribeRecordDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecordDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecordData(request)
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

// DescribeRecordDataRequest is the request struct for api DescribeRecordData
type DescribeRecordDataRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountType          string           `position:"Query" name:"AccountType"`
	Acid                 string           `position:"Query" name:"Acid"`
	AccountId            string           `position:"Query" name:"AccountId"`
	SecLevel             requests.Integer `position:"Query" name:"SecLevel"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ProdCode             string           `position:"Query" name:"ProdCode"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeRecordDataResponse is the response struct for api DescribeRecordData
type DescribeRecordDataResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	OssLink   string `json:"OssLink" xml:"OssLink"`
	Acid      string `json:"Acid" xml:"Acid"`
	AgentId   string `json:"AgentId" xml:"AgentId"`
}

// CreateDescribeRecordDataRequest creates a request to invoke DescribeRecordData API
func CreateDescribeRecordDataRequest() (request *DescribeRecordDataRequest) {
	request = &DescribeRecordDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "DescribeRecordData", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRecordDataResponse creates a response to parse from DescribeRecordData response
func CreateDescribeRecordDataResponse() (response *DescribeRecordDataResponse) {
	response = &DescribeRecordDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
