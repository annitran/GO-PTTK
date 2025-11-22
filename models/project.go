package models

import "time"

type ProjectStatus string

const (
	StatusDraft            ProjectStatus = "draft"             // hồ sơ vừa tạo
	StatusUnderReview      ProjectStatus = "under_review"      // hoàn thiện, hội đồng đang xét
	StatusRevisionRequired ProjectStatus = "revision_required" // yêu cầu chỉnh sửa
	StatusApproved         ProjectStatus = "approved"          // được phê duyệt
	StatusInProgress       ProjectStatus = "in_progress"       // đang thực hiện
	StatusCompleted        ProjectStatus = "completed"         // hoàn thành
	StatusRejected         ProjectStatus = "rejected"          // bị từ chối
)

type Project struct {
	ID            uint64        `json:"id" gorm:"primaryKey;autoIncrement"`
	Title         string        `json:"title" gorm:"not null"`
	ProposerName  string        `json:"proposer_name" gorm:"not null"`
	Email         string        `json:"email" gorm:"not null"`
	Field         string        `json:"field" gorm:"not null"`
	ExpectedStart *time.Time    `json:"expected_start"`
	ExpectedEnd   *time.Time    `json:"expected_end"`
	Status        ProjectStatus `json:"status" gorm:"index"`
	CreatedAt     time.Time     `json:"created_at"`

	Members     []ProjectMember     `json:"members" gorm:"constraint:OnDelete:CASCADE"`
	Attachments []ProjectAttachment `json:"attachments" gorm:"constraint:OnDelete:CASCADE"`
	Reviews     []ProjectReview     `json:"reviews" gorm:"constraint:OnDelete:CASCADE"`
}

type ProjectMember struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	ProjectID uint64    `json:"project_id" gorm:"index"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
type ProjectAttachment struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	ProjectID uint64    `json:"project_id" gorm:"index"`
	FileName  string    `json:"file_name"`
	FileURL   string    `json:"file_url"`
	CreatedAt time.Time `json:"created_at"`
}
type ProjectReview struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	ProjectID uint64    `json:"project_id" gorm:"index"`
	Reviewer  string    `json:"reviewer"`
	Role      string    `json:"role"`
	Comment   string    `json:"comment"`
	Decision  string    `json:"decision"` // "approve", "revise", "reject"
	CreatedAt time.Time `json:"created_at"`
}
