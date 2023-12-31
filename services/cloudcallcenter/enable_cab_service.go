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

// EnableCabService invokes the cloudcallcenter.EnableCabService API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/enablecabservice.html
func (client *Client) EnableCabService(request *EnableCabServiceRequest) (response *EnableCabServiceResponse, err error) {
	response = CreateEnableCabServiceResponse()
	err = client.DoAction(request, response)
	return
}

// EnableCabServiceWithChan invokes the cloudcallcenter.EnableCabService API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/enablecabservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableCabServiceWithChan(request *EnableCabServiceRequest) (<-chan *EnableCabServiceResponse, <-chan error) {
	responseChan := make(chan *EnableCabServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableCabService(request)
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

// EnableCabServiceWithCallback invokes the cloudcallcenter.EnableCabService API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/enablecabservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableCabServiceWithCallback(request *EnableCabServiceRequest, callback func(response *EnableCabServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableCabServiceResponse
		var err error
		defer close(result)
		response, err = client.EnableCabService(request)
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

// EnableCabServiceRequest is the request struct for api EnableCabService
type EnableCabServiceRequest struct {
	*requests.RpcRequest
	Components          string           `position:"Query" name:"Components"`
	OwnerId             requests.Integer `position:"Query" name:"OwnerId"`
	Region              string           `position:"Query" name:"Region"`
	CommodityInstanceId string           `position:"Query" name:"CommodityInstanceId"`
}

// EnableCabServiceResponse is the response struct for api EnableCabService
type EnableCabServiceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateEnableCabServiceRequest creates a request to invoke EnableCabService API
func CreateEnableCabServiceRequest() (request *EnableCabServiceRequest) {
	request = &EnableCabServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "EnableCabService", "", "")
	request.Method = requests.POST
	return
}

// CreateEnableCabServiceResponse creates a response to parse from EnableCabService response
func CreateEnableCabServiceResponse() (response *EnableCabServiceResponse) {
	response = &EnableCabServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
