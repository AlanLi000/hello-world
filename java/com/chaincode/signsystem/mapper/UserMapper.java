package com.chaincode.signsystem.mapper;

import com.chaincode.signsystem.entity.User;
import org.mapstruct.Mapper;
import org.springframework.stereotype.Repository;

@Mapper
@Repository
public interface UserMapper {
    //查询
    User query(long id);

    //注册
    void register(User user);

    //更改密码
    void changePwd(User user);

    //删除用户
    void delete(long id);
}

