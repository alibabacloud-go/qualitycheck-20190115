// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddBusinessCategoryRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AddBusinessCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBusinessCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddBusinessCategoryRequest) SetJsonStr(v string) *AddBusinessCategoryRequest {
	s.JsonStr = &v
	return s
}

type AddBusinessCategoryResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddBusinessCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddBusinessCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddBusinessCategoryResponseBody) SetCode(v string) *AddBusinessCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetData(v string) *AddBusinessCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetMessage(v string) *AddBusinessCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetRequestId(v string) *AddBusinessCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetSuccess(v bool) *AddBusinessCategoryResponseBody {
	s.Success = &v
	return s
}

type AddBusinessCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddBusinessCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddBusinessCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddBusinessCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddBusinessCategoryResponse) SetHeaders(v map[string]*string) *AddBusinessCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddBusinessCategoryResponse) SetStatusCode(v int32) *AddBusinessCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBusinessCategoryResponse) SetBody(v *AddBusinessCategoryResponseBody) *AddBusinessCategoryResponse {
	s.Body = v
	return s
}

type AddRuleCategoryRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AddRuleCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryRequest) SetJsonStr(v string) *AddRuleCategoryRequest {
	s.JsonStr = &v
	return s
}

type AddRuleCategoryResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AddRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRuleCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryResponseBody) SetCode(v string) *AddRuleCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *AddRuleCategoryResponseBody) SetData(v *AddRuleCategoryResponseBodyData) *AddRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *AddRuleCategoryResponseBody) SetMessage(v string) *AddRuleCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *AddRuleCategoryResponseBody) SetRequestId(v string) *AddRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRuleCategoryResponseBody) SetSuccess(v bool) *AddRuleCategoryResponseBody {
	s.Success = &v
	return s
}

type AddRuleCategoryResponseBodyData struct {
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
}

func (s AddRuleCategoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryResponseBodyData) SetSelect(v bool) *AddRuleCategoryResponseBodyData {
	s.Select = &v
	return s
}

type AddRuleCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRuleCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryResponse) SetHeaders(v map[string]*string) *AddRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddRuleCategoryResponse) SetStatusCode(v int32) *AddRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRuleCategoryResponse) SetBody(v *AddRuleCategoryResponseBody) *AddRuleCategoryResponse {
	s.Body = v
	return s
}

type AddThesaurusForApiRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AddThesaurusForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s AddThesaurusForApiRequest) GoString() string {
	return s.String()
}

func (s *AddThesaurusForApiRequest) SetJsonStr(v string) *AddThesaurusForApiRequest {
	s.JsonStr = &v
	return s
}

type AddThesaurusForApiResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddThesaurusForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddThesaurusForApiResponseBody) GoString() string {
	return s.String()
}

func (s *AddThesaurusForApiResponseBody) SetCode(v string) *AddThesaurusForApiResponseBody {
	s.Code = &v
	return s
}

func (s *AddThesaurusForApiResponseBody) SetData(v int64) *AddThesaurusForApiResponseBody {
	s.Data = &v
	return s
}

func (s *AddThesaurusForApiResponseBody) SetMessage(v string) *AddThesaurusForApiResponseBody {
	s.Message = &v
	return s
}

func (s *AddThesaurusForApiResponseBody) SetRequestId(v string) *AddThesaurusForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddThesaurusForApiResponseBody) SetSuccess(v bool) *AddThesaurusForApiResponseBody {
	s.Success = &v
	return s
}

type AddThesaurusForApiResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddThesaurusForApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddThesaurusForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s AddThesaurusForApiResponse) GoString() string {
	return s.String()
}

func (s *AddThesaurusForApiResponse) SetHeaders(v map[string]*string) *AddThesaurusForApiResponse {
	s.Headers = v
	return s
}

func (s *AddThesaurusForApiResponse) SetStatusCode(v int32) *AddThesaurusForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *AddThesaurusForApiResponse) SetBody(v *AddThesaurusForApiResponseBody) *AddThesaurusForApiResponse {
	s.Body = v
	return s
}

type AssignReviewerRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AssignReviewerRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignReviewerRequest) GoString() string {
	return s.String()
}

func (s *AssignReviewerRequest) SetJsonStr(v string) *AssignReviewerRequest {
	s.JsonStr = &v
	return s
}

type AssignReviewerResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssignReviewerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignReviewerResponseBody) GoString() string {
	return s.String()
}

func (s *AssignReviewerResponseBody) SetCode(v string) *AssignReviewerResponseBody {
	s.Code = &v
	return s
}

func (s *AssignReviewerResponseBody) SetMessage(v string) *AssignReviewerResponseBody {
	s.Message = &v
	return s
}

func (s *AssignReviewerResponseBody) SetRequestId(v string) *AssignReviewerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignReviewerResponseBody) SetSuccess(v bool) *AssignReviewerResponseBody {
	s.Success = &v
	return s
}

type AssignReviewerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssignReviewerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssignReviewerResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignReviewerResponse) GoString() string {
	return s.String()
}

func (s *AssignReviewerResponse) SetHeaders(v map[string]*string) *AssignReviewerResponse {
	s.Headers = v
	return s
}

func (s *AssignReviewerResponse) SetStatusCode(v int32) *AssignReviewerResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignReviewerResponse) SetBody(v *AssignReviewerResponseBody) *AssignReviewerResponse {
	s.Body = v
	return s
}

type CreateAsrVocabRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateAsrVocabRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *CreateAsrVocabRequest) SetJsonStr(v string) *CreateAsrVocabRequest {
	s.JsonStr = &v
	return s
}

type CreateAsrVocabResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAsrVocabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsrVocabResponseBody) SetCode(v string) *CreateAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetData(v string) *CreateAsrVocabResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetMessage(v string) *CreateAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetRequestId(v string) *CreateAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAsrVocabResponseBody) SetSuccess(v bool) *CreateAsrVocabResponseBody {
	s.Success = &v
	return s
}

type CreateAsrVocabResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAsrVocabResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *CreateAsrVocabResponse) SetHeaders(v map[string]*string) *CreateAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *CreateAsrVocabResponse) SetStatusCode(v int32) *CreateAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAsrVocabResponse) SetBody(v *CreateAsrVocabResponseBody) *CreateAsrVocabResponse {
	s.Body = v
	return s
}

type CreateSkillGroupConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateSkillGroupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupConfigRequest) SetJsonStr(v string) *CreateSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

type CreateSkillGroupConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSkillGroupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupConfigResponseBody) SetCode(v string) *CreateSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetData(v int64) *CreateSkillGroupConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetMessage(v string) *CreateSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetRequestId(v string) *CreateSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillGroupConfigResponseBody) SetSuccess(v bool) *CreateSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

type CreateSkillGroupConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSkillGroupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupConfigResponse) SetHeaders(v map[string]*string) *CreateSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillGroupConfigResponse) SetStatusCode(v int32) *CreateSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillGroupConfigResponse) SetBody(v *CreateSkillGroupConfigResponseBody) *CreateSkillGroupConfigResponse {
	s.Body = v
	return s
}

type CreateTaskAssignRuleRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateTaskAssignRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskAssignRuleRequest) SetJsonStr(v string) *CreateTaskAssignRuleRequest {
	s.JsonStr = &v
	return s
}

type CreateTaskAssignRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTaskAssignRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskAssignRuleResponseBody) SetCode(v string) *CreateTaskAssignRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetData(v string) *CreateTaskAssignRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetMessage(v string) *CreateTaskAssignRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetRequestId(v string) *CreateTaskAssignRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskAssignRuleResponseBody) SetSuccess(v bool) *CreateTaskAssignRuleResponseBody {
	s.Success = &v
	return s
}

type CreateTaskAssignRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTaskAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTaskAssignRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskAssignRuleResponse) SetHeaders(v map[string]*string) *CreateTaskAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskAssignRuleResponse) SetStatusCode(v int32) *CreateTaskAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskAssignRuleResponse) SetBody(v *CreateTaskAssignRuleResponseBody) *CreateTaskAssignRuleResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetJsonStr(v string) *CreateUserRequest {
	s.JsonStr = &v
	return s
}

type CreateUserResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetCode(v string) *CreateUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserResponseBody) SetMessage(v string) *CreateUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetSuccess(v bool) *CreateUserResponseBody {
	s.Success = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type CreateWarningConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateWarningConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateWarningConfigRequest) SetJsonStr(v string) *CreateWarningConfigRequest {
	s.JsonStr = &v
	return s
}

type CreateWarningConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWarningConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWarningConfigResponseBody) SetCode(v string) *CreateWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetData(v string) *CreateWarningConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetMessage(v string) *CreateWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetRequestId(v string) *CreateWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetSuccess(v bool) *CreateWarningConfigResponseBody {
	s.Success = &v
	return s
}

type CreateWarningConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWarningConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateWarningConfigResponse) SetHeaders(v map[string]*string) *CreateWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateWarningConfigResponse) SetStatusCode(v int32) *CreateWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWarningConfigResponse) SetBody(v *CreateWarningConfigResponseBody) *CreateWarningConfigResponse {
	s.Body = v
	return s
}

type CreateWarningStrategyConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateWarningStrategyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateWarningStrategyConfigRequest) SetJsonStr(v string) *CreateWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

type CreateWarningStrategyConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWarningStrategyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWarningStrategyConfigResponseBody) SetCode(v string) *CreateWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetData(v int64) *CreateWarningStrategyConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetMessage(v string) *CreateWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetRequestId(v string) *CreateWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWarningStrategyConfigResponseBody) SetSuccess(v bool) *CreateWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

type CreateWarningStrategyConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWarningStrategyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *CreateWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateWarningStrategyConfigResponse) SetStatusCode(v int32) *CreateWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWarningStrategyConfigResponse) SetBody(v *CreateWarningStrategyConfigResponseBody) *CreateWarningStrategyConfigResponse {
	s.Body = v
	return s
}

type DelRuleCategoryRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DelRuleCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DelRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryRequest) SetJsonStr(v string) *DelRuleCategoryRequest {
	s.JsonStr = &v
	return s
}

type DelRuleCategoryResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DelRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelRuleCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryResponseBody) SetCode(v string) *DelRuleCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *DelRuleCategoryResponseBody) SetData(v *DelRuleCategoryResponseBodyData) *DelRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *DelRuleCategoryResponseBody) SetMessage(v string) *DelRuleCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *DelRuleCategoryResponseBody) SetRequestId(v string) *DelRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelRuleCategoryResponseBody) SetSuccess(v bool) *DelRuleCategoryResponseBody {
	s.Success = &v
	return s
}

type DelRuleCategoryResponseBodyData struct {
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
}

func (s DelRuleCategoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DelRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryResponseBodyData) SetSelect(v bool) *DelRuleCategoryResponseBodyData {
	s.Select = &v
	return s
}

type DelRuleCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DelRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DelRuleCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DelRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryResponse) SetHeaders(v map[string]*string) *DelRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *DelRuleCategoryResponse) SetStatusCode(v int32) *DelRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DelRuleCategoryResponse) SetBody(v *DelRuleCategoryResponseBody) *DelRuleCategoryResponse {
	s.Body = v
	return s
}

type DelThesaurusForApiRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DelThesaurusForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DelThesaurusForApiRequest) GoString() string {
	return s.String()
}

func (s *DelThesaurusForApiRequest) SetJsonStr(v string) *DelThesaurusForApiRequest {
	s.JsonStr = &v
	return s
}

type DelThesaurusForApiResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelThesaurusForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelThesaurusForApiResponseBody) GoString() string {
	return s.String()
}

func (s *DelThesaurusForApiResponseBody) SetCode(v string) *DelThesaurusForApiResponseBody {
	s.Code = &v
	return s
}

func (s *DelThesaurusForApiResponseBody) SetMessage(v string) *DelThesaurusForApiResponseBody {
	s.Message = &v
	return s
}

func (s *DelThesaurusForApiResponseBody) SetRequestId(v string) *DelThesaurusForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelThesaurusForApiResponseBody) SetSuccess(v bool) *DelThesaurusForApiResponseBody {
	s.Success = &v
	return s
}

type DelThesaurusForApiResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DelThesaurusForApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DelThesaurusForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DelThesaurusForApiResponse) GoString() string {
	return s.String()
}

func (s *DelThesaurusForApiResponse) SetHeaders(v map[string]*string) *DelThesaurusForApiResponse {
	s.Headers = v
	return s
}

func (s *DelThesaurusForApiResponse) SetStatusCode(v int32) *DelThesaurusForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DelThesaurusForApiResponse) SetBody(v *DelThesaurusForApiResponseBody) *DelThesaurusForApiResponse {
	s.Body = v
	return s
}

type DeleteAsrVocabRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteAsrVocabRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *DeleteAsrVocabRequest) SetJsonStr(v string) *DeleteAsrVocabRequest {
	s.JsonStr = &v
	return s
}

type DeleteAsrVocabResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAsrVocabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAsrVocabResponseBody) SetCode(v string) *DeleteAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetData(v string) *DeleteAsrVocabResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetMessage(v string) *DeleteAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetRequestId(v string) *DeleteAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetSuccess(v bool) *DeleteAsrVocabResponseBody {
	s.Success = &v
	return s
}

type DeleteAsrVocabResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAsrVocabResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *DeleteAsrVocabResponse) SetHeaders(v map[string]*string) *DeleteAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *DeleteAsrVocabResponse) SetStatusCode(v int32) *DeleteAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAsrVocabResponse) SetBody(v *DeleteAsrVocabResponseBody) *DeleteAsrVocabResponse {
	s.Body = v
	return s
}

type DeleteBusinessCategoryRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteBusinessCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBusinessCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteBusinessCategoryRequest) SetJsonStr(v string) *DeleteBusinessCategoryRequest {
	s.JsonStr = &v
	return s
}

type DeleteBusinessCategoryResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBusinessCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBusinessCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBusinessCategoryResponseBody) SetCode(v string) *DeleteBusinessCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetData(v string) *DeleteBusinessCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetMessage(v string) *DeleteBusinessCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetRequestId(v string) *DeleteBusinessCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetSuccess(v bool) *DeleteBusinessCategoryResponseBody {
	s.Success = &v
	return s
}

type DeleteBusinessCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBusinessCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBusinessCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBusinessCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteBusinessCategoryResponse) SetHeaders(v map[string]*string) *DeleteBusinessCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteBusinessCategoryResponse) SetStatusCode(v int32) *DeleteBusinessCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBusinessCategoryResponse) SetBody(v *DeleteBusinessCategoryResponseBody) *DeleteBusinessCategoryResponse {
	s.Body = v
	return s
}

type DeleteCustomizationConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteCustomizationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizationConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizationConfigRequest) SetJsonStr(v string) *DeleteCustomizationConfigRequest {
	s.JsonStr = &v
	return s
}

type DeleteCustomizationConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomizationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizationConfigResponseBody) SetCode(v string) *DeleteCustomizationConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetData(v string) *DeleteCustomizationConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetMessage(v string) *DeleteCustomizationConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetRequestId(v string) *DeleteCustomizationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomizationConfigResponseBody) SetSuccess(v bool) *DeleteCustomizationConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteCustomizationConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCustomizationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomizationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizationConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizationConfigResponse) SetHeaders(v map[string]*string) *DeleteCustomizationConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizationConfigResponse) SetStatusCode(v int32) *DeleteCustomizationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizationConfigResponse) SetBody(v *DeleteCustomizationConfigResponseBody) *DeleteCustomizationConfigResponse {
	s.Body = v
	return s
}

type DeleteDataSetRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSetRequest) SetJsonStr(v string) *DeleteDataSetRequest {
	s.JsonStr = &v
	return s
}

type DeleteDataSetResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSetResponseBody) SetCode(v string) *DeleteDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetMessage(v string) *DeleteDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetRequestId(v string) *DeleteDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetSuccess(v bool) *DeleteDataSetResponseBody {
	s.Success = &v
	return s
}

type DeleteDataSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSetResponse) SetHeaders(v map[string]*string) *DeleteDataSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSetResponse) SetStatusCode(v int32) *DeleteDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSetResponse) SetBody(v *DeleteDataSetResponseBody) *DeleteDataSetResponse {
	s.Body = v
	return s
}

type DeletePrecisionTaskRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeletePrecisionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *DeletePrecisionTaskRequest) SetJsonStr(v string) *DeletePrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

type DeletePrecisionTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePrecisionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrecisionTaskResponseBody) SetCode(v string) *DeletePrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePrecisionTaskResponseBody) SetMessage(v string) *DeletePrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePrecisionTaskResponseBody) SetRequestId(v string) *DeletePrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrecisionTaskResponseBody) SetSuccess(v bool) *DeletePrecisionTaskResponseBody {
	s.Success = &v
	return s
}

type DeletePrecisionTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePrecisionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *DeletePrecisionTaskResponse) SetHeaders(v map[string]*string) *DeletePrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *DeletePrecisionTaskResponse) SetStatusCode(v int32) *DeletePrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrecisionTaskResponse) SetBody(v *DeletePrecisionTaskResponseBody) *DeletePrecisionTaskResponse {
	s.Body = v
	return s
}

type DeleteScoreForApiRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteScoreForApiRequest) SetJsonStr(v string) *DeleteScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type DeleteScoreForApiResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScoreForApiResponseBody) SetCode(v string) *DeleteScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScoreForApiResponseBody) SetMessage(v string) *DeleteScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScoreForApiResponseBody) SetRequestId(v string) *DeleteScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScoreForApiResponseBody) SetSuccess(v bool) *DeleteScoreForApiResponseBody {
	s.Success = &v
	return s
}

type DeleteScoreForApiResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteScoreForApiResponse) SetHeaders(v map[string]*string) *DeleteScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteScoreForApiResponse) SetStatusCode(v int32) *DeleteScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScoreForApiResponse) SetBody(v *DeleteScoreForApiResponseBody) *DeleteScoreForApiResponse {
	s.Body = v
	return s
}

type DeleteSkillGroupConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteSkillGroupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupConfigRequest) SetJsonStr(v string) *DeleteSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

type DeleteSkillGroupConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSkillGroupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupConfigResponseBody) SetCode(v string) *DeleteSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSkillGroupConfigResponseBody) SetMessage(v string) *DeleteSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSkillGroupConfigResponseBody) SetRequestId(v string) *DeleteSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillGroupConfigResponseBody) SetSuccess(v bool) *DeleteSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteSkillGroupConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSkillGroupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupConfigResponse) SetHeaders(v map[string]*string) *DeleteSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteSkillGroupConfigResponse) SetStatusCode(v int32) *DeleteSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSkillGroupConfigResponse) SetBody(v *DeleteSkillGroupConfigResponseBody) *DeleteSkillGroupConfigResponse {
	s.Body = v
	return s
}

type DeleteSubScoreForApiRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteSubScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubScoreForApiRequest) SetJsonStr(v string) *DeleteSubScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type DeleteSubScoreForApiResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSubScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubScoreForApiResponseBody) SetCode(v string) *DeleteSubScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSubScoreForApiResponseBody) SetMessage(v string) *DeleteSubScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSubScoreForApiResponseBody) SetRequestId(v string) *DeleteSubScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubScoreForApiResponseBody) SetSuccess(v bool) *DeleteSubScoreForApiResponseBody {
	s.Success = &v
	return s
}

type DeleteSubScoreForApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSubScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSubScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubScoreForApiResponse) SetHeaders(v map[string]*string) *DeleteSubScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubScoreForApiResponse) SetStatusCode(v int32) *DeleteSubScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubScoreForApiResponse) SetBody(v *DeleteSubScoreForApiResponseBody) *DeleteSubScoreForApiResponse {
	s.Body = v
	return s
}

type DeleteTaskAssignRuleRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteTaskAssignRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteTaskAssignRuleRequest) SetJsonStr(v string) *DeleteTaskAssignRuleRequest {
	s.JsonStr = &v
	return s
}

type DeleteTaskAssignRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTaskAssignRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskAssignRuleResponseBody) SetCode(v string) *DeleteTaskAssignRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTaskAssignRuleResponseBody) SetMessage(v string) *DeleteTaskAssignRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTaskAssignRuleResponseBody) SetRequestId(v string) *DeleteTaskAssignRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTaskAssignRuleResponseBody) SetSuccess(v bool) *DeleteTaskAssignRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteTaskAssignRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTaskAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTaskAssignRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteTaskAssignRuleResponse) SetHeaders(v map[string]*string) *DeleteTaskAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteTaskAssignRuleResponse) SetStatusCode(v int32) *DeleteTaskAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTaskAssignRuleResponse) SetBody(v *DeleteTaskAssignRuleResponseBody) *DeleteTaskAssignRuleResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetJsonStr(v string) *DeleteUserRequest {
	s.JsonStr = &v
	return s
}

type DeleteUserResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetCode(v string) *DeleteUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DeleteWarningConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteWarningConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteWarningConfigRequest) SetJsonStr(v string) *DeleteWarningConfigRequest {
	s.JsonStr = &v
	return s
}

type DeleteWarningConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWarningConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWarningConfigResponseBody) SetCode(v string) *DeleteWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWarningConfigResponseBody) SetMessage(v string) *DeleteWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWarningConfigResponseBody) SetRequestId(v string) *DeleteWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWarningConfigResponseBody) SetSuccess(v bool) *DeleteWarningConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteWarningConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWarningConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteWarningConfigResponse) SetHeaders(v map[string]*string) *DeleteWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteWarningConfigResponse) SetStatusCode(v int32) *DeleteWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWarningConfigResponse) SetBody(v *DeleteWarningConfigResponseBody) *DeleteWarningConfigResponse {
	s.Body = v
	return s
}

type DeleteWarningStrategyConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteWarningStrategyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteWarningStrategyConfigRequest) SetJsonStr(v string) *DeleteWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

type DeleteWarningStrategyConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWarningStrategyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWarningStrategyConfigResponseBody) SetCode(v string) *DeleteWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetData(v string) *DeleteWarningStrategyConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetMessage(v string) *DeleteWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetRequestId(v string) *DeleteWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponseBody) SetSuccess(v bool) *DeleteWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteWarningStrategyConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWarningStrategyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *DeleteWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteWarningStrategyConfigResponse) SetStatusCode(v int32) *DeleteWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponse) SetBody(v *DeleteWarningStrategyConfigResponseBody) *DeleteWarningStrategyConfigResponse {
	s.Body = v
	return s
}

type EditThesaurusForApiRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s EditThesaurusForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s EditThesaurusForApiRequest) GoString() string {
	return s.String()
}

func (s *EditThesaurusForApiRequest) SetJsonStr(v string) *EditThesaurusForApiRequest {
	s.JsonStr = &v
	return s
}

type EditThesaurusForApiResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditThesaurusForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditThesaurusForApiResponseBody) GoString() string {
	return s.String()
}

func (s *EditThesaurusForApiResponseBody) SetCode(v string) *EditThesaurusForApiResponseBody {
	s.Code = &v
	return s
}

func (s *EditThesaurusForApiResponseBody) SetData(v int64) *EditThesaurusForApiResponseBody {
	s.Data = &v
	return s
}

func (s *EditThesaurusForApiResponseBody) SetMessage(v string) *EditThesaurusForApiResponseBody {
	s.Message = &v
	return s
}

func (s *EditThesaurusForApiResponseBody) SetRequestId(v string) *EditThesaurusForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditThesaurusForApiResponseBody) SetSuccess(v bool) *EditThesaurusForApiResponseBody {
	s.Success = &v
	return s
}

type EditThesaurusForApiResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EditThesaurusForApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditThesaurusForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s EditThesaurusForApiResponse) GoString() string {
	return s.String()
}

func (s *EditThesaurusForApiResponse) SetHeaders(v map[string]*string) *EditThesaurusForApiResponse {
	s.Headers = v
	return s
}

func (s *EditThesaurusForApiResponse) SetStatusCode(v int32) *EditThesaurusForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *EditThesaurusForApiResponse) SetBody(v *EditThesaurusForApiResponseBody) *EditThesaurusForApiResponse {
	s.Body = v
	return s
}

type GetAsrVocabRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetAsrVocabRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *GetAsrVocabRequest) SetJsonStr(v string) *GetAsrVocabRequest {
	s.JsonStr = &v
	return s
}

type GetAsrVocabResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAsrVocabResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsrVocabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBody) SetCode(v string) *GetAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsrVocabResponseBody) SetData(v *GetAsrVocabResponseBodyData) *GetAsrVocabResponseBody {
	s.Data = v
	return s
}

func (s *GetAsrVocabResponseBody) SetMessage(v string) *GetAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsrVocabResponseBody) SetRequestId(v string) *GetAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsrVocabResponseBody) SetSuccess(v bool) *GetAsrVocabResponseBody {
	s.Success = &v
	return s
}

type GetAsrVocabResponseBodyData struct {
	Name  *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Words *GetAsrVocabResponseBodyDataWords `json:"Words,omitempty" xml:"Words,omitempty" type:"Struct"`
}

func (s GetAsrVocabResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBodyData) SetName(v string) *GetAsrVocabResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAsrVocabResponseBodyData) SetWords(v *GetAsrVocabResponseBodyDataWords) *GetAsrVocabResponseBodyData {
	s.Words = v
	return s
}

type GetAsrVocabResponseBodyDataWords struct {
	Word []*GetAsrVocabResponseBodyDataWordsWord `json:"Word,omitempty" xml:"Word,omitempty" type:"Repeated"`
}

func (s GetAsrVocabResponseBodyDataWords) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabResponseBodyDataWords) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBodyDataWords) SetWord(v []*GetAsrVocabResponseBodyDataWordsWord) *GetAsrVocabResponseBodyDataWords {
	s.Word = v
	return s
}

type GetAsrVocabResponseBodyDataWordsWord struct {
	Weight *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Word   *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s GetAsrVocabResponseBodyDataWordsWord) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabResponseBodyDataWordsWord) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponseBodyDataWordsWord) SetWeight(v int32) *GetAsrVocabResponseBodyDataWordsWord {
	s.Weight = &v
	return s
}

func (s *GetAsrVocabResponseBodyDataWordsWord) SetWord(v string) *GetAsrVocabResponseBodyDataWordsWord {
	s.Word = &v
	return s
}

type GetAsrVocabResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsrVocabResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponse) SetHeaders(v map[string]*string) *GetAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *GetAsrVocabResponse) SetStatusCode(v int32) *GetAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsrVocabResponse) SetBody(v *GetAsrVocabResponseBody) *GetAsrVocabResponse {
	s.Body = v
	return s
}

type GetBusinessCategoryListRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetBusinessCategoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessCategoryListRequest) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListRequest) SetJsonStr(v string) *GetBusinessCategoryListRequest {
	s.JsonStr = &v
	return s
}

type GetBusinessCategoryListResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetBusinessCategoryListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBusinessCategoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessCategoryListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponseBody) SetCode(v string) *GetBusinessCategoryListResponseBody {
	s.Code = &v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetData(v *GetBusinessCategoryListResponseBodyData) *GetBusinessCategoryListResponseBody {
	s.Data = v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetMessage(v string) *GetBusinessCategoryListResponseBody {
	s.Message = &v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetRequestId(v string) *GetBusinessCategoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetSuccess(v bool) *GetBusinessCategoryListResponseBody {
	s.Success = &v
	return s
}

type GetBusinessCategoryListResponseBodyData struct {
	BusinessCategoryBasicInfo []*GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo `json:"BusinessCategoryBasicInfo,omitempty" xml:"BusinessCategoryBasicInfo,omitempty" type:"Repeated"`
}

func (s GetBusinessCategoryListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessCategoryListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponseBodyData) SetBusinessCategoryBasicInfo(v []*GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) *GetBusinessCategoryListResponseBodyData {
	s.BusinessCategoryBasicInfo = v
	return s
}

type GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo struct {
	Bid          *int32  `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	ServiceType  *int32  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) SetBid(v int32) *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo {
	s.Bid = &v
	return s
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) SetBusinessName(v string) *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo {
	s.BusinessName = &v
	return s
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) SetServiceType(v int32) *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo {
	s.ServiceType = &v
	return s
}

type GetBusinessCategoryListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBusinessCategoryListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBusinessCategoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBusinessCategoryListResponse) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponse) SetHeaders(v map[string]*string) *GetBusinessCategoryListResponse {
	s.Headers = v
	return s
}

func (s *GetBusinessCategoryListResponse) SetStatusCode(v int32) *GetBusinessCategoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBusinessCategoryListResponse) SetBody(v *GetBusinessCategoryListResponseBody) *GetBusinessCategoryListResponse {
	s.Body = v
	return s
}

type GetCustomizationConfigListRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetCustomizationConfigListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizationConfigListRequest) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListRequest) SetJsonStr(v string) *GetCustomizationConfigListRequest {
	s.JsonStr = &v
	return s
}

type GetCustomizationConfigListResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCustomizationConfigListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomizationConfigListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizationConfigListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponseBody) SetCode(v string) *GetCustomizationConfigListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetData(v *GetCustomizationConfigListResponseBodyData) *GetCustomizationConfigListResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetMessage(v string) *GetCustomizationConfigListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetRequestId(v string) *GetCustomizationConfigListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomizationConfigListResponseBody) SetSuccess(v bool) *GetCustomizationConfigListResponseBody {
	s.Success = &v
	return s
}

type GetCustomizationConfigListResponseBodyData struct {
	ModelCustomizationDataSetPo []*GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo `json:"ModelCustomizationDataSetPo,omitempty" xml:"ModelCustomizationDataSetPo,omitempty" type:"Repeated"`
}

func (s GetCustomizationConfigListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizationConfigListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponseBodyData) SetModelCustomizationDataSetPo(v []*GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) *GetCustomizationConfigListResponseBodyData {
	s.ModelCustomizationDataSetPo = v
	return s
}

type GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModeCustomizationId *string `json:"ModeCustomizationId,omitempty" xml:"ModeCustomizationId,omitempty"`
	ModelId             *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName           *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelStatus         *int32  `json:"ModelStatus,omitempty" xml:"ModelStatus,omitempty"`
	TaskType            *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetCreateTime(v string) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.CreateTime = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModeCustomizationId(v string) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModeCustomizationId = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModelId(v int64) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModelId = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModelName(v string) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModelName = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetModelStatus(v int32) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.ModelStatus = &v
	return s
}

func (s *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo) SetTaskType(v int32) *GetCustomizationConfigListResponseBodyDataModelCustomizationDataSetPo {
	s.TaskType = &v
	return s
}

type GetCustomizationConfigListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCustomizationConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCustomizationConfigListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizationConfigListResponse) GoString() string {
	return s.String()
}

func (s *GetCustomizationConfigListResponse) SetHeaders(v map[string]*string) *GetCustomizationConfigListResponse {
	s.Headers = v
	return s
}

func (s *GetCustomizationConfigListResponse) SetStatusCode(v int32) *GetCustomizationConfigListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomizationConfigListResponse) SetBody(v *GetCustomizationConfigListResponseBody) *GetCustomizationConfigListResponse {
	s.Body = v
	return s
}

type GetHitResultRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetHitResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHitResultRequest) GoString() string {
	return s.String()
}

func (s *GetHitResultRequest) SetJsonStr(v string) *GetHitResultRequest {
	s.JsonStr = &v
	return s
}

type GetHitResultResponseBody struct {
	Code       *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *int32                        `json:"Count,omitempty" xml:"Count,omitempty"`
	Data       *GetHitResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHitResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHitResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetHitResultResponseBody) SetCode(v string) *GetHitResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetHitResultResponseBody) SetCount(v int32) *GetHitResultResponseBody {
	s.Count = &v
	return s
}

func (s *GetHitResultResponseBody) SetData(v *GetHitResultResponseBodyData) *GetHitResultResponseBody {
	s.Data = v
	return s
}

func (s *GetHitResultResponseBody) SetMessage(v string) *GetHitResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetHitResultResponseBody) SetPageNumber(v int32) *GetHitResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetHitResultResponseBody) SetPageSize(v int32) *GetHitResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetHitResultResponseBody) SetRequestId(v string) *GetHitResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHitResultResponseBody) SetSuccess(v bool) *GetHitResultResponseBody {
	s.Success = &v
	return s
}

type GetHitResultResponseBodyData struct {
	ResultInfo []*GetHitResultResponseBodyDataResultInfo `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty" type:"Repeated"`
}

func (s GetHitResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHitResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHitResultResponseBodyData) SetResultInfo(v []*GetHitResultResponseBodyDataResultInfo) *GetHitResultResponseBodyData {
	s.ResultInfo = v
	return s
}

type GetHitResultResponseBodyDataResultInfo struct {
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetHitResultResponseBodyDataResultInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHitResultResponseBodyDataResultInfo) GoString() string {
	return s.String()
}

func (s *GetHitResultResponseBodyDataResultInfo) SetRid(v int64) *GetHitResultResponseBodyDataResultInfo {
	s.Rid = &v
	return s
}

func (s *GetHitResultResponseBodyDataResultInfo) SetRuleName(v string) *GetHitResultResponseBodyDataResultInfo {
	s.RuleName = &v
	return s
}

type GetHitResultResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHitResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHitResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHitResultResponse) GoString() string {
	return s.String()
}

func (s *GetHitResultResponse) SetHeaders(v map[string]*string) *GetHitResultResponse {
	s.Headers = v
	return s
}

func (s *GetHitResultResponse) SetStatusCode(v int32) *GetHitResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHitResultResponse) SetBody(v *GetHitResultResponseBody) *GetHitResultResponse {
	s.Body = v
	return s
}

type GetNextResultToVerifyRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetNextResultToVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyRequest) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyRequest) SetJsonStr(v string) *GetNextResultToVerifyRequest {
	s.JsonStr = &v
	return s
}

type GetNextResultToVerifyResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetNextResultToVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNextResultToVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBody) SetCode(v string) *GetNextResultToVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetData(v *GetNextResultToVerifyResponseBodyData) *GetNextResultToVerifyResponseBody {
	s.Data = v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetMessage(v string) *GetNextResultToVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetRequestId(v string) *GetNextResultToVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetSuccess(v bool) *GetNextResultToVerifyResponseBody {
	s.Success = &v
	return s
}

type GetNextResultToVerifyResponseBodyData struct {
	AudioScheme    *string                                         `json:"AudioScheme,omitempty" xml:"AudioScheme,omitempty"`
	AudioURL       *string                                         `json:"AudioURL,omitempty" xml:"AudioURL,omitempty"`
	Dialogues      *GetNextResultToVerifyResponseBodyDataDialogues `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Struct"`
	Duration       *int32                                          `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FileId         *string                                         `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileName       *string                                         `json:"FileName,omitempty" xml:"FileName,omitempty"`
	IncorrectWords *int32                                          `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Index          *int32                                          `json:"Index,omitempty" xml:"Index,omitempty"`
	Precision      *float32                                        `json:"Precision,omitempty" xml:"Precision,omitempty"`
	Status         *int32                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalCount     *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UpdateTime     *string                                         `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Verified       *bool                                           `json:"Verified,omitempty" xml:"Verified,omitempty"`
	VerifiedCount  *int32                                          `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyData) SetAudioScheme(v string) *GetNextResultToVerifyResponseBodyData {
	s.AudioScheme = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetAudioURL(v string) *GetNextResultToVerifyResponseBodyData {
	s.AudioURL = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetDialogues(v *GetNextResultToVerifyResponseBodyDataDialogues) *GetNextResultToVerifyResponseBodyData {
	s.Dialogues = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetDuration(v int32) *GetNextResultToVerifyResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetFileId(v string) *GetNextResultToVerifyResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetFileName(v string) *GetNextResultToVerifyResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetIncorrectWords(v int32) *GetNextResultToVerifyResponseBodyData {
	s.IncorrectWords = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetIndex(v int32) *GetNextResultToVerifyResponseBodyData {
	s.Index = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetPrecision(v float32) *GetNextResultToVerifyResponseBodyData {
	s.Precision = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetStatus(v int32) *GetNextResultToVerifyResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetTotalCount(v int32) *GetNextResultToVerifyResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetUpdateTime(v string) *GetNextResultToVerifyResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetVerified(v bool) *GetNextResultToVerifyResponseBodyData {
	s.Verified = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetVerifiedCount(v int32) *GetNextResultToVerifyResponseBodyData {
	s.VerifiedCount = &v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialogues struct {
	Dialogue []*GetNextResultToVerifyResponseBodyDataDialoguesDialogue `json:"Dialogue,omitempty" xml:"Dialogue,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialogues) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialogues) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialogues) SetDialogue(v []*GetNextResultToVerifyResponseBodyDataDialoguesDialogue) *GetNextResultToVerifyResponseBodyDataDialogues {
	s.Dialogue = v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogue struct {
	Begin           *int64                                                        `json:"Begin,omitempty" xml:"Begin,omitempty"`
	BeginTime       *string                                                       `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Deltas          *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas `json:"Deltas,omitempty" xml:"Deltas,omitempty" type:"Struct"`
	EmotionValue    *int32                                                        `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End             *int64                                                        `json:"End,omitempty" xml:"End,omitempty"`
	HourMinSec      *string                                                       `json:"HourMinSec,omitempty" xml:"HourMinSec,omitempty"`
	Identity        *string                                                       `json:"Identity,omitempty" xml:"Identity,omitempty"`
	IncorrectWords  *int32                                                        `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Role            *string                                                       `json:"Role,omitempty" xml:"Role,omitempty"`
	SilenceDuration *int32                                                        `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	SourceRole      *string                                                       `json:"SourceRole,omitempty" xml:"SourceRole,omitempty"`
	SourceWords     *string                                                       `json:"SourceWords,omitempty" xml:"SourceWords,omitempty"`
	SpeechRate      *int32                                                        `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words           *string                                                       `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogue) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetBegin(v int64) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Begin = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetBeginTime(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.BeginTime = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetDeltas(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Deltas = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetEmotionValue(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.EmotionValue = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetEnd(v int64) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.End = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetHourMinSec(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.HourMinSec = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetIdentity(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Identity = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetIncorrectWords(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.IncorrectWords = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetRole(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Role = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSilenceDuration(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SilenceDuration = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSourceRole(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SourceRole = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSourceWords(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SourceWords = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSpeechRate(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SpeechRate = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetWords(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Words = &v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas struct {
	Delta []*GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta `json:"Delta,omitempty" xml:"Delta,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) SetDelta(v []*GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas {
	s.Delta = v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta struct {
	Source *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Target *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	Type   *string                                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) SetSource(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta {
	s.Source = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) SetTarget(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta {
	s.Target = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) SetType(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta {
	s.Type = &v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource struct {
	Line     *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	Position *int32                                                                       `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) SetLine(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource {
	s.Line = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) SetPosition(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource {
	s.Position = &v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) SetLine(v []*string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine {
	s.Line = v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget struct {
	Line     *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	Position *int32                                                                       `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) SetLine(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget {
	s.Line = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) SetPosition(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget {
	s.Position = &v
	return s
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) SetLine(v []*string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine {
	s.Line = v
	return s
}

type GetNextResultToVerifyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNextResultToVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNextResultToVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNextResultToVerifyResponse) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponse) SetHeaders(v map[string]*string) *GetNextResultToVerifyResponse {
	s.Headers = v
	return s
}

func (s *GetNextResultToVerifyResponse) SetStatusCode(v int32) *GetNextResultToVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNextResultToVerifyResponse) SetBody(v *GetNextResultToVerifyResponseBody) *GetNextResultToVerifyResponse {
	s.Body = v
	return s
}

type GetPrecisionTaskRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetPrecisionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskRequest) SetJsonStr(v string) *GetPrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

type GetPrecisionTaskResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetPrecisionTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPrecisionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBody) SetCode(v string) *GetPrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetData(v *GetPrecisionTaskResponseBodyData) *GetPrecisionTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetMessage(v string) *GetPrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetRequestId(v string) *GetPrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetSuccess(v bool) *GetPrecisionTaskResponseBody {
	s.Success = &v
	return s
}

type GetPrecisionTaskResponseBodyData struct {
	DataSetId      *int64                                      `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	DataSetName    *string                                     `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	Duration       *int32                                      `json:"Duration,omitempty" xml:"Duration,omitempty"`
	IncorrectWords *int32                                      `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Name           *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Precisions     *GetPrecisionTaskResponseBodyDataPrecisions `json:"Precisions,omitempty" xml:"Precisions,omitempty" type:"Struct"`
	Source         *int32                                      `json:"Source,omitempty" xml:"Source,omitempty"`
	Status         *int32                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId         *string                                     `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TotalCount     *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UpdateTime     *string                                     `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VerifiedCount  *int32                                      `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s GetPrecisionTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBodyData) SetDataSetId(v int64) *GetPrecisionTaskResponseBodyData {
	s.DataSetId = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetDataSetName(v string) *GetPrecisionTaskResponseBodyData {
	s.DataSetName = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetDuration(v int32) *GetPrecisionTaskResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetIncorrectWords(v int32) *GetPrecisionTaskResponseBodyData {
	s.IncorrectWords = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetName(v string) *GetPrecisionTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetPrecisions(v *GetPrecisionTaskResponseBodyDataPrecisions) *GetPrecisionTaskResponseBodyData {
	s.Precisions = v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetSource(v int32) *GetPrecisionTaskResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetStatus(v int32) *GetPrecisionTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetTaskId(v string) *GetPrecisionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetTotalCount(v int32) *GetPrecisionTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetUpdateTime(v string) *GetPrecisionTaskResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetVerifiedCount(v int32) *GetPrecisionTaskResponseBodyData {
	s.VerifiedCount = &v
	return s
}

type GetPrecisionTaskResponseBodyDataPrecisions struct {
	Precision []*GetPrecisionTaskResponseBodyDataPrecisionsPrecision `json:"Precision,omitempty" xml:"Precision,omitempty" type:"Repeated"`
}

func (s GetPrecisionTaskResponseBodyDataPrecisions) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskResponseBodyDataPrecisions) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBodyDataPrecisions) SetPrecision(v []*GetPrecisionTaskResponseBodyDataPrecisionsPrecision) *GetPrecisionTaskResponseBodyDataPrecisions {
	s.Precision = v
	return s
}

type GetPrecisionTaskResponseBodyDataPrecisionsPrecision struct {
	ModelId   *int64   `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName *string  `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Precision *float32 `json:"Precision,omitempty" xml:"Precision,omitempty"`
	Status    *int32   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId    *string  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetPrecisionTaskResponseBodyDataPrecisionsPrecision) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskResponseBodyDataPrecisionsPrecision) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetModelId(v int64) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.ModelId = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetModelName(v string) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.ModelName = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetPrecision(v float32) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.Precision = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetStatus(v int32) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.Status = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetTaskId(v string) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.TaskId = &v
	return s
}

type GetPrecisionTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPrecisionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponse) SetHeaders(v map[string]*string) *GetPrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetPrecisionTaskResponse) SetStatusCode(v int32) *GetPrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrecisionTaskResponse) SetBody(v *GetPrecisionTaskResponseBody) *GetPrecisionTaskResponse {
	s.Body = v
	return s
}

type GetResultRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResultRequest) GoString() string {
	return s.String()
}

func (s *GetResultRequest) SetJsonStr(v string) *GetResultRequest {
	s.JsonStr = &v
	return s
}

type GetResultResponseBody struct {
	Code          *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Count         *int32                     `json:"Count,omitempty" xml:"Count,omitempty"`
	Data          *GetResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message       *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber    *int32                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCountId *string                    `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	Success       *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetResultResponseBody) SetCode(v string) *GetResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetResultResponseBody) SetCount(v int32) *GetResultResponseBody {
	s.Count = &v
	return s
}

func (s *GetResultResponseBody) SetData(v *GetResultResponseBodyData) *GetResultResponseBody {
	s.Data = v
	return s
}

func (s *GetResultResponseBody) SetMessage(v string) *GetResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetResultResponseBody) SetPageNumber(v int32) *GetResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetResultResponseBody) SetPageSize(v int32) *GetResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetResultResponseBody) SetRequestId(v string) *GetResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResultResponseBody) SetResultCountId(v string) *GetResultResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *GetResultResponseBody) SetSuccess(v bool) *GetResultResponseBody {
	s.Success = &v
	return s
}

type GetResultResponseBodyData struct {
	ResultInfo []*GetResultResponseBodyDataResultInfo `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyData) SetResultInfo(v []*GetResultResponseBodyDataResultInfo) *GetResultResponseBodyData {
	s.ResultInfo = v
	return s
}

type GetResultResponseBodyDataResultInfo struct {
	Agent          *GetResultResponseBodyDataResultInfoAgent          `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	AsrResult      *GetResultResponseBodyDataResultInfoAsrResult      `json:"AsrResult,omitempty" xml:"AsrResult,omitempty" type:"Struct"`
	AssignmentTime *string                                            `json:"AssignmentTime,omitempty" xml:"AssignmentTime,omitempty"`
	Comments       *string                                            `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateTime     *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeLong *string                                            `json:"CreateTimeLong,omitempty" xml:"CreateTimeLong,omitempty"`
	ErrorMessage   *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HitResult      *GetResultResponseBodyDataResultInfoHitResult      `json:"HitResult,omitempty" xml:"HitResult,omitempty" type:"Struct"`
	HitScore       *GetResultResponseBodyDataResultInfoHitScore       `json:"HitScore,omitempty" xml:"HitScore,omitempty" type:"Struct"`
	LastDataId     *string                                            `json:"LastDataId,omitempty" xml:"LastDataId,omitempty"`
	Recording      *GetResultResponseBodyDataResultInfoRecording      `json:"Recording,omitempty" xml:"Recording,omitempty" type:"Struct"`
	Resolver       *string                                            `json:"Resolver,omitempty" xml:"Resolver,omitempty"`
	ReviewResult   *int32                                             `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	ReviewStatus   *int32                                             `json:"ReviewStatus,omitempty" xml:"ReviewStatus,omitempty"`
	ReviewTime     *string                                            `json:"ReviewTime,omitempty" xml:"ReviewTime,omitempty"`
	ReviewTimeLong *string                                            `json:"ReviewTimeLong,omitempty" xml:"ReviewTimeLong,omitempty"`
	ReviewType     *int32                                             `json:"ReviewType,omitempty" xml:"ReviewType,omitempty"`
	Reviewer       *string                                            `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	SchemeIdList   *GetResultResponseBodyDataResultInfoSchemeIdList   `json:"SchemeIdList,omitempty" xml:"SchemeIdList,omitempty" type:"Struct"`
	SchemeNameList *GetResultResponseBodyDataResultInfoSchemeNameList `json:"SchemeNameList,omitempty" xml:"SchemeNameList,omitempty" type:"Struct"`
	Score          *int32                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Status         *int32                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId         *string                                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName       *string                                            `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetResultResponseBodyDataResultInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfo) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfo) SetAgent(v *GetResultResponseBodyDataResultInfoAgent) *GetResultResponseBodyDataResultInfo {
	s.Agent = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetAsrResult(v *GetResultResponseBodyDataResultInfoAsrResult) *GetResultResponseBodyDataResultInfo {
	s.AsrResult = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetAssignmentTime(v string) *GetResultResponseBodyDataResultInfo {
	s.AssignmentTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetComments(v string) *GetResultResponseBodyDataResultInfo {
	s.Comments = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetCreateTime(v string) *GetResultResponseBodyDataResultInfo {
	s.CreateTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetCreateTimeLong(v string) *GetResultResponseBodyDataResultInfo {
	s.CreateTimeLong = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetErrorMessage(v string) *GetResultResponseBodyDataResultInfo {
	s.ErrorMessage = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetHitResult(v *GetResultResponseBodyDataResultInfoHitResult) *GetResultResponseBodyDataResultInfo {
	s.HitResult = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetHitScore(v *GetResultResponseBodyDataResultInfoHitScore) *GetResultResponseBodyDataResultInfo {
	s.HitScore = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetLastDataId(v string) *GetResultResponseBodyDataResultInfo {
	s.LastDataId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetRecording(v *GetResultResponseBodyDataResultInfoRecording) *GetResultResponseBodyDataResultInfo {
	s.Recording = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetResolver(v string) *GetResultResponseBodyDataResultInfo {
	s.Resolver = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewResult(v int32) *GetResultResponseBodyDataResultInfo {
	s.ReviewResult = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewStatus(v int32) *GetResultResponseBodyDataResultInfo {
	s.ReviewStatus = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewTime(v string) *GetResultResponseBodyDataResultInfo {
	s.ReviewTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewTimeLong(v string) *GetResultResponseBodyDataResultInfo {
	s.ReviewTimeLong = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewType(v int32) *GetResultResponseBodyDataResultInfo {
	s.ReviewType = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetReviewer(v string) *GetResultResponseBodyDataResultInfo {
	s.Reviewer = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetSchemeIdList(v *GetResultResponseBodyDataResultInfoSchemeIdList) *GetResultResponseBodyDataResultInfo {
	s.SchemeIdList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetSchemeNameList(v *GetResultResponseBodyDataResultInfoSchemeNameList) *GetResultResponseBodyDataResultInfo {
	s.SchemeNameList = v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetScore(v int32) *GetResultResponseBodyDataResultInfo {
	s.Score = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetStatus(v int32) *GetResultResponseBodyDataResultInfo {
	s.Status = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetTaskId(v string) *GetResultResponseBodyDataResultInfo {
	s.TaskId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfo) SetTaskName(v string) *GetResultResponseBodyDataResultInfo {
	s.TaskName = &v
	return s
}

type GetResultResponseBodyDataResultInfoAgent struct {
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillGroup *string `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoAgent) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoAgent) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoAgent) SetId(v string) *GetResultResponseBodyDataResultInfoAgent {
	s.Id = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAgent) SetName(v string) *GetResultResponseBodyDataResultInfoAgent {
	s.Name = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAgent) SetSkillGroup(v string) *GetResultResponseBodyDataResultInfoAgent {
	s.SkillGroup = &v
	return s
}

type GetResultResponseBodyDataResultInfoAsrResult struct {
	AsrResult []*GetResultResponseBodyDataResultInfoAsrResultAsrResult `json:"AsrResult,omitempty" xml:"AsrResult,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoAsrResult) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoAsrResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoAsrResult) SetAsrResult(v []*GetResultResponseBodyDataResultInfoAsrResultAsrResult) *GetResultResponseBodyDataResultInfoAsrResult {
	s.AsrResult = v
	return s
}

type GetResultResponseBodyDataResultInfoAsrResultAsrResult struct {
	Begin        *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	EmotionValue *int32  `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End          *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SpeechRate   *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words        *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoAsrResultAsrResult) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoAsrResultAsrResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetBegin(v int64) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.Begin = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetEmotionValue(v int32) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.EmotionValue = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetEnd(v int64) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.End = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetRole(v string) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.Role = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetSpeechRate(v int32) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.SpeechRate = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoAsrResultAsrResult) SetWords(v string) *GetResultResponseBodyDataResultInfoAsrResultAsrResult {
	s.Words = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResult struct {
	HitResult []*GetResultResponseBodyDataResultInfoHitResultHitResult `json:"HitResult,omitempty" xml:"HitResult,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResult) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResult) SetHitResult(v []*GetResultResponseBodyDataResultInfoHitResultHitResult) *GetResultResponseBodyDataResultInfoHitResult {
	s.HitResult = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResult struct {
	Hits          *GetResultResponseBodyDataResultInfoHitResultHitResultHits `json:"Hits,omitempty" xml:"Hits,omitempty" type:"Struct"`
	Name          *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	ReviewResult  *int32                                                     `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	Rid           *string                                                    `json:"Rid,omitempty" xml:"Rid,omitempty"`
	SchemeId      *int64                                                     `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	SchemeVersion *int64                                                     `json:"SchemeVersion,omitempty" xml:"SchemeVersion,omitempty"`
	Type          *string                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResult) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResult) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetHits(v *GetResultResponseBodyDataResultInfoHitResultHitResultHits) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Hits = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetName(v string) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Name = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetReviewResult(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.ReviewResult = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetRid(v string) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Rid = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetSchemeId(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.SchemeId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetSchemeVersion(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.SchemeVersion = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResult) SetType(v string) *GetResultResponseBodyDataResultInfoHitResultHitResult {
	s.Type = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHits struct {
	Hit []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit `json:"Hit,omitempty" xml:"Hit,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHits) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHits) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHits) SetHit(v []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) *GetResultResponseBodyDataResultInfoHitResultHitResultHits {
	s.Hit = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit struct {
	Cid      *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid      `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Struct"`
	KeyWords *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Struct"`
	Phrase   *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase   `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) SetCid(v *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit {
	s.Cid = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) SetKeyWords(v *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit {
	s.KeyWords = v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit) SetPhrase(v *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHit {
	s.Phrase = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid struct {
	Cid []*string `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid) SetCid(v []*string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitCid {
	s.Cid = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords struct {
	KeyWord []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord `json:"KeyWord,omitempty" xml:"KeyWord,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords) SetKeyWord(v []*GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWords {
	s.KeyWord = v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord struct {
	Cid  *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	From *int32  `json:"From,omitempty" xml:"From,omitempty"`
	To   *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Val  *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) SetCid(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord {
	s.Cid = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) SetFrom(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord {
	s.From = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) SetTo(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord {
	s.To = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord) SetVal(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitKeyWordsKeyWord {
	s.Val = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase struct {
	Begin        *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	EmotionValue *int32  `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End          *int32  `json:"End,omitempty" xml:"End,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Words        *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) SetBegin(v int64) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	s.Begin = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) SetEmotionValue(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	s.EmotionValue = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) SetEnd(v int32) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	s.End = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) SetRole(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	s.Role = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase) SetWords(v string) *GetResultResponseBodyDataResultInfoHitResultHitResultHitsHitPhrase {
	s.Words = &v
	return s
}

type GetResultResponseBodyDataResultInfoHitScore struct {
	HitScore []*GetResultResponseBodyDataResultInfoHitScoreHitScore `json:"HitScore,omitempty" xml:"HitScore,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoHitScore) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitScore) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitScore) SetHitScore(v []*GetResultResponseBodyDataResultInfoHitScoreHitScore) *GetResultResponseBodyDataResultInfoHitScore {
	s.HitScore = v
	return s
}

type GetResultResponseBodyDataResultInfoHitScoreHitScore struct {
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	ScoreId     *string `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreName   *string `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
	ScoreNumber *string `json:"ScoreNumber,omitempty" xml:"ScoreNumber,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoHitScoreHitScore) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoHitScoreHitScore) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) SetRuleId(v string) *GetResultResponseBodyDataResultInfoHitScoreHitScore {
	s.RuleId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) SetScoreId(v string) *GetResultResponseBodyDataResultInfoHitScoreHitScore {
	s.ScoreId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) SetScoreName(v string) *GetResultResponseBodyDataResultInfoHitScoreHitScore {
	s.ScoreName = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoHitScoreHitScore) SetScoreNumber(v string) *GetResultResponseBodyDataResultInfoHitScoreHitScore {
	s.ScoreNumber = &v
	return s
}

type GetResultResponseBodyDataResultInfoRecording struct {
	Business     *string `json:"Business,omitempty" xml:"Business,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallTime     *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	CallType     *int32  `json:"CallType,omitempty" xml:"CallType,omitempty"`
	Callee       *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	Caller       *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	DataSetName  *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	DialogueSize *int32  `json:"DialogueSize,omitempty" xml:"DialogueSize,omitempty"`
	Duration     *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PrimaryId    *string `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	Remark1      *string `json:"Remark1,omitempty" xml:"Remark1,omitempty"`
	Remark10     *string `json:"Remark10,omitempty" xml:"Remark10,omitempty"`
	Remark11     *string `json:"Remark11,omitempty" xml:"Remark11,omitempty"`
	Remark12     *string `json:"Remark12,omitempty" xml:"Remark12,omitempty"`
	Remark13     *string `json:"Remark13,omitempty" xml:"Remark13,omitempty"`
	Remark2      *string `json:"Remark2,omitempty" xml:"Remark2,omitempty"`
	Remark3      *string `json:"Remark3,omitempty" xml:"Remark3,omitempty"`
	Remark4      *string `json:"Remark4,omitempty" xml:"Remark4,omitempty"`
	Remark5      *int64  `json:"Remark5,omitempty" xml:"Remark5,omitempty"`
	Remark6      *string `json:"Remark6,omitempty" xml:"Remark6,omitempty"`
	Remark7      *string `json:"Remark7,omitempty" xml:"Remark7,omitempty"`
	Remark8      *string `json:"Remark8,omitempty" xml:"Remark8,omitempty"`
	Remark9      *string `json:"Remark9,omitempty" xml:"Remark9,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetResultResponseBodyDataResultInfoRecording) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoRecording) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetBusiness(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Business = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetCallId(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.CallId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetCallTime(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.CallTime = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetCallType(v int32) *GetResultResponseBodyDataResultInfoRecording {
	s.CallType = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetCallee(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Callee = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetCaller(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Caller = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetDataSetName(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.DataSetName = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetDialogueSize(v int32) *GetResultResponseBodyDataResultInfoRecording {
	s.DialogueSize = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetDuration(v int64) *GetResultResponseBodyDataResultInfoRecording {
	s.Duration = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetId(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Id = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetName(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Name = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetPrimaryId(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.PrimaryId = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark1(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark1 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark10(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark10 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark11(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark11 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark12(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark12 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark13(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark13 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark2(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark2 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark3(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark3 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark4(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark4 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark5(v int64) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark5 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark6(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark6 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark7(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark7 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark8(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark8 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetRemark9(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Remark9 = &v
	return s
}

func (s *GetResultResponseBodyDataResultInfoRecording) SetUrl(v string) *GetResultResponseBodyDataResultInfoRecording {
	s.Url = &v
	return s
}

type GetResultResponseBodyDataResultInfoSchemeIdList struct {
	SchemeIdList []*int64 `json:"SchemeIdList,omitempty" xml:"SchemeIdList,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoSchemeIdList) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoSchemeIdList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoSchemeIdList) SetSchemeIdList(v []*int64) *GetResultResponseBodyDataResultInfoSchemeIdList {
	s.SchemeIdList = v
	return s
}

type GetResultResponseBodyDataResultInfoSchemeNameList struct {
	SchemeNameList []*string `json:"SchemeNameList,omitempty" xml:"SchemeNameList,omitempty" type:"Repeated"`
}

func (s GetResultResponseBodyDataResultInfoSchemeNameList) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponseBodyDataResultInfoSchemeNameList) GoString() string {
	return s.String()
}

func (s *GetResultResponseBodyDataResultInfoSchemeNameList) SetSchemeNameList(v []*string) *GetResultResponseBodyDataResultInfoSchemeNameList {
	s.SchemeNameList = v
	return s
}

type GetResultResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponse) GoString() string {
	return s.String()
}

func (s *GetResultResponse) SetHeaders(v map[string]*string) *GetResultResponse {
	s.Headers = v
	return s
}

func (s *GetResultResponse) SetStatusCode(v int32) *GetResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResultResponse) SetBody(v *GetResultResponseBody) *GetResultResponse {
	s.Body = v
	return s
}

type GetResultCallbackRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetResultCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResultCallbackRequest) GoString() string {
	return s.String()
}

func (s *GetResultCallbackRequest) SetJsonStr(v string) *GetResultCallbackRequest {
	s.JsonStr = &v
	return s
}

type GetResultCallbackResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResultCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResultCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *GetResultCallbackResponseBody) SetCode(v string) *GetResultCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *GetResultCallbackResponseBody) SetMessage(v string) *GetResultCallbackResponseBody {
	s.Message = &v
	return s
}

func (s *GetResultCallbackResponseBody) SetRequestId(v string) *GetResultCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResultCallbackResponseBody) SetSuccess(v bool) *GetResultCallbackResponseBody {
	s.Success = &v
	return s
}

type GetResultCallbackResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResultCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResultCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResultCallbackResponse) GoString() string {
	return s.String()
}

func (s *GetResultCallbackResponse) SetHeaders(v map[string]*string) *GetResultCallbackResponse {
	s.Headers = v
	return s
}

func (s *GetResultCallbackResponse) SetStatusCode(v int32) *GetResultCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResultCallbackResponse) SetBody(v *GetResultCallbackResponseBody) *GetResultCallbackResponse {
	s.Body = v
	return s
}

type GetResultToReviewRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetResultToReviewRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewRequest) GoString() string {
	return s.String()
}

func (s *GetResultToReviewRequest) SetJsonStr(v string) *GetResultToReviewRequest {
	s.JsonStr = &v
	return s
}

type GetResultToReviewResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetResultToReviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResultToReviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBody) SetCode(v string) *GetResultToReviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetResultToReviewResponseBody) SetData(v *GetResultToReviewResponseBodyData) *GetResultToReviewResponseBody {
	s.Data = v
	return s
}

func (s *GetResultToReviewResponseBody) SetMessage(v string) *GetResultToReviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetResultToReviewResponseBody) SetRequestId(v string) *GetResultToReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResultToReviewResponseBody) SetSuccess(v bool) *GetResultToReviewResponseBody {
	s.Success = &v
	return s
}

type GetResultToReviewResponseBodyData struct {
	AudioScheme           *string                                                 `json:"AudioScheme,omitempty" xml:"AudioScheme,omitempty"`
	AudioURL              *string                                                 `json:"AudioURL,omitempty" xml:"AudioURL,omitempty"`
	Comments              *string                                                 `json:"Comments,omitempty" xml:"Comments,omitempty"`
	Dialogues             *GetResultToReviewResponseBodyDataDialogues             `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Struct"`
	FileId                *string                                                 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileMergeName         *string                                                 `json:"FileMergeName,omitempty" xml:"FileMergeName,omitempty"`
	HitRuleReviewInfoList *GetResultToReviewResponseBodyDataHitRuleReviewInfoList `json:"HitRuleReviewInfoList,omitempty" xml:"HitRuleReviewInfoList,omitempty" type:"Struct"`
	ManualScoreInfoList   *GetResultToReviewResponseBodyDataManualScoreInfoList   `json:"ManualScoreInfoList,omitempty" xml:"ManualScoreInfoList,omitempty" type:"Struct"`
	ReviewHistoryList     *GetResultToReviewResponseBodyDataReviewHistoryList     `json:"ReviewHistoryList,omitempty" xml:"ReviewHistoryList,omitempty" type:"Struct"`
	Status                *int32                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalScore            *int32                                                  `json:"TotalScore,omitempty" xml:"TotalScore,omitempty"`
	Vid                   *string                                                 `json:"Vid,omitempty" xml:"Vid,omitempty"`
}

func (s GetResultToReviewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyData) SetAudioScheme(v string) *GetResultToReviewResponseBodyData {
	s.AudioScheme = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetAudioURL(v string) *GetResultToReviewResponseBodyData {
	s.AudioURL = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetComments(v string) *GetResultToReviewResponseBodyData {
	s.Comments = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetDialogues(v *GetResultToReviewResponseBodyDataDialogues) *GetResultToReviewResponseBodyData {
	s.Dialogues = v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetFileId(v string) *GetResultToReviewResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetFileMergeName(v string) *GetResultToReviewResponseBodyData {
	s.FileMergeName = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetHitRuleReviewInfoList(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoList) *GetResultToReviewResponseBodyData {
	s.HitRuleReviewInfoList = v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetManualScoreInfoList(v *GetResultToReviewResponseBodyDataManualScoreInfoList) *GetResultToReviewResponseBodyData {
	s.ManualScoreInfoList = v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetReviewHistoryList(v *GetResultToReviewResponseBodyDataReviewHistoryList) *GetResultToReviewResponseBodyData {
	s.ReviewHistoryList = v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetStatus(v int32) *GetResultToReviewResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetTotalScore(v int32) *GetResultToReviewResponseBodyData {
	s.TotalScore = &v
	return s
}

func (s *GetResultToReviewResponseBodyData) SetVid(v string) *GetResultToReviewResponseBodyData {
	s.Vid = &v
	return s
}

type GetResultToReviewResponseBodyDataDialogues struct {
	Dialogue []*GetResultToReviewResponseBodyDataDialoguesDialogue `json:"Dialogue,omitempty" xml:"Dialogue,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataDialogues) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataDialogues) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataDialogues) SetDialogue(v []*GetResultToReviewResponseBodyDataDialoguesDialogue) *GetResultToReviewResponseBodyDataDialogues {
	s.Dialogue = v
	return s
}

type GetResultToReviewResponseBodyDataDialoguesDialogue struct {
	Begin           *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	BeginTime       *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EmotionValue    *int32  `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End             *int64  `json:"End,omitempty" xml:"End,omitempty"`
	HourMinSec      *string `json:"HourMinSec,omitempty" xml:"HourMinSec,omitempty"`
	Identity        *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SilenceDuration *int32  `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	SpeechRate      *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words           *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetResultToReviewResponseBodyDataDialoguesDialogue) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataDialoguesDialogue) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetBegin(v int64) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.Begin = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetBeginTime(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.BeginTime = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetEmotionValue(v int32) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.EmotionValue = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetEnd(v int64) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.End = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetHourMinSec(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.HourMinSec = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetIdentity(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.Identity = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetRole(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.Role = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetSilenceDuration(v int32) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.SilenceDuration = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetSpeechRate(v int32) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.SpeechRate = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataDialoguesDialogue) SetWords(v string) *GetResultToReviewResponseBodyDataDialoguesDialogue {
	s.Words = &v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoList struct {
	HitRuleReviewInfo []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo `json:"HitRuleReviewInfo,omitempty" xml:"HitRuleReviewInfo,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoList) SetHitRuleReviewInfo(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) *GetResultToReviewResponseBodyDataHitRuleReviewInfoList {
	s.HitRuleReviewInfo = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo struct {
	AutoReview           *int32                                                                                       `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	ComplainHistories    *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories    `json:"ComplainHistories,omitempty" xml:"ComplainHistories,omitempty" type:"Struct"`
	Complainable         *bool                                                                                        `json:"Complainable,omitempty" xml:"Complainable,omitempty"`
	ConditionHitInfoList *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList `json:"ConditionHitInfoList,omitempty" xml:"ConditionHitInfoList,omitempty" type:"Struct"`
	ReviewInfo           *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo           `json:"ReviewInfo,omitempty" xml:"ReviewInfo,omitempty" type:"Struct"`
	Rid                  *int64                                                                                       `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName             *string                                                                                      `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	ScoreId              *int64                                                                                       `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreNum             *int32                                                                                       `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreSubId           *int64                                                                                       `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName         *string                                                                                      `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetAutoReview(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.AutoReview = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetComplainHistories(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ComplainHistories = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetComplainable(v bool) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.Complainable = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetConditionHitInfoList(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ConditionHitInfoList = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetReviewInfo(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ReviewInfo = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetRid(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.Rid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetRuleName(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.RuleName = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetScoreId(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ScoreId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetScoreNum(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ScoreNum = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetScoreSubId(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ScoreSubId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo) SetScoreSubName(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfo {
	s.ScoreSubName = &v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories struct {
	ComplainHistories []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories `json:"ComplainHistories,omitempty" xml:"ComplainHistories,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories) SetComplainHistories(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistories {
	s.ComplainHistories = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories struct {
	Comments      *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	OperationType *int32  `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Operator      *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName  *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) SetComments(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	s.Comments = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) SetOperationTime(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	s.OperationTime = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) SetOperationType(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	s.OperationType = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) SetOperator(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	s.Operator = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories) SetOperatorName(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoComplainHistoriesComplainHistories {
	s.OperatorName = &v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList struct {
	ConditionHitInfo []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo `json:"ConditionHitInfo,omitempty" xml:"ConditionHitInfo,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList) SetConditionHitInfo(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoList {
	s.ConditionHitInfo = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo struct {
	Cid      *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid      `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Struct"`
	KeyWords *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Struct"`
	Phrase   *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase   `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) SetCid(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo {
	s.Cid = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) SetKeyWords(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo {
	s.KeyWords = v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo) SetPhrase(v *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfo {
	s.Phrase = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid struct {
	Cid []*string `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid) SetCid(v []*string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoCid {
	s.Cid = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords struct {
	KeyWord []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord `json:"KeyWord,omitempty" xml:"KeyWord,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords) SetKeyWord(v []*GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWords {
	s.KeyWord = v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord struct {
	Cid           *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	CustomizeCode *string `json:"CustomizeCode,omitempty" xml:"CustomizeCode,omitempty"`
	From          *int32  `json:"From,omitempty" xml:"From,omitempty"`
	Pid           *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Tid           *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
	To            *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Val           *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetCid(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.Cid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetCustomizeCode(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.CustomizeCode = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetFrom(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.From = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetPid(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.Pid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetTid(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.Tid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetTo(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.To = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord) SetVal(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoKeyWordsKeyWord {
	s.Val = &v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase struct {
	Begin        *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	EmotionValue *int32  `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End          *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Identity     *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Pid          *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Words        *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetBegin(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.Begin = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetEmotionValue(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.EmotionValue = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetEnd(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.End = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetIdentity(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.Identity = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetPid(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.Pid = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetRole(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.Role = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase) SetWords(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoConditionHitInfoListConditionHitInfoPhrase {
	s.Words = &v
	return s
}

type GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo struct {
	HitId        *string `json:"HitId,omitempty" xml:"HitId,omitempty"`
	ReviewResult *int32  `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	ReviewTime   *string `json:"ReviewTime,omitempty" xml:"ReviewTime,omitempty"`
	Reviewer     *string `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	Rid          *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) SetHitId(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	s.HitId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) SetReviewResult(v int32) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	s.ReviewResult = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) SetReviewTime(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	s.ReviewTime = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) SetReviewer(v string) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	s.Reviewer = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo) SetRid(v int64) *GetResultToReviewResponseBodyDataHitRuleReviewInfoListHitRuleReviewInfoReviewInfo {
	s.Rid = &v
	return s
}

type GetResultToReviewResponseBodyDataManualScoreInfoList struct {
	ManualScoreInfo []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo `json:"ManualScoreInfo,omitempty" xml:"ManualScoreInfo,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoList) SetManualScoreInfo(v []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) *GetResultToReviewResponseBodyDataManualScoreInfoList {
	s.ManualScoreInfo = v
	return s
}

type GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo struct {
	ComplainHistories *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories `json:"ComplainHistories,omitempty" xml:"ComplainHistories,omitempty" type:"Struct"`
	Complainable      *bool                                                                                 `json:"Complainable,omitempty" xml:"Complainable,omitempty"`
	ScoreId           *int64                                                                                `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreNum          *int32                                                                                `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreSubId        *int64                                                                                `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName      *string                                                                               `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetComplainHistories(v *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.ComplainHistories = v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetComplainable(v bool) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.Complainable = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetScoreId(v int64) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.ScoreId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetScoreNum(v int32) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.ScoreNum = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetScoreSubId(v int64) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.ScoreSubId = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo) SetScoreSubName(v string) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfo {
	s.ScoreSubName = &v
	return s
}

type GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories struct {
	ComplainHistories []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories `json:"ComplainHistories,omitempty" xml:"ComplainHistories,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories) SetComplainHistories(v []*GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistories {
	s.ComplainHistories = v
	return s
}

type GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories struct {
	Comments      *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	OperationType *int32  `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Operator      *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName  *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) SetComments(v string) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	s.Comments = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) SetOperationTime(v string) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	s.OperationTime = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) SetOperationType(v int32) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	s.OperationType = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) SetOperator(v int64) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	s.Operator = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories) SetOperatorName(v string) *GetResultToReviewResponseBodyDataManualScoreInfoListManualScoreInfoComplainHistoriesComplainHistories {
	s.OperatorName = &v
	return s
}

type GetResultToReviewResponseBodyDataReviewHistoryList struct {
	ReviewHistory []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory `json:"ReviewHistory,omitempty" xml:"ReviewHistory,omitempty" type:"Repeated"`
}

func (s GetResultToReviewResponseBodyDataReviewHistoryList) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewHistoryList) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryList) SetReviewHistory(v []*GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) *GetResultToReviewResponseBodyDataReviewHistoryList {
	s.ReviewHistory = v
	return s
}

type GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory struct {
	ComplainResult *int32  `json:"ComplainResult,omitempty" xml:"ComplainResult,omitempty"`
	OldScore       *int32  `json:"OldScore,omitempty" xml:"OldScore,omitempty"`
	OperatorName   *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	ReviewResult   *int32  `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	Score          *int32  `json:"Score,omitempty" xml:"Score,omitempty"`
	TimeStr        *string `json:"TimeStr,omitempty" xml:"TimeStr,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetComplainResult(v int32) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.ComplainResult = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetOldScore(v int32) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.OldScore = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetOperatorName(v string) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.OperatorName = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetReviewResult(v int32) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.ReviewResult = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetScore(v int32) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.Score = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetTimeStr(v string) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.TimeStr = &v
	return s
}

func (s *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory) SetType(v int32) *GetResultToReviewResponseBodyDataReviewHistoryListReviewHistory {
	s.Type = &v
	return s
}

type GetResultToReviewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResultToReviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResultToReviewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResultToReviewResponse) GoString() string {
	return s.String()
}

func (s *GetResultToReviewResponse) SetHeaders(v map[string]*string) *GetResultToReviewResponse {
	s.Headers = v
	return s
}

func (s *GetResultToReviewResponse) SetStatusCode(v int32) *GetResultToReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResultToReviewResponse) SetBody(v *GetResultToReviewResponseBody) *GetResultToReviewResponse {
	s.Body = v
	return s
}

type GetRuleRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRuleRequest) SetJsonStr(v string) *GetRuleRequest {
	s.JsonStr = &v
	return s
}

type GetRuleResponseBody struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBody) SetCode(v string) *GetRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleResponseBody) SetData(v *GetRuleResponseBodyData) *GetRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleResponseBody) SetMessage(v string) *GetRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleResponseBody) SetRequestId(v string) *GetRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleResponseBody) SetSuccess(v bool) *GetRuleResponseBody {
	s.Success = &v
	return s
}

type GetRuleResponseBodyData struct {
	Rules *GetRuleResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s GetRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyData) SetRules(v *GetRuleResponseBodyDataRules) *GetRuleResponseBodyData {
	s.Rules = v
	return s
}

type GetRuleResponseBodyDataRules struct {
	RuleInfo []*GetRuleResponseBodyDataRulesRuleInfo `json:"RuleInfo,omitempty" xml:"RuleInfo,omitempty" type:"Repeated"`
}

func (s GetRuleResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataRules) SetRuleInfo(v []*GetRuleResponseBodyDataRulesRuleInfo) *GetRuleResponseBodyDataRules {
	s.RuleInfo = v
	return s
}

type GetRuleResponseBodyDataRulesRuleInfo struct {
	AutoReview               *int32                                                        `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	BusinessCategoryNameList *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Struct"`
	Comments                 *string                                                       `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateEmpid              *string                                                       `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	CreateTime               *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime                  *string                                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsDelete                 *int32                                                        `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	IsOnline                 *int32                                                        `json:"IsOnline,omitempty" xml:"IsOnline,omitempty"`
	LastUpdateEmpid          *string                                                       `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	LastUpdateTime           *string                                                       `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	Name                     *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Rid                      *string                                                       `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleLambda               *string                                                       `json:"RuleLambda,omitempty" xml:"RuleLambda,omitempty"`
	RuleScoreType            *int32                                                        `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	ScoreId                  *int32                                                        `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreName                *string                                                       `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
	ScoreSubId               *int32                                                        `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName             *string                                                       `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
	StartTime                *string                                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                   *int32                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                     *int32                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight                   *string                                                       `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s GetRuleResponseBodyDataRulesRuleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyDataRulesRuleInfo) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetAutoReview(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.AutoReview = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetBusinessCategoryNameList(v *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) *GetRuleResponseBodyDataRulesRuleInfo {
	s.BusinessCategoryNameList = v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetComments(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Comments = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetCreateEmpid(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.CreateEmpid = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetCreateTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.CreateTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetEndTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.EndTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetIsDelete(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.IsDelete = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetIsOnline(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.IsOnline = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetLastUpdateEmpid(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.LastUpdateEmpid = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetLastUpdateTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.LastUpdateTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetName(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Name = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetRid(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Rid = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetRuleLambda(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.RuleLambda = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetRuleScoreType(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.RuleScoreType = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreId(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreId = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreName(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreName = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreSubId(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreSubId = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreSubName(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreSubName = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetStartTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.StartTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetStatus(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Status = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetType(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Type = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetWeight(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Weight = &v
	return s
}

type GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList struct {
	BusinessCategoryNameList []*string `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
}

func (s GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) SetBusinessCategoryNameList(v []*string) *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList {
	s.BusinessCategoryNameList = v
	return s
}

type GetRuleResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRuleResponse) SetHeaders(v map[string]*string) *GetRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRuleResponse) SetStatusCode(v int32) *GetRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleResponse) SetBody(v *GetRuleResponseBody) *GetRuleResponse {
	s.Body = v
	return s
}

type GetRuleCategoryRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetRuleCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryRequest) SetJsonStr(v string) *GetRuleCategoryRequest {
	s.JsonStr = &v
	return s
}

type GetRuleCategoryResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponseBody) SetCode(v string) *GetRuleCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleCategoryResponseBody) SetData(v *GetRuleCategoryResponseBodyData) *GetRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleCategoryResponseBody) SetMessage(v string) *GetRuleCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleCategoryResponseBody) SetRequestId(v string) *GetRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleCategoryResponseBody) SetSuccess(v bool) *GetRuleCategoryResponseBody {
	s.Success = &v
	return s
}

type GetRuleCategoryResponseBodyData struct {
	RuleCountInfo []*GetRuleCategoryResponseBodyDataRuleCountInfo `json:"RuleCountInfo,omitempty" xml:"RuleCountInfo,omitempty" type:"Repeated"`
}

func (s GetRuleCategoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponseBodyData) SetRuleCountInfo(v []*GetRuleCategoryResponseBodyDataRuleCountInfo) *GetRuleCategoryResponseBodyData {
	s.RuleCountInfo = v
	return s
}

type GetRuleCategoryResponseBodyDataRuleCountInfo struct {
	Select   *bool   `json:"Select,omitempty" xml:"Select,omitempty"`
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s GetRuleCategoryResponseBodyDataRuleCountInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleCategoryResponseBodyDataRuleCountInfo) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) SetSelect(v bool) *GetRuleCategoryResponseBodyDataRuleCountInfo {
	s.Select = &v
	return s
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) SetType(v int32) *GetRuleCategoryResponseBodyDataRuleCountInfo {
	s.Type = &v
	return s
}

func (s *GetRuleCategoryResponseBodyDataRuleCountInfo) SetTypeName(v string) *GetRuleCategoryResponseBodyDataRuleCountInfo {
	s.TypeName = &v
	return s
}

type GetRuleCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRuleCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetRuleCategoryResponse) SetHeaders(v map[string]*string) *GetRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetRuleCategoryResponse) SetStatusCode(v int32) *GetRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleCategoryResponse) SetBody(v *GetRuleCategoryResponseBody) *GetRuleCategoryResponse {
	s.Body = v
	return s
}

type GetRuleDetailRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetRuleDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetRuleDetailRequest) SetJsonStr(v string) *GetRuleDetailRequest {
	s.JsonStr = &v
	return s
}

type GetRuleDetailResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetRuleDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBody) SetCode(v string) *GetRuleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleDetailResponseBody) SetData(v *GetRuleDetailResponseBodyData) *GetRuleDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleDetailResponseBody) SetMessage(v string) *GetRuleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleDetailResponseBody) SetRequestId(v string) *GetRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleDetailResponseBody) SetSuccess(v bool) *GetRuleDetailResponseBody {
	s.Success = &v
	return s
}

type GetRuleDetailResponseBodyData struct {
	Conditions *GetRuleDetailResponseBodyDataConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Struct"`
	Count      *int32                                   `json:"Count,omitempty" xml:"Count,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rules      *GetRuleDetailResponseBodyDataRules      `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s GetRuleDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyData) SetConditions(v *GetRuleDetailResponseBodyDataConditions) *GetRuleDetailResponseBodyData {
	s.Conditions = v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetCount(v int32) *GetRuleDetailResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetPageNumber(v int32) *GetRuleDetailResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetPageSize(v int32) *GetRuleDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetRules(v *GetRuleDetailResponseBodyDataRules) *GetRuleDetailResponseBodyData {
	s.Rules = v
	return s
}

type GetRuleDetailResponseBodyDataConditions struct {
	ConditionBasicInfo []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfo `json:"ConditionBasicInfo,omitempty" xml:"ConditionBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditions) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditions) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditions) SetConditionBasicInfo(v []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) *GetRuleDetailResponseBodyDataConditions {
	s.ConditionBasicInfo = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfo struct {
	CheckRange       *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange `json:"CheckRange,omitempty" xml:"CheckRange,omitempty" type:"Struct"`
	ConditionInfoCid *string                                                              `json:"ConditionInfoCid,omitempty" xml:"ConditionInfoCid,omitempty"`
	OperLambda       *string                                                              `json:"OperLambda,omitempty" xml:"OperLambda,omitempty"`
	Operators        *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators  `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Struct"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetCheckRange(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.CheckRange = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetConditionInfoCid(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.ConditionInfoCid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetOperLambda(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.OperLambda = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetOperators(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.Operators = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange struct {
	Absolute *bool                                                                      `json:"Absolute,omitempty" xml:"Absolute,omitempty"`
	Anchor   *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor `json:"Anchor,omitempty" xml:"Anchor,omitempty" type:"Struct"`
	Range    *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange  `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
	Role     *string                                                                    `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetAbsolute(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Absolute = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetAnchor(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Anchor = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetRange(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Range = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetRole(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Role = &v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor struct {
	AnchorCid *string `json:"AnchorCid,omitempty" xml:"AnchorCid,omitempty"`
	HitTime   *int32  `json:"HitTime,omitempty" xml:"HitTime,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) SetAnchorCid(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor {
	s.AnchorCid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) SetHitTime(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor {
	s.HitTime = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) SetLocation(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor {
	s.Location = &v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange struct {
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	To   *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) SetFrom(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange {
	s.From = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) SetTo(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange {
	s.To = &v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators struct {
	OperatorBasicInfo []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo `json:"OperatorBasicInfo,omitempty" xml:"OperatorBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) SetOperatorBasicInfo(v []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators {
	s.OperatorBasicInfo = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo struct {
	Oid      *string                                                                                   `json:"Oid,omitempty" xml:"Oid,omitempty"`
	OperName *string                                                                                   `json:"OperName,omitempty" xml:"OperName,omitempty"`
	Param    *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	Type     *string                                                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetOid(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.Oid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetOperName(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.OperName = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetParam(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.Param = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetType(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.Type = &v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam struct {
	AntModelInfo          *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo       `json:"AntModelInfo,omitempty" xml:"AntModelInfo,omitempty" type:"Struct"`
	Average               *bool                                                                                                       `json:"Average,omitempty" xml:"Average,omitempty"`
	BeginType             *string                                                                                                     `json:"BeginType,omitempty" xml:"BeginType,omitempty"`
	CheckType             *int32                                                                                                      `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	CompareOperator       *string                                                                                                     `json:"CompareOperator,omitempty" xml:"CompareOperator,omitempty"`
	ContextChatMatch      *bool                                                                                                       `json:"ContextChatMatch,omitempty" xml:"ContextChatMatch,omitempty"`
	DelayTime             *int32                                                                                                      `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DifferentRole         *bool                                                                                                       `json:"DifferentRole,omitempty" xml:"DifferentRole,omitempty"`
	Excludes              *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes           `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Struct"`
	From                  *int32                                                                                                      `json:"From,omitempty" xml:"From,omitempty"`
	FromEnd               *bool                                                                                                       `json:"FromEnd,omitempty" xml:"FromEnd,omitempty"`
	HitTime               *int32                                                                                                      `json:"HitTime,omitempty" xml:"HitTime,omitempty"`
	InSentence            *bool                                                                                                       `json:"InSentence,omitempty" xml:"InSentence,omitempty"`
	Interval              *int32                                                                                                      `json:"Interval,omitempty" xml:"Interval,omitempty"`
	KeywordExtension      *bool                                                                                                       `json:"KeywordExtension,omitempty" xml:"KeywordExtension,omitempty"`
	KeywordMatchSize      *int32                                                                                                      `json:"KeywordMatchSize,omitempty" xml:"KeywordMatchSize,omitempty"`
	MaxEmotionChangeValue *int32                                                                                                      `json:"MaxEmotionChangeValue,omitempty" xml:"MaxEmotionChangeValue,omitempty"`
	MinWordSize           *int32                                                                                                      `json:"MinWordSize,omitempty" xml:"MinWordSize,omitempty"`
	NotRegex              *string                                                                                                     `json:"NotRegex,omitempty" xml:"NotRegex,omitempty"`
	OperKeyWords          *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords       `json:"OperKeyWords,omitempty" xml:"OperKeyWords,omitempty" type:"Struct"`
	Phrase                *string                                                                                                     `json:"Phrase,omitempty" xml:"Phrase,omitempty"`
	Pvalues               *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues            `json:"Pvalues,omitempty" xml:"Pvalues,omitempty" type:"Struct"`
	References            *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences         `json:"References,omitempty" xml:"References,omitempty" type:"Struct"`
	Regex                 *string                                                                                                     `json:"Regex,omitempty" xml:"Regex,omitempty"`
	Score                 *int32                                                                                                      `json:"Score,omitempty" xml:"Score,omitempty"`
	SimilarityThreshold   *float32                                                                                                    `json:"Similarity_threshold,omitempty" xml:"Similarity_threshold,omitempty"`
	SimilarlySentences    *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences `json:"SimilarlySentences,omitempty" xml:"SimilarlySentences,omitempty" type:"Struct"`
	Target                *int32                                                                                                      `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetRole            *string                                                                                                     `json:"TargetRole,omitempty" xml:"TargetRole,omitempty"`
	Threshold             *float32                                                                                                    `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	VelocityInMint        *int32                                                                                                      `json:"VelocityInMint,omitempty" xml:"VelocityInMint,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetAntModelInfo(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.AntModelInfo = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetAverage(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Average = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetBeginType(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.BeginType = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetCheckType(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.CheckType = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetCompareOperator(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.CompareOperator = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetContextChatMatch(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.ContextChatMatch = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetDelayTime(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.DelayTime = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetDifferentRole(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.DifferentRole = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetExcludes(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Excludes = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetFrom(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.From = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetFromEnd(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.FromEnd = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetHitTime(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.HitTime = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetInSentence(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.InSentence = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetInterval(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Interval = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetKeywordExtension(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.KeywordExtension = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetKeywordMatchSize(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.KeywordMatchSize = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetMaxEmotionChangeValue(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.MaxEmotionChangeValue = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetMinWordSize(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.MinWordSize = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetNotRegex(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.NotRegex = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetOperKeyWords(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.OperKeyWords = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetPhrase(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Phrase = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetPvalues(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Pvalues = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetReferences(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.References = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetRegex(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Regex = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetScore(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Score = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetSimilarityThreshold(v float32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.SimilarityThreshold = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetSimilarlySentences(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.SimilarlySentences = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetTarget(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Target = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetTargetRole(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.TargetRole = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetThreshold(v float32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Threshold = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetVelocityInMint(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.VelocityInMint = &v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo struct {
	AntModelInfo []*string `json:"AntModelInfo,omitempty" xml:"AntModelInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) SetAntModelInfo(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo {
	s.AntModelInfo = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes struct {
	Excludes []*string `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) SetExcludes(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes {
	s.Excludes = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords struct {
	OperKeyWord []*string `json:"OperKeyWord,omitempty" xml:"OperKeyWord,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) SetOperKeyWord(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords {
	s.OperKeyWord = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues struct {
	Pvalues []*string `json:"Pvalues,omitempty" xml:"Pvalues,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) SetPvalues(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues {
	s.Pvalues = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences struct {
	Reference []*string `json:"Reference,omitempty" xml:"Reference,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) SetReference(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences {
	s.Reference = v
	return s
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences struct {
	SimilarlySentence []*string `json:"SimilarlySentence,omitempty" xml:"SimilarlySentence,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) SetSimilarlySentence(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences {
	s.SimilarlySentence = v
	return s
}

type GetRuleDetailResponseBodyDataRules struct {
	RuleBasicInfo []*GetRuleDetailResponseBodyDataRulesRuleBasicInfo `json:"RuleBasicInfo,omitempty" xml:"RuleBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRules) SetRuleBasicInfo(v []*GetRuleDetailResponseBodyDataRulesRuleBasicInfo) *GetRuleDetailResponseBodyDataRules {
	s.RuleBasicInfo = v
	return s
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfo struct {
	BusinessCategories *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories `json:"BusinessCategories,omitempty" xml:"BusinessCategories,omitempty" type:"Struct"`
	Rid                *string                                                            `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleLambda         *string                                                            `json:"RuleLambda,omitempty" xml:"RuleLambda,omitempty"`
	Triggers           *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers           `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Struct"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetBusinessCategories(v *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.BusinessCategories = v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetRid(v string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.Rid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetRuleLambda(v string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.RuleLambda = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetTriggers(v *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.Triggers = v
	return s
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories struct {
	BusinessCategoryBasicInfo []*GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo `json:"BusinessCategoryBasicInfo,omitempty" xml:"BusinessCategoryBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) SetBusinessCategoryBasicInfo(v []*GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories {
	s.BusinessCategoryBasicInfo = v
	return s
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo struct {
	Bid          *int32  `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	ServiceType  *int32  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) SetBid(v int32) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo {
	s.Bid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) SetBusinessName(v string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo {
	s.BusinessName = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) SetServiceType(v int32) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo {
	s.ServiceType = &v
	return s
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers struct {
	Trigger []*string `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) SetTrigger(v []*string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers {
	s.Trigger = v
	return s
}

type GetRuleDetailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRuleDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponse) SetHeaders(v map[string]*string) *GetRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetRuleDetailResponse) SetStatusCode(v int32) *GetRuleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleDetailResponse) SetBody(v *GetRuleDetailResponseBody) *GetRuleDetailResponse {
	s.Body = v
	return s
}

type GetScoreInfoRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetScoreInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoRequest) GoString() string {
	return s.String()
}

func (s *GetScoreInfoRequest) SetJsonStr(v string) *GetScoreInfoRequest {
	s.JsonStr = &v
	return s
}

type GetScoreInfoResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetScoreInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScoreInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBody) SetCode(v string) *GetScoreInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetScoreInfoResponseBody) SetData(v *GetScoreInfoResponseBodyData) *GetScoreInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetScoreInfoResponseBody) SetMessage(v string) *GetScoreInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetScoreInfoResponseBody) SetRequestId(v string) *GetScoreInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScoreInfoResponseBody) SetSuccess(v bool) *GetScoreInfoResponseBody {
	s.Success = &v
	return s
}

type GetScoreInfoResponseBodyData struct {
	ScorePo []*GetScoreInfoResponseBodyDataScorePo `json:"ScorePo,omitempty" xml:"ScorePo,omitempty" type:"Repeated"`
}

func (s GetScoreInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyData) SetScorePo(v []*GetScoreInfoResponseBodyDataScorePo) *GetScoreInfoResponseBodyData {
	s.ScorePo = v
	return s
}

type GetScoreInfoResponseBodyDataScorePo struct {
	ScoreId    *int32                                         `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreInfos *GetScoreInfoResponseBodyDataScorePoScoreInfos `json:"ScoreInfos,omitempty" xml:"ScoreInfos,omitempty" type:"Struct"`
	ScoreName  *string                                        `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
}

func (s GetScoreInfoResponseBodyDataScorePo) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponseBodyDataScorePo) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyDataScorePo) SetScoreId(v int32) *GetScoreInfoResponseBodyDataScorePo {
	s.ScoreId = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePo) SetScoreInfos(v *GetScoreInfoResponseBodyDataScorePoScoreInfos) *GetScoreInfoResponseBodyDataScorePo {
	s.ScoreInfos = v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePo) SetScoreName(v string) *GetScoreInfoResponseBodyDataScorePo {
	s.ScoreName = &v
	return s
}

type GetScoreInfoResponseBodyDataScorePoScoreInfos struct {
	ScoreParam []*GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam `json:"ScoreParam,omitempty" xml:"ScoreParam,omitempty" type:"Repeated"`
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfos) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfos) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfos) SetScoreParam(v []*GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) *GetScoreInfoResponseBodyDataScorePoScoreInfos {
	s.ScoreParam = v
	return s
}

type GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam struct {
	ScoreNum     *int32  `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreSubId   *int32  `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName *string `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
	ScoreType    *int32  `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreNum(v int32) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreNum = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreSubId(v int32) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreSubId = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreSubName(v string) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreSubName = &v
	return s
}

func (s *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam) SetScoreType(v int32) *GetScoreInfoResponseBodyDataScorePoScoreInfosScoreParam {
	s.ScoreType = &v
	return s
}

type GetScoreInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetScoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScoreInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScoreInfoResponse) GoString() string {
	return s.String()
}

func (s *GetScoreInfoResponse) SetHeaders(v map[string]*string) *GetScoreInfoResponse {
	s.Headers = v
	return s
}

func (s *GetScoreInfoResponse) SetStatusCode(v int32) *GetScoreInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScoreInfoResponse) SetBody(v *GetScoreInfoResponseBody) *GetScoreInfoResponse {
	s.Body = v
	return s
}

type GetSkillGroupConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetSkillGroupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigRequest) SetJsonStr(v string) *GetSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

type GetSkillGroupConfigResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetSkillGroupConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSkillGroupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBody) SetCode(v string) *GetSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetData(v *GetSkillGroupConfigResponseBodyData) *GetSkillGroupConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetMessage(v string) *GetSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetRequestId(v string) *GetSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetSuccess(v bool) *GetSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

type GetSkillGroupConfigResponseBodyData struct {
	AllContentQualityCheck *int32                                          `json:"AllContentQualityCheck,omitempty" xml:"AllContentQualityCheck,omitempty"`
	AllRids                *string                                         `json:"AllRids,omitempty" xml:"AllRids,omitempty"`
	AllRuleList            *GetSkillGroupConfigResponseBodyDataAllRuleList `json:"AllRuleList,omitempty" xml:"AllRuleList,omitempty" type:"Struct"`
	CreateTime             *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id                     *int64                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId             *string                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModelId                *int64                                          `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName              *string                                         `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Name                   *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	QualityCheckType       *int32                                          `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	Rid                    *string                                         `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleList               *GetSkillGroupConfigResponseBodyDataRuleList    `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	SkillGroupFrom         *int32                                          `json:"SkillGroupFrom,omitempty" xml:"SkillGroupFrom,omitempty"`
	SkillGroupId           *string                                         `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName         *string                                         `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	Status                 *int32                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                   *int32                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime             *string                                         `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VocabId                *int64                                          `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
	VocabName              *string                                         `json:"VocabName,omitempty" xml:"VocabName,omitempty"`
}

func (s GetSkillGroupConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyData) SetAllContentQualityCheck(v int32) *GetSkillGroupConfigResponseBodyData {
	s.AllContentQualityCheck = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetAllRids(v string) *GetSkillGroupConfigResponseBodyData {
	s.AllRids = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetAllRuleList(v *GetSkillGroupConfigResponseBodyDataAllRuleList) *GetSkillGroupConfigResponseBodyData {
	s.AllRuleList = v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetCreateTime(v string) *GetSkillGroupConfigResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetId(v int64) *GetSkillGroupConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetInstanceId(v string) *GetSkillGroupConfigResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetModelId(v int64) *GetSkillGroupConfigResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetModelName(v string) *GetSkillGroupConfigResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetName(v string) *GetSkillGroupConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetQualityCheckType(v int32) *GetSkillGroupConfigResponseBodyData {
	s.QualityCheckType = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetRid(v string) *GetSkillGroupConfigResponseBodyData {
	s.Rid = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetRuleList(v *GetSkillGroupConfigResponseBodyDataRuleList) *GetSkillGroupConfigResponseBodyData {
	s.RuleList = v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetSkillGroupFrom(v int32) *GetSkillGroupConfigResponseBodyData {
	s.SkillGroupFrom = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetSkillGroupId(v string) *GetSkillGroupConfigResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetSkillGroupName(v string) *GetSkillGroupConfigResponseBodyData {
	s.SkillGroupName = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetStatus(v int32) *GetSkillGroupConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetType(v int32) *GetSkillGroupConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetUpdateTime(v string) *GetSkillGroupConfigResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetVocabId(v int64) *GetSkillGroupConfigResponseBodyData {
	s.VocabId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetVocabName(v string) *GetSkillGroupConfigResponseBodyData {
	s.VocabName = &v
	return s
}

type GetSkillGroupConfigResponseBodyDataAllRuleList struct {
	RuleNameInfo []*GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleList) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleList) SetRuleNameInfo(v []*GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) *GetSkillGroupConfigResponseBodyDataAllRuleList {
	s.RuleNameInfo = v
	return s
}

type GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo struct {
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) SetRid(v int64) *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) SetRuleName(v string) *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

type GetSkillGroupConfigResponseBodyDataRuleList struct {
	RuleNameInfo []*GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s GetSkillGroupConfigResponseBodyDataRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataRuleList) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataRuleList) SetRuleNameInfo(v []*GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) *GetSkillGroupConfigResponseBodyDataRuleList {
	s.RuleNameInfo = v
	return s
}

type GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo struct {
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) SetRid(v int64) *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) SetRuleName(v string) *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

type GetSkillGroupConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSkillGroupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponse) SetHeaders(v map[string]*string) *GetSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupConfigResponse) SetStatusCode(v int32) *GetSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillGroupConfigResponse) SetBody(v *GetSkillGroupConfigResponseBody) *GetSkillGroupConfigResponse {
	s.Body = v
	return s
}

type GetSyncResultRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetSyncResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetSyncResultRequest) SetJsonStr(v string) *GetSyncResultRequest {
	s.JsonStr = &v
	return s
}

type GetSyncResultResponseBody struct {
	Code          *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Count         *int32                           `json:"Count,omitempty" xml:"Count,omitempty"`
	Data          []*GetSyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message       *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber    *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCountId *string                          `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	Success       *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSyncResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBody) SetCode(v string) *GetSyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetSyncResultResponseBody) SetCount(v int32) *GetSyncResultResponseBody {
	s.Count = &v
	return s
}

func (s *GetSyncResultResponseBody) SetData(v []*GetSyncResultResponseBodyData) *GetSyncResultResponseBody {
	s.Data = v
	return s
}

func (s *GetSyncResultResponseBody) SetMessage(v string) *GetSyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSyncResultResponseBody) SetPageNumber(v int32) *GetSyncResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetSyncResultResponseBody) SetPageSize(v int32) *GetSyncResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSyncResultResponseBody) SetRequestId(v string) *GetSyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSyncResultResponseBody) SetResultCountId(v string) *GetSyncResultResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *GetSyncResultResponseBody) SetSuccess(v bool) *GetSyncResultResponseBody {
	s.Success = &v
	return s
}

type GetSyncResultResponseBodyData struct {
	Agent        *GetSyncResultResponseBodyDataAgent       `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	AsrResult    []*GetSyncResultResponseBodyDataAsrResult `json:"AsrResult,omitempty" xml:"AsrResult,omitempty" type:"Repeated"`
	Comments     *string                                   `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateTime   *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HitResult    []*GetSyncResultResponseBodyDataHitResult `json:"HitResult,omitempty" xml:"HitResult,omitempty" type:"Repeated"`
	Recording    *GetSyncResultResponseBodyDataRecording   `json:"Recording,omitempty" xml:"Recording,omitempty" type:"Struct"`
	Resolver     *string                                   `json:"Resolver,omitempty" xml:"Resolver,omitempty"`
	ReviewResult *int32                                    `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	ReviewStatus *int32                                    `json:"ReviewStatus,omitempty" xml:"ReviewStatus,omitempty"`
	Reviewer     *string                                   `json:"Reviewer,omitempty" xml:"Reviewer,omitempty"`
	Score        *int32                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	Status       *int32                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId       *string                                   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName     *string                                   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetSyncResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyData) SetAgent(v *GetSyncResultResponseBodyDataAgent) *GetSyncResultResponseBodyData {
	s.Agent = v
	return s
}

func (s *GetSyncResultResponseBodyData) SetAsrResult(v []*GetSyncResultResponseBodyDataAsrResult) *GetSyncResultResponseBodyData {
	s.AsrResult = v
	return s
}

func (s *GetSyncResultResponseBodyData) SetComments(v string) *GetSyncResultResponseBodyData {
	s.Comments = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetCreateTime(v string) *GetSyncResultResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetErrorMessage(v string) *GetSyncResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetHitResult(v []*GetSyncResultResponseBodyDataHitResult) *GetSyncResultResponseBodyData {
	s.HitResult = v
	return s
}

func (s *GetSyncResultResponseBodyData) SetRecording(v *GetSyncResultResponseBodyDataRecording) *GetSyncResultResponseBodyData {
	s.Recording = v
	return s
}

func (s *GetSyncResultResponseBodyData) SetResolver(v string) *GetSyncResultResponseBodyData {
	s.Resolver = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetReviewResult(v int32) *GetSyncResultResponseBodyData {
	s.ReviewResult = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetReviewStatus(v int32) *GetSyncResultResponseBodyData {
	s.ReviewStatus = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetReviewer(v string) *GetSyncResultResponseBodyData {
	s.Reviewer = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetScore(v int32) *GetSyncResultResponseBodyData {
	s.Score = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetStatus(v int32) *GetSyncResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetTaskId(v string) *GetSyncResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetSyncResultResponseBodyData) SetTaskName(v string) *GetSyncResultResponseBodyData {
	s.TaskName = &v
	return s
}

type GetSyncResultResponseBodyDataAgent struct {
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillGroup *string `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
}

func (s GetSyncResultResponseBodyDataAgent) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataAgent) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataAgent) SetId(v string) *GetSyncResultResponseBodyDataAgent {
	s.Id = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAgent) SetName(v string) *GetSyncResultResponseBodyDataAgent {
	s.Name = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAgent) SetSkillGroup(v string) *GetSyncResultResponseBodyDataAgent {
	s.SkillGroup = &v
	return s
}

type GetSyncResultResponseBodyDataAsrResult struct {
	Begin           *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	EmotionValue    *int32  `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End             *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SilenceDuration *int32  `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	SpeechRate      *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words           *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetSyncResultResponseBodyDataAsrResult) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataAsrResult) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetBegin(v int64) *GetSyncResultResponseBodyDataAsrResult {
	s.Begin = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetEmotionValue(v int32) *GetSyncResultResponseBodyDataAsrResult {
	s.EmotionValue = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetEnd(v int64) *GetSyncResultResponseBodyDataAsrResult {
	s.End = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetRole(v string) *GetSyncResultResponseBodyDataAsrResult {
	s.Role = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetSilenceDuration(v int32) *GetSyncResultResponseBodyDataAsrResult {
	s.SilenceDuration = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetSpeechRate(v int32) *GetSyncResultResponseBodyDataAsrResult {
	s.SpeechRate = &v
	return s
}

func (s *GetSyncResultResponseBodyDataAsrResult) SetWords(v string) *GetSyncResultResponseBodyDataAsrResult {
	s.Words = &v
	return s
}

type GetSyncResultResponseBodyDataHitResult struct {
	Hits         []*GetSyncResultResponseBodyDataHitResultHits `json:"Hits,omitempty" xml:"Hits,omitempty" type:"Repeated"`
	Name         *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	ReviewResult *int32                                        `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
	Rid          *string                                       `json:"Rid,omitempty" xml:"Rid,omitempty"`
	Type         *string                                       `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSyncResultResponseBodyDataHitResult) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResult) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResult) SetHits(v []*GetSyncResultResponseBodyDataHitResultHits) *GetSyncResultResponseBodyDataHitResult {
	s.Hits = v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResult) SetName(v string) *GetSyncResultResponseBodyDataHitResult {
	s.Name = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResult) SetReviewResult(v int32) *GetSyncResultResponseBodyDataHitResult {
	s.ReviewResult = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResult) SetRid(v string) *GetSyncResultResponseBodyDataHitResult {
	s.Rid = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResult) SetType(v string) *GetSyncResultResponseBodyDataHitResult {
	s.Type = &v
	return s
}

type GetSyncResultResponseBodyDataHitResultHits struct {
	Cid      []*string                                             `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
	KeyWords []*GetSyncResultResponseBodyDataHitResultHitsKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	Phrase   *GetSyncResultResponseBodyDataHitResultHitsPhrase     `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s GetSyncResultResponseBodyDataHitResultHits) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResultHits) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResultHits) SetCid(v []*string) *GetSyncResultResponseBodyDataHitResultHits {
	s.Cid = v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHits) SetKeyWords(v []*GetSyncResultResponseBodyDataHitResultHitsKeyWords) *GetSyncResultResponseBodyDataHitResultHits {
	s.KeyWords = v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHits) SetPhrase(v *GetSyncResultResponseBodyDataHitResultHitsPhrase) *GetSyncResultResponseBodyDataHitResultHits {
	s.Phrase = v
	return s
}

type GetSyncResultResponseBodyDataHitResultHitsKeyWords struct {
	Cid  *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	From *int32  `json:"From,omitempty" xml:"From,omitempty"`
	To   *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Val  *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s GetSyncResultResponseBodyDataHitResultHitsKeyWords) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResultHitsKeyWords) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) SetCid(v string) *GetSyncResultResponseBodyDataHitResultHitsKeyWords {
	s.Cid = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) SetFrom(v int32) *GetSyncResultResponseBodyDataHitResultHitsKeyWords {
	s.From = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) SetTo(v int32) *GetSyncResultResponseBodyDataHitResultHitsKeyWords {
	s.To = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsKeyWords) SetVal(v string) *GetSyncResultResponseBodyDataHitResultHitsKeyWords {
	s.Val = &v
	return s
}

type GetSyncResultResponseBodyDataHitResultHitsPhrase struct {
	Begin           *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	EmotionValue    *int32  `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End             *int32  `json:"End,omitempty" xml:"End,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SilenceDuration *int32  `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	SpeechRate      *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words           *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetSyncResultResponseBodyDataHitResultHitsPhrase) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataHitResultHitsPhrase) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetBegin(v int64) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.Begin = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetEmotionValue(v int32) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.EmotionValue = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetEnd(v int32) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.End = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetRole(v string) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.Role = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetSilenceDuration(v int32) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.SilenceDuration = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetSpeechRate(v int32) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.SpeechRate = &v
	return s
}

func (s *GetSyncResultResponseBodyDataHitResultHitsPhrase) SetWords(v string) *GetSyncResultResponseBodyDataHitResultHitsPhrase {
	s.Words = &v
	return s
}

type GetSyncResultResponseBodyDataRecording struct {
	Business      *string `json:"Business,omitempty" xml:"Business,omitempty"`
	CallId        *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallTime      *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	CallType      *int32  `json:"CallType,omitempty" xml:"CallType,omitempty"`
	Callee        *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	Caller        *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	DataSetName   *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationAudio *int64  `json:"DurationAudio,omitempty" xml:"DurationAudio,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PrimaryId     *string `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	Remark1       *string `json:"Remark1,omitempty" xml:"Remark1,omitempty"`
	Remark2       *string `json:"Remark2,omitempty" xml:"Remark2,omitempty"`
	Remark3       *string `json:"Remark3,omitempty" xml:"Remark3,omitempty"`
	Url           *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetSyncResultResponseBodyDataRecording) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponseBodyDataRecording) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponseBodyDataRecording) SetBusiness(v string) *GetSyncResultResponseBodyDataRecording {
	s.Business = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetCallId(v string) *GetSyncResultResponseBodyDataRecording {
	s.CallId = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetCallTime(v string) *GetSyncResultResponseBodyDataRecording {
	s.CallTime = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetCallType(v int32) *GetSyncResultResponseBodyDataRecording {
	s.CallType = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetCallee(v string) *GetSyncResultResponseBodyDataRecording {
	s.Callee = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetCaller(v string) *GetSyncResultResponseBodyDataRecording {
	s.Caller = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetDataSetName(v string) *GetSyncResultResponseBodyDataRecording {
	s.DataSetName = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetDuration(v int64) *GetSyncResultResponseBodyDataRecording {
	s.Duration = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetDurationAudio(v int64) *GetSyncResultResponseBodyDataRecording {
	s.DurationAudio = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetId(v string) *GetSyncResultResponseBodyDataRecording {
	s.Id = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetName(v string) *GetSyncResultResponseBodyDataRecording {
	s.Name = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetPrimaryId(v string) *GetSyncResultResponseBodyDataRecording {
	s.PrimaryId = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetRemark1(v string) *GetSyncResultResponseBodyDataRecording {
	s.Remark1 = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetRemark2(v string) *GetSyncResultResponseBodyDataRecording {
	s.Remark2 = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetRemark3(v string) *GetSyncResultResponseBodyDataRecording {
	s.Remark3 = &v
	return s
}

func (s *GetSyncResultResponseBodyDataRecording) SetUrl(v string) *GetSyncResultResponseBodyDataRecording {
	s.Url = &v
	return s
}

type GetSyncResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSyncResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSyncResultResponse) GoString() string {
	return s.String()
}

func (s *GetSyncResultResponse) SetHeaders(v map[string]*string) *GetSyncResultResponse {
	s.Headers = v
	return s
}

func (s *GetSyncResultResponse) SetStatusCode(v int32) *GetSyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSyncResultResponse) SetBody(v *GetSyncResultResponseBody) *GetSyncResultResponse {
	s.Body = v
	return s
}

type GetThesaurusBySynonymForApiRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetThesaurusBySynonymForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s GetThesaurusBySynonymForApiRequest) GoString() string {
	return s.String()
}

func (s *GetThesaurusBySynonymForApiRequest) SetJsonStr(v string) *GetThesaurusBySynonymForApiRequest {
	s.JsonStr = &v
	return s
}

type GetThesaurusBySynonymForApiResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetThesaurusBySynonymForApiResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetThesaurusBySynonymForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetThesaurusBySynonymForApiResponseBody) GoString() string {
	return s.String()
}

func (s *GetThesaurusBySynonymForApiResponseBody) SetCode(v string) *GetThesaurusBySynonymForApiResponseBody {
	s.Code = &v
	return s
}

func (s *GetThesaurusBySynonymForApiResponseBody) SetData(v *GetThesaurusBySynonymForApiResponseBodyData) *GetThesaurusBySynonymForApiResponseBody {
	s.Data = v
	return s
}

func (s *GetThesaurusBySynonymForApiResponseBody) SetMessage(v string) *GetThesaurusBySynonymForApiResponseBody {
	s.Message = &v
	return s
}

func (s *GetThesaurusBySynonymForApiResponseBody) SetRequestId(v string) *GetThesaurusBySynonymForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetThesaurusBySynonymForApiResponseBody) SetSuccess(v bool) *GetThesaurusBySynonymForApiResponseBody {
	s.Success = &v
	return s
}

type GetThesaurusBySynonymForApiResponseBodyData struct {
	ThesaurusPo []*GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo `json:"ThesaurusPo,omitempty" xml:"ThesaurusPo,omitempty" type:"Repeated"`
}

func (s GetThesaurusBySynonymForApiResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetThesaurusBySynonymForApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetThesaurusBySynonymForApiResponseBodyData) SetThesaurusPo(v []*GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo) *GetThesaurusBySynonymForApiResponseBodyData {
	s.ThesaurusPo = v
	return s
}

type GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo struct {
	Business    *string                                                            `json:"Business,omitempty" xml:"Business,omitempty"`
	Id          *int64                                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	SynonymList *GetThesaurusBySynonymForApiResponseBodyDataThesaurusPoSynonymList `json:"SynonymList,omitempty" xml:"SynonymList,omitempty" type:"Struct"`
}

func (s GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo) String() string {
	return tea.Prettify(s)
}

func (s GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo) GoString() string {
	return s.String()
}

func (s *GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo) SetBusiness(v string) *GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo {
	s.Business = &v
	return s
}

func (s *GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo) SetId(v int64) *GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo {
	s.Id = &v
	return s
}

func (s *GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo) SetSynonymList(v *GetThesaurusBySynonymForApiResponseBodyDataThesaurusPoSynonymList) *GetThesaurusBySynonymForApiResponseBodyDataThesaurusPo {
	s.SynonymList = v
	return s
}

type GetThesaurusBySynonymForApiResponseBodyDataThesaurusPoSynonymList struct {
	SynonymList []*string `json:"SynonymList,omitempty" xml:"SynonymList,omitempty" type:"Repeated"`
}

func (s GetThesaurusBySynonymForApiResponseBodyDataThesaurusPoSynonymList) String() string {
	return tea.Prettify(s)
}

func (s GetThesaurusBySynonymForApiResponseBodyDataThesaurusPoSynonymList) GoString() string {
	return s.String()
}

func (s *GetThesaurusBySynonymForApiResponseBodyDataThesaurusPoSynonymList) SetSynonymList(v []*string) *GetThesaurusBySynonymForApiResponseBodyDataThesaurusPoSynonymList {
	s.SynonymList = v
	return s
}

type GetThesaurusBySynonymForApiResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetThesaurusBySynonymForApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetThesaurusBySynonymForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s GetThesaurusBySynonymForApiResponse) GoString() string {
	return s.String()
}

func (s *GetThesaurusBySynonymForApiResponse) SetHeaders(v map[string]*string) *GetThesaurusBySynonymForApiResponse {
	s.Headers = v
	return s
}

func (s *GetThesaurusBySynonymForApiResponse) SetStatusCode(v int32) *GetThesaurusBySynonymForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *GetThesaurusBySynonymForApiResponse) SetBody(v *GetThesaurusBySynonymForApiResponseBody) *GetThesaurusBySynonymForApiResponse {
	s.Body = v
	return s
}

type GetWarningStrategyConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetWarningStrategyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigRequest) SetJsonStr(v string) *GetWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

type GetWarningStrategyConfigResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetWarningStrategyConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWarningStrategyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBody) SetCode(v string) *GetWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetData(v *GetWarningStrategyConfigResponseBodyData) *GetWarningStrategyConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetMessage(v string) *GetWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetRequestId(v string) *GetWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBody) SetSuccess(v bool) *GetWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

type GetWarningStrategyConfigResponseBodyData struct {
	Id                  *int64                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	IntervalTime        *int64                                                       `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	Lambda              *string                                                      `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Level               *int64                                                       `json:"Level,omitempty" xml:"Level,omitempty"`
	MaxNumber           *int64                                                       `json:"MaxNumber,omitempty" xml:"MaxNumber,omitempty"`
	Name                *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	WarningStrategyList *GetWarningStrategyConfigResponseBodyDataWarningStrategyList `json:"WarningStrategyList,omitempty" xml:"WarningStrategyList,omitempty" type:"Struct"`
}

func (s GetWarningStrategyConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyData) SetId(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetIntervalTime(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.IntervalTime = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetLambda(v string) *GetWarningStrategyConfigResponseBodyData {
	s.Lambda = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetLevel(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.Level = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetMaxNumber(v int64) *GetWarningStrategyConfigResponseBodyData {
	s.MaxNumber = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetName(v string) *GetWarningStrategyConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyData) SetWarningStrategyList(v *GetWarningStrategyConfigResponseBodyDataWarningStrategyList) *GetWarningStrategyConfigResponseBodyData {
	s.WarningStrategyList = v
	return s
}

type GetWarningStrategyConfigResponseBodyDataWarningStrategyList struct {
	WarningStrategyList []*GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList `json:"warningStrategyList,omitempty" xml:"warningStrategyList,omitempty" type:"Repeated"`
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyList) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyList) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyList) SetWarningStrategyList(v []*GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) *GetWarningStrategyConfigResponseBodyDataWarningStrategyList {
	s.WarningStrategyList = v
	return s
}

type GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList struct {
	Code                *string                                                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Duration            *int64                                                                               `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationExpression  *int64                                                                               `json:"DurationExpression,omitempty" xml:"DurationExpression,omitempty"`
	HitNumber           *int64                                                                               `json:"HitNumber,omitempty" xml:"HitNumber,omitempty"`
	HitNumberExpression *int64                                                                               `json:"HitNumberExpression,omitempty" xml:"HitNumberExpression,omitempty"`
	HitRuleList         *string                                                                              `json:"HitRuleList,omitempty" xml:"HitRuleList,omitempty"`
	HitType             *int64                                                                               `json:"HitType,omitempty" xml:"HitType,omitempty"`
	Id                  *int64                                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Range               *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
	Status              *int64                                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetCode(v string) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Code = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetDuration(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Duration = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetDurationExpression(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.DurationExpression = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitNumber(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitNumber = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitNumberExpression(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitNumberExpression = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitRuleList(v string) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitRuleList = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetHitType(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.HitType = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetId(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Id = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetRange(v *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Range = v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList) SetStatus(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyList {
	s.Status = &v
	return s
}

type GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange struct {
	RangeNum *int64 `json:"RangeNum,omitempty" xml:"RangeNum,omitempty"`
	Type     *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) SetRangeNum(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange {
	s.RangeNum = &v
	return s
}

func (s *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange) SetType(v int64) *GetWarningStrategyConfigResponseBodyDataWarningStrategyListWarningStrategyListRange {
	s.Type = &v
	return s
}

type GetWarningStrategyConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWarningStrategyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *GetWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *GetWarningStrategyConfigResponse) SetStatusCode(v int32) *GetWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWarningStrategyConfigResponse) SetBody(v *GetWarningStrategyConfigResponseBody) *GetWarningStrategyConfigResponse {
	s.Body = v
	return s
}

type HandleComplaintRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s HandleComplaintRequest) String() string {
	return tea.Prettify(s)
}

func (s HandleComplaintRequest) GoString() string {
	return s.String()
}

func (s *HandleComplaintRequest) SetJsonStr(v string) *HandleComplaintRequest {
	s.JsonStr = &v
	return s
}

type HandleComplaintResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HandleComplaintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandleComplaintResponseBody) GoString() string {
	return s.String()
}

func (s *HandleComplaintResponseBody) SetCode(v string) *HandleComplaintResponseBody {
	s.Code = &v
	return s
}

func (s *HandleComplaintResponseBody) SetData(v string) *HandleComplaintResponseBody {
	s.Data = &v
	return s
}

func (s *HandleComplaintResponseBody) SetMessage(v string) *HandleComplaintResponseBody {
	s.Message = &v
	return s
}

func (s *HandleComplaintResponseBody) SetRequestId(v string) *HandleComplaintResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleComplaintResponseBody) SetSuccess(v bool) *HandleComplaintResponseBody {
	s.Success = &v
	return s
}

type HandleComplaintResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HandleComplaintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HandleComplaintResponse) String() string {
	return tea.Prettify(s)
}

func (s HandleComplaintResponse) GoString() string {
	return s.String()
}

func (s *HandleComplaintResponse) SetHeaders(v map[string]*string) *HandleComplaintResponse {
	s.Headers = v
	return s
}

func (s *HandleComplaintResponse) SetStatusCode(v int32) *HandleComplaintResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleComplaintResponse) SetBody(v *HandleComplaintResponseBody) *HandleComplaintResponse {
	s.Body = v
	return s
}

type InsertScoreForApiRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s InsertScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *InsertScoreForApiRequest) SetJsonStr(v string) *InsertScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type InsertScoreForApiResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *InsertScoreForApiResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *InsertScoreForApiResponseBody) SetCode(v string) *InsertScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *InsertScoreForApiResponseBody) SetData(v *InsertScoreForApiResponseBodyData) *InsertScoreForApiResponseBody {
	s.Data = v
	return s
}

func (s *InsertScoreForApiResponseBody) SetMessage(v string) *InsertScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *InsertScoreForApiResponseBody) SetRequestId(v string) *InsertScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertScoreForApiResponseBody) SetSuccess(v bool) *InsertScoreForApiResponseBody {
	s.Success = &v
	return s
}

type InsertScoreForApiResponseBodyData struct {
	ScoreId   *int64  `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreName *string `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
}

func (s InsertScoreForApiResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InsertScoreForApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertScoreForApiResponseBodyData) SetScoreId(v int64) *InsertScoreForApiResponseBodyData {
	s.ScoreId = &v
	return s
}

func (s *InsertScoreForApiResponseBodyData) SetScoreName(v string) *InsertScoreForApiResponseBodyData {
	s.ScoreName = &v
	return s
}

type InsertScoreForApiResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InsertScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *InsertScoreForApiResponse) SetHeaders(v map[string]*string) *InsertScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *InsertScoreForApiResponse) SetStatusCode(v int32) *InsertScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertScoreForApiResponse) SetBody(v *InsertScoreForApiResponseBody) *InsertScoreForApiResponse {
	s.Body = v
	return s
}

type InsertSubScoreForApiRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s InsertSubScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertSubScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *InsertSubScoreForApiRequest) SetJsonStr(v string) *InsertSubScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type InsertSubScoreForApiResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *InsertSubScoreForApiResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertSubScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertSubScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *InsertSubScoreForApiResponseBody) SetCode(v string) *InsertSubScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *InsertSubScoreForApiResponseBody) SetData(v *InsertSubScoreForApiResponseBodyData) *InsertSubScoreForApiResponseBody {
	s.Data = v
	return s
}

func (s *InsertSubScoreForApiResponseBody) SetMessage(v string) *InsertSubScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *InsertSubScoreForApiResponseBody) SetRequestId(v string) *InsertSubScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertSubScoreForApiResponseBody) SetSuccess(v bool) *InsertSubScoreForApiResponseBody {
	s.Success = &v
	return s
}

type InsertSubScoreForApiResponseBodyData struct {
	ScoreSubId   *int64  `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName *string `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
}

func (s InsertSubScoreForApiResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InsertSubScoreForApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertSubScoreForApiResponseBodyData) SetScoreSubId(v int64) *InsertSubScoreForApiResponseBodyData {
	s.ScoreSubId = &v
	return s
}

func (s *InsertSubScoreForApiResponseBodyData) SetScoreSubName(v string) *InsertSubScoreForApiResponseBodyData {
	s.ScoreSubName = &v
	return s
}

type InsertSubScoreForApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InsertSubScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertSubScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertSubScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *InsertSubScoreForApiResponse) SetHeaders(v map[string]*string) *InsertSubScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *InsertSubScoreForApiResponse) SetStatusCode(v int32) *InsertSubScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertSubScoreForApiResponse) SetBody(v *InsertSubScoreForApiResponseBody) *InsertSubScoreForApiResponse {
	s.Body = v
	return s
}

type InvalidRuleRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s InvalidRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s InvalidRuleRequest) GoString() string {
	return s.String()
}

func (s *InvalidRuleRequest) SetJsonStr(v string) *InvalidRuleRequest {
	s.JsonStr = &v
	return s
}

type InvalidRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InvalidRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvalidRuleResponseBody) GoString() string {
	return s.String()
}

func (s *InvalidRuleResponseBody) SetCode(v string) *InvalidRuleResponseBody {
	s.Code = &v
	return s
}

func (s *InvalidRuleResponseBody) SetData(v bool) *InvalidRuleResponseBody {
	s.Data = &v
	return s
}

func (s *InvalidRuleResponseBody) SetMessage(v string) *InvalidRuleResponseBody {
	s.Message = &v
	return s
}

func (s *InvalidRuleResponseBody) SetRequestId(v string) *InvalidRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvalidRuleResponseBody) SetSuccess(v bool) *InvalidRuleResponseBody {
	s.Success = &v
	return s
}

type InvalidRuleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvalidRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvalidRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s InvalidRuleResponse) GoString() string {
	return s.String()
}

func (s *InvalidRuleResponse) SetHeaders(v map[string]*string) *InvalidRuleResponse {
	s.Headers = v
	return s
}

func (s *InvalidRuleResponse) SetStatusCode(v int32) *InvalidRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *InvalidRuleResponse) SetBody(v *InvalidRuleResponseBody) *InvalidRuleResponse {
	s.Body = v
	return s
}

type ListAsrVocabRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListAsrVocabRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *ListAsrVocabRequest) SetJsonStr(v string) *ListAsrVocabRequest {
	s.JsonStr = &v
	return s
}

type ListAsrVocabResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListAsrVocabResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAsrVocabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponseBody) SetCode(v string) *ListAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *ListAsrVocabResponseBody) SetData(v *ListAsrVocabResponseBodyData) *ListAsrVocabResponseBody {
	s.Data = v
	return s
}

func (s *ListAsrVocabResponseBody) SetMessage(v string) *ListAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *ListAsrVocabResponseBody) SetRequestId(v string) *ListAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsrVocabResponseBody) SetSuccess(v bool) *ListAsrVocabResponseBody {
	s.Success = &v
	return s
}

type ListAsrVocabResponseBodyData struct {
	AsrVocab []*ListAsrVocabResponseBodyDataAsrVocab `json:"AsrVocab,omitempty" xml:"AsrVocab,omitempty" type:"Repeated"`
}

func (s ListAsrVocabResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAsrVocabResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponseBodyData) SetAsrVocab(v []*ListAsrVocabResponseBodyDataAsrVocab) *ListAsrVocabResponseBodyData {
	s.AsrVocab = v
	return s
}

type ListAsrVocabResponseBodyDataAsrVocab struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UpdateTime   *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s ListAsrVocabResponseBodyDataAsrVocab) String() string {
	return tea.Prettify(s)
}

func (s ListAsrVocabResponseBodyDataAsrVocab) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetCreateTime(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.CreateTime = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetId(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.Id = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetName(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.Name = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetUpdateTime(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.UpdateTime = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetVocabularyId(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.VocabularyId = &v
	return s
}

type ListAsrVocabResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAsrVocabResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponse) SetHeaders(v map[string]*string) *ListAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *ListAsrVocabResponse) SetStatusCode(v int32) *ListAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsrVocabResponse) SetBody(v *ListAsrVocabResponseBody) *ListAsrVocabResponse {
	s.Body = v
	return s
}

type ListBusinessSpacesRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListBusinessSpacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessSpacesRequest) GoString() string {
	return s.String()
}

func (s *ListBusinessSpacesRequest) SetJsonStr(v string) *ListBusinessSpacesRequest {
	s.JsonStr = &v
	return s
}

type ListBusinessSpacesResponseBody struct {
	Code        *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*ListBusinessSpacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message     *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBusinessSpacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBusinessSpacesResponseBody) SetCode(v string) *ListBusinessSpacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListBusinessSpacesResponseBody) SetCurrentPage(v int32) *ListBusinessSpacesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListBusinessSpacesResponseBody) SetData(v []*ListBusinessSpacesResponseBodyData) *ListBusinessSpacesResponseBody {
	s.Data = v
	return s
}

func (s *ListBusinessSpacesResponseBody) SetMessage(v string) *ListBusinessSpacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListBusinessSpacesResponseBody) SetPageNumber(v int32) *ListBusinessSpacesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBusinessSpacesResponseBody) SetPageSize(v int32) *ListBusinessSpacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBusinessSpacesResponseBody) SetRequestId(v string) *ListBusinessSpacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBusinessSpacesResponseBody) SetSuccess(v bool) *ListBusinessSpacesResponseBody {
	s.Success = &v
	return s
}

type ListBusinessSpacesResponseBodyData struct {
	AliUid              *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	BusinessSpaceCode   *string `json:"BusinessSpaceCode,omitempty" xml:"BusinessSpaceCode,omitempty"`
	BusinessSpaceName   *string `json:"BusinessSpaceName,omitempty" xml:"BusinessSpaceName,omitempty"`
	CurrentStatus       *int32  `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Language            *string `json:"Language,omitempty" xml:"Language,omitempty"`
	LoginUserType       *int32  `json:"LoginUserType,omitempty" xml:"LoginUserType,omitempty"`
	OrderInstanceId     *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	ProductType         *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RoleName            *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SubAliUid           *int64  `json:"SubAliUid,omitempty" xml:"SubAliUid,omitempty"`
	XspaceCommodityCode *string `json:"XspaceCommodityCode,omitempty" xml:"XspaceCommodityCode,omitempty"`
	XspaceProductCode   *string `json:"XspaceProductCode,omitempty" xml:"XspaceProductCode,omitempty"`
}

func (s ListBusinessSpacesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessSpacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBusinessSpacesResponseBodyData) SetAliUid(v int64) *ListBusinessSpacesResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetBusinessSpaceCode(v string) *ListBusinessSpacesResponseBodyData {
	s.BusinessSpaceCode = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetBusinessSpaceName(v string) *ListBusinessSpacesResponseBodyData {
	s.BusinessSpaceName = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetCurrentStatus(v int32) *ListBusinessSpacesResponseBodyData {
	s.CurrentStatus = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetEndTime(v string) *ListBusinessSpacesResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetId(v int64) *ListBusinessSpacesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetLanguage(v string) *ListBusinessSpacesResponseBodyData {
	s.Language = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetLoginUserType(v int32) *ListBusinessSpacesResponseBodyData {
	s.LoginUserType = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetOrderInstanceId(v string) *ListBusinessSpacesResponseBodyData {
	s.OrderInstanceId = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetProductType(v string) *ListBusinessSpacesResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetRoleName(v string) *ListBusinessSpacesResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetStartTime(v string) *ListBusinessSpacesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetSubAliUid(v int64) *ListBusinessSpacesResponseBodyData {
	s.SubAliUid = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetXspaceCommodityCode(v string) *ListBusinessSpacesResponseBodyData {
	s.XspaceCommodityCode = &v
	return s
}

func (s *ListBusinessSpacesResponseBodyData) SetXspaceProductCode(v string) *ListBusinessSpacesResponseBodyData {
	s.XspaceProductCode = &v
	return s
}

type ListBusinessSpacesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBusinessSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBusinessSpacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessSpacesResponse) GoString() string {
	return s.String()
}

func (s *ListBusinessSpacesResponse) SetHeaders(v map[string]*string) *ListBusinessSpacesResponse {
	s.Headers = v
	return s
}

func (s *ListBusinessSpacesResponse) SetStatusCode(v int32) *ListBusinessSpacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBusinessSpacesResponse) SetBody(v *ListBusinessSpacesResponseBody) *ListBusinessSpacesResponse {
	s.Body = v
	return s
}

type ListHotWordsTasksRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListHotWordsTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotWordsTasksRequest) GoString() string {
	return s.String()
}

func (s *ListHotWordsTasksRequest) SetJsonStr(v string) *ListHotWordsTasksRequest {
	s.JsonStr = &v
	return s
}

type ListHotWordsTasksResponseBody struct {
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *int32                             `json:"Count,omitempty" xml:"Count,omitempty"`
	Data       *ListHotWordsTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHotWordsTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotWordsTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotWordsTasksResponseBody) SetCode(v string) *ListHotWordsTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotWordsTasksResponseBody) SetCount(v int32) *ListHotWordsTasksResponseBody {
	s.Count = &v
	return s
}

func (s *ListHotWordsTasksResponseBody) SetData(v *ListHotWordsTasksResponseBodyData) *ListHotWordsTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListHotWordsTasksResponseBody) SetMessage(v string) *ListHotWordsTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotWordsTasksResponseBody) SetPageNumber(v int32) *ListHotWordsTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHotWordsTasksResponseBody) SetPageSize(v int32) *ListHotWordsTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHotWordsTasksResponseBody) SetRequestId(v string) *ListHotWordsTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotWordsTasksResponseBody) SetSuccess(v bool) *ListHotWordsTasksResponseBody {
	s.Success = &v
	return s
}

type ListHotWordsTasksResponseBodyData struct {
	HotWordsTaskPo []*ListHotWordsTasksResponseBodyDataHotWordsTaskPo `json:"HotWordsTaskPo,omitempty" xml:"HotWordsTaskPo,omitempty" type:"Repeated"`
}

func (s ListHotWordsTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotWordsTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotWordsTasksResponseBodyData) SetHotWordsTaskPo(v []*ListHotWordsTasksResponseBodyDataHotWordsTaskPo) *ListHotWordsTasksResponseBodyData {
	s.HotWordsTaskPo = v
	return s
}

type ListHotWordsTasksResponseBodyDataHotWordsTaskPo struct {
	DialogueParam     *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam `json:"DialogueParam,omitempty" xml:"DialogueParam,omitempty" type:"Struct"`
	EndTime           *string                                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceStatus    *int32                                                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	LastExecutionTime *string                                                       `json:"LastExecutionTime,omitempty" xml:"LastExecutionTime,omitempty"`
	Message           *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Name              *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	StartTime         *string                                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status            *int32                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskConfigId      *int64                                                        `json:"TaskConfigId,omitempty" xml:"TaskConfigId,omitempty"`
	TimeInterval      *int32                                                        `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	TimeUnit          *int32                                                        `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	Type              *int32                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	WordsParam        *ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam    `json:"WordsParam,omitempty" xml:"WordsParam,omitempty" type:"Struct"`
}

func (s ListHotWordsTasksResponseBodyDataHotWordsTaskPo) String() string {
	return tea.Prettify(s)
}

func (s ListHotWordsTasksResponseBodyDataHotWordsTaskPo) GoString() string {
	return s.String()
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetDialogueParam(v *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.DialogueParam = v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetEndTime(v string) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.EndTime = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetInstanceStatus(v int32) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.InstanceStatus = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetLastExecutionTime(v string) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.LastExecutionTime = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetMessage(v string) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.Message = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetName(v string) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.Name = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetStartTime(v string) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.StartTime = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetStatus(v int32) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.Status = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetTaskConfigId(v int64) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.TaskConfigId = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetTimeInterval(v int32) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.TimeInterval = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetTimeUnit(v int32) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.TimeUnit = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetType(v int32) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.Type = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPo) SetWordsParam(v *ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam) *ListHotWordsTasksResponseBodyDataHotWordsTaskPo {
	s.WordsParam = v
	return s
}

type ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam struct {
	DataSetIds *string `json:"DataSetIds,omitempty" xml:"DataSetIds,omitempty"`
	DialogueId *int64  `json:"DialogueId,omitempty" xml:"DialogueId,omitempty"`
	EndIndex   *int32  `json:"EndIndex,omitempty" xml:"EndIndex,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Role       *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	SourceType *int32  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartIndex *int32  `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) String() string {
	return tea.Prettify(s)
}

func (s ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) GoString() string {
	return s.String()
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) SetDataSetIds(v string) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam {
	s.DataSetIds = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) SetDialogueId(v int64) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam {
	s.DialogueId = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) SetEndIndex(v int32) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam {
	s.EndIndex = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) SetEndTime(v string) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam {
	s.EndTime = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) SetRole(v int32) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam {
	s.Role = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) SetSourceType(v int32) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam {
	s.SourceType = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) SetStartIndex(v int32) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam {
	s.StartIndex = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam) SetStartTime(v string) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoDialogueParam {
	s.StartTime = &v
	return s
}

type ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam struct {
	Excludes      *string `json:"Excludes,omitempty" xml:"Excludes,omitempty"`
	ExtraConfigId *int64  `json:"ExtraConfigId,omitempty" xml:"ExtraConfigId,omitempty"`
	Includes      *string `json:"Includes,omitempty" xml:"Includes,omitempty"`
}

func (s ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam) String() string {
	return tea.Prettify(s)
}

func (s ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam) GoString() string {
	return s.String()
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam) SetExcludes(v string) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam {
	s.Excludes = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam) SetExtraConfigId(v int64) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam {
	s.ExtraConfigId = &v
	return s
}

func (s *ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam) SetIncludes(v string) *ListHotWordsTasksResponseBodyDataHotWordsTaskPoWordsParam {
	s.Includes = &v
	return s
}

type ListHotWordsTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHotWordsTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotWordsTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotWordsTasksResponse) GoString() string {
	return s.String()
}

func (s *ListHotWordsTasksResponse) SetHeaders(v map[string]*string) *ListHotWordsTasksResponse {
	s.Headers = v
	return s
}

func (s *ListHotWordsTasksResponse) SetStatusCode(v int32) *ListHotWordsTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotWordsTasksResponse) SetBody(v *ListHotWordsTasksResponseBody) *ListHotWordsTasksResponse {
	s.Body = v
	return s
}

type ListPrecisionTaskRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListPrecisionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskRequest) SetJsonStr(v string) *ListPrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

type ListPrecisionTaskResponseBody struct {
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *int32                             `json:"Count,omitempty" xml:"Count,omitempty"`
	Data       *ListPrecisionTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPrecisionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBody) SetCode(v string) *ListPrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetCount(v int32) *ListPrecisionTaskResponseBody {
	s.Count = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetData(v *ListPrecisionTaskResponseBodyData) *ListPrecisionTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetMessage(v string) *ListPrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetPageNumber(v int32) *ListPrecisionTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetPageSize(v int32) *ListPrecisionTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetRequestId(v string) *ListPrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetSuccess(v bool) *ListPrecisionTaskResponseBody {
	s.Success = &v
	return s
}

type ListPrecisionTaskResponseBodyData struct {
	PrecisionTask []*ListPrecisionTaskResponseBodyDataPrecisionTask `json:"PrecisionTask,omitempty" xml:"PrecisionTask,omitempty" type:"Repeated"`
}

func (s ListPrecisionTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyData) SetPrecisionTask(v []*ListPrecisionTaskResponseBodyDataPrecisionTask) *ListPrecisionTaskResponseBodyData {
	s.PrecisionTask = v
	return s
}

type ListPrecisionTaskResponseBodyDataPrecisionTask struct {
	CreateTime     *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataSetId      *int64                                                    `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	DataSetName    *string                                                   `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	Duration       *int32                                                    `json:"Duration,omitempty" xml:"Duration,omitempty"`
	IncorrectWords *int32                                                    `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Name           *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Precisions     *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions `json:"Precisions,omitempty" xml:"Precisions,omitempty" type:"Struct"`
	Source         *int32                                                    `json:"Source,omitempty" xml:"Source,omitempty"`
	Status         *int32                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId         *string                                                   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TotalCount     *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UpdateTime     *string                                                   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VerifiedCount  *int32                                                    `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTask) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTask) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetCreateTime(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.CreateTime = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetDataSetId(v int64) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.DataSetId = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetDataSetName(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.DataSetName = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetDuration(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Duration = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetIncorrectWords(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.IncorrectWords = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetName(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Name = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetPrecisions(v *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Precisions = v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetSource(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Source = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetStatus(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Status = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetTaskId(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.TaskId = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetTotalCount(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.TotalCount = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetUpdateTime(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.UpdateTime = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetVerifiedCount(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.VerifiedCount = &v
	return s
}

type ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions struct {
	Precision []*ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision `json:"Precision,omitempty" xml:"Precision,omitempty" type:"Repeated"`
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) SetPrecision(v []*ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions {
	s.Precision = v
	return s
}

type ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision struct {
	CreateTime *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModelId    *int64   `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName  *string  `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Precision  *float32 `json:"Precision,omitempty" xml:"Precision,omitempty"`
	Status     *int32   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId     *string  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetCreateTime(v string) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.CreateTime = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetModelId(v int64) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.ModelId = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetModelName(v string) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.ModelName = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetPrecision(v float32) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.Precision = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetStatus(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.Status = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetTaskId(v string) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.TaskId = &v
	return s
}

type ListPrecisionTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPrecisionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponse) SetHeaders(v map[string]*string) *ListPrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *ListPrecisionTaskResponse) SetStatusCode(v int32) *ListPrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrecisionTaskResponse) SetBody(v *ListPrecisionTaskResponseBody) *ListPrecisionTaskResponse {
	s.Body = v
	return s
}

type ListQualityCheckSchemeRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListQualityCheckSchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeRequest) SetJsonStr(v string) *ListQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

type ListQualityCheckSchemeResponseBody struct {
	Code          *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Count         *int32                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	Data          []*ListQualityCheckSchemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message       *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber    *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCountId *string                                   `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	Success       *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQualityCheckSchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBody) SetCode(v string) *ListQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetCount(v int32) *ListQualityCheckSchemeResponseBody {
	s.Count = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetData(v []*ListQualityCheckSchemeResponseBodyData) *ListQualityCheckSchemeResponseBody {
	s.Data = v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetMessage(v string) *ListQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetPageNumber(v int32) *ListQualityCheckSchemeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetPageSize(v int32) *ListQualityCheckSchemeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetRequestId(v string) *ListQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetResultCountId(v string) *ListQualityCheckSchemeResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetSuccess(v bool) *ListQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

type ListQualityCheckSchemeResponseBodyData struct {
	CreateTime          *string                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserName      *string                                                      `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DataType            *int32                                                       `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description         *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleList            []*ListQualityCheckSchemeResponseBodyDataRuleList            `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	SchemeCheckTypeList []*ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList `json:"SchemeCheckTypeList,omitempty" xml:"SchemeCheckTypeList,omitempty" type:"Repeated"`
	SchemeId            *int64                                                       `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	Status              *int32                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateType        *int32                                                       `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	Type                *int32                                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime          *string                                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUserName      *string                                                      `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
	Version             *int64                                                       `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListQualityCheckSchemeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyData) SetCreateTime(v string) *ListQualityCheckSchemeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetCreateUserName(v string) *ListQualityCheckSchemeResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetDataType(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.DataType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetDescription(v string) *ListQualityCheckSchemeResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetName(v string) *ListQualityCheckSchemeResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetRuleList(v []*ListQualityCheckSchemeResponseBodyDataRuleList) *ListQualityCheckSchemeResponseBodyData {
	s.RuleList = v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetSchemeCheckTypeList(v []*ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) *ListQualityCheckSchemeResponseBodyData {
	s.SchemeCheckTypeList = v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetSchemeId(v int64) *ListQualityCheckSchemeResponseBodyData {
	s.SchemeId = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetStatus(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetTemplateType(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.TemplateType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetType(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetUpdateTime(v string) *ListQualityCheckSchemeResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetUpdateUserName(v string) *ListQualityCheckSchemeResponseBodyData {
	s.UpdateUserName = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetVersion(v int64) *ListQualityCheckSchemeResponseBodyData {
	s.Version = &v
	return s
}

type ListQualityCheckSchemeResponseBodyDataRuleList struct {
	Rules []*ListQualityCheckSchemeResponseBodyDataRuleListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ListQualityCheckSchemeResponseBodyDataRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyDataRuleList) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleList) SetRules(v []*ListQualityCheckSchemeResponseBodyDataRuleListRules) *ListQualityCheckSchemeResponseBodyDataRuleList {
	s.Rules = v
	return s
}

type ListQualityCheckSchemeResponseBodyDataRuleListRules struct {
	CheckType     *int32  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Rid           *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleScoreType *int32  `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	ScoreNum      *int32  `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	ScoreNumType  *int32  `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	ScoreType     *int32  `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	TargetType    *int32  `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListQualityCheckSchemeResponseBodyDataRuleListRules) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyDataRuleListRules) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetCheckType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.CheckType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetName(v string) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.Name = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetRid(v int64) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.Rid = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetRuleScoreType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.RuleScoreType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetScoreNum(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.ScoreNum = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetScoreNumType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.ScoreNumType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetScoreType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.ScoreType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetTargetType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.TargetType = &v
	return s
}

type ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList struct {
	CheckName  *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckType  *int32  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Enable     *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Score      *int32  `json:"Score,omitempty" xml:"Score,omitempty"`
	TargetType *int32  `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetCheckName(v string) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.CheckName = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetCheckType(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.CheckType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetEnable(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.Enable = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetScore(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.Score = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetTargetType(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.TargetType = &v
	return s
}

type ListQualityCheckSchemeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQualityCheckSchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *ListQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *ListQualityCheckSchemeResponse) SetStatusCode(v int32) *ListQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityCheckSchemeResponse) SetBody(v *ListQualityCheckSchemeResponseBody) *ListQualityCheckSchemeResponse {
	s.Body = v
	return s
}

type ListRolesRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) SetJsonStr(v string) *ListRolesRequest {
	s.JsonStr = &v
	return s
}

type ListRolesResponseBody struct {
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListRolesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) SetCode(v string) *ListRolesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRolesResponseBody) SetData(v *ListRolesResponseBodyData) *ListRolesResponseBody {
	s.Data = v
	return s
}

func (s *ListRolesResponseBody) SetMessage(v string) *ListRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) SetSuccess(v bool) *ListRolesResponseBody {
	s.Success = &v
	return s
}

type ListRolesResponseBodyData struct {
	Role []*ListRolesResponseBodyDataRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyData) SetRole(v []*ListRolesResponseBodyDataRole) *ListRolesResponseBodyData {
	s.Role = v
	return s
}

type ListRolesResponseBodyDataRole struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Level       *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListRolesResponseBodyDataRole) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRole) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRole) SetCreateTime(v string) *ListRolesResponseBodyDataRole {
	s.CreateTime = &v
	return s
}

func (s *ListRolesResponseBodyDataRole) SetDisplayName(v string) *ListRolesResponseBodyDataRole {
	s.DisplayName = &v
	return s
}

func (s *ListRolesResponseBodyDataRole) SetId(v int64) *ListRolesResponseBodyDataRole {
	s.Id = &v
	return s
}

func (s *ListRolesResponseBodyDataRole) SetLevel(v int32) *ListRolesResponseBodyDataRole {
	s.Level = &v
	return s
}

func (s *ListRolesResponseBodyDataRole) SetName(v string) *ListRolesResponseBodyDataRole {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyDataRole) SetUpdateTime(v string) *ListRolesResponseBodyDataRole {
	s.UpdateTime = &v
	return s
}

type ListRolesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) SetHeaders(v map[string]*string) *ListRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRolesResponse) SetStatusCode(v int32) *ListRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRolesResponse) SetBody(v *ListRolesResponseBody) *ListRolesResponse {
	s.Body = v
	return s
}

type ListRulesRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) SetJsonStr(v string) *ListRulesRequest {
	s.JsonStr = &v
	return s
}

type ListRulesResponseBody struct {
	Code       *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *int32                       `json:"Count,omitempty" xml:"Count,omitempty"`
	Data       []*ListRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) SetCode(v string) *ListRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRulesResponseBody) SetCount(v int32) *ListRulesResponseBody {
	s.Count = &v
	return s
}

func (s *ListRulesResponseBody) SetData(v []*ListRulesResponseBodyData) *ListRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListRulesResponseBody) SetMessage(v string) *ListRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRulesResponseBody) SetPageNumber(v int32) *ListRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRulesResponseBody) SetPageSize(v int32) *ListRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetSuccess(v bool) *ListRulesResponseBody {
	s.Success = &v
	return s
}

type ListRulesResponseBodyData struct {
	BusinessCategoryNameList []*string `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
	Comments                 *string   `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateTime               *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Name                     *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Rid                      *int64    `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleType                 *int32    `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Type                     *int32    `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName                 *string   `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyData) SetBusinessCategoryNameList(v []*string) *ListRulesResponseBodyData {
	s.BusinessCategoryNameList = v
	return s
}

func (s *ListRulesResponseBodyData) SetComments(v string) *ListRulesResponseBodyData {
	s.Comments = &v
	return s
}

func (s *ListRulesResponseBodyData) SetCreateTime(v string) *ListRulesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListRulesResponseBodyData) SetName(v string) *ListRulesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListRulesResponseBodyData) SetRid(v int64) *ListRulesResponseBodyData {
	s.Rid = &v
	return s
}

func (s *ListRulesResponseBodyData) SetRuleType(v int32) *ListRulesResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *ListRulesResponseBodyData) SetType(v int32) *ListRulesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListRulesResponseBodyData) SetTypeName(v string) *ListRulesResponseBodyData {
	s.TypeName = &v
	return s
}

type ListRulesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetStatusCode(v int32) *ListRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

type ListSkillGroupConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListSkillGroupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigRequest) SetJsonStr(v string) *ListSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

type ListSkillGroupConfigResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListSkillGroupConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSkillGroupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBody) SetCode(v string) *ListSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetData(v *ListSkillGroupConfigResponseBodyData) *ListSkillGroupConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetMessage(v string) *ListSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetRequestId(v string) *ListSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetSuccess(v bool) *ListSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

type ListSkillGroupConfigResponseBodyData struct {
	SkillGroupConfig []*ListSkillGroupConfigResponseBodyDataSkillGroupConfig `json:"SkillGroupConfig,omitempty" xml:"SkillGroupConfig,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyData) SetSkillGroupConfig(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfig) *ListSkillGroupConfigResponseBodyData {
	s.SkillGroupConfig = v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfig struct {
	AllContentQualityCheck *int32                                                                 `json:"AllContentQualityCheck,omitempty" xml:"AllContentQualityCheck,omitempty"`
	AllRids                *string                                                                `json:"AllRids,omitempty" xml:"AllRids,omitempty"`
	AllRuleList            *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList       `json:"AllRuleList,omitempty" xml:"AllRuleList,omitempty" type:"Struct"`
	CreateTime             *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id                     *int64                                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId             *string                                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModelId                *int64                                                                 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName              *string                                                                `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Name                   *string                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	QualityCheckType       *int32                                                                 `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	Rid                    *string                                                                `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleList               *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList          `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	ScreenSwitch           *bool                                                                  `json:"ScreenSwitch,omitempty" xml:"ScreenSwitch,omitempty"`
	SkillGroupFrom         *int32                                                                 `json:"SkillGroupFrom,omitempty" xml:"SkillGroupFrom,omitempty"`
	SkillGroupId           *string                                                                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName         *string                                                                `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	SkillGroupScreens      *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens `json:"SkillGroupScreens,omitempty" xml:"SkillGroupScreens,omitempty" type:"Struct"`
	Status                 *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                   *int32                                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime             *string                                                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VocabId                *int64                                                                 `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
	VocabName              *string                                                                `json:"VocabName,omitempty" xml:"VocabName,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetAllContentQualityCheck(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.AllContentQualityCheck = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetAllRids(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.AllRids = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetAllRuleList(v *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.AllRuleList = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetCreateTime(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.CreateTime = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetId(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Id = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetInstanceId(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetModelId(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.ModelId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetModelName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.ModelName = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Name = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetQualityCheckType(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.QualityCheckType = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetRid(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Rid = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetRuleList(v *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.RuleList = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetScreenSwitch(v bool) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.ScreenSwitch = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupFrom(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupFrom = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupId(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupName = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupScreens(v *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupScreens = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetStatus(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Status = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetType(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Type = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetUpdateTime(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.UpdateTime = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetVocabId(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.VocabId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetVocabName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.VocabName = &v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList struct {
	RuleNameInfo []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) SetRuleNameInfo(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList {
	s.RuleNameInfo = v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo struct {
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) SetRid(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) SetRuleName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList struct {
	RuleNameInfo []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) SetRuleNameInfo(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList {
	s.RuleNameInfo = v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo struct {
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) SetRid(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) SetRuleName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens struct {
	SkillGroupScreen []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen `json:"SkillGroupScreen,omitempty" xml:"SkillGroupScreen,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) SetSkillGroupScreen(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens {
	s.SkillGroupScreen = v
	return s
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen struct {
	DataType *int32  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Symbol   *int32  `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetDataType(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.DataType = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.Name = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetSymbol(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.Symbol = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetValue(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.Value = &v
	return s
}

type ListSkillGroupConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSkillGroupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponse) SetHeaders(v map[string]*string) *ListSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupConfigResponse) SetStatusCode(v int32) *ListSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillGroupConfigResponse) SetBody(v *ListSkillGroupConfigResponseBody) *ListSkillGroupConfigResponse {
	s.Body = v
	return s
}

type ListTaskAssignRulesRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListTaskAssignRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesRequest) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesRequest) SetJsonStr(v string) *ListTaskAssignRulesRequest {
	s.JsonStr = &v
	return s
}

type ListTaskAssignRulesResponseBody struct {
	Code       *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *int32                               `json:"Count,omitempty" xml:"Count,omitempty"`
	Data       *ListTaskAssignRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTaskAssignRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBody) SetCode(v string) *ListTaskAssignRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetCount(v int32) *ListTaskAssignRulesResponseBody {
	s.Count = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetData(v *ListTaskAssignRulesResponseBodyData) *ListTaskAssignRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetMessage(v string) *ListTaskAssignRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetPageNumber(v int32) *ListTaskAssignRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetPageSize(v int32) *ListTaskAssignRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetRequestId(v string) *ListTaskAssignRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBody) SetSuccess(v bool) *ListTaskAssignRulesResponseBody {
	s.Success = &v
	return s
}

type ListTaskAssignRulesResponseBodyData struct {
	TaskAssignRuleInfo []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo `json:"TaskAssignRuleInfo,omitempty" xml:"TaskAssignRuleInfo,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyData) SetTaskAssignRuleInfo(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) *ListTaskAssignRulesResponseBodyData {
	s.TaskAssignRuleInfo = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo struct {
	Agents         *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents       `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Struct"`
	AgentsStr      *string                                                            `json:"AgentsStr,omitempty" xml:"AgentsStr,omitempty"`
	AssignmentType *int32                                                             `json:"AssignmentType,omitempty" xml:"AssignmentType,omitempty"`
	CallTimeEnd    *int64                                                             `json:"CallTimeEnd,omitempty" xml:"CallTimeEnd,omitempty"`
	CallTimeStart  *int64                                                             `json:"CallTimeStart,omitempty" xml:"CallTimeStart,omitempty"`
	CallType       *int32                                                             `json:"CallType,omitempty" xml:"CallType,omitempty"`
	CreateTime     *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DurationMax    *int32                                                             `json:"DurationMax,omitempty" xml:"DurationMax,omitempty"`
	DurationMin    *int32                                                             `json:"DurationMin,omitempty" xml:"DurationMin,omitempty"`
	Enabled        *int32                                                             `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Priority       *int32                                                             `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Reviewers      *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers    `json:"Reviewers,omitempty" xml:"Reviewers,omitempty" type:"Struct"`
	RuleId         *int64                                                             `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName       *string                                                            `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Rules          *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules        `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	SamplingMode   *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode `json:"SamplingMode,omitempty" xml:"SamplingMode,omitempty" type:"Struct"`
	SkillGroups    *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups  `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty" type:"Struct"`
	SkillGroupsStr *string                                                            `json:"SkillGroupsStr,omitempty" xml:"SkillGroupsStr,omitempty"`
	UpdateTime     *string                                                            `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetAgents(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Agents = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetAgentsStr(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.AgentsStr = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetAssignmentType(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.AssignmentType = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCallTimeEnd(v int64) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CallTimeEnd = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCallTimeStart(v int64) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CallTimeStart = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCallType(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CallType = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetCreateTime(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.CreateTime = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetDurationMax(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.DurationMax = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetDurationMin(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.DurationMin = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetEnabled(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Enabled = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetPriority(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Priority = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetReviewers(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Reviewers = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetRuleId(v int64) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.RuleId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetRuleName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.RuleName = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetRules(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.Rules = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetSamplingMode(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.SamplingMode = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetSkillGroups(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.SkillGroups = v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetSkillGroupsStr(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.SkillGroupsStr = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo) SetUpdateTime(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfo {
	s.UpdateTime = &v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents struct {
	Agent []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents) SetAgent(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgents {
	s.Agent = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent struct {
	AgentId   *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) SetAgentId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent {
	s.AgentId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent) SetAgentName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoAgentsAgent {
	s.AgentName = &v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers struct {
	Reviewer []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer `json:"Reviewer,omitempty" xml:"Reviewer,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers) SetReviewer(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewers {
	s.Reviewer = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer struct {
	ReviewerId   *string `json:"ReviewerId,omitempty" xml:"ReviewerId,omitempty"`
	ReviewerName *string `json:"ReviewerName,omitempty" xml:"ReviewerName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) SetReviewerId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer {
	s.ReviewerId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer) SetReviewerName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoReviewersReviewer {
	s.ReviewerName = &v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules struct {
	RuleBasicInfo []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo `json:"RuleBasicInfo,omitempty" xml:"RuleBasicInfo,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules) SetRuleBasicInfo(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRules {
	s.RuleBasicInfo = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Rid  *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) SetName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo {
	s.Name = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo) SetRid(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoRulesRuleBasicInfo {
	s.Rid = &v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode struct {
	AnyNumberOfDraws       *int32                                                                               `json:"AnyNumberOfDraws,omitempty" xml:"AnyNumberOfDraws,omitempty"`
	Designated             *bool                                                                                `json:"Designated,omitempty" xml:"Designated,omitempty"`
	Dimension              *int32                                                                               `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	Limit                  *int32                                                                               `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NumberOfDraws          *int32                                                                               `json:"NumberOfDraws,omitempty" xml:"NumberOfDraws,omitempty"`
	Proportion             *float32                                                                             `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	RandomInspectionNumber *int32                                                                               `json:"RandomInspectionNumber,omitempty" xml:"RandomInspectionNumber,omitempty"`
	SamplingModeAgents     *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents `json:"SamplingModeAgents,omitempty" xml:"SamplingModeAgents,omitempty" type:"Struct"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetAnyNumberOfDraws(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.AnyNumberOfDraws = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetDesignated(v bool) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Designated = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetDimension(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Dimension = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetLimit(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Limit = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetNumberOfDraws(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.NumberOfDraws = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetProportion(v float32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.Proportion = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetRandomInspectionNumber(v int32) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.RandomInspectionNumber = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode) SetSamplingModeAgents(v *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingMode {
	s.SamplingModeAgents = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents struct {
	SamplingModeAgent []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent `json:"SamplingModeAgent,omitempty" xml:"SamplingModeAgent,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents) SetSamplingModeAgent(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgents {
	s.SamplingModeAgent = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent struct {
	AgentId   *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) SetAgentId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent {
	s.AgentId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent) SetAgentName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSamplingModeSamplingModeAgentsSamplingModeAgent {
	s.AgentName = &v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups struct {
	SkillGroup []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty" type:"Repeated"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups) SetSkillGroup(v []*ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroups {
	s.SkillGroup = v
	return s
}

type ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup struct {
	SkillId   *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) SetSkillId(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup {
	s.SkillId = &v
	return s
}

func (s *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup) SetSkillName(v string) *ListTaskAssignRulesResponseBodyDataTaskAssignRuleInfoSkillGroupsSkillGroup {
	s.SkillName = &v
	return s
}

type ListTaskAssignRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTaskAssignRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTaskAssignRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskAssignRulesResponse) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponse) SetHeaders(v map[string]*string) *ListTaskAssignRulesResponse {
	s.Headers = v
	return s
}

func (s *ListTaskAssignRulesResponse) SetStatusCode(v int32) *ListTaskAssignRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskAssignRulesResponse) SetBody(v *ListTaskAssignRulesResponseBody) *ListTaskAssignRulesResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetJsonStr(v string) *ListUsersRequest {
	s.JsonStr = &v
	return s
}

type ListUsersResponseBody struct {
	Code       *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *int32                     `json:"Count,omitempty" xml:"Count,omitempty"`
	Data       *ListUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetCount(v int32) *ListUsersResponseBody {
	s.Count = &v
	return s
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

type ListUsersResponseBodyData struct {
	User []*ListUsersResponseBodyDataUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetUser(v []*ListUsersResponseBodyDataUser) *ListUsersResponseBodyData {
	s.User = v
	return s
}

type ListUsersResponseBodyDataUser struct {
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LoginUserType *int32  `json:"LoginUserType,omitempty" xml:"LoginUserType,omitempty"`
	RoleName      *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyDataUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataUser) SetAliUid(v string) *ListUsersResponseBodyDataUser {
	s.AliUid = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetCreateTime(v string) *ListUsersResponseBodyDataUser {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetDescription(v string) *ListUsersResponseBodyDataUser {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetDisplayName(v string) *ListUsersResponseBodyDataUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetId(v int64) *ListUsersResponseBodyDataUser {
	s.Id = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetLoginUserType(v int32) *ListUsersResponseBodyDataUser {
	s.LoginUserType = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetRoleName(v string) *ListUsersResponseBodyDataUser {
	s.RoleName = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetUpdateTime(v string) *ListUsersResponseBodyDataUser {
	s.UpdateTime = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetUserName(v string) *ListUsersResponseBodyDataUser {
	s.UserName = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListWarningConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListWarningConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *ListWarningConfigRequest) SetJsonStr(v string) *ListWarningConfigRequest {
	s.JsonStr = &v
	return s
}

type ListWarningConfigResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListWarningConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWarningConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBody) SetCode(v string) *ListWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListWarningConfigResponseBody) SetData(v *ListWarningConfigResponseBodyData) *ListWarningConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListWarningConfigResponseBody) SetMessage(v string) *ListWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListWarningConfigResponseBody) SetRequestId(v string) *ListWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarningConfigResponseBody) SetSuccess(v bool) *ListWarningConfigResponseBody {
	s.Success = &v
	return s
}

type ListWarningConfigResponseBodyData struct {
	WarningConfigInfo []*ListWarningConfigResponseBodyDataWarningConfigInfo `json:"WarningConfigInfo,omitempty" xml:"WarningConfigInfo,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyData) SetWarningConfigInfo(v []*ListWarningConfigResponseBodyDataWarningConfigInfo) *ListWarningConfigResponseBodyData {
	s.WarningConfigInfo = v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfo struct {
	Channels   *ListWarningConfigResponseBodyDataWarningConfigInfoChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	ConfigId   *int64                                                      `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigName *string                                                     `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	CreateTime *string                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RidList    *ListWarningConfigResponseBodyDataWarningConfigInfoRidList  `json:"RidList,omitempty" xml:"RidList,omitempty" type:"Struct"`
	RuleList   *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	Status     *int32                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime *string                                                     `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfo) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfo) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetChannels(v *ListWarningConfigResponseBodyDataWarningConfigInfoChannels) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.Channels = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetConfigId(v int64) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.ConfigId = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetConfigName(v string) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.ConfigName = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetCreateTime(v string) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.CreateTime = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetRidList(v *ListWarningConfigResponseBodyDataWarningConfigInfoRidList) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.RidList = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetRuleList(v *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.RuleList = v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetStatus(v int32) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.Status = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfo) SetUpdateTime(v string) *ListWarningConfigResponseBodyDataWarningConfigInfo {
	s.UpdateTime = &v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfoChannels struct {
	Channel []*ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannels) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannels) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannels) SetChannel(v []*ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) *ListWarningConfigResponseBodyDataWarningConfigInfoChannels {
	s.Channel = v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel struct {
	Type *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) SetType(v int32) *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel {
	s.Type = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel) SetUrl(v string) *ListWarningConfigResponseBodyDataWarningConfigInfoChannelsChannel {
	s.Url = &v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfoRidList struct {
	RidList []*string `json:"RidList,omitempty" xml:"RidList,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRidList) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRidList) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRidList) SetRidList(v []*string) *ListWarningConfigResponseBodyDataWarningConfigInfoRidList {
	s.RidList = v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfoRuleList struct {
	WarningRule []*ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule `json:"WarningRule,omitempty" xml:"WarningRule,omitempty" type:"Repeated"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList) SetWarningRule(v []*ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) *ListWarningConfigResponseBodyDataWarningConfigInfoRuleList {
	s.WarningRule = v
	return s
}

type ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule struct {
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) SetRid(v int64) *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule {
	s.Rid = &v
	return s
}

func (s *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule) SetRuleName(v string) *ListWarningConfigResponseBodyDataWarningConfigInfoRuleListWarningRule {
	s.RuleName = &v
	return s
}

type ListWarningConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWarningConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *ListWarningConfigResponse) SetHeaders(v map[string]*string) *ListWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *ListWarningConfigResponse) SetStatusCode(v int32) *ListWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarningConfigResponse) SetBody(v *ListWarningConfigResponseBody) *ListWarningConfigResponse {
	s.Body = v
	return s
}

type ListWarningStrategyConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListWarningStrategyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigRequest) SetJsonStr(v string) *ListWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

type ListWarningStrategyConfigResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Count     *int32                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      *ListWarningStrategyConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWarningStrategyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponseBody) SetCode(v string) *ListWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetCount(v int32) *ListWarningStrategyConfigResponseBody {
	s.Count = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetData(v *ListWarningStrategyConfigResponseBodyData) *ListWarningStrategyConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetMessage(v string) *ListWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetRequestId(v string) *ListWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBody) SetSuccess(v bool) *ListWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

type ListWarningStrategyConfigResponseBodyData struct {
	Data []*ListWarningStrategyConfigResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s ListWarningStrategyConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWarningStrategyConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponseBodyData) SetData(v []*ListWarningStrategyConfigResponseBodyDataData) *ListWarningStrategyConfigResponseBodyData {
	s.Data = v
	return s
}

type ListWarningStrategyConfigResponseBodyDataData struct {
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IntervalTime *int64  `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	Lambda       *string `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	Level        *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	MaxNumber    *int64  `json:"MaxNumber,omitempty" xml:"MaxNumber,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListWarningStrategyConfigResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s ListWarningStrategyConfigResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetId(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.Id = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetIntervalTime(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.IntervalTime = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetLambda(v string) *ListWarningStrategyConfigResponseBodyDataData {
	s.Lambda = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetLevel(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.Level = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetMaxNumber(v int64) *ListWarningStrategyConfigResponseBodyDataData {
	s.MaxNumber = &v
	return s
}

func (s *ListWarningStrategyConfigResponseBodyDataData) SetName(v string) *ListWarningStrategyConfigResponseBodyDataData {
	s.Name = &v
	return s
}

type ListWarningStrategyConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWarningStrategyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *ListWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *ListWarningStrategyConfigResponse) SetStatusCode(v int32) *ListWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarningStrategyConfigResponse) SetBody(v *ListWarningStrategyConfigResponseBody) *ListWarningStrategyConfigResponse {
	s.Body = v
	return s
}

type RestartAsrTaskRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s RestartAsrTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartAsrTaskRequest) GoString() string {
	return s.String()
}

func (s *RestartAsrTaskRequest) SetJsonStr(v string) *RestartAsrTaskRequest {
	s.JsonStr = &v
	return s
}

type RestartAsrTaskResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RestartAsrTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartAsrTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartAsrTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RestartAsrTaskResponseBody) SetCode(v string) *RestartAsrTaskResponseBody {
	s.Code = &v
	return s
}

func (s *RestartAsrTaskResponseBody) SetData(v *RestartAsrTaskResponseBodyData) *RestartAsrTaskResponseBody {
	s.Data = v
	return s
}

func (s *RestartAsrTaskResponseBody) SetMessage(v string) *RestartAsrTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RestartAsrTaskResponseBody) SetRequestId(v string) *RestartAsrTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartAsrTaskResponseBody) SetSuccess(v bool) *RestartAsrTaskResponseBody {
	s.Success = &v
	return s
}

type RestartAsrTaskResponseBodyData struct {
	RestartResult []*RestartAsrTaskResponseBodyDataRestartResult `json:"RestartResult,omitempty" xml:"RestartResult,omitempty" type:"Repeated"`
}

func (s RestartAsrTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RestartAsrTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartAsrTaskResponseBodyData) SetRestartResult(v []*RestartAsrTaskResponseBodyDataRestartResult) *RestartAsrTaskResponseBodyData {
	s.RestartResult = v
	return s
}

type RestartAsrTaskResponseBodyDataRestartResult struct {
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartAsrTaskResponseBodyDataRestartResult) String() string {
	return tea.Prettify(s)
}

func (s RestartAsrTaskResponseBodyDataRestartResult) GoString() string {
	return s.String()
}

func (s *RestartAsrTaskResponseBodyDataRestartResult) SetData(v string) *RestartAsrTaskResponseBodyDataRestartResult {
	s.Data = &v
	return s
}

func (s *RestartAsrTaskResponseBodyDataRestartResult) SetMessage(v string) *RestartAsrTaskResponseBodyDataRestartResult {
	s.Message = &v
	return s
}

func (s *RestartAsrTaskResponseBodyDataRestartResult) SetSuccess(v bool) *RestartAsrTaskResponseBodyDataRestartResult {
	s.Success = &v
	return s
}

type RestartAsrTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartAsrTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartAsrTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartAsrTaskResponse) GoString() string {
	return s.String()
}

func (s *RestartAsrTaskResponse) SetHeaders(v map[string]*string) *RestartAsrTaskResponse {
	s.Headers = v
	return s
}

func (s *RestartAsrTaskResponse) SetStatusCode(v int32) *RestartAsrTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartAsrTaskResponse) SetBody(v *RestartAsrTaskResponseBody) *RestartAsrTaskResponse {
	s.Body = v
	return s
}

type SaveConfigDataSetRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SaveConfigDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveConfigDataSetRequest) GoString() string {
	return s.String()
}

func (s *SaveConfigDataSetRequest) SetJsonStr(v string) *SaveConfigDataSetRequest {
	s.JsonStr = &v
	return s
}

type SaveConfigDataSetResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveConfigDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveConfigDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *SaveConfigDataSetResponseBody) SetCode(v string) *SaveConfigDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *SaveConfigDataSetResponseBody) SetMessage(v string) *SaveConfigDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *SaveConfigDataSetResponseBody) SetRequestId(v string) *SaveConfigDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveConfigDataSetResponseBody) SetSuccess(v bool) *SaveConfigDataSetResponseBody {
	s.Success = &v
	return s
}

type SaveConfigDataSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveConfigDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveConfigDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveConfigDataSetResponse) GoString() string {
	return s.String()
}

func (s *SaveConfigDataSetResponse) SetHeaders(v map[string]*string) *SaveConfigDataSetResponse {
	s.Headers = v
	return s
}

func (s *SaveConfigDataSetResponse) SetStatusCode(v int32) *SaveConfigDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveConfigDataSetResponse) SetBody(v *SaveConfigDataSetResponseBody) *SaveConfigDataSetResponse {
	s.Body = v
	return s
}

type SubmitComplaintRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitComplaintRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitComplaintRequest) GoString() string {
	return s.String()
}

func (s *SubmitComplaintRequest) SetJsonStr(v string) *SubmitComplaintRequest {
	s.JsonStr = &v
	return s
}

type SubmitComplaintResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitComplaintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitComplaintResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitComplaintResponseBody) SetCode(v string) *SubmitComplaintResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetData(v string) *SubmitComplaintResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetMessage(v string) *SubmitComplaintResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetRequestId(v string) *SubmitComplaintResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitComplaintResponseBody) SetSuccess(v bool) *SubmitComplaintResponseBody {
	s.Success = &v
	return s
}

type SubmitComplaintResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitComplaintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitComplaintResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitComplaintResponse) GoString() string {
	return s.String()
}

func (s *SubmitComplaintResponse) SetHeaders(v map[string]*string) *SubmitComplaintResponse {
	s.Headers = v
	return s
}

func (s *SubmitComplaintResponse) SetStatusCode(v int32) *SubmitComplaintResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitComplaintResponse) SetBody(v *SubmitComplaintResponseBody) *SubmitComplaintResponse {
	s.Body = v
	return s
}

type SubmitPrecisionTaskRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitPrecisionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitPrecisionTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitPrecisionTaskRequest) SetJsonStr(v string) *SubmitPrecisionTaskRequest {
	s.JsonStr = &v
	return s
}

type SubmitPrecisionTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitPrecisionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitPrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPrecisionTaskResponseBody) SetCode(v string) *SubmitPrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetData(v string) *SubmitPrecisionTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetMessage(v string) *SubmitPrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetRequestId(v string) *SubmitPrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetSuccess(v bool) *SubmitPrecisionTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitPrecisionTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitPrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitPrecisionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitPrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitPrecisionTaskResponse) SetHeaders(v map[string]*string) *SubmitPrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitPrecisionTaskResponse) SetStatusCode(v int32) *SubmitPrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPrecisionTaskResponse) SetBody(v *SubmitPrecisionTaskResponseBody) *SubmitPrecisionTaskResponse {
	s.Body = v
	return s
}

type SubmitQualityCheckTaskRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitQualityCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitQualityCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitQualityCheckTaskRequest) SetJsonStr(v string) *SubmitQualityCheckTaskRequest {
	s.JsonStr = &v
	return s
}

type SubmitQualityCheckTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitQualityCheckTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitQualityCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitQualityCheckTaskResponseBody) SetCode(v string) *SubmitQualityCheckTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetData(v string) *SubmitQualityCheckTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetMessage(v string) *SubmitQualityCheckTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetRequestId(v string) *SubmitQualityCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitQualityCheckTaskResponseBody) SetSuccess(v bool) *SubmitQualityCheckTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitQualityCheckTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitQualityCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitQualityCheckTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitQualityCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitQualityCheckTaskResponse) SetHeaders(v map[string]*string) *SubmitQualityCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitQualityCheckTaskResponse) SetStatusCode(v int32) *SubmitQualityCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitQualityCheckTaskResponse) SetBody(v *SubmitQualityCheckTaskResponseBody) *SubmitQualityCheckTaskResponse {
	s.Body = v
	return s
}

type SubmitReviewInfoRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitReviewInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitReviewInfoRequest) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoRequest) SetJsonStr(v string) *SubmitReviewInfoRequest {
	s.JsonStr = &v
	return s
}

type SubmitReviewInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitReviewInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitReviewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoResponseBody) SetCode(v string) *SubmitReviewInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetData(v string) *SubmitReviewInfoResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetMessage(v string) *SubmitReviewInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetRequestId(v string) *SubmitReviewInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetSuccess(v bool) *SubmitReviewInfoResponseBody {
	s.Success = &v
	return s
}

type SubmitReviewInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitReviewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitReviewInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitReviewInfoResponse) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoResponse) SetHeaders(v map[string]*string) *SubmitReviewInfoResponse {
	s.Headers = v
	return s
}

func (s *SubmitReviewInfoResponse) SetStatusCode(v int32) *SubmitReviewInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitReviewInfoResponse) SetBody(v *SubmitReviewInfoResponseBody) *SubmitReviewInfoResponse {
	s.Body = v
	return s
}

type SyncQualityCheckRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SyncQualityCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckRequest) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckRequest) SetJsonStr(v string) *SyncQualityCheckRequest {
	s.JsonStr = &v
	return s
}

type SyncQualityCheckResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SyncQualityCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncQualityCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBody) SetCode(v string) *SyncQualityCheckResponseBody {
	s.Code = &v
	return s
}

func (s *SyncQualityCheckResponseBody) SetData(v *SyncQualityCheckResponseBodyData) *SyncQualityCheckResponseBody {
	s.Data = v
	return s
}

func (s *SyncQualityCheckResponseBody) SetMessage(v string) *SyncQualityCheckResponseBody {
	s.Message = &v
	return s
}

func (s *SyncQualityCheckResponseBody) SetRequestId(v string) *SyncQualityCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncQualityCheckResponseBody) SetSuccess(v bool) *SyncQualityCheckResponseBody {
	s.Success = &v
	return s
}

type SyncQualityCheckResponseBodyData struct {
	BeginTime *int64                                   `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Rules     []*SyncQualityCheckResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Score     *int32                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	TaskId    *string                                  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Tid       *string                                  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SyncQualityCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyData) SetBeginTime(v int64) *SyncQualityCheckResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *SyncQualityCheckResponseBodyData) SetRules(v []*SyncQualityCheckResponseBodyDataRules) *SyncQualityCheckResponseBodyData {
	s.Rules = v
	return s
}

func (s *SyncQualityCheckResponseBodyData) SetScore(v int32) *SyncQualityCheckResponseBodyData {
	s.Score = &v
	return s
}

func (s *SyncQualityCheckResponseBodyData) SetTaskId(v string) *SyncQualityCheckResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SyncQualityCheckResponseBodyData) SetTid(v string) *SyncQualityCheckResponseBodyData {
	s.Tid = &v
	return s
}

type SyncQualityCheckResponseBodyDataRules struct {
	Hit      []*SyncQualityCheckResponseBodyDataRulesHit `json:"Hit,omitempty" xml:"Hit,omitempty" type:"Repeated"`
	Rid      *string                                     `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleName *string                                     `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s SyncQualityCheckResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRules) SetHit(v []*SyncQualityCheckResponseBodyDataRulesHit) *SyncQualityCheckResponseBodyDataRules {
	s.Hit = v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRules) SetRid(v string) *SyncQualityCheckResponseBodyDataRules {
	s.Rid = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRules) SetRuleName(v string) *SyncQualityCheckResponseBodyDataRules {
	s.RuleName = &v
	return s
}

type SyncQualityCheckResponseBodyDataRulesHit struct {
	HitKeyWords []*SyncQualityCheckResponseBodyDataRulesHitHitKeyWords `json:"HitKeyWords,omitempty" xml:"HitKeyWords,omitempty" type:"Repeated"`
	Phrase      *SyncQualityCheckResponseBodyDataRulesHitPhrase        `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s SyncQualityCheckResponseBodyDataRulesHit) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRulesHit) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRulesHit) SetHitKeyWords(v []*SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) *SyncQualityCheckResponseBodyDataRulesHit {
	s.HitKeyWords = v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHit) SetPhrase(v *SyncQualityCheckResponseBodyDataRulesHitPhrase) *SyncQualityCheckResponseBodyDataRulesHit {
	s.Phrase = v
	return s
}

type SyncQualityCheckResponseBodyDataRulesHitHitKeyWords struct {
	Cid  *int32  `json:"Cid,omitempty" xml:"Cid,omitempty"`
	From *int32  `json:"From,omitempty" xml:"From,omitempty"`
	Pid  *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	To   *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Val  *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) SetCid(v int32) *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	s.Cid = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) SetFrom(v int32) *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	s.From = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) SetPid(v int32) *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	s.Pid = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) SetTo(v int32) *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	s.To = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords) SetVal(v string) *SyncQualityCheckResponseBodyDataRulesHitHitKeyWords {
	s.Val = &v
	return s
}

type SyncQualityCheckResponseBodyDataRulesHitPhrase struct {
	Begin           *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	EmotionValue    *int32  `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End             *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Identity        *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SilenceDuration *int32  `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	SpeechRate      *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words           *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s SyncQualityCheckResponseBodyDataRulesHitPhrase) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponseBodyDataRulesHitPhrase) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetBegin(v int64) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.Begin = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetEmotionValue(v int32) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.EmotionValue = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetEnd(v int64) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.End = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetIdentity(v string) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.Identity = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetRole(v string) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.Role = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetSilenceDuration(v int32) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.SilenceDuration = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetSpeechRate(v int32) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.SpeechRate = &v
	return s
}

func (s *SyncQualityCheckResponseBodyDataRulesHitPhrase) SetWords(v string) *SyncQualityCheckResponseBodyDataRulesHitPhrase {
	s.Words = &v
	return s
}

type SyncQualityCheckResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SyncQualityCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncQualityCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncQualityCheckResponse) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponse) SetHeaders(v map[string]*string) *SyncQualityCheckResponse {
	s.Headers = v
	return s
}

func (s *SyncQualityCheckResponse) SetStatusCode(v int32) *SyncQualityCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncQualityCheckResponse) SetBody(v *SyncQualityCheckResponseBody) *SyncQualityCheckResponse {
	s.Body = v
	return s
}

type UpdateAsrVocabRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateAsrVocabRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *UpdateAsrVocabRequest) SetJsonStr(v string) *UpdateAsrVocabRequest {
	s.JsonStr = &v
	return s
}

type UpdateAsrVocabResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAsrVocabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAsrVocabResponseBody) SetCode(v string) *UpdateAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetData(v string) *UpdateAsrVocabResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetMessage(v string) *UpdateAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetRequestId(v string) *UpdateAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAsrVocabResponseBody) SetSuccess(v bool) *UpdateAsrVocabResponseBody {
	s.Success = &v
	return s
}

type UpdateAsrVocabResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAsrVocabResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *UpdateAsrVocabResponse) SetHeaders(v map[string]*string) *UpdateAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *UpdateAsrVocabResponse) SetStatusCode(v int32) *UpdateAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAsrVocabResponse) SetBody(v *UpdateAsrVocabResponseBody) *UpdateAsrVocabResponse {
	s.Body = v
	return s
}

type UpdateQualityCheckDataRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateQualityCheckDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQualityCheckDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckDataRequest) SetJsonStr(v string) *UpdateQualityCheckDataRequest {
	s.JsonStr = &v
	return s
}

type UpdateQualityCheckDataResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQualityCheckDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQualityCheckDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckDataResponseBody) SetCode(v string) *UpdateQualityCheckDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateQualityCheckDataResponseBody) SetMessage(v string) *UpdateQualityCheckDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQualityCheckDataResponseBody) SetRequestId(v string) *UpdateQualityCheckDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQualityCheckDataResponseBody) SetSuccess(v bool) *UpdateQualityCheckDataResponseBody {
	s.Success = &v
	return s
}

type UpdateQualityCheckDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateQualityCheckDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateQualityCheckDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQualityCheckDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckDataResponse) SetHeaders(v map[string]*string) *UpdateQualityCheckDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateQualityCheckDataResponse) SetStatusCode(v int32) *UpdateQualityCheckDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQualityCheckDataResponse) SetBody(v *UpdateQualityCheckDataResponseBody) *UpdateQualityCheckDataResponse {
	s.Body = v
	return s
}

type UpdateRuleRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleRequest) SetJsonStr(v string) *UpdateRuleRequest {
	s.JsonStr = &v
	return s
}

type UpdateRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponseBody) SetCode(v string) *UpdateRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleResponseBody) SetData(v string) *UpdateRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRuleResponseBody) SetMessage(v string) *UpdateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleResponseBody) SetRequestId(v string) *UpdateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleResponseBody) SetSuccess(v bool) *UpdateRuleResponseBody {
	s.Success = &v
	return s
}

type UpdateRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponse) SetHeaders(v map[string]*string) *UpdateRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleResponse) SetStatusCode(v int32) *UpdateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleResponse) SetBody(v *UpdateRuleResponseBody) *UpdateRuleResponse {
	s.Body = v
	return s
}

type UpdateScoreForApiRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateScoreForApiRequest) SetJsonStr(v string) *UpdateScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type UpdateScoreForApiResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScoreForApiResponseBody) SetCode(v string) *UpdateScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateScoreForApiResponseBody) SetMessage(v string) *UpdateScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateScoreForApiResponseBody) SetRequestId(v string) *UpdateScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScoreForApiResponseBody) SetSuccess(v bool) *UpdateScoreForApiResponseBody {
	s.Success = &v
	return s
}

type UpdateScoreForApiResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateScoreForApiResponse) SetHeaders(v map[string]*string) *UpdateScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateScoreForApiResponse) SetStatusCode(v int32) *UpdateScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScoreForApiResponse) SetBody(v *UpdateScoreForApiResponseBody) *UpdateScoreForApiResponse {
	s.Body = v
	return s
}

type UpdateSkillGroupConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateSkillGroupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupConfigRequest) SetJsonStr(v string) *UpdateSkillGroupConfigRequest {
	s.JsonStr = &v
	return s
}

type UpdateSkillGroupConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSkillGroupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupConfigResponseBody) SetCode(v string) *UpdateSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSkillGroupConfigResponseBody) SetMessage(v string) *UpdateSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSkillGroupConfigResponseBody) SetRequestId(v string) *UpdateSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillGroupConfigResponseBody) SetSuccess(v bool) *UpdateSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateSkillGroupConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSkillGroupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupConfigResponse) SetHeaders(v map[string]*string) *UpdateSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillGroupConfigResponse) SetStatusCode(v int32) *UpdateSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSkillGroupConfigResponse) SetBody(v *UpdateSkillGroupConfigResponseBody) *UpdateSkillGroupConfigResponse {
	s.Body = v
	return s
}

type UpdateSubScoreForApiRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateSubScoreForApiRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubScoreForApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubScoreForApiRequest) SetJsonStr(v string) *UpdateSubScoreForApiRequest {
	s.JsonStr = &v
	return s
}

type UpdateSubScoreForApiResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSubScoreForApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubScoreForApiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubScoreForApiResponseBody) SetCode(v string) *UpdateSubScoreForApiResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubScoreForApiResponseBody) SetMessage(v string) *UpdateSubScoreForApiResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubScoreForApiResponseBody) SetRequestId(v string) *UpdateSubScoreForApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubScoreForApiResponseBody) SetSuccess(v bool) *UpdateSubScoreForApiResponseBody {
	s.Success = &v
	return s
}

type UpdateSubScoreForApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSubScoreForApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSubScoreForApiResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubScoreForApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubScoreForApiResponse) SetHeaders(v map[string]*string) *UpdateSubScoreForApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubScoreForApiResponse) SetStatusCode(v int32) *UpdateSubScoreForApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubScoreForApiResponse) SetBody(v *UpdateSubScoreForApiResponseBody) *UpdateSubScoreForApiResponse {
	s.Body = v
	return s
}

type UpdateSyncQualityCheckDataRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateSyncQualityCheckDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncQualityCheckDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataRequest) SetJsonStr(v string) *UpdateSyncQualityCheckDataRequest {
	s.JsonStr = &v
	return s
}

type UpdateSyncQualityCheckDataResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UpdateSyncQualityCheckDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSyncQualityCheckDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncQualityCheckDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetCode(v string) *UpdateSyncQualityCheckDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetData(v *UpdateSyncQualityCheckDataResponseBodyData) *UpdateSyncQualityCheckDataResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetMessage(v string) *UpdateSyncQualityCheckDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetRequestId(v string) *UpdateSyncQualityCheckDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBody) SetSuccess(v bool) *UpdateSyncQualityCheckDataResponseBody {
	s.Success = &v
	return s
}

type UpdateSyncQualityCheckDataResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Tid    *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateSyncQualityCheckDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncQualityCheckDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataResponseBodyData) SetTaskId(v string) *UpdateSyncQualityCheckDataResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponseBodyData) SetTid(v string) *UpdateSyncQualityCheckDataResponseBodyData {
	s.Tid = &v
	return s
}

type UpdateSyncQualityCheckDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSyncQualityCheckDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSyncQualityCheckDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncQualityCheckDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataResponse) SetHeaders(v map[string]*string) *UpdateSyncQualityCheckDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateSyncQualityCheckDataResponse) SetStatusCode(v int32) *UpdateSyncQualityCheckDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSyncQualityCheckDataResponse) SetBody(v *UpdateSyncQualityCheckDataResponseBody) *UpdateSyncQualityCheckDataResponse {
	s.Body = v
	return s
}

type UpdateTaskAssignRuleRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateTaskAssignRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskAssignRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskAssignRuleRequest) SetJsonStr(v string) *UpdateTaskAssignRuleRequest {
	s.JsonStr = &v
	return s
}

type UpdateTaskAssignRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskAssignRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskAssignRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskAssignRuleResponseBody) SetCode(v string) *UpdateTaskAssignRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTaskAssignRuleResponseBody) SetMessage(v string) *UpdateTaskAssignRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTaskAssignRuleResponseBody) SetRequestId(v string) *UpdateTaskAssignRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskAssignRuleResponseBody) SetSuccess(v bool) *UpdateTaskAssignRuleResponseBody {
	s.Success = &v
	return s
}

type UpdateTaskAssignRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskAssignRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskAssignRuleResponse) SetHeaders(v map[string]*string) *UpdateTaskAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskAssignRuleResponse) SetStatusCode(v int32) *UpdateTaskAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskAssignRuleResponse) SetBody(v *UpdateTaskAssignRuleResponseBody) *UpdateTaskAssignRuleResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetJsonStr(v string) *UpdateUserRequest {
	s.JsonStr = &v
	return s
}

type UpdateUserResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetCode(v string) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetSuccess(v bool) *UpdateUserResponseBody {
	s.Success = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

type UpdateUserConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateUserConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserConfigRequest) SetJsonStr(v string) *UpdateUserConfigRequest {
	s.JsonStr = &v
	return s
}

type UpdateUserConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserConfigResponseBody) SetCode(v string) *UpdateUserConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserConfigResponseBody) SetMessage(v string) *UpdateUserConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserConfigResponseBody) SetRequestId(v string) *UpdateUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserConfigResponseBody) SetSuccess(v bool) *UpdateUserConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateUserConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserConfigResponse) SetHeaders(v map[string]*string) *UpdateUserConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserConfigResponse) SetStatusCode(v int32) *UpdateUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserConfigResponse) SetBody(v *UpdateUserConfigResponseBody) *UpdateUserConfigResponse {
	s.Body = v
	return s
}

type UpdateWarningConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateWarningConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateWarningConfigRequest) SetJsonStr(v string) *UpdateWarningConfigRequest {
	s.JsonStr = &v
	return s
}

type UpdateWarningConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWarningConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWarningConfigResponseBody) SetCode(v string) *UpdateWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWarningConfigResponseBody) SetMessage(v string) *UpdateWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWarningConfigResponseBody) SetRequestId(v string) *UpdateWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWarningConfigResponseBody) SetSuccess(v bool) *UpdateWarningConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateWarningConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWarningConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateWarningConfigResponse) SetHeaders(v map[string]*string) *UpdateWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateWarningConfigResponse) SetStatusCode(v int32) *UpdateWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWarningConfigResponse) SetBody(v *UpdateWarningConfigResponseBody) *UpdateWarningConfigResponse {
	s.Body = v
	return s
}

type UpdateWarningStrategyConfigRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateWarningStrategyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningStrategyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateWarningStrategyConfigRequest) SetJsonStr(v string) *UpdateWarningStrategyConfigRequest {
	s.JsonStr = &v
	return s
}

type UpdateWarningStrategyConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWarningStrategyConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningStrategyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWarningStrategyConfigResponseBody) SetCode(v string) *UpdateWarningStrategyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetData(v string) *UpdateWarningStrategyConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetMessage(v string) *UpdateWarningStrategyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetRequestId(v string) *UpdateWarningStrategyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponseBody) SetSuccess(v bool) *UpdateWarningStrategyConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateWarningStrategyConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWarningStrategyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *UpdateWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateWarningStrategyConfigResponse) SetStatusCode(v int32) *UpdateWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponse) SetBody(v *UpdateWarningStrategyConfigResponseBody) *UpdateWarningStrategyConfigResponse {
	s.Body = v
	return s
}

type UploadAudioDataRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadAudioDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadAudioDataRequest) GoString() string {
	return s.String()
}

func (s *UploadAudioDataRequest) SetJsonStr(v string) *UploadAudioDataRequest {
	s.JsonStr = &v
	return s
}

type UploadAudioDataResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadAudioDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadAudioDataResponseBody) GoString() string {
	return s.String()
}

func (s *UploadAudioDataResponseBody) SetCode(v string) *UploadAudioDataResponseBody {
	s.Code = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetData(v string) *UploadAudioDataResponseBody {
	s.Data = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetMessage(v string) *UploadAudioDataResponseBody {
	s.Message = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetRequestId(v string) *UploadAudioDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadAudioDataResponseBody) SetSuccess(v bool) *UploadAudioDataResponseBody {
	s.Success = &v
	return s
}

type UploadAudioDataResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadAudioDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadAudioDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadAudioDataResponse) GoString() string {
	return s.String()
}

func (s *UploadAudioDataResponse) SetHeaders(v map[string]*string) *UploadAudioDataResponse {
	s.Headers = v
	return s
}

func (s *UploadAudioDataResponse) SetStatusCode(v int32) *UploadAudioDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadAudioDataResponse) SetBody(v *UploadAudioDataResponseBody) *UploadAudioDataResponse {
	s.Body = v
	return s
}

type UploadDataRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDataRequest) GoString() string {
	return s.String()
}

func (s *UploadDataRequest) SetJsonStr(v string) *UploadDataRequest {
	s.JsonStr = &v
	return s
}

type UploadDataResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDataResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDataResponseBody) SetCode(v string) *UploadDataResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDataResponseBody) SetData(v string) *UploadDataResponseBody {
	s.Data = &v
	return s
}

func (s *UploadDataResponseBody) SetMessage(v string) *UploadDataResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDataResponseBody) SetRequestId(v string) *UploadDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDataResponseBody) SetSuccess(v bool) *UploadDataResponseBody {
	s.Success = &v
	return s
}

type UploadDataResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadDataResponse) GoString() string {
	return s.String()
}

func (s *UploadDataResponse) SetHeaders(v map[string]*string) *UploadDataResponse {
	s.Headers = v
	return s
}

func (s *UploadDataResponse) SetStatusCode(v int32) *UploadDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDataResponse) SetBody(v *UploadDataResponseBody) *UploadDataResponse {
	s.Body = v
	return s
}

type UploadDataSyncRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadDataSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncRequest) GoString() string {
	return s.String()
}

func (s *UploadDataSyncRequest) SetJsonStr(v string) *UploadDataSyncRequest {
	s.JsonStr = &v
	return s
}

type UploadDataSyncResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UploadDataSyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDataSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBody) SetCode(v string) *UploadDataSyncResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDataSyncResponseBody) SetData(v *UploadDataSyncResponseBodyData) *UploadDataSyncResponseBody {
	s.Data = v
	return s
}

func (s *UploadDataSyncResponseBody) SetMessage(v string) *UploadDataSyncResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDataSyncResponseBody) SetRequestId(v string) *UploadDataSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDataSyncResponseBody) SetSuccess(v bool) *UploadDataSyncResponseBody {
	s.Success = &v
	return s
}

type UploadDataSyncResponseBodyData struct {
	ResultInfo []*UploadDataSyncResponseBodyDataResultInfo `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyData) SetResultInfo(v []*UploadDataSyncResponseBodyDataResultInfo) *UploadDataSyncResponseBodyData {
	s.ResultInfo = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfo struct {
	HandScoreIdList *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList `json:"HandScoreIdList,omitempty" xml:"HandScoreIdList,omitempty" type:"Struct"`
	Rules           *UploadDataSyncResponseBodyDataResultInfoRules           `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	Score           *int32                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfo) SetHandScoreIdList(v *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) *UploadDataSyncResponseBodyDataResultInfo {
	s.HandScoreIdList = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfo) SetRules(v *UploadDataSyncResponseBodyDataResultInfoRules) *UploadDataSyncResponseBodyDataResultInfo {
	s.Rules = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfo) SetScore(v int32) *UploadDataSyncResponseBodyDataResultInfo {
	s.Score = &v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoHandScoreIdList struct {
	HandScoreIdList []*string `json:"HandScoreIdList,omitempty" xml:"HandScoreIdList,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList) SetHandScoreIdList(v []*string) *UploadDataSyncResponseBodyDataResultInfoHandScoreIdList {
	s.HandScoreIdList = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRules struct {
	RuleHitInfo []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo `json:"RuleHitInfo,omitempty" xml:"RuleHitInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRules) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRules) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRules) SetRuleHitInfo(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) *UploadDataSyncResponseBodyDataResultInfoRules {
	s.RuleHitInfo = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo struct {
	ConditionInfo *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo `json:"ConditionInfo,omitempty" xml:"ConditionInfo,omitempty" type:"Struct"`
	Hit           *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit           `json:"Hit,omitempty" xml:"Hit,omitempty" type:"Struct"`
	Rid           *string                                                                `json:"Rid,omitempty" xml:"Rid,omitempty"`
	Tid           *string                                                                `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) SetConditionInfo(v *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo {
	s.ConditionInfo = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) SetHit(v *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo {
	s.Hit = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) SetRid(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo {
	s.Rid = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo) SetTid(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfo {
	s.Tid = &v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo struct {
	ConditionBasicInfo []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo `json:"ConditionBasicInfo,omitempty" xml:"ConditionBasicInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo) SetConditionBasicInfo(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfo {
	s.ConditionBasicInfo = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo struct {
	ConditionInfoCid *string `json:"ConditionInfoCid,omitempty" xml:"ConditionInfoCid,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo) SetConditionInfoCid(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoConditionInfoConditionBasicInfo {
	s.ConditionInfoCid = &v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit struct {
	ConditionHitInfo []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo `json:"ConditionHitInfo,omitempty" xml:"ConditionHitInfo,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit) SetConditionHitInfo(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHit {
	s.ConditionHitInfo = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo struct {
	HitCids     *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids     `json:"HitCids,omitempty" xml:"HitCids,omitempty" type:"Struct"`
	HitKeyWords *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords `json:"HitKeyWords,omitempty" xml:"HitKeyWords,omitempty" type:"Struct"`
	Phrase      *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase      `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) SetHitCids(v *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	s.HitCids = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) SetHitKeyWords(v *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	s.HitKeyWords = v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo) SetPhrase(v *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfo {
	s.Phrase = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids struct {
	CidItem []*string `json:"CidItem,omitempty" xml:"CidItem,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids) SetCidItem(v []*string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitCids {
	s.CidItem = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords struct {
	HitKeyWord []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord `json:"HitKeyWord,omitempty" xml:"HitKeyWord,omitempty" type:"Repeated"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords) SetHitKeyWord(v []*UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWords {
	s.HitKeyWord = v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord struct {
	From *int32  `json:"From,omitempty" xml:"From,omitempty"`
	Pid  *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Tid  *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
	To   *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Val  *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetFrom(v int32) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.From = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetPid(v int32) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.Pid = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetTid(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.Tid = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetTo(v int32) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.To = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord) SetVal(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoHitKeyWordsHitKeyWord {
	s.Val = &v
	return s
}

type UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase struct {
	Begin     *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	End       *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Identity  *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Words     *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetBegin(v int64) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Begin = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetBeginTime(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.BeginTime = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetEnd(v int64) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.End = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetIdentity(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Identity = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetRole(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Role = &v
	return s
}

func (s *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase) SetWords(v string) *UploadDataSyncResponseBodyDataResultInfoRulesRuleHitInfoHitConditionHitInfoPhrase {
	s.Words = &v
	return s
}

type UploadDataSyncResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadDataSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadDataSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadDataSyncResponse) GoString() string {
	return s.String()
}

func (s *UploadDataSyncResponse) SetHeaders(v map[string]*string) *UploadDataSyncResponse {
	s.Headers = v
	return s
}

func (s *UploadDataSyncResponse) SetStatusCode(v int32) *UploadDataSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDataSyncResponse) SetBody(v *UploadDataSyncResponseBody) *UploadDataSyncResponse {
	s.Body = v
	return s
}

type UploadRuleRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadRuleRequest) GoString() string {
	return s.String()
}

func (s *UploadRuleRequest) SetJsonStr(v string) *UploadRuleRequest {
	s.JsonStr = &v
	return s
}

type UploadRuleResponseBody struct {
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UploadRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRuleResponseBody) SetCode(v string) *UploadRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UploadRuleResponseBody) SetData(v *UploadRuleResponseBodyData) *UploadRuleResponseBody {
	s.Data = v
	return s
}

func (s *UploadRuleResponseBody) SetMessage(v string) *UploadRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UploadRuleResponseBody) SetRequestId(v string) *UploadRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadRuleResponseBody) SetSuccess(v bool) *UploadRuleResponseBody {
	s.Success = &v
	return s
}

type UploadRuleResponseBodyData struct {
	RidInfo []*string `json:"RidInfo,omitempty" xml:"RidInfo,omitempty" type:"Repeated"`
}

func (s UploadRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UploadRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadRuleResponseBodyData) SetRidInfo(v []*string) *UploadRuleResponseBodyData {
	s.RidInfo = v
	return s
}

type UploadRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadRuleResponse) GoString() string {
	return s.String()
}

func (s *UploadRuleResponse) SetHeaders(v map[string]*string) *UploadRuleResponse {
	s.Headers = v
	return s
}

func (s *UploadRuleResponse) SetStatusCode(v int32) *UploadRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadRuleResponse) SetBody(v *UploadRuleResponseBody) *UploadRuleResponse {
	s.Body = v
	return s
}

type VerifyFileRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s VerifyFileRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyFileRequest) GoString() string {
	return s.String()
}

func (s *VerifyFileRequest) SetJsonStr(v string) *VerifyFileRequest {
	s.JsonStr = &v
	return s
}

type VerifyFileResponseBody struct {
	Code      *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *float32 `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyFileResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyFileResponseBody) SetCode(v string) *VerifyFileResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyFileResponseBody) SetData(v float32) *VerifyFileResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyFileResponseBody) SetMessage(v string) *VerifyFileResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyFileResponseBody) SetRequestId(v string) *VerifyFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyFileResponseBody) SetSuccess(v bool) *VerifyFileResponseBody {
	s.Success = &v
	return s
}

type VerifyFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyFileResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyFileResponse) GoString() string {
	return s.String()
}

