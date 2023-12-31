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

// OemMarketingPotential invokes the cloudwf.OemMarketingPotential API synchronously
// api document: https://help.aliyun.com/api/cloudwf/oemmarketingpotential.html
func (client *Client) OemMarketingPotential(request *OemMarketingPotentialRequest) (response *OemMarketingPotentialResponse, err error) {
	response = CreateOemMarketingPotentialResponse()
	err = client.DoAction(request, response)
	return
}

// OemMarketingPotentialWithChan invokes the cloudwf.OemMarketingPotential API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/oemmarketingpotential.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OemMarketingPotentialWithChan(request *OemMarketingPotentialRequest) (<-chan *OemMarketingPotentialResponse, <-chan error) {
	responseChan := make(chan *OemMarketingPotentialResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OemMarketingPotential(request)
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

// OemMarketingPotentialWithCallback invokes the cloudwf.OemMarketingPotential API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/oemmarketingpotential.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OemMarketingPotentialWithCallback(request *OemMarketingPotentialRequest, callback func(response *OemMarketingPotentialResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OemMarketingPotentialResponse
		var err error
		defer close(result)
		response, err = client.OemMarketingPotential(request)
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

// OemMarketingPotentialRequest is the request struct for api OemMarketingPotential
type OemMarketingPotentialRequest struct {
	*requests.RpcRequest
	Bid requests.Integer `position:"Query" name:"Bid"`
}

// OemMarketingPotentialResponse is the response struct for api OemMarketingPotential
type OemMarketingPotentialResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateOemMarketingPotentialRequest creates a request to invoke OemMarketingPotential API
func CreateOemMarketingPotentialRequest() (request *OemMarketingPotentialRequest) {
	request = &OemMarketingPotentialRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "OemMarketingPotential", "cloudwf", "openAPI")
	return
}

// CreateOemMarketingPotentialResponse creates a response to parse from OemMarketingPotential response
func CreateOemMarketingPotentialResponse() (response *OemMarketingPotentialResponse) {
	response = &OemMarketingPotentialResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
