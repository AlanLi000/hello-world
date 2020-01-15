package com.chaincode.signsystem.config;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import springfox.documentation.builders.ApiInfoBuilder;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.service.ApiInfo;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

//http://localhost:8080/swagger-ui.html D:\idea_workspace\sign
@Configuration  //配置文件，让Spring来加载改类配置，再通过@EnableSwagger2注解来启用Swagger2.用@Configuration注解该类，等价于XML中配置beans；用@Bean标注方法等价于XML中配置bean。
@EnableSwagger2
public class SwaggerConfig{
    /*通过creatRestApi()函数创建Docker的Bean之后，apiInfo()用来创建该Api的基本信息（这些基本信息会展现文档页面中）
    */
    @Bean
    public Docket creatRestApi() {
        return new Docket(DocumentationType.SWAGGER_2)
                .apiInfo(apiInfo())
                .select()
                //指定接口的位置
                .apis(RequestHandlerSelectors.basePackage("com.chaincode.signsystem.controller"))
                .paths(PathSelectors.any())
                .build();
    }

    private ApiInfo apiInfo() {
        return new ApiInfoBuilder()
                .title("后台接口文档")
                .description("手签合同的测试方法")
                .termsOfServiceUrl("")
                .contact("xnsydx_lcl")
                .version("1.0")
                .build();
    }

}