func (s *VerifyFileResponse) SetHeaders(v map[string]*string) *VerifyFileResponse {
	s.Headers = v
	return s
}

func (s *VerifyFileResponse) SetStatusCode(v int32) *VerifyFileResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyFileResponse) SetBody(v *VerifyFileResponseBody) *VerifyFileResponse {
	s.Body = v
	return s
}

type VerifySentenceRequest struct {
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s VerifySentenceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceRequest) GoString() string {
	return s.String()
}

func (s *VerifySentenceRequest) SetJsonStr(v string) *VerifySentenceRequest {
	s.JsonStr = &v
	return s
}

type VerifySentenceResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *VerifySentenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	IncorrectWords *int32                          `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceRole     *int32                          `json:"SourceRole,omitempty" xml:"SourceRole,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TargetRole     *int32                          `json:"TargetRole,omitempty" xml:"TargetRole,omitempty"`
}

func (s VerifySentenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBody) SetCode(v string) *VerifySentenceResponseBody {
	s.Code = &v
	return s
}

func (s *VerifySentenceResponseBody) SetData(v *VerifySentenceResponseBodyData) *VerifySentenceResponseBody {
	s.Data = v
	return s
}

func (s *VerifySentenceResponseBody) SetIncorrectWords(v int32) *VerifySentenceResponseBody {
	s.IncorrectWords = &v
	return s
}

