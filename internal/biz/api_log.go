package biz

func (u *Url) ApiLOG(ip, userAgent, api, reqUrl string) error {
	err := u.dbClient.AddAccess(ip, userAgent, api, reqUrl)
	if err != nil {
		return err
	}
	return nil
}
