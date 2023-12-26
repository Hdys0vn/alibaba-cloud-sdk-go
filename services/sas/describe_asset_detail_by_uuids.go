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

// DescribeAssetDetailByUuids invokes the sas.DescribeAssetDetailByUuids API synchronously
func (client *Client) DescribeAssetDetailByUuids(request *DescribeAssetDetailByUuidsRequest) (response *DescribeAssetDetailByUuidsResponse, err error) {
	response = CreateDescribeAssetDetailByUuidsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAssetDetailByUuidsWithChan invokes the sas.DescribeAssetDetailByUuids API asynchronously
func (client *Client) DescribeAssetDetailByUuidsWithChan(request *DescribeAssetDetailByUuidsRequest) (<-chan *DescribeAssetDetailByUuidsResponse, <-chan error) {
	responseChan := make(chan *DescribeAssetDetailByUuidsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAssetDetailByUuids(request)
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

// DescribeAssetDetailByUuidsWithCallback invokes the sas.DescribeAssetDetailByUuids API asynchronously
func (client *Client) DescribeAssetDetailByUuidsWithCallback(request *DescribeAssetDetailByUuidsRequest, callback func(response *DescribeAssetDetailByUuidsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAssetDetailByUuidsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAssetDetailByUuids(request)
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

// DescribeAssetDetailByUuidsRequest is the request struct for api DescribeAssetDetailByUuids
type DescribeAssetDetailByUuidsRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Uuids    string `position:"Query" name:"Uuids"`
}

// DescribeAssetDetailByUuidsResponse is the response struct for api DescribeAssetDetailByUuids
type DescribeAssetDetailByUuidsResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	AssetList []Asset `json:"AssetList" xml:"AssetList"`
}

// CreateDescribeAssetDetailByUuidsRequest creates a request to invoke DescribeAssetDetailByUuids API
func CreateDescribeAssetDetailByUuidsRequest() (request *DescribeAssetDetailByUuidsRequest) {
	request = &DescribeAssetDetailByUuidsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeAssetDetailByUuids", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAssetDetailByUuidsResponse creates a response to parse from DescribeAssetDetailByUuids response
func CreateDescribeAssetDetailByUuidsResponse() (response *DescribeAssetDetailByUuidsResponse) {
	response = &DescribeAssetDetailByUuidsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}