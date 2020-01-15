/*
problems
	1.换回值信息修改
	2.Printf 和 Sprintf
	3.string 和 []byte
*/

//包名
//一个chaincode通常是一个golang源码文件，这个包名必须是main
package main

//导入包
//chaincode需要引入一些Fabric提供的系统包，这些系统包提供了chaincode和Fabirc进行通信的接口。
import (
	"bytes"
	"encoding/json"
	"fmt"

	//"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("sign")

//定义chaincode主结构体
//每个chaincode都需要定义个结构体，结构体的名字可以是任意符合golang命名规范的字符串。chaincode的住结构体必须实现Chaincode接口
//   type Chaincode interface {
//	        Init(stub ChaincodeStubInterface) pb.Response
//	        Invoke(stub ChaincodeStubInterface) pb.Response
//   }
type SimpleChaincode struct{}

//返回结果信息
type ResultInfo struct {
	Code  interface{} `json:"code"`
	State interface{} `json:"state"`
	Msg   interface{} `json:"msg"`
}

type User struct {
	Truename  string     `json:"name"`      //姓名
	Sex       string     `json:"sex"`       //性别
	Tel       string     `json:"tel"`       //电话
	Idcard    string     `json:"idcard"`    //身份证号
	ID        string     `json:"id"`        //用户名*****java中的username
	Password  string     `json:"password"`  //密码
	Contracts []Contract `json:"contracts"` //合同
}
type Contract struct {
	CreatorTel string   `json:"creatorid"` //合同创建人
	ConID     string   `json:"conid"`     //合同id
	Conname   string   `json:"conname"`   //合同名称
	Contype   string   `json:"contype"`   //合同类型
	Context   string   `json:"context"`   //合同文本
	TimeStamp string   `json:"timeStamp"` //时间戳
	Signatory []string `json:"signatory"` //签署人
}
/*
状态码：
		01.参数个数不对
        02.从链上获取信息错误 stub.GetState()
	    03.发布到链上错误 stub.PutState()
		04.编码失败 json.Marshal()
		05.json.Unmarshal失败
		06.stub.DelState()失败
		07.stub.GetHistoryForKey错误
		08.getHistoryListResult错误
		---------------
		0.成功
		1.用户已存在
		2.合同已存在
		3.用户不存在
		4.密码错误
		5.合同不存在
*/
func ErrorResult(msg string,n string) []byte {
	var result ResultInfo
	result.Code = n
	result.State = "ERROR"
	result.Msg = msg
	logger.Info(msg)

	rbytes, _ := json.Marshal(result) //rbytes是[]byte类型

	return rbytes
}

func SuccessResult(msg string) []byte {
	var result ResultInfo
	result.Code = 0
	result.State = "SUCCESS"
	result.Msg = msg
	logger.Info(msg)

	rbytes, _ := json.Marshal(result) //rbytes是[]byte类型

	return rbytes
}

//^==============================================================C:\Users\MogicBook\.GoLand2019.2\config\scratches\signtest.go

//安装Chaincode(链码初始化)
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("############# sign Init ##############")
	return shim.Success(nil)
}

//Invoke interface链码交互的具体方法
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	logger.Info("################# Register Invoke ###############")
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "regist": //注册
		return t.regist(stub, args)
	case "creation": //创建合同
		return t.creation(stub, args)
	case "login": //登录
		return t.login(stub, args)
	case "changePwd": //修改密码
		return t.changePwd(stub, args)
	case "userquery": //查询用户信息
		return t.userquery(stub, args)
	case "delete": //删除用户
		return t.delete(stub, args)
	case "contractquery": //查询合同信息
		return t.contractquery(stub, args)
	//case "add": //添加合同
	//	return t.add(stub, args)
	case "sign": //签名
		return t.sign(stub, args)
	case "getHistoryForKey": //查询user历史记录
		return t.getHistoryForKey(stub, args)
	default:
		return shim.Error(fmt.Sprintf("unsupported function :%s ", function))
	}

	//return shim.Success(nil)
}

