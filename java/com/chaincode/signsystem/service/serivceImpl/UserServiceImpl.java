package com.chaincode.signsystem.service.serivceImpl;

import com.chaincode.signsystem.mapper.UserMapper;
import com.chaincode.signsystem.entity.User;
import com.chaincode.signsystem.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.RequestParam;

@Service
public class UserServiceImpl implements UserService {

    // 注入mapper类
    @Autowired
    private UserMapper ump;

    //按id查询
    public User query(long id){
        User user = ump.query(id);
        if (null != user) {
            return user;
        }
        return null;
    }

    //注册
    public User register(User user){
        ump.register(user);
        return user;
    }


    //更改密码
    public int changePwd(User user)
    {
        ump.changePwd(user);
        return 1;

    }

    //删除用户
    public int delete(long id)
    {
        ump.delete(id);
        return 1;
    }



}
