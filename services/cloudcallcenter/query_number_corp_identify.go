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

// QueryNumberCorpIdentify invokes the cloudcallcenter.QueryNumberCorpIdentify API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/querynumbercorpidentify.html
func (client *Client) QueryNumberCorpIdentify(request *QueryNumberCorpIdentifyRequest) (response *QueryNumberCorpIdentifyResponse, err error) {
	response = CreateQueryNumberCorpIdentifyResponse()
	err = client.DoAction(request, response)
	return
}

// QueryNumberCorpIdentifyWithChan invokes the cloudcallcenter.QueryNumberCorpIdentify API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/querynumbercorpidentify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryNumberCorpIdentifyWithChan(request *QueryNumberCorpIdentifyRequest) (<-chan *QueryNumberCorpIdentifyResponse, <-chan error) {
	responseChan := make(chan *QueryNumberCorpIdentifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryNumberCorpIdentify(request)
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

// QueryNumberCorpIdentifyWithCallback invokes the cloudcallcenter.QueryNumberCorpIdentify API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/querynumbercorpidentify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryNumberCorpIdentifyWithCallback(request *QueryNumberCorpIdentifyRequest, callback func(response *QueryNumberCorpIdentifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryNumberCorpIdentifyResponse
		var err error
		defer close(result)
		response, err = client.QueryNumberCorpIdentify(request)
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

// QueryNumberCorpIdentifyRequest is the request struct for api QueryNumberCorpIdentify
type QueryNumberCorpIdentifyRequest struct {
	*requests.RpcRequest
	WorkOrderId string           `position:"Query" name:"WorkOrderId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Status      string           `position:"Query" name:"Status"`
}

// QueryNumberCorpIdentifyResponse is the response struct for api QueryNumberCorpIdentify
type QueryNumberCorpIdentifyResponse struct {
	*responses.BaseResponse
	RequestId      string                        `json:"RequestId" xml:"RequestId"`
	Success        bool                          `json:"Success" xml:"Success"`
	Code           string                        `json:"Code" xml:"Code"`
	Message        string                        `json:"Message" xml:"Message"`
	HttpStatusCode int                           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInQueryNumberCorpIdentify `json:"Data" xml:"Data"`
}

// CreateQueryNumberCorpIdentifyRequest creates a request to invoke QueryNumberCorpIdentify API
func CreateQueryNumberCorpIdentifyRequest() (request *QueryNumberCorpIdentifyRequest) {
	request = &QueryNumberCorpIdentifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "QueryNumberCorpIdentify", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryNumberCorpIdentifyResponse creates a response to parse from QueryNumberCorpIdentify response
func CreateQueryNumberCorpIdentifyResponse() (response *QueryNumberCorpIdentifyResponse) {
	response = &QueryNumberCorpIdentifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}