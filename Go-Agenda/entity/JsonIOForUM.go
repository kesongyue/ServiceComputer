package entity

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

/*
 * 将error检测接口，可以检测是哪个函数出的问题
 */
type Check interface {
	CheckError(err error)
}

type mystring struct {
	name string
}

/*
 * 用自己的结构体mystring，实现CheckError方法
 */

func (funcName mystring) CheckError(err error) {
	if err != nil {
		fmt.Println("Can not run %s", funcName.name)
		panic(err)
	}
}

/*
 * 将json字符流，解析成User对象
 */
func UserJsonDecode(userByte []byte) User {
	var user User
	err := json.Unmarshal(userByte, &user)
	var check Check = mystring{"UserJsonDecode"}
	check.CheckError(err)
	return user
}

/*
 * 从json文件，读取到[]User。多个用户信息
 */
func UReadFromJsonFile(userFile string) []User {
	var users []User
	inputFile, err := os.Open(userFile)
	var check Check = mystring{"UReadFromJsonFile"}
	check.CheckError(err)
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	lineCounter := 0
	for {
		line, err := inputReader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		lineCounter++
		users = append(users, UserJsonDecode([]byte(line)))
	}
	return users
}

/*
 * 将User对象，转换成字符流
 */
func UserJsonEncode(user User) []byte {
	var userByte []byte
	userByte, err := json.Marshal(user)
	var check Check = mystring{"UserJsonEncode"}
	check.CheckError(err)
	return userByte
}

/*
 * 将[]User对象，写到指定json文件
 */
func UWriteToJsonFile(users []User, userFile string) {
	outputFile, err := os.OpenFile(userFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	var check Check = mystring{"UWriteToJsonFile"}
	check.CheckError(err)
	for num := range users {
		var userByte []byte = UserJsonEncode(users[num])
		userByte = append(userByte, '\n')
		outputFile.Write(userByte)
	}
	outputFile.Close()
}

/*
 * 将Meeting字符流，解析成User对象
 */
func MeetingJsonDecode(meetingByte []byte) Meeting {
	var meeting Meeting
	err := json.Unmarshal(meetingByte, &meeting)
	var check Check = mystring{"MeetingJsonDecode"}
	check.CheckError(err)
	return meeting
}

/*
 * 从json文件，读取到[]Meeting，多个会议信息
 */
func MReadFromJsonFile(meetingFile string) []Meeting {
	var meetings []Meeting
	inputFile, err := os.Open(meetingFile)
	var check Check = mystring{"MReadFromJsonFile"}
	check.CheckError(err)
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	lineCounter := 0
	for {
		line, err := inputReader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		lineCounter++
		meetings = append(meetings, MeetingJsonDecode([]byte(line)))
	}
	return meetings
}

/*
 * 将Meeting对象，转换成字符流
 */
func MeetingJsonEncode(meeting Meeting) []byte {
	var meetingByte []byte
	meetingByte, err := json.Marshal(meeting)
	var check Check = mystring{"MeetingJsonEncode"}
	check.CheckError(err)
	return meetingByte
}

/*
 * 将[]Meeting对象，写到指定json文件
 */
func MWriteToJsonFile(meetings []Meeting, meetingFile string) {
	outputFile, err := os.OpenFile(meetingFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	var check Check = mystring{"MWriteToJsonFile"}
	check.CheckError(err)
	for num := range meetings {
		var meetingByte []byte = MeetingJsonEncode(meetings[num])
		meetingByte = append(meetingByte, '\n')
		outputFile.Write(meetingByte)
	}
	outputFile.Close()
}
