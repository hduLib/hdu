package zjooc

func allPages[T any](u *User, getNextURL func() string) ([]T, error) {
	var items []T
	for no := 1; ; no++ {
		url := getNextURL()
		resp := new(Resp[[]T])
		err := u.get(url, resp)
		if err != nil {
			return nil, err
		}
		for _, v := range resp.Data {
			items = append(items, v)
		}
		if len(items) < no*pageSize {
			break
		}
	}
	return items, nil
}
