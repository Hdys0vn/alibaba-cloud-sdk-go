package dds

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

// DescribeUserEncryptionKeyList invokes the dds.DescribeUserEncryptionKeyList API synchronously
func (client *Client) DescribeUserEncryptionKeyList(request *DescribeUserEncryptionKeyListRequest) (response *DescribeUserEncryptionKeyListResponse, err error) {
	response = CreateDescribeUserEncryptionKeyListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserEncryptionKeyListWithChan invokes the dds.DescribeUserEncryptionKeyList API asynchronously
func (client *Client) DescribeUserEncryptionKeyListWithChan(request *DescribeUserEncryptionKeyListRequest) (<-chan *DescribeUserEncryptionKeyListResponse, <-chan error) {
	responseChan := make(chan *DescribeUserEncryptionKeyListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserEncryptionKeyList(request)
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

// DescribeUserEncryptionKeyListWithCallback invokes the dds.DescribeUserEncryptionKeyList API asynchronously
func (client *Client) DescribeUserEncryptionKeyListWithCallback(request *DescribeUserEncryptionKeyListRequest, callback func(response *DescribeUserEncryptionKeyListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserEncryptionKeyListResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserEncryptionKeyList(request)
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

// DescribeUserEncryptionKeyListRequest is the request struct for api DescribeUserEncryptionKeyList
type DescribeUserEncryptionKeyListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TargetRegionId       string           `position:"Query" name:"TargetRegionId"`
}

// DescribeUserEncryptionKeyListResponse is the response struct for api DescribeUserEncryptionKeyList
type DescribeUserEncryptionKeyListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	KeyIds    KeyIds `json:"KeyIds" xml:"KeyIds"`
}

// CreateDescribeUserEncryptionKeyListRequest creates a request to invoke DescribeUserEncryptionKeyList API
func CreateDescribeUserEncryptionKeyListRequest() (request *DescribeUserEncryptionKeyListRequest) {
	request = &DescribeUserEncryptionKeyListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeUserEncryptionKeyList", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUserEncryptionKeyListResponse creates a response to parse from DescribeUserEncryptionKeyList response
func CreateDescribeUserEncryptionKeyListResponse() (response *DescribeUserEncryptionKeyListResponse) {
	response = &DescribeUserEncryptionKeyListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
