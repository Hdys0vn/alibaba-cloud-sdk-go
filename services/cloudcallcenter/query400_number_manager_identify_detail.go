package cloudcallcenter

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

// Query400NumberManagerIdentifyDetail invokes the cloudcallcenter.Query400NumberManagerIdentifyDetail API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/query400numbermanageridentifydetail.html
func (client *Client) Query400NumberManagerIdentifyDetail(request *Query400NumberManagerIdentifyDetailRequest) (response *Query400NumberManagerIdentifyDetailResponse, err error) {
	response = CreateQuery400NumberManagerIdentifyDetailResponse()
	err = client.DoAction(request, response)
	return
}

// Query400NumberManagerIdentifyDetailWithChan invokes the cloudcallcenter.Query400NumberManagerIdentifyDetail API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/query400numbermanageridentifydetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) Query400NumberManagerIdentifyDetailWithChan(request *Query400NumberManagerIdentifyDetailRequest) (<-chan *Query400NumberManagerIdentifyDetailResponse, <-chan error) {
	responseChan := make(chan *Query400NumberManagerIdentifyDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Query400NumberManagerIdentifyDetail(request)
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

// Query400NumberManagerIdentifyDetailWithCallback invokes the cloudcallcenter.Query400NumberManagerIdentifyDetail API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/query400numbermanageridentifydetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) Query400NumberManagerIdentifyDetailWithCallback(request *Query400NumberManagerIdentifyDetailRequest, callback func(response *Query400NumberManagerIdentifyDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *Query400NumberManagerIdentifyDetailResponse
		var err error
		defer close(result)
		response, err = client.Query400NumberManagerIdentifyDetail(request)
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

// Query400NumberManagerIdentifyDetailRequest is the request struct for api Query400NumberManagerIdentifyDetail
type Query400NumberManagerIdentifyDetailRequest struct {
	*requests.RpcRequest
	OrderId string `position:"Query" name:"OrderId"`
}

// Query400NumberManagerIdentifyDetailResponse is the response struct for api Query400NumberManagerIdentifyDetail
type Query400NumberManagerIdentifyDetailResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateQuery400NumberManagerIdentifyDetailRequest creates a request to invoke Query400NumberManagerIdentifyDetail API
func CreateQuery400NumberManagerIdentifyDetailRequest() (request *Query400NumberManagerIdentifyDetailRequest) {
	request = &Query400NumberManagerIdentifyDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "Query400NumberManagerIdentifyDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateQuery400NumberManagerIdentifyDetailResponse creates a response to parse from Query400NumberManagerIdentifyDetail response
func CreateQuery400NumberManagerIdentifyDetailResponse() (response *Query400NumberManagerIdentifyDetailResponse) {
	response = &Query400NumberManagerIdentifyDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
