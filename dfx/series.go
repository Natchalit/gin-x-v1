package dfx

func (s *SeriesX) IsNan() []bool {
	return s.Series.IsNaN()
}

func (s *SeriesX) HasNan() bool {
	return s.Series.HasNaN()
}

func (s *SeriesX) Mean() float64 {
	return s.Series.Mean()
}

func (s *SeriesX) Copy() (*SeriesX, error) {
	cp := s.Series.Copy()
	if cp.Err != nil {
		return nil, cp.Err
	}

	return &SeriesX{
		Series: cp,
	}, nil
}

func (s *SeriesX) Records() []string {
	return s.Series.Records()
}
