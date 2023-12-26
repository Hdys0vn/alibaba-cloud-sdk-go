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

// BusinessCreate invokes the cloudwf.BusinessCreate API synchronously
// api document: https://help.aliyun.com/api/cloudwf/businesscreate.html
func (client *Client) BusinessCreate(request *BusinessCreateRequest) (response *BusinessCreateResponse, err error) {
	response = CreateBusinessCreateResponse()
	err = client.DoAction(request, response)
	return
}

// BusinessCreateWithChan invokes the cloudwf.BusinessCreate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/businesscreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BusinessCreateWithChan(request *BusinessCreateRequest) (<-chan *BusinessCreateResponse, <-chan error) {
	responseChan := make(chan *BusinessCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BusinessCreate(request)
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

// BusinessCreateWithCallback invokes the cloudwf.BusinessCreate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/businesscreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BusinessCreateWithCallback(request *BusinessCreateRequest, callback func(response *BusinessCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BusinessCreateResponse
		var err error
		defer close(result)
		response, err = client.BusinessCreate(request)
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

// BusinessCreateRequest is the request struct for api BusinessCreate
type BusinessCreateRequest struct {
	*requests.RpcRequest
	BusinessCity     string           `position:"Query" name:"BusinessCity"`
	Combo            string           `position:"Query" name:"Combo"`
	WarnEmail        string           `position:"Query" name:"WarnEmail"`
	BusinessManager  string           `position:"Query" name:"BusinessManager"`
	BusinessType     requests.Integer `position:"Query" name:"BusinessType"`
	Warn             requests.Integer `position:"Query" name:"Warn"`
	BusinessName     string           `position:"Query" name:"BusinessName"`
	BusinessTopType  requests.Integer `position:"Query" name:"BusinessTopType"`
	BusinessAddress  string           `position:"Query" name:"BusinessAddress"`
	BusinessTel      string           `position:"Query" name:"BusinessTel"`
	BusinessProvince string           `position:"Query" name:"BusinessProvince"`
	BusinessSubtype  requests.Integer `position:"Query" name:"BusinessSubtype"`
}

// BusinessCreateResponse is the response struct for api BusinessCreate
type BusinessCreateResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateBusinessCreateRequest creates a request to invoke BusinessCreate API
func CreateBusinessCreateRequest() (request *BusinessCreateRequest) {
	request = &BusinessCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "BusinessCreate", "cloudwf", "openAPI")
	return
}

// CreateBusinessCreateResponse creates a response to parse from BusinessCreate response
func CreateBusinessCreateResponse() (response *BusinessCreateResponse) {
	response = &BusinessCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}