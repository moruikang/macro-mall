// @Author moruikang
// @Date 2025/3/16 17:35:00
// @Desc

package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"macro-mall/internal/admin/dto"
	"macro-mall/internal/admin/store"
	"macro-mall/internal/admin/store/models"
	"macro-mall/internal/pkg/common/bcryptx"
	"math"
)

type UmsAdminService interface {
	GetByUsername(g *gin.Context, username string) (*models.UmsAdmin, error)
	GetAdminInfo(g *gin.Context, username string) (map[string]interface{}, error)
	Register(g *gin.Context, dto *dto.AdminRegisterDTO) error
	Get(g *gin.Context, id int64) (*models.UmsAdmin, error)
	Update(g *gin.Context, dto *dto.AdminUpdateDTO) error
	Delete(g *gin.Context, ids []int64) error
	List(g *gin.Context, dto *dto.PublicPageQuery) (*dto.PageResult, error)
	UpdateStatus(g *gin.Context, userId int64, dto *dto.AdminUpdateStatusDTO) error
	UpdatePassword(g *gin.Context, dto *dto.AdminUpdatePasswordDTO) error
	UpdateRoles(g *gin.Context, dto *dto.AdminRoleRelationDTO) error
	GetUserRoleList(g *gin.Context, id int64) ([]*models.UmsRole, error)
}

type umsAdminService struct {
	store store.Factory
}

var _ UmsAdminService = (*umsAdminService)(nil)

func NewUmsAdminService(store store.Factory) UmsAdminService {
	return &umsAdminService{
		store: store,
	}
}

func (svc *umsAdminService) GetByUsername(g *gin.Context, username string) (*models.UmsAdmin, error) {
	return svc.store.UmsAdmins().GetByUserName(g, username)
}

func (svc *umsAdminService) GetAdminInfo(g *gin.Context, username string) (map[string]interface{}, error) {
	admin, err := svc.store.UmsAdmins().GetByUserName(g, username)
	if err != nil {
		return nil, err
	}

	userRoleList, err := svc.GetUserRoleList(g, admin.Id)
	if err != nil {
		return nil, err
	}
	roleNames := make([]string, 0)
	for _, role := range userRoleList {
		roleNames = append(roleNames, role.Name)
	}
	menus, err := svc.store.UmsRoles().GetMenuList(g, admin.Id)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"username": admin.Username,
		"icon":     admin.Icon,
		"roles":    roleNames,
		"menus":    menus,
	}, nil

}

func (svc *umsAdminService) Register(g *gin.Context, dto *dto.AdminRegisterDTO) error {

	u, err := svc.store.UmsAdmins().GetByUserName(g, dto.Username)
	if err != nil {
		return err
	}
	if u != nil {
		return errors.New("user already exists")
	}

	password, err := bcryptx.EncodingPassword(dto.Password)
	if err != nil {
		return err
	}

	admin := &models.UmsAdmin{
		Username: dto.Username,
		Password: password,
		Icon:     dto.Icon,
		Nickname: dto.Nickname,
		Email:    dto.Email,
		Status:   1,
		Note:     dto.Note,
	}
	return svc.store.UmsAdmins().Create(g, admin)
}

func (svc *umsAdminService) Get(g *gin.Context, id int64) (*models.UmsAdmin, error) {
	return svc.store.UmsAdmins().GetById(g, id)
}

func (svc *umsAdminService) Update(g *gin.Context, dto *dto.AdminUpdateDTO) error {
	admin, err := svc.store.UmsAdmins().GetById(g.Request.Context(), dto.Id)
	if err != nil {
		return err
	}
	admin.Username = dto.Username
	admin.Email = dto.Email
	admin.Icon = dto.Icon
	admin.Nickname = dto.Nickname
	admin.Note = dto.Note
	return svc.store.UmsAdmins().Update(g, admin)
}

func (svc *umsAdminService) Delete(g *gin.Context, ids []int64) error {
	return svc.store.UmsAdmins().DeleteCollection(g, ids)
}

func (svc *umsAdminService) List(g *gin.Context, query *dto.PublicPageQuery) (*dto.PageResult, error) {
	totalCount, admins, err := svc.store.UmsAdmins().List(g, query.Keyword, query.PageSize, query.PageNum)
	if err != nil {
		return nil, err
	}
	totalPage := int(math.Ceil(float64(totalCount) / float64(query.PageSize)))
	return &dto.PageResult{
		List: admins,
		Pagination: dto.Pagination{
			Total:     totalCount,
			PageNum:   query.PageNum,
			PageSize:  query.PageSize,
			TotalPage: totalPage,
		},
	}, nil
}

func (svc *umsAdminService) UpdateStatus(g *gin.Context, userId int64, dto *dto.AdminUpdateStatusDTO) error {
	return svc.store.UmsAdmins().UpdateStatus(g, userId, dto.Status)
}

func (svc *umsAdminService) UpdatePassword(g *gin.Context, dto *dto.AdminUpdatePasswordDTO) error {
	admin, err := svc.store.UmsAdmins().GetByUserName(g, dto.Username)
	if err != nil {
		return err
	}
	if admin == nil {
		return errors.New(fmt.Sprintf("user[%s] don't exists", dto.Username))
	}

	if !bcryptx.VerifyPassword(admin.Password, dto.OldPassword) {
		return errors.New("invalid old password")
	}

	password, err := bcryptx.EncodingPassword(dto.NewPassword)
	if err != nil {
		return err
	}
	admin.Password = password
	return svc.store.UmsAdmins().Update(g, admin)
}

func (svc *umsAdminService) UpdateRoles(g *gin.Context, dto *dto.AdminRoleRelationDTO) error {

	roleRelations := make([]*models.UmsAdminRoleRelation, 0)
	for _, roleId := range dto.RoleIds {
		roleRela := &models.UmsAdminRoleRelation{
			AdminId: roleId,
			RoleId:  roleId,
		}
		roleRelations = append(roleRelations, roleRela)

	}
	return svc.store.UmsAdmins().UpdateRoles(g, dto.AdminId, roleRelations)
}

func (svc *umsAdminService) GetUserRoleList(g *gin.Context, id int64) ([]*models.UmsRole, error) {
	return svc.store.UmsAdmins().GetUserRoles(g, id)
}
