package request

type ExtendBayarApprovalRequest struct {
	KdUserLf    string               `json:"kd_user_lf"`
	Datas       []ExtendBayarRequest `json:"dtas"`
	StsApproval string               `json:"sts_approval"`
}