//=====================================================
//注册 args:Name| Sex | Tel | Idcard | ID | Password
//=====================================================
func (t *SimpleChaincode) regist(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Info("regist args: ", args)
	if len(args) != 6 {
		return shim.Success(ErrorResult("function need six args","01"))
	}
	//定义接受参数的变量
	var Name, Sex string
	var Password, Tel string
	var Idcard, ID string
	var err error

	ID = args[0]
	Password = args[1]
	Name = args[2]
	Idcard = args[3]
	Sex = args[4]
	Tel = args[5]
	//定义结构体
	var user User
	user.Truename = Name
	user.ID = ID
	user.Sex = Sex
	user.Tel = Tel
	user.Idcard = Idcard
	user.Password = Password

	uBytes, _ := json.Marshal(user) //将数据编码成Json字符串
	//get the state from ledger
	ubytes, err := stub.GetState(Tel)
	fmt.Println(string(ubytes))
	if err != nil {
		return shim.Success(ErrorResult("注册失败Failed to get state","02"))
	}
	if ubytes != nil {
		return shim.Success(ErrorResult("用户已存在this user already exist","1"))
	}
	//Write the state back to the ledger(写入账本)
	err = stub.PutState(Tel, uBytes)
	if err != nil {
		return shim.Success(ErrorResult("注册失败" + err.Error(),"03"))
	}
	//return shim.Success([]byte(ID + "帐号创建成功! "))
	return shim.Success(SuccessResult("注册成功Registered successfully"))
}

//=======================================
//验证账号密码是否匹配,登录 args:Tel|Password
//=======================================
func (t *SimpleChaincode) login(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Info("login args:", args)
	if len(args) != 2 {
		return shim.Success(ErrorResult("function need two args","01"))
	}
	Tel := args[0]
	password := args[1]
	//query the ledger
	Bytes, err := stub.GetState(Tel)
	if err != nil {
		return shim.Success(ErrorResult("登陆失败Failed to get account" + err.Error(),"02"))
	}
	if Bytes == nil {
		return shim.Success(ErrorResult("用户不存在This id does not exists","3"))
	}

	var user User
	//反序列化json,将bytes数据传到user中
	err = json.Unmarshal(Bytes, &user)

	if err != nil {
		return shim.Success(ErrorResult("登陆失败Failed to unmarshal user" + err.Error(),"05"))
	}
	if user.Password == password {
		return shim.Success(SuccessResult("登陆成功Log in successfully"))
	} else {
		return shim.Success(ErrorResult("密码错误Wrong password!","4"))
	}
}

//==============================
//更改用户密码 args:Tel| OldPassword |newPassword
//==============================
func (t *SimpleChaincode) changePwd(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Info("changing password args :", args)
	if len(args) != 3 {
		return shim.Success(ErrorResult("function need three args","01"))
	}
	Tel := args[0]
	oldPassword := args[1]
	newPassword := args[2]
	var err error
	Bytes, err := stub.GetState(Tel)
	if err != nil {
		return shim.Success(ErrorResult("修改失败Failed get account" + err.Error(),"02"))
	}
	if Bytes == nil {
		return shim.Success(ErrorResult("用户不存在user does not exist","3"))
	}
	var user User
	err = json.Unmarshal(Bytes, &user)
	if err != nil {
		return shim.Success(ErrorResult("修改失败Failed to unmarshal user" + err.Error(),"05"))
	}
	if user.Password == oldPassword {
		user.Password = newPassword
	} else {
		return shim.Success(ErrorResult("密码错误password wrong!","4"))
	}
	ubytes, _ := json.Marshal(user)
	err = stub.PutState(user.Tel, ubytes)
	if err != nil {
		return shim.Success(ErrorResult("修改失败put to state error!" + err.Error(),"03"))
	}
	return shim.Success(SuccessResult("修改成功modify successfully"))
}

