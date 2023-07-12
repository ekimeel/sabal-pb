package pb

func (x *Metric) Unix() int64 {
	if x.Timestamp == nil {
		return 0
	}
	return x.Timestamp.GetSeconds()
}
