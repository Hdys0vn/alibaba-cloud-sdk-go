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

// GetDeviceInfoByMac invokes the cloudwf.GetDeviceInfoByMac API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getdeviceinfobymac.html
func (client *Client) GetDeviceInfoByMac(request *GetDeviceInfoByMacRequest) (response *GetDeviceInfoByMacResponse, err error) {
	response = CreateGetDeviceInfoByMacResponse()
	err = client.DoAction(request, response)
	return
}

// GetDeviceInfoByMacWithChan invokes the cloudwf.GetDeviceInfoByMac API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getdeviceinfobymac.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceInfoByMacWithChan(request *GetDeviceInfoByMacRequest) (<-chan *GetDeviceInfoByMacResponse, <-chan error) {
	responseChan := make(chan *GetDeviceInfoByMacResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDeviceInfoByMac(request)
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

// GetDeviceInfoByMacWithCallback invokes the cloudwf.GetDeviceInfoByMac API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getdeviceinfobymac.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceInfoByMacWithCallback(request *GetDeviceInfoByMacRequest, callback func(response *GetDeviceInfoByMacResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDeviceInfoByMacResponse
		var err error
		defer close(result)
		response, err = client.GetDeviceInfoByMac(request)
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

// GetDeviceInfoByMacRequest is the request struct for api GetDeviceInfoByMac
type GetDeviceInfoByMacRequest struct {
	*requests.RpcRequest
	Mac string `position:"Query" name:"Mac"`
}

// GetDeviceInfoByMacResponse is the response struct for api GetDeviceInfoByMac
type GetDeviceInfoByMacResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetDeviceInfoByMacRequest creates a request to invoke GetDeviceInfoByMac API
func CreateGetDeviceInfoByMacRequest() (request *GetDeviceInfoByMacRequest) {
	request = &GetDeviceInfoByMacRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetDeviceInfoByMac", "cloudwf", "openAPI")
	return
}

// CreateGetDeviceInfoByMacResponse creates a response to parse from GetDeviceInfoByMac response
func CreateGetDeviceInfoByMacResponse() (response *GetDeviceInfoByMacResponse) {
	response = &GetDeviceInfoByMacResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
