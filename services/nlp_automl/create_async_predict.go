package nlp_automl

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

// CreateAsyncPredict invokes the nlp_automl.CreateAsyncPredict API synchronously
// api document: https://help.aliyun.com/api/nlp-automl/createasyncpredict.html
func (client *Client) CreateAsyncPredict(request *CreateAsyncPredictRequest) (response *CreateAsyncPredictResponse, err error) {
	response = CreateCreateAsyncPredictResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAsyncPredictWithChan invokes the nlp_automl.CreateAsyncPredict API asynchronously
// api document: https://help.aliyun.com/api/nlp-automl/createasyncpredict.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAsyncPredictWithChan(request *CreateAsyncPredictRequest) (<-chan *CreateAsyncPredictResponse, <-chan error) {
	responseChan := make(chan *CreateAsyncPredictResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAsyncPredict(request)
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

// CreateAsyncPredictWithCallback invokes the nlp_automl.CreateAsyncPredict API asynchronously
// api document: https://help.aliyun.com/api/nlp-automl/createasyncpredict.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAsyncPredictWithCallback(request *CreateAsyncPredictRequest, callback func(response *CreateAsyncPredictResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAsyncPredictResponse
		var err error
		defer close(result)
		response, err = client.CreateAsyncPredict(request)
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

// CreateAsyncPredictRequest is the request struct for api CreateAsyncPredict
type CreateAsyncPredictRequest struct {
	*requests.RpcRequest
	TopK         requests.Integer `position:"Body" name:"TopK"`
	FileType     string           `position:"Body" name:"FileType"`
	DetailTag    string           `position:"Body" name:"DetailTag"`
	Content      string           `position:"Body" name:"Content"`
	FileContent  string           `position:"Body" name:"FileContent"`
	ModelId      requests.Integer `position:"Body" name:"ModelId"`
	FileUrl      string           `position:"Body" name:"FileUrl"`
	ModelVersion string           `position:"Body" name:"ModelVersion"`
}

// CreateAsyncPredictResponse is the response struct for api CreateAsyncPredict
type CreateAsyncPredictResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	AsyncPredictId int64  `json:"AsyncPredictId" xml:"AsyncPredictId"`
}

// CreateCreateAsyncPredictRequest creates a request to invoke CreateAsyncPredict API
func CreateCreateAsyncPredictRequest() (request *CreateAsyncPredictRequest) {
	request = &CreateAsyncPredictRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("nlp-automl", "2019-11-11", "CreateAsyncPredict", "nlpautoml", "openAPI")
	return
}

// CreateCreateAsyncPredictResponse creates a response to parse from CreateAsyncPredict response
func CreateCreateAsyncPredictResponse() (response *CreateAsyncPredictResponse) {
	response = &CreateAsyncPredictResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
