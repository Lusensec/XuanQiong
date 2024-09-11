package controllers

import (
    "strconv"

    "xuanqiong/config"
    "xuanqiong/types"
    "xuanqiong/models"
    "github.com/gin-gonic/gin"
)

var (
    cfg = config.Config
    maxAttempts = cfg.Login.MaxAttempts
)

// 登录
func Login(c *gin.Context) {
    if models.IsLocked(c.ClientIP()) {
        c.JSON(429, gin.H{"error": "Too many login attempts. Please try again later."})
        return
    }
    var logindata types.LoginData
    if err := c.ShouldBindJSON(&logindata); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }
    loginUser := models.CheckLogin(logindata.Username, logindata.Password)
    if loginUser != nil {
        c.JSON(200, gin.H{"msg":"Login Successful", "Username":loginUser.Username, "token": loginUser.Token})
    } else {
        maxAttempts--
        if maxAttempts == 0 {
            models.LockIP(c.ClientIP(), cfg.Login.LockoutDuration)
            c.JSON(429, gin.H{"error": "Too many login attempts. Please try again later."})
            maxAttempts = cfg.Login.MaxAttempts
            return
        }
        c.JSON(401, gin.H{"error": "Invalid username or password, you can try " + strconv.FormatInt(maxAttempts, 10) + " times."})
        return
    }
}

// 退出
func Logout(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        err := models.CleanToken(currentUser.Username)
        if err == nil {
            c.JSON(200, gin.H{"已退出": "显示未登录页面"})
            return
        } else {
            c.JSON(200, gin.H{"error": err})
        }
    }
    c.JSON(200, gin.H{"未登录": "显示未登录页面"})
}

// 创建用户
func CreateUser(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"error": "Permission denied"})
            return
        }
        var logindata types.LoginData
        if err := c.ShouldBindJSON(&logindata); err != nil {
            c.JSON(400, gin.H{"error": "Invalid input"})
            return
        }
        err := models.CreateUser(logindata.Username, logindata.Password, 0)
        if err == nil {
            c.JSON(200, gin.H{"Successful": "创建用户成功"})
            return
        } else {
            c.JSON(200, gin.H{"error": err.Error()})
            return
        }
    }
    c.JSON(200, gin.H{"未登录": "显示未登录页面"})
}

// 删除用户
func DeleteUser(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"error": "Permission denied"})
            return
        }
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(200, gin.H{"error": "Invalid input"})
            return
        }
        username, ok := data["username"].(string)
        if !ok {
            c.JSON(200, gin.H{"error": "Invalid input"})
            return
        }
        err := models.DeleteUser(username)
        if err != nil {
            c.JSON(200, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": "删除用户" + username + "成功"})
        return
    }
    c.JSON(200, gin.H{"未登录": "显示未登录页面"})
}

// 启用或禁用用户
func SetUserStatus(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"error": "Permission denied"})
            return
        }
        // 设置状态
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(200, gin.H{"error": "Invalid input"})
            return
        }
        username, ok := data["username"].(string)
        if !ok {
            c.JSON(200, gin.H{"error": "Invalid input"})
            return
        }
        err := models.SetUserStatus(username)
        if err != nil {
            c.JSON(200, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": "修改用户" + username + "状态成功"})
        return
    }
    c.JSON(200, gin.H{"未登录": "显示未登录页面"})
}

// 获取所有用户信息
func GetUsers(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"error": "Permission denied"})
            return
        }
        users := models.GetUsers()
        c.JSON(200, gin.H{"message": users})
        return
    }
    c.JSON(200, gin.H{"message": "未登录"})
}

// 修改用户信息
func UpdateUser(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    currentUser := models.GetUserByToken(token)
    if currentUser != nil {
        if currentUser.Role != 1 {
            c.JSON(200, gin.H{"error": "Permission denied"})
            return
        }
        var data map[string]interface{}
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(200, gin.H{"error": "Invalid input"})
            return
        }
        username, ok := data["username"].(string)
        if !ok {
            c.JSON(200, gin.H{"error": "Invalid input"})
            return
        }
        password, _ := data["password"].(string)
        role, ok := data["role"].(int64)
        if !ok {
            role = -1
        }
        status, ok := data["status"].(int64)
        if !ok {
            status = -1
        }
        err := models.UpdateUser(username, password, role, status)
        if err != nil {
            c.JSON(200, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": "修改成功"})
        return
    }
    c.JSON(200, gin.H{"message": "未登录"})
}
