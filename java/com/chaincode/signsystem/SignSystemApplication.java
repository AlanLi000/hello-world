package com.chaincode.signsystem;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Configuration;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@SpringBootApplication
@EnableSwagger2
@Configuration
//指定要扫描的mybatis映射类的路径,放在应用类的前面
@MapperScan("com.chaincode.signsystem.mapper")
public class SignSystemApplication {

    public static void main(String[] args) {
        SpringApplication.run(SignSystemApplication.class, args);
    }
}
