package trademark

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

// CombineLoa invokes the trademark.CombineLoa API synchronously
// api document: https://help.aliyun.com/api/trademark/combineloa.html
func (client *Client) CombineLoa(request *CombineLoaRequest) (response *CombineLoaResponse, err error) {
	response = CreateCombineLoaResponse()
	err = client.DoAction(request, response)
	return
}

// CombineLoaWithChan invokes the trademark.CombineLoa API asynchronously
// api document: https://help.aliyun.com/api/trademark/combineloa.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CombineLoaWithChan(request *CombineLoaRequest) (<-chan *CombineLoaResponse, <-chan error) {
	responseChan := make(chan *CombineLoaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CombineLoa(request)
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

// CombineLoaWithCallback invokes the trademark.CombineLoa API asynchronously
// api document: https://help.aliyun.com/api/trademark/combineloa.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CombineLoaWithCallback(request *CombineLoaRequest, callback func(response *CombineLoaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CombineLoaResponse
		var err error
		defer close(result)
		response, err = client.CombineLoa(request)
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

// CombineLoaRequest is the request struct for api CombineLoa
type CombineLoaRequest struct {
	*requests.RpcRequest
	TrademarkName string `position:"Query" name:"TrademarkName"`
	MaterialName  string `position:"Query" name:"MaterialName"`
	Address       string `position:"Query" name:"Address"`
	Nationality   string `position:"Query" name:"Nationality"`
	TmProduceType string `position:"Query" name:"TmProduceType"`
	MaterialId    string `position:"Query" name:"MaterialId"`
}

// CombineLoaResponse is the response struct for api CombineLoa
type CombineLoaResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	TemplateCombineUrl string `json:"TemplateCombineUrl" xml:"TemplateCombineUrl"`
}

// CreateCombineLoaRequest creates a request to invoke CombineLoa API
func CreateCombineLoaRequest() (request *CombineLoaRequest) {
	request = &CombineLoaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "CombineLoa", "trademark", "openAPI")
	return
}

// CreateCombineLoaResponse creates a response to parse from CombineLoa response
func CreateCombineLoaResponse() (response *CombineLoaResponse) {
	response = &CombineLoaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