//================================
//查询账号 args:Tel
//================================
func (t *SimpleChaincode) userquery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Success(ErrorResult("function need one args","01"))
	}
	var Tel string
	var err error
	Tel = args[0]
	Avalbytes, err := stub.GetState(Tel)
	if err != nil {
		return shim.Success(ErrorResult("Failed to get state","02"))
	}
	if Avalbytes == nil {
		return shim.Success(ErrorResult("用户不存在no account", "3"))
	}
	jsonResp := string(Avalbytes)
	fmt.Printf("Query Response:%s \n", jsonResp) //???????????????????
	//return shim.Success(Avalbytes)
	return shim.Success(SuccessResult("user query result:" + string(Avalbytes))) //????????????????
}

//====================
//删除账号 args：Tel
//====================
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Info("delete user args:", args)
	if len(args) != 1 {
		return shim.Success(ErrorResult("function need one args","01"))
	}
	Tel := args[0]
	ubytes, err := stub.GetState(Tel)
	if err != nil {
		return shim.Success(ErrorResult("删除失败Failed to get user from state" + err.Error(),"02"))
	}

	if ubytes == nil {
		return shim.Success(ErrorResult("用户不存在User does not exist","3"))
	}
	var user User
	user = User{}
	Bytes, err := json.Marshal(user)
	if err != nil {
		return shim.Success(ErrorResult("删除失败Failed to marshal user" + err.Error(),"04"))
	}
	err = stub.PutState(Tel, Bytes)
	if err != nil {
		return shim.Success(ErrorResult("删除失败Failed to put user to state" + err.Error(),"03"))
	}
	err = stub.DelState(Tel)
	if err != nil {
		return shim.Success(ErrorResult("删除失败Failed to delete user" + err.Error(),"06"))
	}
	return shim.Success(SuccessResult("账号已被删除The account has been deleted"))
}

//======================
//查询合同信息 args: ConID
//======================
func (t *SimpleChaincode) contractquery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Success(ErrorResult("function need one args","01"))
	}
	ConID := args[0]
	//stub.GetState方法，查询数据，获取指定资产key的value，访问账本的状态
	cBytes, err := stub.GetState(ConID)
	if err != nil {
		return shim.Success(ErrorResult("查询失败Failed to get state for ConID","01"))
	}

	if cBytes == nil {

		return shim.Success(ErrorResult("合同不存在Does not exist ConID","5"))
	}
	jsonResp := string(cBytes)
	fmt.Printf("Query response:%s\n", jsonResp) //Spintf 和 Printf的区别？？？？？？？？
	//return shim.Success(dBytes)
	return shim.Success(SuccessResult("contract query result:" + string(cBytes)))
}

//==============================================================================================
//创建合同并上链 5个参数args:Tel | ConID | Conname | Contype | Context
//==============================================================================================
func (t *SimpleChaincode) creation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Success(ErrorResult("function need five args","01"))
	}
	var CreatorTel string
	var ConID, Conname string
	var Contype string
	var Context string

	var err error
	CreatorTel = args[0]
	ConID = args[1]
	Conname = args[2]
	Contype = args[3]
	Context = args[4]
	//合同创建
	var contract Contract
	contract.CreatorTel = CreatorTel
	contract.ConID = ConID
	contract.Conname = Conname
	contract.Contype = Contype
	contract.Context = Context
	//获取时间戳
	currentTime := time.Unix(time.Now().UnixNano()/1e9, 0)
	contract.TimeStamp = currentTime.String()
	//验证合同是否存在
	cBytes, err := stub.GetState(ConID)
	if err != nil {
		return shim.Success(ErrorResult("合同创建失败Failed to get state","02"))
	}
	if cBytes != nil {
		return shim.Success(ErrorResult("合同已存在this contract already exist","2"))
	}
	//获取用户信息
	uBytes, err := stub.GetState(CreatorTel)
	var user User
	err = json.Unmarshal(uBytes, &user)
	if err != nil {
		return shim.Success(ErrorResult("Failed to unmarshal user and: " + err.Error(),"05"))
	}
	//上链，调用此方法
	i := add(stub,contract,user)
	if i != "0"{
		return shim.Success(ErrorResult("上链失败Failed to linked the contract","i="+i))
	}
	return shim.Success(SuccessResult("合同上链成功The contract has been linked"))
}

