package biz

func (u *Url) ApiLOG(ip, userAgent, reqUrl, method, full string) error {
	err := u.dbClient.AddAccess(ip, userAgent, reqUrl, method, full)
	if err != nil {
		return err
	}
	return nil
}
