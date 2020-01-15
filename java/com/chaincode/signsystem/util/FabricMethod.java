package com.chaincode.signsystem.util;

import com.chaincode.signsystem.fabric.ChaincodeManager;
import com.chaincode.signsystem.entity.Contract;
import com.chaincode.signsystem.entity.User;
import com.chaincode.signsystem.test.FabricManagerTest;
import com.alibaba.fastjson.JSONObject;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.CrossOrigin;


@Slf4j
@CrossOrigin
public class FabricMethod {
    private static ChaincodeManager testManager;
    private static ChaincodeManager signManager;

    static {
        try {
            testManager = FabricManagerTest.obtain().getManager(1);
            signManager = FabricManagerTest.obtain().getManager(2);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    //功能
    //1.注册,参数args:ID| Password | Name | Idcard | Sex | Tel
    public static String register(User user) {
        String[] args = {user.getUsername(), user.getPassword(), user.getTruename(), user.getIdcard(), user.getSex(), user.getTel()};
        JSONObject jsonObject;
        try {
            jsonObject = testManager.invoke("regist", args);
            System.out.println(jsonObject);
        } catch (Exception e) {
            return "register error!!";
        }
        String result = jsonObject.get("data").toString();

        return result;
    }

    //2.登录 args:Tel | Password
    public static String login(User user) {

        String[] args = {user.getTel(), user.getPassword()};
        JSONObject jsonObject;
        try {
            jsonObject = testManager.invoke("login", args);
            System.out.println(jsonObject);
        } catch (Exception e) {
            return "login error!!";
        }
        String result = jsonObject.get("data").toString();
        return result;
    }

    //3.更改用户密码 args:Tel| OldPassword |newPassword
    public static String changePwd(User user) {
        String[] args = {user.getTel(), user.getPassword(),user.getNewPassword()};
        JSONObject jsonObject;
        try {
            jsonObject = testManager.invoke("changePwd", args);
            System.out.println(jsonObject);
        } catch (Exception e) {
            return "change error!!";
        }
        String result = jsonObject.get("data").toString();
        return result;
    }

    //4.查询账号 args:Tel
    public static String userquery(User user) {
        String[] args = {user.getTel()};
        JSONObject jsonObject;
        try {
            jsonObject = testManager.invoke("userquery", args);
            System.out.println(jsonObject);
        } catch (Exception e) {
            return "userquery error!!";
        }
        String result = jsonObject.get("data").toString();
        return result;
    }

    //5.删除账号 args：Tel
    public static String delete(User user) {
        String[] args = {user.getTel()};
        JSONObject jsonObject;
        try {
            jsonObject = testManager.invoke("delete", args);
            System.out.println(jsonObject);
        } catch (Exception e) {
            return "delete error!!";
        }
        String result = jsonObject.get("data").toString();
        return result;
    }

    //6.创建合同并上链 5个参数args:CreatorTel | ConID | Conname | Contype | Context
    public static String creation(Contract contract) {
        String[] args = {contract.getCreatorTel(),contract.getConID(),contract.getConname(),contract.getContype(),contract.getContext()};
        JSONObject jsonObject;
        try {
            jsonObject = testManager.invoke("creation", args);
            System.out.println(jsonObject);
        } catch (Exception e) {
            return "creation error!!";
        }
        String result = jsonObject.get("data").toString();
        return result;
    }

    //7.添加合同信息 args: ConID
    public static String add(Contract contract) {
        String[] args = {contract.getConID()};
        JSONObject jsonObject;
        try {
            jsonObject = testManager.invoke("add", args);
            System.out.println(jsonObject);
        } catch (Exception e) {
            return "add error!!";
        }
        String result = jsonObject.get("data").toString();

        return result;
    }

    //8.合同签名 args:ConID | Tel
    public static String sign(User user,Contract contract) {
        String[] args = {contract.getConID(),user.getTel()};
        JSONObject jsonObject;
        try {
            jsonObject = testManager.invoke("sign", args);
            System.out.println(jsonObject);
        } catch (Exception e) {
            return "sign error!!";
        }
        String result = jsonObject.get("data").toString();

        return result;
    }

    //9.查询合同信息 args: ConID
    public static String contractquery(Contract contract) {
        String[] args = {contract.getConID()};
        JSONObject jsonObject;
        try {
            jsonObject = testManager.invoke("contractquery", args);
            System.out.println(jsonObject);
        } catch (Exception e) {
            return "contractquery error!!";
        }
        String result = jsonObject.get("data").toString();

        return result;
    }


    //10.通过key查看历史记录 args: Tel
    public static String getHistoryForKey(User user) {
        String[] args = {user.getTel()};
        JSONObject jsonObject;
        try {
            jsonObject = testManager.invoke("getHistoryForKey", args);
            System.out.println(jsonObject);
        } catch (Exception e) {
            return "getHistoryForKey error!!";
        }
        String result = jsonObject.get("data").toString();

        return result;
    }

}
