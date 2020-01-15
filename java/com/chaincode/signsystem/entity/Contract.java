package com.chaincode.signsystem.entity;

import lombok.Data;

@Data
public class Contract {

    private String creatorTel;
    private String conID;
    private String conname;
    private String contype;
    private String context;
    private String timeStamp;
    private String[] signatorys;

    public String getCreatorTel() {
        return creatorTel;
    }

    public String getConID() {
        return conID;
    }

    public String getConname() {
        return conname;
    }

    public String getContype() {
        return contype;
    }

    public String getContext() {
        return context;
    }

    public String getTimeStamp() {
        return timeStamp;
    }

    public String[] getSignatorys() {
        return signatorys;
    }

    public void setCreatorTel(String creatorTel) {
        this.creatorTel = creatorTel;
    }

    public void setConID(String conID) {
        this.conID = conID;
    }

    public void setConname(String conname) {
        this.conname = conname;
    }

    public void setContype(String contype) {
        this.contype = contype;
    }

    public void setContext(String context) {
        this.context = context;
    }

    public void setTimeStamp(String timeStamp) {
        this.timeStamp = timeStamp;
    }

    public void setSignatorys(String[] signatorys) {
        this.signatorys = signatorys;
    }


}