func (s *VerifySentenceResponseBody) SetMessage(v string) *VerifySentenceResponseBody {
	s.Message = &v
	return s
}

func (s *VerifySentenceResponseBody) SetRequestId(v string) *VerifySentenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifySentenceResponseBody) SetSourceRole(v int32) *VerifySentenceResponseBody {
	s.SourceRole = &v
	return s
}

func (s *VerifySentenceResponseBody) SetSuccess(v bool) *VerifySentenceResponseBody {
	s.Success = &v
	return s
}

func (s *VerifySentenceResponseBody) SetTargetRole(v int32) *VerifySentenceResponseBody {
	s.TargetRole = &v
	return s
}

type VerifySentenceResponseBodyData struct {
	Delta []*VerifySentenceResponseBodyDataDelta `json:"Delta,omitempty" xml:"Delta,omitempty" type:"Repeated"`
}

func (s VerifySentenceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyData) SetDelta(v []*VerifySentenceResponseBodyDataDelta) *VerifySentenceResponseBodyData {
	s.Delta = v
	return s
}

type VerifySentenceResponseBodyDataDelta struct {
	Source *VerifySentenceResponseBodyDataDeltaSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Target *VerifySentenceResponseBodyDataDeltaTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	Type   *string                                    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s VerifySentenceResponseBodyDataDelta) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDelta) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDelta) SetSource(v *VerifySentenceResponseBodyDataDeltaSource) *VerifySentenceResponseBodyDataDelta {
	s.Source = v
	return s
}

func (s *VerifySentenceResponseBodyDataDelta) SetTarget(v *VerifySentenceResponseBodyDataDeltaTarget) *VerifySentenceResponseBodyDataDelta {
	s.Target = v
	return s
}

func (s *VerifySentenceResponseBodyDataDelta) SetType(v string) *VerifySentenceResponseBodyDataDelta {
	s.Type = &v
	return s
}

type VerifySentenceResponseBodyDataDeltaSource struct {
	Line     *VerifySentenceResponseBodyDataDeltaSourceLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	Position *int32                                         `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s VerifySentenceResponseBodyDataDeltaSource) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaSource) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaSource) SetLine(v *VerifySentenceResponseBodyDataDeltaSourceLine) *VerifySentenceResponseBodyDataDeltaSource {
	s.Line = v
	return s
}

func (s *VerifySentenceResponseBodyDataDeltaSource) SetPosition(v int32) *VerifySentenceResponseBodyDataDeltaSource {
	s.Position = &v
	return s
}

type VerifySentenceResponseBodyDataDeltaSourceLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s VerifySentenceResponseBodyDataDeltaSourceLine) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaSourceLine) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaSourceLine) SetLine(v []*string) *VerifySentenceResponseBodyDataDeltaSourceLine {
	s.Line = v
	return s
}

type VerifySentenceResponseBodyDataDeltaTarget struct {
	Line     *VerifySentenceResponseBodyDataDeltaTargetLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	Position *int32                                         `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s VerifySentenceResponseBodyDataDeltaTarget) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaTarget) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaTarget) SetLine(v *VerifySentenceResponseBodyDataDeltaTargetLine) *VerifySentenceResponseBodyDataDeltaTarget {
	s.Line = v
	return s
}

func (s *VerifySentenceResponseBodyDataDeltaTarget) SetPosition(v int32) *VerifySentenceResponseBodyDataDeltaTarget {
	s.Position = &v
	return s
}

type VerifySentenceResponseBodyDataDeltaTargetLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s VerifySentenceResponseBodyDataDeltaTargetLine) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponseBodyDataDeltaTargetLine) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponseBodyDataDeltaTargetLine) SetLine(v []*string) *VerifySentenceResponseBodyDataDeltaTargetLine {
	s.Line = v
	return s
}

type VerifySentenceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifySentenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifySentenceResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifySentenceResponse) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponse) SetHeaders(v map[string]*string) *VerifySentenceResponse {
	s.Headers = v
	return s
}

func (s *VerifySentenceResponse) SetStatusCode(v int32) *VerifySentenceResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifySentenceResponse) SetBody(v *VerifySentenceResponseBody) *VerifySentenceResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("qualitycheck"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddBusinessCategoryWithOptions(request *AddBusinessCategoryRequest, runtime *util.RuntimeOptions) (_result *AddBusinessCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddBusinessCategory"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddBusinessCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddBusinessCategory(request *AddBusinessCategoryRequest) (_result *AddBusinessCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddBusinessCategoryResponse{}
	_body, _err := client.AddBusinessCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRuleCategoryWithOptions(request *AddRuleCategoryRequest, runtime *util.RuntimeOptions) (_result *AddRuleCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRuleCategory"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRuleCategory(request *AddRuleCategoryRequest) (_result *AddRuleCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRuleCategoryResponse{}
	_body, _err := client.AddRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddThesaurusForApiWithOptions(request *AddThesaurusForApiRequest, runtime *util.RuntimeOptions) (_result *AddThesaurusForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddThesaurusForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddThesaurusForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddThesaurusForApi(request *AddThesaurusForApiRequest) (_result *AddThesaurusForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddThesaurusForApiResponse{}
	_body, _err := client.AddThesaurusForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssignReviewerWithOptions(request *AssignReviewerRequest, runtime *util.RuntimeOptions) (_result *AssignReviewerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssignReviewer"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignReviewerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssignReviewer(request *AssignReviewerRequest) (_result *AssignReviewerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssignReviewerResponse{}
	_body, _err := client.AssignReviewerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAsrVocabWithOptions(request *CreateAsrVocabRequest, runtime *util.RuntimeOptions) (_result *CreateAsrVocabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAsrVocab"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAsrVocab(request *CreateAsrVocabRequest) (_result *CreateAsrVocabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAsrVocabResponse{}
	_body, _err := client.CreateAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSkillGroupConfigWithOptions(request *CreateSkillGroupConfigRequest, runtime *util.RuntimeOptions) (_result *CreateSkillGroupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSkillGroupConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSkillGroupConfig(request *CreateSkillGroupConfigRequest) (_result *CreateSkillGroupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSkillGroupConfigResponse{}
	_body, _err := client.CreateSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTaskAssignRuleWithOptions(request *CreateTaskAssignRuleRequest, runtime *util.RuntimeOptions) (_result *CreateTaskAssignRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTaskAssignRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTaskAssignRule(request *CreateTaskAssignRuleRequest) (_result *CreateTaskAssignRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTaskAssignRuleResponse{}
	_body, _err := client.CreateTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWarningConfigWithOptions(request *CreateWarningConfigRequest, runtime *util.RuntimeOptions) (_result *CreateWarningConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWarningConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWarningConfig(request *CreateWarningConfigRequest) (_result *CreateWarningConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWarningConfigResponse{}
	_body, _err := client.CreateWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWarningStrategyConfigWithOptions(request *CreateWarningStrategyConfigRequest, runtime *util.RuntimeOptions) (_result *CreateWarningStrategyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWarningStrategyConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWarningStrategyConfig(request *CreateWarningStrategyConfigRequest) (_result *CreateWarningStrategyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWarningStrategyConfigResponse{}
	_body, _err := client.CreateWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DelRuleCategoryWithOptions(request *DelRuleCategoryRequest, runtime *util.RuntimeOptions) (_result *DelRuleCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DelRuleCategory"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DelRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DelRuleCategory(request *DelRuleCategoryRequest) (_result *DelRuleCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DelRuleCategoryResponse{}
	_body, _err := client.DelRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DelThesaurusForApiWithOptions(request *DelThesaurusForApiRequest, runtime *util.RuntimeOptions) (_result *DelThesaurusForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DelThesaurusForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DelThesaurusForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DelThesaurusForApi(request *DelThesaurusForApiRequest) (_result *DelThesaurusForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DelThesaurusForApiResponse{}
	_body, _err := client.DelThesaurusForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAsrVocabWithOptions(request *DeleteAsrVocabRequest, runtime *util.RuntimeOptions) (_result *DeleteAsrVocabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAsrVocab"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAsrVocab(request *DeleteAsrVocabRequest) (_result *DeleteAsrVocabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAsrVocabResponse{}
	_body, _err := client.DeleteAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBusinessCategoryWithOptions(request *DeleteBusinessCategoryRequest, runtime *util.RuntimeOptions) (_result *DeleteBusinessCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBusinessCategory"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBusinessCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBusinessCategory(request *DeleteBusinessCategoryRequest) (_result *DeleteBusinessCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBusinessCategoryResponse{}
	_body, _err := client.DeleteBusinessCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomizationConfigWithOptions(request *DeleteCustomizationConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomizationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomizationConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomizationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomizationConfig(request *DeleteCustomizationConfigRequest) (_result *DeleteCustomizationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomizationConfigResponse{}
	_body, _err := client.DeleteCustomizationConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataSetWithOptions(request *DeleteDataSetRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSet"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataSet(request *DeleteDataSetRequest) (_result *DeleteDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.DeleteDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePrecisionTaskWithOptions(request *DeletePrecisionTaskRequest, runtime *util.RuntimeOptions) (_result *DeletePrecisionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePrecisionTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePrecisionTask(request *DeletePrecisionTaskRequest) (_result *DeletePrecisionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePrecisionTaskResponse{}
	_body, _err := client.DeletePrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScoreForApiWithOptions(request *DeleteScoreForApiRequest, runtime *util.RuntimeOptions) (_result *DeleteScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScoreForApi(request *DeleteScoreForApiRequest) (_result *DeleteScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScoreForApiResponse{}
	_body, _err := client.DeleteScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSkillGroupConfigWithOptions(request *DeleteSkillGroupConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteSkillGroupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSkillGroupConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSkillGroupConfig(request *DeleteSkillGroupConfigRequest) (_result *DeleteSkillGroupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSkillGroupConfigResponse{}
	_body, _err := client.DeleteSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSubScoreForApiWithOptions(request *DeleteSubScoreForApiRequest, runtime *util.RuntimeOptions) (_result *DeleteSubScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSubScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSubScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSubScoreForApi(request *DeleteSubScoreForApiRequest) (_result *DeleteSubScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSubScoreForApiResponse{}
	_body, _err := client.DeleteSubScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTaskAssignRuleWithOptions(request *DeleteTaskAssignRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteTaskAssignRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTaskAssignRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTaskAssignRule(request *DeleteTaskAssignRuleRequest) (_result *DeleteTaskAssignRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTaskAssignRuleResponse{}
	_body, _err := client.DeleteTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWarningConfigWithOptions(request *DeleteWarningConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteWarningConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWarningConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWarningConfig(request *DeleteWarningConfigRequest) (_result *DeleteWarningConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWarningConfigResponse{}
	_body, _err := client.DeleteWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWarningStrategyConfigWithOptions(request *DeleteWarningStrategyConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteWarningStrategyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWarningStrategyConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWarningStrategyConfig(request *DeleteWarningStrategyConfigRequest) (_result *DeleteWarningStrategyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWarningStrategyConfigResponse{}
	_body, _err := client.DeleteWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditThesaurusForApiWithOptions(request *EditThesaurusForApiRequest, runtime *util.RuntimeOptions) (_result *EditThesaurusForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditThesaurusForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditThesaurusForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditThesaurusForApi(request *EditThesaurusForApiRequest) (_result *EditThesaurusForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditThesaurusForApiResponse{}
	_body, _err := client.EditThesaurusForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsrVocabWithOptions(request *GetAsrVocabRequest, runtime *util.RuntimeOptions) (_result *GetAsrVocabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsrVocab"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsrVocab(request *GetAsrVocabRequest) (_result *GetAsrVocabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsrVocabResponse{}
	_body, _err := client.GetAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBusinessCategoryListWithOptions(request *GetBusinessCategoryListRequest, runtime *util.RuntimeOptions) (_result *GetBusinessCategoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBusinessCategoryList"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBusinessCategoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBusinessCategoryList(request *GetBusinessCategoryListRequest) (_result *GetBusinessCategoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBusinessCategoryListResponse{}
	_body, _err := client.GetBusinessCategoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomizationConfigListWithOptions(request *GetCustomizationConfigListRequest, runtime *util.RuntimeOptions) (_result *GetCustomizationConfigListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomizationConfigList"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomizationConfigListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomizationConfigList(request *GetCustomizationConfigListRequest) (_result *GetCustomizationConfigListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomizationConfigListResponse{}
	_body, _err := client.GetCustomizationConfigListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHitResultWithOptions(request *GetHitResultRequest, runtime *util.RuntimeOptions) (_result *GetHitResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHitResult"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHitResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHitResult(request *GetHitResultRequest) (_result *GetHitResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHitResultResponse{}
	_body, _err := client.GetHitResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNextResultToVerifyWithOptions(request *GetNextResultToVerifyRequest, runtime *util.RuntimeOptions) (_result *GetNextResultToVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNextResultToVerify"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNextResultToVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNextResultToVerify(request *GetNextResultToVerifyRequest) (_result *GetNextResultToVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNextResultToVerifyResponse{}
	_body, _err := client.GetNextResultToVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPrecisionTaskWithOptions(request *GetPrecisionTaskRequest, runtime *util.RuntimeOptions) (_result *GetPrecisionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPrecisionTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPrecisionTask(request *GetPrecisionTaskRequest) (_result *GetPrecisionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPrecisionTaskResponse{}
	_body, _err := client.GetPrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResultWithOptions(request *GetResultRequest, runtime *util.RuntimeOptions) (_result *GetResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResult"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResult(request *GetResultRequest) (_result *GetResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResultResponse{}
	_body, _err := client.GetResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResultCallbackWithOptions(request *GetResultCallbackRequest, runtime *util.RuntimeOptions) (_result *GetResultCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResultCallback"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResultCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResultCallback(request *GetResultCallbackRequest) (_result *GetResultCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResultCallbackResponse{}
	_body, _err := client.GetResultCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResultToReviewWithOptions(request *GetResultToReviewRequest, runtime *util.RuntimeOptions) (_result *GetResultToReviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResultToReview"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResultToReviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResultToReview(request *GetResultToReviewRequest) (_result *GetResultToReviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResultToReviewResponse{}
	_body, _err := client.GetResultToReviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRuleWithOptions(request *GetRuleRequest, runtime *util.RuntimeOptions) (_result *GetRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRule(request *GetRuleRequest) (_result *GetRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleResponse{}
	_body, _err := client.GetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRuleCategoryWithOptions(request *GetRuleCategoryRequest, runtime *util.RuntimeOptions) (_result *GetRuleCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRuleCategory"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRuleCategory(request *GetRuleCategoryRequest) (_result *GetRuleCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleCategoryResponse{}
	_body, _err := client.GetRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRuleDetailWithOptions(request *GetRuleDetailRequest, runtime *util.RuntimeOptions) (_result *GetRuleDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRuleDetail"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRuleDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRuleDetail(request *GetRuleDetailRequest) (_result *GetRuleDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleDetailResponse{}
	_body, _err := client.GetRuleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScoreInfoWithOptions(request *GetScoreInfoRequest, runtime *util.RuntimeOptions) (_result *GetScoreInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScoreInfo"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScoreInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScoreInfo(request *GetScoreInfoRequest) (_result *GetScoreInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScoreInfoResponse{}
	_body, _err := client.GetScoreInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSkillGroupConfigWithOptions(request *GetSkillGroupConfigRequest, runtime *util.RuntimeOptions) (_result *GetSkillGroupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSkillGroupConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSkillGroupConfig(request *GetSkillGroupConfigRequest) (_result *GetSkillGroupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSkillGroupConfigResponse{}
	_body, _err := client.GetSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSyncResultWithOptions(request *GetSyncResultRequest, runtime *util.RuntimeOptions) (_result *GetSyncResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSyncResult"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSyncResult(request *GetSyncResultRequest) (_result *GetSyncResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSyncResultResponse{}
	_body, _err := client.GetSyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetThesaurusBySynonymForApiWithOptions(request *GetThesaurusBySynonymForApiRequest, runtime *util.RuntimeOptions) (_result *GetThesaurusBySynonymForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetThesaurusBySynonymForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetThesaurusBySynonymForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetThesaurusBySynonymForApi(request *GetThesaurusBySynonymForApiRequest) (_result *GetThesaurusBySynonymForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetThesaurusBySynonymForApiResponse{}
	_body, _err := client.GetThesaurusBySynonymForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWarningStrategyConfigWithOptions(request *GetWarningStrategyConfigRequest, runtime *util.RuntimeOptions) (_result *GetWarningStrategyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWarningStrategyConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWarningStrategyConfig(request *GetWarningStrategyConfigRequest) (_result *GetWarningStrategyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWarningStrategyConfigResponse{}
	_body, _err := client.GetWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandleComplaintWithOptions(request *HandleComplaintRequest, runtime *util.RuntimeOptions) (_result *HandleComplaintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("HandleComplaint"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HandleComplaintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HandleComplaint(request *HandleComplaintRequest) (_result *HandleComplaintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HandleComplaintResponse{}
	_body, _err := client.HandleComplaintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertScoreForApiWithOptions(request *InsertScoreForApiRequest, runtime *util.RuntimeOptions) (_result *InsertScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertScoreForApi(request *InsertScoreForApiRequest) (_result *InsertScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertScoreForApiResponse{}
	_body, _err := client.InsertScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertSubScoreForApiWithOptions(request *InsertSubScoreForApiRequest, runtime *util.RuntimeOptions) (_result *InsertSubScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertSubScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertSubScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertSubScoreForApi(request *InsertSubScoreForApiRequest) (_result *InsertSubScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertSubScoreForApiResponse{}
	_body, _err := client.InsertSubScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvalidRuleWithOptions(request *InvalidRuleRequest, runtime *util.RuntimeOptions) (_result *InvalidRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvalidRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvalidRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvalidRule(request *InvalidRuleRequest) (_result *InvalidRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvalidRuleResponse{}
	_body, _err := client.InvalidRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAsrVocabWithOptions(request *ListAsrVocabRequest, runtime *util.RuntimeOptions) (_result *ListAsrVocabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAsrVocab"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAsrVocab(request *ListAsrVocabRequest) (_result *ListAsrVocabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAsrVocabResponse{}
	_body, _err := client.ListAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBusinessSpacesWithOptions(request *ListBusinessSpacesRequest, runtime *util.RuntimeOptions) (_result *ListBusinessSpacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBusinessSpaces"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBusinessSpacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBusinessSpaces(request *ListBusinessSpacesRequest) (_result *ListBusinessSpacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBusinessSpacesResponse{}
	_body, _err := client.ListBusinessSpacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotWordsTasksWithOptions(request *ListHotWordsTasksRequest, runtime *util.RuntimeOptions) (_result *ListHotWordsTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotWordsTasks"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotWordsTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotWordsTasks(request *ListHotWordsTasksRequest) (_result *ListHotWordsTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotWordsTasksResponse{}
	_body, _err := client.ListHotWordsTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPrecisionTaskWithOptions(request *ListPrecisionTaskRequest, runtime *util.RuntimeOptions) (_result *ListPrecisionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrecisionTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPrecisionTask(request *ListPrecisionTaskRequest) (_result *ListPrecisionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrecisionTaskResponse{}
	_body, _err := client.ListPrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQualityCheckSchemeWithOptions(request *ListQualityCheckSchemeRequest, runtime *util.RuntimeOptions) (_result *ListQualityCheckSchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQualityCheckScheme"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQualityCheckScheme(request *ListQualityCheckSchemeRequest) (_result *ListQualityCheckSchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQualityCheckSchemeResponse{}
	_body, _err := client.ListQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRolesWithOptions(request *ListRolesRequest, runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoles"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *util.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRules"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSkillGroupConfigWithOptions(request *ListSkillGroupConfigRequest, runtime *util.RuntimeOptions) (_result *ListSkillGroupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSkillGroupConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSkillGroupConfig(request *ListSkillGroupConfigRequest) (_result *ListSkillGroupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSkillGroupConfigResponse{}
	_body, _err := client.ListSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTaskAssignRulesWithOptions(request *ListTaskAssignRulesRequest, runtime *util.RuntimeOptions) (_result *ListTaskAssignRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTaskAssignRules"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTaskAssignRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTaskAssignRules(request *ListTaskAssignRulesRequest) (_result *ListTaskAssignRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskAssignRulesResponse{}
	_body, _err := client.ListTaskAssignRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWarningConfigWithOptions(request *ListWarningConfigRequest, runtime *util.RuntimeOptions) (_result *ListWarningConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWarningConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWarningConfig(request *ListWarningConfigRequest) (_result *ListWarningConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWarningConfigResponse{}
	_body, _err := client.ListWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWarningStrategyConfigWithOptions(request *ListWarningStrategyConfigRequest, runtime *util.RuntimeOptions) (_result *ListWarningStrategyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWarningStrategyConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWarningStrategyConfig(request *ListWarningStrategyConfigRequest) (_result *ListWarningStrategyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWarningStrategyConfigResponse{}
	_body, _err := client.ListWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartAsrTaskWithOptions(request *RestartAsrTaskRequest, runtime *util.RuntimeOptions) (_result *RestartAsrTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartAsrTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartAsrTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartAsrTask(request *RestartAsrTaskRequest) (_result *RestartAsrTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartAsrTaskResponse{}
	_body, _err := client.RestartAsrTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveConfigDataSetWithOptions(request *SaveConfigDataSetRequest, runtime *util.RuntimeOptions) (_result *SaveConfigDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveConfigDataSet"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveConfigDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveConfigDataSet(request *SaveConfigDataSetRequest) (_result *SaveConfigDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveConfigDataSetResponse{}
	_body, _err := client.SaveConfigDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitComplaintWithOptions(request *SubmitComplaintRequest, runtime *util.RuntimeOptions) (_result *SubmitComplaintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitComplaint"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitComplaintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitComplaint(request *SubmitComplaintRequest) (_result *SubmitComplaintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitComplaintResponse{}
	_body, _err := client.SubmitComplaintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitPrecisionTaskWithOptions(request *SubmitPrecisionTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitPrecisionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitPrecisionTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitPrecisionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitPrecisionTask(request *SubmitPrecisionTaskRequest) (_result *SubmitPrecisionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitPrecisionTaskResponse{}
	_body, _err := client.SubmitPrecisionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitQualityCheckTaskWithOptions(request *SubmitQualityCheckTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitQualityCheckTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitQualityCheckTask"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitQualityCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitQualityCheckTask(request *SubmitQualityCheckTaskRequest) (_result *SubmitQualityCheckTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitQualityCheckTaskResponse{}
	_body, _err := client.SubmitQualityCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitReviewInfoWithOptions(request *SubmitReviewInfoRequest, runtime *util.RuntimeOptions) (_result *SubmitReviewInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitReviewInfo"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitReviewInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitReviewInfo(request *SubmitReviewInfoRequest) (_result *SubmitReviewInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitReviewInfoResponse{}
	_body, _err := client.SubmitReviewInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncQualityCheckWithOptions(request *SyncQualityCheckRequest, runtime *util.RuntimeOptions) (_result *SyncQualityCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncQualityCheck"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncQualityCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncQualityCheck(request *SyncQualityCheckRequest) (_result *SyncQualityCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncQualityCheckResponse{}
	_body, _err := client.SyncQualityCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAsrVocabWithOptions(request *UpdateAsrVocabRequest, runtime *util.RuntimeOptions) (_result *UpdateAsrVocabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAsrVocab"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAsrVocab(request *UpdateAsrVocabRequest) (_result *UpdateAsrVocabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAsrVocabResponse{}
	_body, _err := client.UpdateAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateQualityCheckDataWithOptions(request *UpdateQualityCheckDataRequest, runtime *util.RuntimeOptions) (_result *UpdateQualityCheckDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQualityCheckData"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQualityCheckDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateQualityCheckData(request *UpdateQualityCheckDataRequest) (_result *UpdateQualityCheckDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateQualityCheckDataResponse{}
	_body, _err := client.UpdateQualityCheckDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRuleWithOptions(request *UpdateRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRule(request *UpdateRuleRequest) (_result *UpdateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRuleResponse{}
	_body, _err := client.UpdateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateScoreForApiWithOptions(request *UpdateScoreForApiRequest, runtime *util.RuntimeOptions) (_result *UpdateScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateScoreForApi(request *UpdateScoreForApiRequest) (_result *UpdateScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateScoreForApiResponse{}
	_body, _err := client.UpdateScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSkillGroupConfigWithOptions(request *UpdateSkillGroupConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateSkillGroupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSkillGroupConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSkillGroupConfig(request *UpdateSkillGroupConfigRequest) (_result *UpdateSkillGroupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSkillGroupConfigResponse{}
	_body, _err := client.UpdateSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSubScoreForApiWithOptions(request *UpdateSubScoreForApiRequest, runtime *util.RuntimeOptions) (_result *UpdateSubScoreForApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSubScoreForApi"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSubScoreForApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSubScoreForApi(request *UpdateSubScoreForApiRequest) (_result *UpdateSubScoreForApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSubScoreForApiResponse{}
	_body, _err := client.UpdateSubScoreForApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSyncQualityCheckDataWithOptions(request *UpdateSyncQualityCheckDataRequest, runtime *util.RuntimeOptions) (_result *UpdateSyncQualityCheckDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSyncQualityCheckData"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSyncQualityCheckDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSyncQualityCheckData(request *UpdateSyncQualityCheckDataRequest) (_result *UpdateSyncQualityCheckDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSyncQualityCheckDataResponse{}
	_body, _err := client.UpdateSyncQualityCheckDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskAssignRuleWithOptions(request *UpdateTaskAssignRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateTaskAssignRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskAssignRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskAssignRule(request *UpdateTaskAssignRuleRequest) (_result *UpdateTaskAssignRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTaskAssignRuleResponse{}
	_body, _err := client.UpdateTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserConfigWithOptions(request *UpdateUserConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateUserConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserConfig(request *UpdateUserConfigRequest) (_result *UpdateUserConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserConfigResponse{}
	_body, _err := client.UpdateUserConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWarningConfigWithOptions(request *UpdateWarningConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateWarningConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWarningConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWarningConfig(request *UpdateWarningConfigRequest) (_result *UpdateWarningConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWarningConfigResponse{}
	_body, _err := client.UpdateWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWarningStrategyConfigWithOptions(request *UpdateWarningStrategyConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateWarningStrategyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWarningStrategyConfig"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWarningStrategyConfig(request *UpdateWarningStrategyConfigRequest) (_result *UpdateWarningStrategyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWarningStrategyConfigResponse{}
	_body, _err := client.UpdateWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadAudioDataWithOptions(request *UploadAudioDataRequest, runtime *util.RuntimeOptions) (_result *UploadAudioDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadAudioData"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadAudioDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadAudioData(request *UploadAudioDataRequest) (_result *UploadAudioDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadAudioDataResponse{}
	_body, _err := client.UploadAudioDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadDataWithOptions(request *UploadDataRequest, runtime *util.RuntimeOptions) (_result *UploadDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadData"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadData(request *UploadDataRequest) (_result *UploadDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadDataResponse{}
	_body, _err := client.UploadDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadDataSyncWithOptions(request *UploadDataSyncRequest, runtime *util.RuntimeOptions) (_result *UploadDataSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadDataSync"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadDataSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadDataSync(request *UploadDataSyncRequest) (_result *UploadDataSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadDataSyncResponse{}
	_body, _err := client.UploadDataSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadRuleWithOptions(request *UploadRuleRequest, runtime *util.RuntimeOptions) (_result *UploadRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadRule"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadRule(request *UploadRuleRequest) (_result *UploadRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadRuleResponse{}
	_body, _err := client.UploadRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyFileWithOptions(request *VerifyFileRequest, runtime *util.RuntimeOptions) (_result *VerifyFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyFile"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyFile(request *VerifyFileRequest) (_result *VerifyFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyFileResponse{}
	_body, _err := client.VerifyFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifySentenceWithOptions(request *VerifySentenceRequest, runtime *util.RuntimeOptions) (_result *VerifySentenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonStr)) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifySentence"),
		Version:     tea.String("2019-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifySentenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifySentence(request *VerifySentenceRequest) (_result *VerifySentenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifySentenceResponse{}
	_body, _err := client.VerifySentenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
