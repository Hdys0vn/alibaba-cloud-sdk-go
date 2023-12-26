package sas_api

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

// GetAccountProfile invokes the sas_api.GetAccountProfile API synchronously
// api document: https://help.aliyun.com/api/sas-api/getaccountprofile.html
func (client *Client) GetAccountProfile(request *GetAccountProfileRequest) (response *GetAccountProfileResponse, err error) {
	response = CreateGetAccountProfileResponse()
	err = client.DoAction(request, response)
	return
}

// GetAccountProfileWithChan invokes the sas_api.GetAccountProfile API asynchronously
// api document: https://help.aliyun.com/api/sas-api/getaccountprofile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAccountProfileWithChan(request *GetAccountProfileRequest) (<-chan *GetAccountProfileResponse, <-chan error) {
	responseChan := make(chan *GetAccountProfileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccountProfile(request)
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

// GetAccountProfileWithCallback invokes the sas_api.GetAccountProfile API asynchronously
// api document: https://help.aliyun.com/api/sas-api/getaccountprofile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAccountProfileWithCallback(request *GetAccountProfileRequest, callback func(response *GetAccountProfileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAccountProfileResponse
		var err error
		defer close(result)
		response, err = client.GetAccountProfile(request)
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

// GetAccountProfileRequest is the request struct for api GetAccountProfile
type GetAccountProfileRequest struct {
	*requests.RpcRequest
	DeviceIdMd5     string           `position:"Query" name:"DeviceIdMd5"`
	Carrier         requests.Integer `position:"Query" name:"Carrier"`
	Os              string           `position:"Query" name:"Os"`
	Phone           string           `position:"Query" name:"Phone"`
	RequestUrl      string           `position:"Query" name:"RequestUrl"`
	Ip              string           `position:"Query" name:"Ip"`
	UserAgent       string           `position:"Query" name:"UserAgent"`
	ConnectionType  requests.Integer `position:"Query" name:"ConnectionType"`
	SensType        requests.Integer `position:"Query" name:"SensType"`
	DeviceType      requests.Integer `position:"Query" name:"DeviceType"`
	AccessTimestamp requests.Integer `position:"Query" name:"AccessTimestamp"`
	BusinessType    requests.Integer `position:"Query" name:"BusinessType"`
}

// GetAccountProfileResponse is the response struct for api GetAccountProfile
type GetAccountProfileResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetAccountProfileRequest creates a request to invoke GetAccountProfile API
func CreateGetAccountProfileRequest() (request *GetAccountProfileRequest) {
	request = &GetAccountProfileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas-api", "2017-07-05", "GetAccountProfile", "sas-api", "openAPI")
	return
}

// CreateGetAccountProfileResponse creates a response to parse from GetAccountProfile response
func CreateGetAccountProfileResponse() (response *GetAccountProfileResponse) {
	response = &GetAccountProfileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}