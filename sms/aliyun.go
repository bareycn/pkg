package sms

type AliyunConfiguration struct {
	AccessKeyId     string `mapstructure:"access_id"`
	AccessKeySecret string `mapstructure:"access_secret"`
	Endpoint        string `mapstructure:"endpoint"`  // 短信服务节点
	IsGlobe         bool   `mapstructure:"is_globe"`  // 是否是国际短信
	SenderID        string `mapstructure:"sender_id"` // 国际短信需要发送者ID
}

type Aliyun struct {
	conf *AliyunConfiguration
}

func NewAliyun(config *AliyunConfiguration) (Sms, error) {
	return nil, nil
}

// Send 发送国内短信
func (a *Aliyun) Send(phone, template string, data map[string]interface{}) error {
	return nil
}

// SendGlobe 发送国际短信
func (a *Aliyun) SendGlobe(phone, template string, data map[string]interface{}) error {
	return nil
}
