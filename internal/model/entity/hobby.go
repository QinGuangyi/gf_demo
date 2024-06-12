// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Hobby is the golang structure for table hobby.
type Hobby struct {
	Id    uint   `json:"id"    orm:"id"     ` // ID
	EmpId uint   `json:"empId" orm:"emp_id" ` // EmpID
	Hobby string `json:"hobby" orm:"hobby"  ` // 爱好

	
}