//======================
//添加合同信息 args: ConID | Tel
//======================
func add(stub shim.ChaincodeStubInterface, arg1 Contract,arg2 User) string {
	contract := arg1
	user := arg2

	//将用户Tel装入Signatory中（切片）
	contract.Signatory = append(contract.Signatory,user.Tel)
	cJsonbytes, err := json.Marshal(contract)
	if err != nil {
		return "04"
	}
	//将合同信息添加到用户信息中（切片）
	user.Contracts = append(user.Contracts, contract)
	uJsonbytes, err := json.Marshal(user)
	if err != nil {
		return "04"
	}
	//上传已更新的用户信息
	err = stub.PutState(user.Tel, uJsonbytes)
	if err != nil {
		return "03"
	}
	//上传已更新的合同信息
	err = stub.PutState(contract.ConID, cJsonbytes)
	if err != nil {
		return "03"
	}
	return "0"
}

//===========================
//合同签名 args:ConID | Tel
//===========================
func (t *SimpleChaincode) sign(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Success(ErrorResult("function need two args","01"))
	}
	ConID := args[0]
	Tel := args[1]
	//验证合同是否存在
	cBytes, err := stub.GetState(ConID)
	if err != nil {
		return shim.Success(ErrorResult("签名失败Failed to get contract from state" + err.Error(),"02"))
	}
	if cBytes == nil {
		return shim.Success(ErrorResult("合同不存在does not exist contract","5"))
	}
	//合同存在，接受合同信息
	var contract Contract
	err = json.Unmarshal(cBytes, &contract)
	if err != nil {
		return shim.Success(ErrorResult("签名失败Failed to unmarshal contarct" + err.Error(),"05"))
	}
	//获取用户信息
	uBytes, err := stub.GetState(Tel)
	var user User
	err = json.Unmarshal(uBytes, &user)
	if err != nil {
		return shim.Success(ErrorResult("Failed to unmarshal user and: " + err.Error(),"05"))
	}
	//时间戳
	currentTime := time.Unix(time.Now().UnixNano()/1e9, 0)
	contract.TimeStamp = currentTime.String()
	//上链，调用此方法
	i := add(stub,contract,user)
	if i != "0"{
		return shim.Success(ErrorResult("签名失败Failed to sign the contract","i="+i))
	}
	return shim.Success(SuccessResult("签名成功sign successfully"))
}

//===========================
//通过key查看历史记录 args: Tel
//===========================
func (t *SimpleChaincode) getHistoryForKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Success(ErrorResult("function need one args","01"))
	}
	Tel := args[0]
	HisInterface, err := stub.GetHistoryForKey(Tel)
	if err != nil {
		return shim.Success(ErrorResult("查看历史记录失败Failed to get hisInterface!","07"))
	}
	hbytes, err := getHistoryListResult(HisInterface)
	if err != nil {
		return shim.Success(ErrorResult("查看历史记录失败Failed to get historyListResult!","08"))
	}
	//return shim.Success(hbytes)//???????????????????
	return shim.Success(SuccessResult("history query:" + string(hbytes))) //??????????????????
}

//==============================
//获取列表结果，接getHistoryForKey
//==============================
func getHistoryListResult(resultsIterator shim.HistoryQueryIteratorInterface) ([]byte, error) {

	defer resultsIterator.Close()
	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		item, _ := json.Marshal(queryResponse)
		buffer.Write(item)
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	fmt.Printf("queryResult:\n%s\n", buffer.String())
	return buffer.Bytes(), nil
}

//主函数
//main函数
//调用shim包的Start方法，启动chaincode，如果启动成功，这个函数会一直阻塞在这个地方，不会退出。
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting SimpleChaincode:%s", err)
	}
}
