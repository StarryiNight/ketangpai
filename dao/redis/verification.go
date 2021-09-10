package redis

import "time"

const VerificationTime = 30 * 60 * time.Second
func SetVerificationCode(username string, code string) error {
	pipeline := client.TxPipeline()
	pipeline.Set(username,code, VerificationTime)
	_, err := pipeline.Exec()
	if err != nil {
		return err
	}
	return nil
}

func GetVerificationCode(username string) (code string , err error  ){
	code = client.Get(username).Val()
	return code,nil
}

func DelVerificationCode(username string)error{
	client.Del(username)
	return nil
}