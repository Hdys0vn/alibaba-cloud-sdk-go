package cloudesl

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

// DescribeStoreByTemplateVersion invokes the cloudesl.DescribeStoreByTemplateVersion API synchronously
func (client *Client) DescribeStoreByTemplateVersion(request *DescribeStoreByTemplateVersionRequest) (response *DescribeStoreByTemplateVersionResponse, err error) {
	response = CreateDescribeStoreByTemplateVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStoreByTemplateVersionWithChan invokes the cloudesl.DescribeStoreByTemplateVersion API asynchronously
func (client *Client) DescribeStoreByTemplateVersionWithChan(request *DescribeStoreByTemplateVersionRequest) (<-chan *DescribeStoreByTemplateVersionResponse, <-chan error) {
	responseChan := make(chan *DescribeStoreByTemplateVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStoreByTemplateVersion(request)
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

// DescribeStoreByTemplateVersionWithCallback invokes the cloudesl.DescribeStoreByTemplateVersion API asynchronously
func (client *Client) DescribeStoreByTemplateVersionWithCallback(request *DescribeStoreByTemplateVersionRequest, callback func(response *DescribeStoreByTemplateVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStoreByTemplateVersionResponse
		var err error
		defer close(result)
		response, err = client.DescribeStoreByTemplateVersion(request)
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

// DescribeStoreByTemplateVersionRequest is the request struct for api DescribeStoreByTemplateVersion
type DescribeStoreByTemplateVersionRequest struct {
	*requests.RpcRequest
	TemplateVersion string `position:"Body" name:"TemplateVersion"`
}

// DescribeStoreByTemplateVersionResponse is the response struct for api DescribeStoreByTemplateVersion
type DescribeStoreByTemplateVersionResponse struct {
	*responses.BaseResponse
	RequestId      string           `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string           `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool             `json:"Success" xml:"Success"`
	ErrorCode      string           `json:"ErrorCode" xml:"ErrorCode"`
	Code           string           `json:"Code" xml:"Code"`
	Message        string           `json:"Message" xml:"Message"`
	DynamicMessage string           `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode    string           `json:"DynamicCode" xml:"DynamicCode"`
	Stores         []SelectItemInfo `json:"Stores" xml:"Stores"`
}

// CreateDescribeStoreByTemplateVersionRequest creates a request to invoke DescribeStoreByTemplateVersion API
func CreateDescribeStoreByTemplateVersionRequest() (request *DescribeStoreByTemplateVersionRequest) {
	request = &DescribeStoreByTemplateVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribeStoreByTemplateVersion", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeStoreByTemplateVersionResponse creates a response to parse from DescribeStoreByTemplateVersion response
func CreateDescribeStoreByTemplateVersionResponse() (response *DescribeStoreByTemplateVersionResponse) {
	response = &DescribeStoreByTemplateVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}