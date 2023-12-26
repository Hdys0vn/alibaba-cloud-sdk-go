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

// QueryAliyunCorpNumber invokes the cloudcallcenter.QueryAliyunCorpNumber API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryaliyuncorpnumber.html
func (client *Client) QueryAliyunCorpNumber(request *QueryAliyunCorpNumberRequest) (response *QueryAliyunCorpNumberResponse, err error) {
	response = CreateQueryAliyunCorpNumberResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAliyunCorpNumberWithChan invokes the cloudcallcenter.QueryAliyunCorpNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryaliyuncorpnumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAliyunCorpNumberWithChan(request *QueryAliyunCorpNumberRequest) (<-chan *QueryAliyunCorpNumberResponse, <-chan error) {
	responseChan := make(chan *QueryAliyunCorpNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAliyunCorpNumber(request)
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

// QueryAliyunCorpNumberWithCallback invokes the cloudcallcenter.QueryAliyunCorpNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryaliyuncorpnumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAliyunCorpNumberWithCallback(request *QueryAliyunCorpNumberRequest, callback func(response *QueryAliyunCorpNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAliyunCorpNumberResponse
		var err error
		defer close(result)
		response, err = client.QueryAliyunCorpNumber(request)
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

// QueryAliyunCorpNumberRequest is the request struct for api QueryAliyunCorpNumber
type QueryAliyunCorpNumberRequest struct {
	*requests.RpcRequest
	AvailableOnly   requests.Boolean `position:"Query" name:"AvailableOnly"`
	NumberStatus    string           `position:"Query" name:"NumberStatus"`
	TaobaoUid       requests.Integer `position:"Query" name:"TaobaoUid"`
	NumberType      string           `position:"Query" name:"NumberType"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ProviderChannel requests.Integer `position:"Query" name:"ProviderChannel"`
	Number          string           `position:"Query" name:"Number"`
	Provider        string           `position:"Query" name:"Provider"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// QueryAliyunCorpNumberResponse is the response struct for api QueryAliyunCorpNumber
type QueryAliyunCorpNumberResponse struct {
	*responses.BaseResponse
	RequestId      string                      `json:"RequestId" xml:"RequestId"`
	Success        bool                        `json:"Success" xml:"Success"`
	Code           string                      `json:"Code" xml:"Code"`
	Message        string                      `json:"Message" xml:"Message"`
	HttpStatusCode int                         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInQueryAliyunCorpNumber `json:"Data" xml:"Data"`
}

// CreateQueryAliyunCorpNumberRequest creates a request to invoke QueryAliyunCorpNumber API
func CreateQueryAliyunCorpNumberRequest() (request *QueryAliyunCorpNumberRequest) {
	request = &QueryAliyunCorpNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "QueryAliyunCorpNumber", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryAliyunCorpNumberResponse creates a response to parse from QueryAliyunCorpNumber response
func CreateQueryAliyunCorpNumberResponse() (response *QueryAliyunCorpNumberResponse) {
	response = &QueryAliyunCorpNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
