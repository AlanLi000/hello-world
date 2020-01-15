package com.chaincode.signsystem.entity;

import lombok.Data;

@Data
public class User {
    private int    id;
    private String username;
    private String password;
    private String newPassword;     //辅助字段
    private String truename;
    private String idcard;
    private String sex;
    private String tel;
    private Contract[] contracts;

    public User(){

    }


    public String getUsername() {
        return username;
    }

    public String getPassword() {
        return password;
    }

    public String getNewPassword() {
        return newPassword;
    }

    public String getTruename() {
        return truename;
    }

    public String getIdcard() {
        return idcard;
    }

    public String getSex() {
        return sex;
    }

    public String getTel() {
        return tel;
    }

    public Contract[] getContracts() {
        return contracts;
    }

    public void setUsername(String Username) {
        this.username = Username;
    }

    public void setPassword(String Password) {
        this.password = Password;
    }

    public void setNewPassword(String newPassword) {
        this.newPassword = newPassword;
    }

    public void setTruename(String Truename) {
        this.truename = Truename;
    }

    public void setIdcard(String Idcard) {
        this.idcard = Idcard;
    }

    public void setSex(String Sex) {
        this.sex = Sex;
    }

    public void setTel(String Tel) {
        this.tel = Tel;
    }

    public void setContracts(Contract[] contracts) {
        this.contracts = contracts;
    }
}