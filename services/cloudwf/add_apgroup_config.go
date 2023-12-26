package cloudwf

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

// AddApgroupConfig invokes the cloudwf.AddApgroupConfig API synchronously
// api document: https://help.aliyun.com/api/cloudwf/addapgroupconfig.html
func (client *Client) AddApgroupConfig(request *AddApgroupConfigRequest) (response *AddApgroupConfigResponse, err error) {
	response = CreateAddApgroupConfigResponse()
	err = client.DoAction(request, response)
	return
}

// AddApgroupConfigWithChan invokes the cloudwf.AddApgroupConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/addapgroupconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddApgroupConfigWithChan(request *AddApgroupConfigRequest) (<-chan *AddApgroupConfigResponse, <-chan error) {
	responseChan := make(chan *AddApgroupConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddApgroupConfig(request)
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

// AddApgroupConfigWithCallback invokes the cloudwf.AddApgroupConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/addapgroupconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddApgroupConfigWithCallback(request *AddApgroupConfigRequest, callback func(response *AddApgroupConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddApgroupConfigResponse
		var err error
		defer close(result)
		response, err = client.AddApgroupConfig(request)
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

// AddApgroupConfigRequest is the request struct for api AddApgroupConfig
type AddApgroupConfigRequest struct {
	*requests.RpcRequest
	ParentApgroupId requests.Integer `position:"Query" name:"ParentApgroupId"`
	Name            string           `position:"Query" name:"Name"`
	Description     string           `position:"Query" name:"Description"`
}

// AddApgroupConfigResponse is the response struct for api AddApgroupConfig
type AddApgroupConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateAddApgroupConfigRequest creates a request to invoke AddApgroupConfig API
func CreateAddApgroupConfigRequest() (request *AddApgroupConfigRequest) {
	request = &AddApgroupConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "AddApgroupConfig", "cloudwf", "openAPI")
	return
}

// CreateAddApgroupConfigResponse creates a response to parse from AddApgroupConfig response
func CreateAddApgroupConfigResponse() (response *AddApgroupConfigResponse) {
	response = &AddApgroupConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
