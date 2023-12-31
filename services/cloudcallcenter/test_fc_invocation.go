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

// TestFcInvocation invokes the cloudcallcenter.TestFcInvocation API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/testfcinvocation.html
func (client *Client) TestFcInvocation(request *TestFcInvocationRequest) (response *TestFcInvocationResponse, err error) {
	response = CreateTestFcInvocationResponse()
	err = client.DoAction(request, response)
	return
}

// TestFcInvocationWithChan invokes the cloudcallcenter.TestFcInvocation API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/testfcinvocation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TestFcInvocationWithChan(request *TestFcInvocationRequest) (<-chan *TestFcInvocationResponse, <-chan error) {
	responseChan := make(chan *TestFcInvocationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TestFcInvocation(request)
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

// TestFcInvocationWithCallback invokes the cloudcallcenter.TestFcInvocation API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/testfcinvocation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TestFcInvocationWithCallback(request *TestFcInvocationRequest, callback func(response *TestFcInvocationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TestFcInvocationResponse
		var err error
		defer close(result)
		response, err = client.TestFcInvocation(request)
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

// TestFcInvocationRequest is the request struct for api TestFcInvocation
type TestFcInvocationRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	Role         string `position:"Query" name:"Role"`
	Service      string `position:"Query" name:"Service"`
	FunctionName string `position:"Query" name:"FunctionName"`
	Parameter    string `position:"Query" name:"Parameter"`
	Region       string `position:"Query" name:"Region"`
}

// TestFcInvocationResponse is the response struct for api TestFcInvocation
type TestFcInvocationResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Payload        string `json:"Payload" xml:"Payload"`
}

// CreateTestFcInvocationRequest creates a request to invoke TestFcInvocation API
func CreateTestFcInvocationRequest() (request *TestFcInvocationRequest) {
	request = &TestFcInvocationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "TestFcInvocation", "", "")
	request.Method = requests.POST
	return
}

// CreateTestFcInvocationResponse creates a response to parse from TestFcInvocation response
func CreateTestFcInvocationResponse() (response *TestFcInvocationResponse) {
	response = &TestFcInvocationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
