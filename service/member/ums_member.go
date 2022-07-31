package member

import (
	"github.com/tongxz/xs-admin-vue/global"
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/member"
	memberReq "github.com/tongxz/xs-admin-vue/model/member/request"
)

type MemberService struct {
}

// CreateMember 创建Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) CreateMember(member member.Member) (err error) {
	err = global.GVA_DB.Create(&member).Error
	return err
}

// DeleteMember 删除Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) DeleteMember(member member.Member) (err error) {
	err = global.GVA_DB.Delete(&member).Error
	return err
}

// DeleteMemberByIds 批量删除Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) DeleteMemberByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]member.Member{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMember 更新Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) UpdateMember(member member.Member) (err error) {
	err = global.GVA_DB.Save(&member).Error
	return err
}

// GetMember 根据id获取Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) GetMember(id uint) (member member.Member, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&member).Error
	return
}

// GetMemberInfoList 分页获取Member记录
// Author [piexlmax](https://github.com/piexlmax)
func (memberService *MemberService) GetMemberInfoList(info memberReq.MemberSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&member.Member{})
	var members []member.Member
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Mobile != nil {
		db = db.Where("mobile = ?", info.Mobile)
	}
	if info.Gender != nil {
		db = db.Where("gender = ?", info.Gender)
	}
	if info.MemberCard != nil {
		db = db.Where("member_card = ?", info.MemberCard)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&members).Error
	return members, total, err
}
