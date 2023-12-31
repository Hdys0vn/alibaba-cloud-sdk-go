package afs

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

// DescribeCaptchaDay invokes the afs.DescribeCaptchaDay API synchronously
// api document: https://help.aliyun.com/api/afs/describecaptchaday.html
func (client *Client) DescribeCaptchaDay(request *DescribeCaptchaDayRequest) (response *DescribeCaptchaDayResponse, err error) {
	response = CreateDescribeCaptchaDayResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCaptchaDayWithChan invokes the afs.DescribeCaptchaDay API asynchronously
// api document: https://help.aliyun.com/api/afs/describecaptchaday.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCaptchaDayWithChan(request *DescribeCaptchaDayRequest) (<-chan *DescribeCaptchaDayResponse, <-chan error) {
	responseChan := make(chan *DescribeCaptchaDayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCaptchaDay(request)
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

// DescribeCaptchaDayWithCallback invokes the afs.DescribeCaptchaDay API asynchronously
// api document: https://help.aliyun.com/api/afs/describecaptchaday.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCaptchaDayWithCallback(request *DescribeCaptchaDayRequest, callback func(response *DescribeCaptchaDayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCaptchaDayResponse
		var err error
		defer close(result)
		response, err = client.DescribeCaptchaDay(request)
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

// DescribeCaptchaDayRequest is the request struct for api DescribeCaptchaDay
type DescribeCaptchaDayRequest struct {
	*requests.RpcRequest
	SourceIp   string `position:"Query" name:"SourceIp"`
	ConfigName string `position:"Query" name:"ConfigName"`
	RefExtId   string `position:"Query" name:"RefExtId"`
	Time       string `position:"Query" name:"Time"`
	Type       string `position:"Query" name:"Type"`
}

// DescribeCaptchaDayResponse is the response struct for api DescribeCaptchaDay
type DescribeCaptchaDayResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	BizCode    string     `json:"BizCode" xml:"BizCode"`
	HasData    bool       `json:"HasData" xml:"HasData"`
	CaptchaDay CaptchaDay `json:"CaptchaDay" xml:"CaptchaDay"`
}

// CreateDescribeCaptchaDayRequest creates a request to invoke DescribeCaptchaDay API
func CreateDescribeCaptchaDayRequest() (request *DescribeCaptchaDayRequest) {
	request = &DescribeCaptchaDayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("afs", "2018-01-12", "DescribeCaptchaDay", "afs", "openAPI")
	return
}

// CreateDescribeCaptchaDayResponse creates a response to parse from DescribeCaptchaDay response
func CreateDescribeCaptchaDayResponse() (response *DescribeCaptchaDayResponse) {
	response = &DescribeCaptchaDayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
