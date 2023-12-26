package companyreg

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

// OperateProduceForPartner invokes the companyreg.OperateProduceForPartner API synchronously
func (client *Client) OperateProduceForPartner(request *OperateProduceForPartnerRequest) (response *OperateProduceForPartnerResponse, err error) {
	response = CreateOperateProduceForPartnerResponse()
	err = client.DoAction(request, response)
	return
}

// OperateProduceForPartnerWithChan invokes the companyreg.OperateProduceForPartner API asynchronously
func (client *Client) OperateProduceForPartnerWithChan(request *OperateProduceForPartnerRequest) (<-chan *OperateProduceForPartnerResponse, <-chan error) {
	responseChan := make(chan *OperateProduceForPartnerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OperateProduceForPartner(request)
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

// OperateProduceForPartnerWithCallback invokes the companyreg.OperateProduceForPartner API asynchronously
func (client *Client) OperateProduceForPartnerWithCallback(request *OperateProduceForPartnerRequest, callback func(response *OperateProduceForPartnerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OperateProduceForPartnerResponse
		var err error
		defer close(result)
		response, err = client.OperateProduceForPartner(request)
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

// OperateProduceForPartnerRequest is the request struct for api OperateProduceForPartner
type OperateProduceForPartnerRequest struct {
	*requests.RpcRequest
	BizType     string `position:"Query" name:"BizType"`
	ExtInfo     string `position:"Query" name:"ExtInfo"`
	BizId       string `position:"Query" name:"BizId"`
	OperateType string `position:"Query" name:"OperateType"`
}

// OperateProduceForPartnerResponse is the response struct for api OperateProduceForPartner
type OperateProduceForPartnerResponse struct {
	*responses.BaseResponse
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateOperateProduceForPartnerRequest creates a request to invoke OperateProduceForPartner API
func CreateOperateProduceForPartnerRequest() (request *OperateProduceForPartnerRequest) {
	request = &OperateProduceForPartnerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "OperateProduceForPartner", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOperateProduceForPartnerResponse creates a response to parse from OperateProduceForPartner response
func CreateOperateProduceForPartnerResponse() (response *OperateProduceForPartnerResponse) {
	response = &OperateProduceForPartnerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
