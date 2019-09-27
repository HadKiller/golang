package model

type Request struct {
	URL string
	ParserFuc func(*[]byte)ParserResult
}
const (
男  =iota
女
)
type ParserResult struct {
	NextRequest []Request
	Data []interface{}
}
type User struct {
	Age,Height,IsRecommend,MemberID,Sex int
	AvatarURL,Car,Children,Constellation,Education,House,IntroduceContent,LastModTime,Marriage string
	NickName,Occupation,Salary,WorkCity string
}

