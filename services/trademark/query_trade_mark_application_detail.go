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

// QueryTradeMarkApplicationDetail invokes the trademark.QueryTradeMarkApplicationDetail API synchronously
// api document: https://help.aliyun.com/api/trademark/querytrademarkapplicationdetail.html
func (client *Client) QueryTradeMarkApplicationDetail(request *QueryTradeMarkApplicationDetailRequest) (response *QueryTradeMarkApplicationDetailResponse, err error) {
	response = CreateQueryTradeMarkApplicationDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTradeMarkApplicationDetailWithChan invokes the trademark.QueryTradeMarkApplicationDetail API asynchronously
// api document: https://help.aliyun.com/api/trademark/querytrademarkapplicationdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTradeMarkApplicationDetailWithChan(request *QueryTradeMarkApplicationDetailRequest) (<-chan *QueryTradeMarkApplicationDetailResponse, <-chan error) {
	responseChan := make(chan *QueryTradeMarkApplicationDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTradeMarkApplicationDetail(request)
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

// QueryTradeMarkApplicationDetailWithCallback invokes the trademark.QueryTradeMarkApplicationDetail API asynchronously
// api document: https://help.aliyun.com/api/trademark/querytrademarkapplicationdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTradeMarkApplicationDetailWithCallback(request *QueryTradeMarkApplicationDetailRequest, callback func(response *QueryTradeMarkApplicationDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTradeMarkApplicationDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryTradeMarkApplicationDetail(request)
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

// QueryTradeMarkApplicationDetailRequest is the request struct for api QueryTradeMarkApplicationDetail
type QueryTradeMarkApplicationDetailRequest struct {
	*requests.RpcRequest
	BizId string `position:"Query" name:"BizId"`
}

// QueryTradeMarkApplicationDetailResponse is the response struct for api QueryTradeMarkApplicationDetail
type QueryTradeMarkApplicationDetailResponse struct {
	*responses.BaseResponse
	Status              int                                                  `json:"Status" xml:"Status"`
	RequestId           string                                               `json:"RequestId" xml:"RequestId"`
	LoaUrl              string                                               `json:"LoaUrl" xml:"LoaUrl"`
	OrderPrice          int                                                  `json:"OrderPrice" xml:"OrderPrice"`
	TmIcon              string                                               `json:"TmIcon" xml:"TmIcon"`
	OrderId             string                                               `json:"OrderId" xml:"OrderId"`
	Type                int                                                  `json:"Type" xml:"Type"`
	TmNameType          int                                                  `json:"TmNameType" xml:"TmNameType"`
	TmName              string                                               `json:"TmName" xml:"TmName"`
	BizId               string                                               `json:"BizId" xml:"BizId"`
	GrayIconUrl         string                                               `json:"GrayIconUrl" xml:"GrayIconUrl"`
	Note                string                                               `json:"Note" xml:"Note"`
	TmNumber            string                                               `json:"TmNumber" xml:"TmNumber"`
	AcceptUrl           string                                               `json:"AcceptUrl" xml:"AcceptUrl"`
	CreateTime          int64                                                `json:"CreateTime" xml:"CreateTime"`
	UpdateTime          int64                                                `json:"UpdateTime" xml:"UpdateTime"`
	ExtendInfo          map[string]interface{}                               `json:"ExtendInfo" xml:"ExtendInfo"`
	SendSbjLogistics    string                                               `json:"SendSbjLogistics" xml:"SendSbjLogistics"`
	SendUserLogistics   string                                               `json:"SendUserLogistics" xml:"SendUserLogistics"`
	RecvUserLogistics   string                                               `json:"RecvUserLogistics" xml:"RecvUserLogistics"`
	MaterialId          int64                                                `json:"MaterialId" xml:"MaterialId"`
	ReceiptUrl          ReceiptUrl                                           `json:"ReceiptUrl" xml:"ReceiptUrl"`
	JudgeResultUrl      JudgeResultUrl                                       `json:"JudgeResultUrl" xml:"JudgeResultUrl"`
	AdminUploads        AdminUploads                                         `json:"AdminUploads" xml:"AdminUploads"`
	MaterialDetail      MaterialDetail                                       `json:"MaterialDetail" xml:"MaterialDetail"`
	FirstClassification FirstClassification                                  `json:"FirstClassification" xml:"FirstClassification"`
	RenewResponse       RenewResponse                                        `json:"RenewResponse" xml:"RenewResponse"`
	ThirdClassification ThirdClassificationInQueryTradeMarkApplicationDetail `json:"ThirdClassification" xml:"ThirdClassification"`
	Supplements         Supplements                                          `json:"Supplements" xml:"Supplements"`
}

// CreateQueryTradeMarkApplicationDetailRequest creates a request to invoke QueryTradeMarkApplicationDetail API
func CreateQueryTradeMarkApplicationDetailRequest() (request *QueryTradeMarkApplicationDetailRequest) {
	request = &QueryTradeMarkApplicationDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "QueryTradeMarkApplicationDetail", "trademark", "openAPI")
	return
}

// CreateQueryTradeMarkApplicationDetailResponse creates a response to parse from QueryTradeMarkApplicationDetail response
func CreateQueryTradeMarkApplicationDetailResponse() (response *QueryTradeMarkApplicationDetailResponse) {
	response = &QueryTradeMarkApplicationDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}