package com.chaincode.signsystem.controller;

import com.chaincode.signsystem.entity.User;
import com.chaincode.signsystem.service.UserService;
import io.swagger.annotations.Api;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@Api(tags = "后端调用")
@RestController
@RequestMapping("/user")
public class UserController {

    // 注入mapper类
    @Autowired
    private UserService userService;

    //注册
    @RequestMapping(value="/register", method= RequestMethod.POST)
    public User register(User user){
        return userService.register(user);
    }

    //查询用户信息
    @RequestMapping(value="/query", method= RequestMethod.GET, produces="application/json")
    public User getUser(long id) {

        User user = this.userService.query(id);

        return user;
    }

    //更改密码
    @RequestMapping(value="/changePwd", method= RequestMethod.POST)
    public String changePwd(User user){
        int i= userService.changePwd(user);
        if(i==1)
        return "修改成功";
        else
            return "修改失败";
    }

    //删除用户
    @RequestMapping(value="/delete", method= RequestMethod.POST)
    public String  delete (long id){
        int i = userService.delete(id);
        if(i==1)
            return "修改成功";
        else
            return "修改失败";
    }
}
