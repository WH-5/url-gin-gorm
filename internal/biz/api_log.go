package biz

func (u *Url) ApiLOG(ip, userAgent, reqUrl, method string) error {
	err := u.dbClient.AddAccess(ip, userAgent, reqUrl, method)
	if err != nil {
		return err
	}
	return nil
}
