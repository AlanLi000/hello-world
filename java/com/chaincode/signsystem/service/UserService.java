package com.chaincode.signsystem.service;

import com.chaincode.signsystem.entity.User;

public interface UserService {
    //查询
    User query(long id);

    //注册
    User register(User user);

    //更改密码
    int changePwd(User user);

    //删除用户
    int delete(long id);
}
