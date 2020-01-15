package com.chaincode.signsystem.controller;

import com.chaincode.signsystem.entity.Contract;
import com.chaincode.signsystem.entity.User;
import com.chaincode.signsystem.util.FabricMethod;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import org.springframework.web.bind.annotation.*;

@CrossOrigin
@Api(tags = "链码调用")
@RestController
@RequestMapping("/test")
public class TestController {
    //1.注册
    @ApiOperation(value = "用户注册" ,notes = "这是一个说明")
    @PostMapping("/register")
    public Object register(User user){

        return FabricMethod.register(user);
    }

    //2.登录
    @ApiOperation("登录")
    @PostMapping("/login")
    public Object login(User user){

        return FabricMethod.login(user);
    }

    //3.查询用户
    @ApiOperation("查询")
    @PostMapping("/userquery")
    public Object userquery(User user){

        return FabricMethod.userquery(user);
    }

    //4.更改密码
    @ApiOperation("更改")
    @PostMapping("/changePwd")
    public Object changePwd(User user){

        return FabricMethod.changePwd(user);
    }

    //5.创建合同
    @ApiOperation("创建")
    @PostMapping("/creation")
    public Object creation(Contract contract){

        return FabricMethod.creation(contract);
    }

    //6.合同查询
    @ApiOperation("合同查询")
    @PostMapping("/contractquery")
    public Object contractquery(Contract contract){

        return FabricMethod.contractquery(contract);
    }


    //7.添加
    @ApiOperation("添加")
    @PostMapping("/add")
    public Object add(Contract contract){

        return FabricMethod.add(contract);
    }

    //8.签名
    @ApiOperation("签名")
    @PostMapping("/sign")
    public Object sign(User user,Contract contract){

        return FabricMethod.sign(user,contract);
    }

    //9.查询历史信息
    @ApiOperation("历史记录")
    @PostMapping("/getHistoryForKey")
    public Object getHistoryForKey(User user){

        return FabricMethod.getHistoryForKey(user);
    }

    //10.删除用户
    @ApiOperation("删除")
    @PostMapping("/delete")
    public Object delete(User user){

        return FabricMethod.delete(user);
    }

}


