package userHandler

import (
	"app/internal/core/ports"
	"app/internal/dto"
	jwtPkg "app/pkg/jwt"
	"app/pkg/standard"
	"app/pkg/validate"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type userHdl struct {
	userService ports.UserService
}

type UserHdl interface {
	Register(g *gin.Context)
	UpdateUser(g *gin.Context)
	DeleteUser(g *gin.Context)
	GetUser(g *gin.Context)
	Login(g *gin.Context)
	LogOut(g *gin.Context)
	AuthMiddleware(g *gin.Context)
}

func NewAuthHanderler(userService ports.UserService) UserHdl {
	return userHdl{
		userService: userService,
	}
}

func (uh userHdl) Register(g *gin.Context) {
	ctx := g.Request.Context()
	req := dto.CreateUserRequest{}
	err := g.Bind(&req)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = validate.Struct(req)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := uh.userService.CreateUser(ctx, req.UserProperty)
	if err != nil {
		g.JSON(http.StatusOK, dto.BaseOKResponse{
			Code:    standard.ResponseErrorCode,
			Message: standard.ResponseErrorMessage,
		})
		return
	}
	g.JSON(http.StatusOK, dto.BaseOKResponse{
		Code:    standard.ResponseSuccessCode,
		Message: standard.ResponseSuccessMessage,
		Data:    user,
	})
}

func (uh userHdl) UpdateUser(g *gin.Context) {
	ctx := g.Request.Context()
	req := dto.UpdateUserRequest{}
	id := g.GetString("userid")
	err := g.Bind(&req)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = uh.userService.UpdateUserById(ctx, id, req.UserMetadata)
	if err != nil {
		g.JSON(http.StatusOK, dto.BaseOKResponse{
			Code:    standard.ResponseErrorCode,
			Message: standard.ResponseErrorMessage,
		})
		return
	}
	g.JSON(http.StatusOK, dto.BaseOKResponse{
		Code:    standard.ResponseSuccessCode,
		Message: standard.ResponseSuccessMessage,
	})
}

func (uh userHdl) DeleteUser(g *gin.Context) {
	ctx := g.Request.Context()

	id := g.GetString("userid")
	err := uh.userService.DeleteUserById(ctx, id)
	if err != nil {
		g.JSON(http.StatusOK, dto.BaseOKResponse{
			Code:    standard.ResponseErrorCode,
			Message: standard.ResponseErrorMessage,
		})
		return
	}
	g.JSON(http.StatusOK, dto.BaseOKResponse{
		Code:    standard.ResponseSuccessCode,
		Message: standard.ResponseSuccessMessage,
	})
}

func (uh userHdl) GetUser(g *gin.Context) {
	ctx := g.Request.Context()
	id := g.GetString("userid")
	user, err := uh.userService.GetUserById(ctx, id)
	user.Password = ""
	if err != nil {
		g.JSON(http.StatusOK, dto.BaseOKResponse{
			Code:    standard.ResponseErrorCode,
			Message: standard.ResponseErrorMessage,
		})
		return
	}
	g.JSON(http.StatusOK, dto.BaseOKResponse{
		Code:    standard.ResponseSuccessCode,
		Message: standard.ResponseSuccessMessage,
		Data:    user,
	})
}

func (uh userHdl) Login(g *gin.Context) {
	ctx := g.Request.Context()
	req := dto.LoginRequest{}
	err := g.Bind(&req)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}
	token, err := uh.userService.Login(ctx, req.UserNamePassword)
	if err != nil {
		g.JSON(http.StatusOK, dto.BaseOKResponse{
			Code:    standard.ResponseErrorCode,
			Message: err.Error(),
		})
		return
	}
	g.JSON(http.StatusOK, dto.BaseOKResponse{
		Code:    standard.ResponseSuccessCode,
		Message: standard.ResponseSuccessMessage,
		Data:    dto.LoginResponse{ Token: token },
	})
}

func (uh userHdl) LogOut(g *gin.Context) {
	ctx := g.Request.Context()
	id := g.GetString("userid")
	err := uh.userService.LogOut(ctx, id)
	if err != nil {
		g.JSON(http.StatusOK, dto.BaseOKResponse{
			Code:    standard.ResponseErrorCode,
			Message: standard.ResponseErrorMessage,
		})
		return
	}
	g.JSON(http.StatusOK, dto.BaseOKResponse{
		Code:    standard.ResponseSuccessCode,
		Message: standard.ResponseSuccessMessage,
		Data:    nil,
	})
}

func (uh userHdl) AuthMiddleware(g *gin.Context) {
	berarerToken := g.GetHeader("Authorization")
	berarerTokens := strings.Split(berarerToken, "Bearer ")
	if len(berarerTokens) != 2 {
		g.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	token := strings.ReplaceAll(berarerTokens[1], " ", "")

	//validate token
	claim, ok := jwtPkg.ValidAndGetClaims(token)
	if !ok {
		g.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	
	g.Set("userid",claim.Id)
	g.Set("token",token)
	g.Next()
}
