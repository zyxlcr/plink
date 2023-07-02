package user

import (
	"chatcser/config"
	"chatcser/pkg/model"
	"chatcser/pkg/plink/iface"
	"chatcser/pkg/utils"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func (b BaseUser) Ping(res iface.ResponseWriter, req *iface.Request) {
	fmt.Println("ping")

}

func (b BaseUser) Reg(reqBody BaseUser, req *iface.Request) error {
	fmt.Println("Reg")
	m := BaseUser{Name: reqBody.Name}
	mapper := model.NewMapper(m, nil)
	_, err := mapper.SelectOne()
	if err == nil {
		return errors.Errorf("用户已存在: %s", reqBody.Name)
	}
	user := BaseUser{
		Name:     reqBody.Name,
		Password: utils.BcryptHash(reqBody.Password),
		Email:    reqBody.Email,
	}
	err = mapper.Insert(&user)
	return err

}

func (b BaseUser) Login(reqBody LoginReq) (*LoginRes, error) {
	//var req LoginReq
	//err := c.ShouldBind(&req)

	m := BaseUser{Name: reqBody.Username}
	mapper := model.NewMapper(m, nil)
	user, err := mapper.SelectOne()
	if err != nil {
		println("用户不存在")
		utils.CheckError(err)
		return &LoginRes{}, errors.Errorf("用户不存在: %s", reqBody.Username)
	}
	if ok := utils.BcryptCheck(reqBody.Password, user.Password); !ok {
		println("用户名或密码错误")
		return &LoginRes{}, errors.New("用户名或密码错误")
	}
	println("ok")
	res := &LoginRes{
		Username: reqBody.Username,
		Uid:      cast.ToString(user.ID),
		Token:    b.CreatToken(user),
	}
	return res, nil //b.tokenNext(c, user)
}

func (b BaseUser) ApiLogin(c *gin.Context) (string, error) {
	var req LoginReq
	var u UserInfo
	err := c.ShouldBind(&req)
	utils.CheckError(err)
	m := BaseUser{Name: req.Username}
	mapper := model.NewMapper(m, nil)
	user, err := mapper.SelectOne()
	if err != nil {
		return u, errors.Errorf("用户不存在: %s", req.Username)
	}
	if ok := utils.BcryptCheck(req.Password, user.Password); !ok {
		return u, errors.New("用户名或密码错误")
	}
	j := utils.NewJWT()
	claims := j.CreateClaims(utils.BaseClaims{
		UID:      user.ID,
		Username: user.Name,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		config.GVA_LOG.Error("获取token失败!", zap.Error(err))
		//response.FailWithMessage("获取token失败", c)
		return u, err
	}
	u.UserId = fmt.Sprint(user.ID)
	u.Token = token
	return u, nil
}

func (b BaseUser) CreatToken(user BaseUser) (string, error) {
	j := utils.NewJWT()
	claims := j.CreateClaims(utils.BaseClaims{
		UID:      user.ID,
		Username: user.Name,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		config.GVA_LOG.Error("获取token失败!", zap.Error(err))
		//response.FailWithMessage("获取token失败", c)
		return "", err
	}

	return token, nil
}

func (b BaseUser) Auth(token string) (int64, error) {

	j := utils.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		return 0, err
	}
	zap.Errors(claims.Username, nil)
	return claims.UID, nil
}
