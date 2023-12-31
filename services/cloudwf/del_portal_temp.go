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

// DelPortalTemp invokes the cloudwf.DelPortalTemp API synchronously
// api document: https://help.aliyun.com/api/cloudwf/delportaltemp.html
func (client *Client) DelPortalTemp(request *DelPortalTempRequest) (response *DelPortalTempResponse, err error) {
	response = CreateDelPortalTempResponse()
	err = client.DoAction(request, response)
	return
}

// DelPortalTempWithChan invokes the cloudwf.DelPortalTemp API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/delportaltemp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DelPortalTempWithChan(request *DelPortalTempRequest) (<-chan *DelPortalTempResponse, <-chan error) {
	responseChan := make(chan *DelPortalTempResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DelPortalTemp(request)
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

// DelPortalTempWithCallback invokes the cloudwf.DelPortalTemp API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/delportaltemp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DelPortalTempWithCallback(request *DelPortalTempRequest, callback func(response *DelPortalTempResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DelPortalTempResponse
		var err error
		defer close(result)
		response, err = client.DelPortalTemp(request)
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

// DelPortalTempRequest is the request struct for api DelPortalTemp
type DelPortalTempRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// DelPortalTempResponse is the response struct for api DelPortalTemp
type DelPortalTempResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateDelPortalTempRequest creates a request to invoke DelPortalTemp API
func CreateDelPortalTempRequest() (request *DelPortalTempRequest) {
	request = &DelPortalTempRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "DelPortalTemp", "cloudwf", "openAPI")
	return
}

// CreateDelPortalTempResponse creates a response to parse from DelPortalTemp response
func CreateDelPortalTempResponse() (response *DelPortalTempResponse) {
	response = &DelPortalTempResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
